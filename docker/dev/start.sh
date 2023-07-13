#!/bin/ash

if [ -f "go.mod" ]; then
    go mod tidy
fi

echo "Starting"

while true; do
sleep 10
done