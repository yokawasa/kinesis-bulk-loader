# Change Log

All notable changes to the "kinesis-bulk-loader" will be documented in this file.

## v0.0.3

- Add `--endpoint-url` to specify the URL to send the API request to. For most cases, the AWS SDK automatically determines the URL based on the selected service and the specified AWS Region

**NOTE**
From v0.0.3, you can specify the url to send API requests to. For example, you can give the localstack endpoint (http://localhost:4566) when you develop kinesis application using localstack like this:
```
kinesis-bulk-loader -stream foo-test01 -region ap-northeast-1 -k 123 -m xxx -endpoint-url http://localhost:4566
``` 

## v0.0.2

- Add Verbose message on PutRecord result that includes SequenceNumber and ShardId

## v0.0.1

- Initial release
