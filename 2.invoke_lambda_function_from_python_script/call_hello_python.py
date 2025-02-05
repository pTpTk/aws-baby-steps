import boto3

session = boto3.Session(profile_name='my-dev-profile')
client = session.client('lambda')
response = client.invoke(
    FunctionName='hello-world-python',
    # InvocationType='Event', # async invocation
    InvocationType='RequestResponse', # sync invocation
    Payload='{"key1": "vallll", "key2": "vaaaallll", "key3": "vvvvvvallll"}'
)
payload = response['Payload'].read()
print(payload)
