@echo off
echo "Starting DB Service"
go run database.go middleware.go mux.go tablecreate.go types.go utility.go --dev=1