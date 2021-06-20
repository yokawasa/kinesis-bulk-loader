#!/bin/bash

stream_name=test-kds01
region=ap-northeast-1
connections=10
numcalls=10
retry=1

./dist/kinesis-bulk-loader_darwin \
  -stream ${stream_name} \
  -region ${region} \
  -c ${connections} \
  -n ${numcalls} \
  -r ${retry} \
  -k testkey \
  -m testbody \
  -verbose \

