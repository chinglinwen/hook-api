#!/bin/sh
GOOS=linux go build
tar -czf hook-api.tar.gz hook-api
curl -s fs.qianbao-inc.com/k8s/soft/uploadapi -F file=@hook-api.tar.gz -F truncate=yes
cksum ./hook-api