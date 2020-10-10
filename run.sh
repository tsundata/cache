#!/usr/bin/env
trap "rm server;kill 0" EXIT

go build -o server cmd/main.go
./server -port=8001 &
./server -port=8002 &
./server -port=8003 -api=1 &

sleep 2
echo ">>> start test"
curl "http://localhost:5000/api?key=Tom" &
curl "http://localhost:5000/api?key=Tom" &
curl "http://localhost:5000/api?key=Tom" &

wait
