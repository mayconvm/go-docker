#!/bin/ash

echo "Start project ""github.com/${NAME_PROJECT}"

if [ ! -f "go.mod" ]; then
    go mod init "github.com/${NAME_PROJECT}"
    go mod tidy

    exit 0
fi

echo "Starting"

echo "Download dependencies"
go mod download

echo "Starting"
go run *.go

# while true; do
# sleep 2
# done
