# Toolsgo
## Commands:
### reset database
    $ ~/go/bin/waynego db:reset
    
### start http server
    $ ~/go/bin/waynego server

### start nsq consuming worker
    $ ~/go/bin/waynego nsq <topic-name>

### create ju match data & insert into redis
    $ ~/go/bin/waynego ju <match_id>

### insert testing data to redis
    $ ~/go/bin/waynego msg match|offer|bp|bo

### add topics to nsqd
    $ ~/go/bin/waynego nsq:topic name1,name2,...

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