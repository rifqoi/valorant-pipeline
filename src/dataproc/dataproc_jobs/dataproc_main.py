import os
import re
import json
from typing import Text

from google.cloud import storage
import pyspark
from pyspark.sql import SparkSession
import pyspark.sql.functions as F


class ValorantTransform:
    def __init__(
        self,
        landing_bucket: str,
        processed_bucket: str,
        spark: SparkSession,
        target_deployment: str,
    ):
        self.spark = spark
        self.load_path = f"gs://{landing_bucket}"
        self.save_path = f"gs://{processed_bucket}"
        self.target_deployment = target_deployment

    def transform_data_array(self, json_blob: storage.Blob) -> None:
        # Checking multiline json or not
        with storage.fileio.BlobReader(blob=json_blob) as f:
            if len(f.readlines()) != 1:
                df = self.spark.read.option("multiline", "true").json(
                    f"{self.load_path}/{json_blob.name}"
                )
            else:
                df = self.spark.read.json(f"{self.load_path}/{json_blob.name}")

        data = df.withColumn("data", F.explode(df.data)).select("data.*")
        players = (
            data.withColumn("all_players", F.explode("players.all_players"))
            .select("all_players", "Metadata")
            .select("all_players.*", "Metadata.*")
        )
        self.current_json = json_blob.name
        self.data = data
        self.players = players

    def validate_file(self, json_blob: storage.Blob) -> None:
        try:
            if self.current_json != json_blob.name:
                self.transform_data_array(json_blob)
        except:
            self.transform_data_array(json_blob)

    def write_csv(
        self,
        df: pyspark.sql.DataFrame,
        target_path: str,
    ) -> None:
        print(f"{self.save_path}/{target_path}")
        if self.target_deployment == "local":
            df.write.csv(
                target_path,
                mode="append",
                header=True,
                timestampFormat="yyyy-MM-dd HH:mm:ss",
                quote='"',
                escape='"',
            )
        elif self.target_deployment == "gcp":
            df.write.csv(
                f"{self.save_path}/{target_path}",
                mode="append",
                header=True,
                timestampFormat="yyyy-MM-dd HH:mm:ss",
                quote='"',
                escape='"',
            )

    def transform_matches_details(self, json_blob: storage.Blob) -> None:
        self.validate_file(json_blob)

        matches_details = (
            self.data.select("Metadata", "teams.*")
            .selectExpr(
                "Metadata", "stack(2, 'blue', blue, 'red', red) as (team, stats)"
            )
            .selectExpr("Metadata.*", "team", "stats.*")
            .withColumn(
                "game_start",
                F.from_utc_timestamp(F.from_unixtime("game_start"), "Asia/Jakarta"),
            )
            .withColumn(
                "game_length",
                F.from_unixtime(F.col("game_length").cast("int") / 1000, "HH:mm:ss"),
            )
            .select(
                "matchid",
                "season_id",
                "region",
                "game_start",
                "game_length",
                "mode",
                "team",
                "has_won",
                "rounds_played",
                "rounds_lost",
                "rounds_won",
            )
        )

        self.write_csv(matches_details, "match_details")
        matches_details.show()

    def transform_parties(self, json_blob: storage.Blob) -> None:
        self.validate_file(json_blob)

        parties = (
            self.players.groupBy("party_id")
            .agg(
                F.collect_list(
                    F.struct(
                        F.col("matchid"),
                        F.col("puuid"),
                        F.col("name"),
                        F.col("tag"),
                        F.col("team"),
                    )
                ).alias("new")
            )
            .select(
                "party_id",
                F.explode("new"),
            )
            .select("party_id", "col.*")
        )

        self.write_csv(parties, "parties")
        parties.show()

    def transform_players(self, json_blob: storage.Blob) -> None:
        self.validate_file(json_blob)

        players = (
            self.data.withColumn("players", F.explode("players.all_players"))
            .select("players.*", "Metadata.*")
            .select(
                "puuid",
                "name",
                "tag",
            )
            .dropDuplicates()
            .orderBy("name")
        )

        self.write_csv(players, "players")
        players.show()

    def transform_player_stats(self, json_blob: storage.Blob) -> None:
        self.validate_file(json_blob)

        player_stats = self.players.select("*", "behaviour.*", "stats.*").select(
            "matchid",
            "puuid",
            "party_id",
            "name",
            "tag",
            "character",
            "currenttier",
            "currenttier_patched",
            "damage_made",
            "damage_received",
            "score",
            "kills",
            "deaths",
            "assists",
            "afk_rounds",
        )
        self.write_csv(player_stats, "player_stats")
        player_stats.show()


def main():
    landing_bucket = "valorant_landing_bucket_erudite-bonbon-352111"
    processed_bucket = "valorant_process_bucket_erudite-bonbon-352111"
    data_lake_bucket = "valorant_data_lake_erudite-bonbon-352111"

    storage_client = storage.Client()
    prefix = "Player/"

    landing_blobs = [
        (blob, blob.updated)
        for blob in storage_client.list_blobs(
            landing_bucket,
            prefix=prefix,
        )
        if re.search(r"Player\/.*\/matches\/.*.json", blob.name)
    ]

    spark = (
        SparkSession.builder.master("local[*]")
        .appName("ValorantTransform")
        .config(
            "spark.jars.packages",
            "com.google.cloud.bigdataoss:gcs-connector:hadoop2-1.9.17",
        )
        .config(
            "spark.jars.excludes",
            "javax.jms:jms,com.sun.jdmk:jmxtools,com.sun.jmx:jmxri",
        )
        .config("spark.driver.userClassPathFirst", "true")
        .config("spark.executor.userClassPathFirst", "true")
        .config(
            "spark.hadoop.fs.gs.impl",
            "com.google.cloud.hadoop.fs.gcs.GoogleHadoopFileSystem",
        )
        .config(
            "google.cloud.auth.service.account.json.keyfile",
            os.getenv("HOME") + "/.google/credentials/google_credentials.json",
        )
        .config("google.cloud.auth.service.account.", "true")
        .config("spark.sql.session.timeZone", "UTC")
        .getOrCreate()
    )

    vtransform = ValorantTransform(
        landing_bucket=landing_bucket,
        processed_bucket=processed_bucket,
        spark=spark,
        target_deployment="gcp",
    )

    count = 0
    for blob, datetime in landing_blobs:
        if count == 2:
            break
        print(blob.name)
        if blob.name.strip().endswith(".json"):
            vtransform.transform_matches_details(blob)
            vtransform.transform_parties(blob)
            vtransform.transform_players(blob)
            vtransform.transform_player_stats(blob)


if __name__ == "__main__":
    main()
