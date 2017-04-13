#!/bin/bash
echo Starting WebFront
cd "$(dirname $(readlink -f $0))"
# now we change directory to the proper directory to run a live webserver
cd ./../../frontend
npm start