#!/bin/bash

CLUSTER_NAME=$1
gcloud dataproc clusters create $1\
    --region asia-southeast1 \
    --zone asia-southeast1-a \
    --single-node \
    --master-machine-type e2-standard-2 \
    --master-boot-disk-size 31 \
    --image-version 2.0-debian10 \
    --metadata PIP_PACKAGES="google-cloud-storage" \
    --initialization-actions 'gs://goog-dataproc-initialization-actions-asia-southeast1/python/pip-install.sh' \
    --project erudite-bonbon-352111
