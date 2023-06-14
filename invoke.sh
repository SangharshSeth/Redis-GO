#!/bin/bash

echo "...........Running Server............"

go run src/main.go

redis-cli SET name $1

