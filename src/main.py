import os
import re
from google.cloud import storage
from pyspark.sql import SparkSession
from valorant_transform.transform import ValorantTransform


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
        .config("google.cloud.auth.service.account.", "true")
        .config(
            "google.cloud.auth.service.account.json.keyfile",
            os.getenv("HOME") + "/.google/credentials/google_credentials.json",
        )
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
