#!/usr/bin/env bash



mkdir -p /root/www/go_www/src/zuji
cd /root/www/go_www/src/zuji
git clone ssh://git@git-dp.nqyong.com:3022/Zuji/goCommon.git common

go get -v github.com/gin-gonic/gin
go get -v github.com/google/uuid