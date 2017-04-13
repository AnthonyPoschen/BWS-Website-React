echo 'Starting database'
cd "$(dirname $(readlink -f $0))"
go run database.go middleware.go mux.go tablecreate.go types.go utility.go --dev=1