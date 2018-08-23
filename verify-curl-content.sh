#!/bin/sh
#verify curl content @123
# kubectl delete pod `kubectl get pods -n qb-qa-10 | grep ops-fs | awk '{ print $1 }' |grep -v NAME | head -1 ` -n  qb-qa-10

cd /root/t/ops-fs-test
while true; do 
  out="$( curl -s ops-fs-test.wk.qianbao-inc.com )"
  if [ "x$out" != "x$prev" ]; then
    extra="changed"
  else
    extra=""
  fi
  echo "$( date +%F_%T.%N )" $out $extra
  prev="$out"
  sleep 1
done
