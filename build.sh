#!/bin/sh
set -ex
go install -v -mod vendor -ldflags '-s -w -extldflags "-static"' -trimpath
upx -9 $(which xmlq)

# go build -v -mod vendor -ldflags '-s -w -linkmode "external" -extldflags "-static"' -trimpath
# upx -9 xmlq
