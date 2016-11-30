@echo off

docker build -t api ./api
docker build -t balancer ./balancer
start cmd /c "echo balancer &&   docker run --publish 8000:8000 --name balancer --rm balancer"
start cmd /c "echo api &&        docker run --publish 8001:8001 --name api --rm api"
echo "Giving time for servers to be booted before linking"
timeout /t 10
docker network connect mynetwork api
docker network connect mynetwork balancer