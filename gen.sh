#!/bin/bash

basepath=$(cd `dirname $0`; pwd)
echo $basepath
mkdir -p $basepath/c2s
protoc -I $basepath --go_out=./c2s $basepath/protobuf/protofiles/c2s.proto
mkdir -p $basepath/cs
protoc -I $basepath --go_out=./cs $basepath/protobuf/protofiles/cs.proto

#protoc --go_out=./c2s --go_opt=paths=source_relative $basepath/c2s/c2s.proto
egf-pb generate
