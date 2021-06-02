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

## client js pb
```
sudo npm install protobufjs -g
sudo npm install semver -g
sudo npm --silent install chalk@^4.0.0 jsdoc@^3.6.3 minimist@^1.2.0 uglify-js@^3.7.7 -g
升级node
sudo npm install -g n
sudo n 8.16.2

```
```
sudo npm install egf-protobuf -g
egf-pb init
egf-pb generate
```
