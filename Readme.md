# how to gen pb code

## win10

```
go get -u github.com/golang/protobuf/protoc-gen-go
download protoc from https://github.com/protocolbuffers/protobuf/releases

```

## ubuntu
```
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./c2s/c2s.proto

```
