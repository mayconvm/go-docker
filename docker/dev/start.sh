#!/bin/ash

echo "Start project ""github.com/${NAME_PROJECT}"

if [ ! -f "go.mod" ]; then
    go mod init "github.com/${NAME_PROJECT}"
    go mod tidy

    exit 0
fi

echo "Starting"

while true; do
sleep 2
done