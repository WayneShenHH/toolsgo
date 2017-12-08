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



























## Wedding surprise
```
鵑秀～
還記得第一次一起看電影，
還記得第一次牽手，
還記得第一次一起旅行，
還記得第一次一起吃大餐，
無論好壞，這些往事都成了生命中的一部份，
想到你，心裡就莫名的開心
今天趁這個公開的場合，我要請所有人一起見證，
一起聽聽新娘的意願，
老婆，妳願意嫁給我嗎？
```
