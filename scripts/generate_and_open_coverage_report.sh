#!/usr/bin/env sh
# author: dijitol

go test -coverprofile=coverage.out ../... 
go tool cover -html=coverage.out
