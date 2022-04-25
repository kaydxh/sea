# Created by kayxhding on 2020-10-11 19:34:08
#!/usr/bin/env bash

# Fail on any error
set -euo pipefail

# Get the git commit
GIT_COMMIT=$(git rev-parse HEAD)
GIT_TAG=$(git describe --long --tags --dirty --tags --always)
GIT_BUILD_TIME=$(TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

export CGO_ENABLED=0
export GOFLAGS="-mod=readonly"

OUT_PUT_PATH=output
OUT_PUT_BIN_PATH=${OUT_PUT_PATH}/bin
echo "==> Building..."
rm -rf ${OUT_PUT_PATH}/*
#go build -mod=vendor -o ${OUT_PUT_PATH}/sealet ./cmd/sealet
go mod tidy
go build -o ${OUT_PUT_BIN_PATH}/sealet ./cmd/sealet
cp -rf conf ${OUT_PUT_PATH}

# Done!
echo
echo "==> Results:"
ls -hl ${OUT_PUT_PATH}/
