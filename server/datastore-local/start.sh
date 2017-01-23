IP=`ifconfig eth0 | grep "inet addr" | cut -d ':' -f 2 | cut -d ' ' -f 1`
echo $IP
tcp-proxy -l="$IP:8400" -r "127.0.0.1:8400" &
./gcd/gcd.sh "$@"
pkill -P $$