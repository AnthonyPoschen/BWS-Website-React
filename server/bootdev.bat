@echo off

robocopy ./../frontend/dist ./webfront/www /COPYALL /E


:: Boot all services in seperate prompts which will recieve the stdout for monitoring.
:: start cmd /c "echo datastore &&	docker run --network=mynetwork --publish 8400:8400 --name datastore brainwave-studios/datastore" 
start sh ./api/run.sh
start sh ./Balancer/run.sh
start sh ./database/run.sh
start sh ./webfront/run.sh
start cmd /c "echo DynamoDB && 	docker run --network=mynetwork --publish 8400:8000 --name dynamodb --rm forty8bit/dynamodb-local -sharedDb

:: Notes
:: Create Docker network as a bridged network.
:: be careful messing with datastore that was painful to solve.

::docker run -p 8400:8000 --name dynamodb forty8bit/dynamodb-local -sharedDb