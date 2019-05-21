# Toolsgo
## Commands:
### reset database
    $ ~/go/bin/waynego db:reset
    
### start http server
    $ ~/go/bin/waynego server

### create ju match data & insert into redis
    $ ~/go/bin/waynego ju <match_id>

### insert testing data to redis
    $ ~/go/bin/waynego msg match|offer|bp|bo

### start nsq consuming worker
    $ ./toolsgo nsq <topic-name> <channel-name>

### produce a message to nsqd
    $ ./toolsgo nsq:msg <topic-name> <message>

### add topics to nsqd
    $ ./toolsgo topic:add name1,name2,...

### get all topics
    $ ./toolsgo topic:all

## start nsq services
    
1. nsqlookupd
```
$ /path-to-nsq-bin/nsqlookupd
```
2. nsqd
```
$ /path-to-nsq-bin/nsqd --lookupd-tcp-address=127.0.0.1:4160 --broadcast-address=127.0.0.1
```
3. nsqadmin
```
$ /path-to-nsq-bin/nsqadmin --lookupd-http-address=127.0.0.1:4161
```
4. nsq_to_file
```
$ /path-to-nsq-bin/nsq_to_file --topic=<topic-name> --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
```

## start gRPC services

1. install package
```
download plugin: https://github.com/protocolbuffers/protobuf/releases
$ go get github.com/golang/protobuf@master
```

2. generate *.pb.go
```
protoc ./pb/*.proto --go_out=plugins=grpc:.
```