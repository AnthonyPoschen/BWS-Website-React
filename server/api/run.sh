echo 'Starting Api'
cd "${0%/*}"
go run api.go middleware.go types.go