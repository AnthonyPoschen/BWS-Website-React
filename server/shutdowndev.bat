@echo off

docker stop api balancer webfront
docker rm api balancer webfront