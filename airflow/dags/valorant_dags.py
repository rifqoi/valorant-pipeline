import logging

from airflow import DAG
from airflow.utils.dates import days_ago
from airflow.operators.bash import BashOperator
from airflow.operators.python import PythonOperator
from airflow.operators.dummy_operator import DummyOperator

from airflow.providers.google.cloud.operators.bigquery import (
    BigQueryCreateExternalTableOperator,
)
from airflow.providers.google.cloud.operators.dataproc import (
    DataprocSubmitPySparkJobOperator,
)
from google.cloud import storage


storage_client = storage.Client()
landing_bucket = "valorant_landing_bucket_erudite-bonbon-352111"
data_lake_bucket = "valorant_data_lake_erudite-bonbon-352111"
process_bucket = "valorant_process_bucket_erudite-bonbon-352111"

default_args = {
    "owner": "airflow",
    "start_date": days_ago(1),
    "depends_on_past": False,
    "retries": 1,
}

with DAG(
    dag_id="valorant_dag",
    description="Valorant Data Pipeline DAG",
    schedule_interval="@daily",
    default_args=default_args,
    catchup=False,
    max_active_runs=1,
    tags=["valorant"],
) as dag:

    start_operator = DummyOperator(
        task_id="Begin_execution",
    )

    ingest_data_from_api = BashOperator(
        task_id="ingest_data_from_api",
        bash_command="/usr/local/bin/ingest",
    )

    spark_process_json_data = DataprocSubmitPySparkJobOperator(
        main="gs://dtc_data_lake_erudite-bonbon-352111/dataproc_main.py",
        task_id="process_json_data_with_spark",
        cluster_name="valorant-cluster",
        region="asia-southeast1",
        job_name="transform_valorant_json_files",
    )
    from gcs_tasks import move_blob

    json_pattern = r"Player\/.*/matches\/.*.json"
    move_json_to_data_lake = PythonOperator(
        task_id="move_json_from_landing_to_data_lake",
        python_callable=move_blob,
        op_args=[
            storage_client,
            landing_bucket,
            data_lake_bucket,
            json_pattern,
        ],
    )

    csv_pattern = r".*\/.*\.csv"
    move_processed_csv_to_data_lake = PythonOperator(
        task_id="move_processed_csv_to_data_lake",
        python_callable=move_blob,
        op_args=[
            storage_client,
            process_bucket,
            data_lake_bucket,
            csv_pattern,
        ],
    )
    end_operator = DummyOperator(
        task_id="End_execution",
    )

    (
        start_operator
        >> ingest_data_from_api
        >> spark_process_json_data
        >> move_json_to_data_lake
        >> move_processed_csv_to_data_lake
        >> end_operator
    )
