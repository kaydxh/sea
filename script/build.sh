# Created by kayxhding on 2020-10-11 19:34:08
#!/usr/bin/env bash

# exit by command return non-zero exit code
set -o errexit
# Indicate an error when it encounters an undefined variable
set -o nounset
# Fail on any error.
set -o pipefail
# set -o xtrace

#SEA_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)

# Get the git commit
# GIT="git --work-tree"
GIT="git"
GIT_COMMIT=$("${GIT}" rev-parse "HEAD^{commit}")
GIT_TAG=$(git describe --long --tags --dirty --tags --always)
GIT_BUILD_TIME=$(TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')
GIT_DIRTY=$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)

GIT_TREE_STATE=${GIT_TREE_STATE-}
if git_status=$("${GIT}" status --porcelain) && [[ -z ${git_status} ]]; then
   GIT_TREE_STATE="clean"
else
   GIT_TREE_STATE="dirty"
fi


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
