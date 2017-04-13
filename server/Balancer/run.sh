echo 'Starting Balancer'
cd "$(dirname $(readlink -f $0))"
go run balancer.go