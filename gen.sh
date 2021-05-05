#!/bin/bash
basepath=$(cd `dirname $0`; pwd)
echo $basepath
#protoc -I $basepath --go_out=$basepath $basepath/c2s.proto
protoc --go_out=plugins=grpc:./c2s --go_opt=paths=source_relative $basepath/c2s/c2s.proto
