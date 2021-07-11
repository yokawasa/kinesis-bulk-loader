# kinesis-bulk-loader

[![Upload Release Asset](https://github.com/yokawasa/kinesis-bulk-loader/actions/workflows/release.yml/badge.svg)](https://github.com/yokawasa/kinesis-bulk-loader/actions/workflows/release.yml)

A Golang tool that put bulk messages in parallel to AWS Kinesis Data Stream.

## Usage

```
kinesis-bulk-loader [options...]

Options:
-stream string       (Required) Kinesis stream name
-region string       Region for Kinesis stream
                     By default "ap-northeast-1"
-k string            (Required) Partition key
-k string            (Required) Partition key
-m string            (Required) Message payload to put into the stream
-c connections       Number of parallel simultaneous Kinesis session
                     By default 1; Must be more than 0
-n num-calls         Run for exactly this number of calls by each Kinesis session
                     By default 1; Must be more than 0
-r retry-num         Number fo Retry in each message send
                     By default 1; Must be more than 0
-endpoint-url string The URL to send the API request to
                     By default "", which mean the AWS SDK automatically determines the URL
-version             Prints out build version information
-verbose             Verbose option
-h                   help message
```

## Download

You can download the compiled command with [downloader](https://github.com/yokawasa/kinesis-bulk-loader/blob/main/downloader) like this:

```
# Download latest command
./downloader

# Download the command with a specified version
./downloader v0.0.3
```
Or you can download it on the fly with the following commmand:

```
curl -sS https://raw.githubusercontent.com/yokawasa/kinesis-bulk-loader/main/downloader | bash --
```


Output would be like this:
```
Downloading kinesis-bulk-loader_darwin_amd64 ...
Archive:  kinesis-bulk-loader.zip
  inflating: kinesis-bulk-loader_darwin_amd64
kinesis-bulk-loader
Downloaded into kinesis-bulk-loader
Please add kinesis-bulk-loader to your path; e.g copy paste in your shell and/or ~/.profile
```

## Execute the command

```bash
stream_name=test-kds01
region=ap-northeast-1
connections=10
numcalls=10
retry=1

kinesis-bulk-loader \
  -stream ${stream_name} \
  -region ${region} \
  -c ${connections} \
  -n ${numcalls} \
  -r ${retry} \
  -k testkey \
  -m testbody \
  -verbose \
```

Sample output would be like this:
```
[Verbose] Mssage: PartitionKey testkey Data testbody9IVUWSP2Nc
[Verbose] Mssage: PartitionKey testkey Data testbodyBpLnfgDsc2
[Verbose] Mssage: PartitionKey testkey Data testbodyDkh9h2fhfU
[Verbose] Mssage: PartitionKey testkey Data testbodyUsaD6HEdz0
[Verbose] PutRecord Result: PartitionKey testkey SequenceNumber 49619437532912338680531404070433076316966140521717170194 ShardId shardId-000000000001
[Verbose] PutRecord Result: PartitionKey testkey SequenceNumber 49619437532912338680531404070434285242785755150891876370 ShardId shardId-000000000001
[Verbose] Mssage: PartitionKey testkey Data testbodyDkh9h2fhfU
[Verbose] PutRecord Result: PartitionKey testkey SequenceNumber 49619437532912338680531404070436703094424984409241288722 ShardId shardId-000000000001
[Verbose] PutRecord Result: PartitionKey testkey SequenceNumber 49619437532912338680531404070439120946064213736310177810 ShardId shardId-000000000001
[Verbose] Mssage: PartitionKey testkey Data testbodyWD8F2qNfHK
[Verbose] PutRecord Result: PartitionKey testkey SequenceNumber 49619437532912338680531404070443956649342672253009002514 ShardId shardId-000000000001

...snip...

-----------------------
Kinesis Bulk Loader Summary
-----------------------
Sent messages: 100
Errors: 0
Duration (sec): 2.349315041
Average (ms): 23
```

## Build and Run (For Developer)

To build, simply run `make` like below
```
make

golint /Users/yoichika/dev/github/kinesis-bulk-loader
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-bulk-loader/dist/kinesis-bulk-loader_linux_amd64 kinesis-bulk-loader/src
GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-bulk-loader/dist/kinesis-bulk-loader_darwin_amd64 kinesis-bulk-loader/src
GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o kinesis-bulk-loader/dist/kinesis-bulk-loader_windows_amd64 kinesis-bulk-loader/src
```

Suppose you are using macOS, run the `kinesis-bulk-loader_darwin` (while `kinesis-bulk-loader_linux` if you are using Linux, or `kinesis-bulk-loader_windows` if using Windows) like below

```bash
./dist/kinesis-bulk-loader_darwin_amd64 -stream test-kds01 -k hoge -m test -c 10 -n 100 -verbose
```

Finally clean built commands

```
make clean
```

## Relevant project

The Kinesis Consumer side can be tested with [yokawasa/amazon-kinesis-consumer](https://github.com/yokawasa/amazon-kinesis-consumer)


## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yokawasa/kinesis-bulk-loader.
