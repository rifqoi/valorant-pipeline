import os
import logging

from airflow import DAG
from airflow.utils.dates import days_ago
from airflow.operators.bash import BashOperator
from airflow.operators.python import PythonOperator
from airflow.operators.dummy_operator import DummyOperator

from google.cloud import storage
from airflow.providers.google.cloud.operators.bigquery import (
    BigQueryCreateExternalTableOperator,
)

import pyarrow.csv as pv
import pyarrow.parquet as pq

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

    end_operator = DummyOperator(
        task_id="End_execution",
    )

    start_operator >> ingest_data_from_api >> end_operator
