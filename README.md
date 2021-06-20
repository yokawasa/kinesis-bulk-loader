# kinesis-bulk-loader

A Golang tool that put bulk messages in parallel to AWS Kinesis Data Stream.

## Usage

```
kinesis-bulk-loader [options...]

Options:
-stream string       (Required) Kinesis stream name
-region string       Region for Kinesis stream
                     By default "ap-northeast-1"
-k string            (Required) Partition key
-m string            (Required) Message payload to put into the stream
-c connections       Number of parallel simultaneous Kinesis session
                     By default 1; Must be more than 0
-n num-calls         Run for exactly this number of calls by each Kinesis session
                     By default 1; Must be more than 0
-r retry-num         Number fo Retry in each message send
                     By default 1; Must be more than 0
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
./downloader v0.0.1
```

Output would be like this:
```
Downloading kinesis-bulk-loader_darwin ...
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
[Verbose] Mssage Partition Key testkey Body testbodyBpLnfgDsc2
[Verbose] Mssage Partition Key testkey Body testbodyWD8F2qNfHK
... snip ...
[Verbose] Mssage Partition Key testkey Body testbody9IVUWSP2Nc
[Verbose] Mssage Partition Key testkey Body testbodyWD8F2qNfHK
[Verbose] Mssage Partition Key testkey Body testbodyhV3vC5AWX3

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
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o /Users/yoichika/dev/github/kinesis-bulk-loader/dist/kinesis-bulk-loader_linux /Users/yoichika/dev/github/kinesis-bulk-loader/src
GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=v0.0.1" -o /Users/yoichika/dev/github/kinesis-bulk-loader/dist/kinesis-bulk-loader_darwin /Users/yoichika/dev/github/kinesis-bulk-loader/src
```

Suppose you are using macOS, run the `kinesis-bulk-loader_darwin` (while `kinesis-bulk-loader_linux` if you are using Linux) like below

```bash
./dist/kinesis-bulk-loader_darwin -stream test-kds01 -k hoge -m test -c 10 -n 100 -verbose
```

Finally clean built commands

```
make clean
```

## Relevant project

The Kinesis Consumer side can be tested with [yokawasa/amazon-kinesis-consumer](https://github.com/yokawasa/amazon-kinesis-consumer)


## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/yokawasa/kinesis-bulk-loader.
