cd ~/hook
rm -f hook-api
wget http://fs.qianbao-inc.com/k8s/soft/hook-api.tar.gz
tar -xzf hook-api.tar.gz
rm -f hook-api.tar.gz
chmod +x hook-api
cksum hook-api
./stop.sh && ./start.sh