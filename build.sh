#!/bin/bash
export PWD=`pwd`
export EXECNAME="route7"
export TIME=`date +%Y%m%d%H%M%S`
export GOPATH=`pwd`

go build -o $PWD/bin/$EXECNAME-$TIME
ln -nfs $PWD/bin/$EXECNAME-$TIME $PWD/bin/$EXECNAME