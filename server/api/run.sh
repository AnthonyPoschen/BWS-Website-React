echo 'Starting Api'
cd "$(dirname $(readlink -f $0))"
go run api.go middleware.go types.go