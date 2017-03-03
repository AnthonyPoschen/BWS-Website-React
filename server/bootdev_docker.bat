@echo off

robocopy ./../frontend/dist ./webfront/www /COPYALL /E


:: Boot all services in seperate prompts which will recieve the stdout for monitoring.
:: start cmd /c "echo datastore &&	docker run --network=mynetwork --publish 8400:8400 --name datastore brainwave-studios/datastore" 
start cmd /c "echo balancer &&  docker run --network=mynetwork --publish 8000:8000 --name balancer brainwave-studios/balancer"
start cmd /c "echo api &&       docker run --network=mynetwork --publish 8001:8001 -e DATASTORE_EMULATOR_HOST='datastore:8400' --name api brainwave-studios/api"
start cmd /c "echo webfront &&	docker run --network=mynetwork --publish 8002:8002 --name webfront brainwave-studios/webfront"
start cmd /c "echo DynamoDB && 	docker run --network=mynetwork --publish 8400:8000 --name dynamodb forty8bit/dynamodb-local -sharedDb

:: Notes
:: Create Docker network as a bridged network.
:: be careful messing with datastore that was painful to solve.

docker run -p 8400:8000 --name dynamodb forty8bit/dynamodb-local -sharedDb