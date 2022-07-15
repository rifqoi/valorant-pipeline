import pyspark
import os
from google.cloud import storage
import json
import pyspark.sql.functions as F
from pyspark.sql import SparkSession


class ValorantTransform:
    def __init__(self, landing_bucket: str, processed_bucket: str, spark: SparkSession):
        self.spark = spark
        self.load_path = landing_bucket
        self.save_path = processed_bucket

    def transform_data_array(self, json_blob: str) -> None:
        df = self.spark.read.json(f"gs://{self.load_path}/{json_blob}")
        data = df.withColumn("data", F.explode(df.data)).select("data.*")
        self.data = data

    def transform_matches_details(self, json_blob: str):
        # I don't know how to check whether an attribute is exist...
        try:
            print(self.data)
        except:
            self.transform_data_array(json_blob)

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
        matches_details.show()
