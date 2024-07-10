#!/usr/bin/env sh
# author: dijitol
# description: Generate and open coverage report in browser. Should be run from the root of the project

go test -coverprofile=coverage.out ./... 
go tool cover -html=coverage.out
