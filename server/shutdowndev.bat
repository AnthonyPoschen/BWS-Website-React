@echo off

docker stop api balancer webfront datastore
docker rm api balancer webfront datastore