#!/bin/bash

yum install -y upx

#mpath=$(dirname $0)
#cd $mpath

## build macos
# CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "${LDFLAGS}" -o bin/deploy

# ## build linux client
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o bin/deploy_client
# upx bin/deploy_client
# mv bin/deploy_client ${docker_path}/deploy_client

## build linux cgo with xgo
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -o crontab_worker
upx crontab_worker
