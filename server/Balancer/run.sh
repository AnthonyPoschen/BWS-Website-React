echo 'Starting Balancer'
echo "$(dirname $(readlink -f $0))"
go run ./balancer.go