#!/bin/sh
#verify upstream content @172.28.40.80
# kubectl delete pod `kubectl get pods -n qb-qa-10 | grep ops-fs | awk '{ print $1 }' |grep -v NAME | head -1 ` -n  qb-qa-10
# kubectl scale deploy ops-fs -n qb-qa-10 --replicas=2; 

file=/apps/soft/nginx/conf.d/upstream/ops_fs_upstream_qb-qa-10.conf
while true; do 
  out="$( cat $file | grep server | awk '{ print $2 }' )"
  if [ "x$out" != "x$prev" ]; then
    extra="changed"
  else
    extra=""
  fi
  echo "$( date +%F_%T.%N )" $out $extra
  prev="$out"
  sleep 1
done
