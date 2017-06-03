@echo off
echo "Starting Webfront Service"

::call "bash ./watch-frontend.bsh"
go run webfront.go -ext="./../../frontend/src"