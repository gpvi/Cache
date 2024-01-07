#!/bin/bash
#trap "rm server;kill 0" EXIT
#
#go build -buildvcs=false -o server
#./server -port=8001 &
##./server -port=8002 &
#./server -port=8003 -api=1 &
#
#sleep 2
#echo ">>> start test"
#curl "http://localhost:9999/api?key=Tom" &
#curl "http://localhost:9999/api?key=Tom" &
#curl "http://localhost:9999/api?key=Tom" &
#
#wait
# Function to clean up and terminate the script
cleanup() {
    rm server
    pkill -P $$  # Send signal to all child processes of the current process
    exit
}

# Set up trap to call cleanup function on EXIT
trap cleanup EXIT

go build -buildvcs=false -o server
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &
curl "http://localhost:9999/api?key=Tom" &

wait
