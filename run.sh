#! /bin/bash

export loglevel="info"
export servicename="zWebsite"
export servicePort=":8081"
export port="8081"

export DBUSER="root"
export DBPW="root"
export DBHOST="127.0.0.1"
export DBPORT=30306

export DBNAME="zwebsite"

export REDISHOST="127.0.0.1"
export REDISPORT="6379"

go run *.go

