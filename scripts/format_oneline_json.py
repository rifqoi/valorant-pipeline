from google.cloud import storage
import json


# first establish your client
storage_client = storage.Client()
bucket = "valorant_landing_bucket_erudite-bonbon-352111"
bucket2 = "valorant_data_lake_erudite-bonbon-352111"
prefix = "Player/"

blobs = [
    (blob, blob.updated)
    for blob in storage_client.list_blobs(
        bucket2,
        prefix=prefix,
    )
]


for blob, datetime in blobs:
    if blob.name.strip().endswith(".json"):
        print(blob.name)
        with storage.fileio.BlobReader(blob=blob) as f:
            json_string = json.loads(f.read())
            text = json.dumps(json_string, separators=(",", ":")).encode("gbk")

        with storage.fileio.BlobWriter(blob=blob) as f:
            f.write(text)
