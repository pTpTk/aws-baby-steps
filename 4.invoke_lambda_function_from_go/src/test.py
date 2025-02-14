import boto3

client = boto3.client('lambda')
response = client.invoke(
    FunctionName='myFunction',
    # InvocationType='Event', # async invocation
    InvocationType='RequestResponse', # sync invocation
    Payload='{"key1": "vallll", "key2": "vaaaallll", "key3": "vvvvvvallll"}'
)
payload = response['Payload'].read()
print(payload)
