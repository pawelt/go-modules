# Golang 1.11+ modules test

First, get all dependencies

    go mod vendor

This should fetch all dependencies into the `./vendor` folder and not install them globally.

Then run

    go run cmd/cmd1/main.go

or build

    go build -ldflags "-s -w"  -o ./bin/cmd1 ./cmd/cmd1

