from google.cloud import storage
import re
from typing import Text


def move_blob(
    client: storage.Client,
    source_bucket_name: Text,
    target_bucket_name: Text,
    pattern: Text,
) -> None:

    source_bucket = client.bucket(source_bucket_name)
    target_bucket = client.bucket(target_bucket_name)

    blobs = [
        blob
        for blob in client.list_blobs(
            source_bucket,
        )
        if re.search(pattern, blob.name)
    ]
    # Move blob by copy the blobs from source bucket to target bucket
    # and then delete the blobs from the source bucket
    for blob in blobs:
        # Copy the blob
        blob_copy = source_bucket.copy_blob(
            blob,
            target_bucket,
            blob.name,
        )
        # Delete the blob
        source_bucket.delete_blob(blob.name)
        print(
            f"Blob: {source_bucket.name}/{blob.name} \n moved to \n {target_bucket.name}/{blob_copy.name}\n"
        )
