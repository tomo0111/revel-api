#!/bin/bash

set -e -u -x

go get -u github.com/golang/dep/cmd/dep

cp -r repository /go/src/github.com/tomoyane/

cd /go/src/github.com/tomoyane/

dep ensure

pwd
ls -la

go test -v go test -v github.com/tomoyane/repository/test/...
