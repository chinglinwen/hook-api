#!/bin/sh
# example start for pre-env
mv h.log{,.bak}
./hook-api  -n qb-pre -nginx-grp BJ-M7 -env pre -upstream http://upstream-pre.sched.qianbao-inc.com &> h.log &