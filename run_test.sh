#!/bin/bash

export GIN_MODE=test
go test -coverpkg="Decision-Tool" -c -tags testrunmain && ./Decision-Tool.test -test.coverprofile=system.out && go tool cover -html=system.out -o system.html
