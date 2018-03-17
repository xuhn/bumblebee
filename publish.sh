#!/usr/bin/env bash
MYDATE=`date +%d%m%y%H%M`
FILE=mytest-$MYDATE
go build -o ../$FILE main.go
if [ -e ../$FILE ]; then
  echo '文件存在';
  mv ../$FILE /data/mytest
  cd /data/mytest
  ln -sf $FILE mytest
  pm2 start mytest
fi
