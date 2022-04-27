# Created by kayxhding on 2020-10-11 19:34:08
#!/usr/bin/env bash

# exit by command return non-zero exit code
set -o errexit
# Indicate an error when it encounters an undefined variable
set -o nounset
# Fail on any error.
set -o pipefail
# set -o xtrace

SEA_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
SEA_OUTPUT_SUB_PATH=${SEA_OUTPUT_SUB_PATH:-output}
SEA_OUTPUT_PATH=${SEA_ROOT}/${SEA_OUTPUT_SUB_PATH}
SEA_OUTPUT_BIN_PATH=${SEA_OUTPUT_PATH}/bin

source "${SEA_ROOT}/script/version.sh"

export CGO_ENABLED=0
export GOFLAGS="-mod=readonly"

function platform() {
  echo "$(go env GOHOSTOS)/$(go env GOHOSTARCH)"
}

echo "==> Building in $(platform)"

function build() {
  local goldflags
  goldflags="all=$(ldflags) ${GOLDFLAGS:-}"
  echo "==${goldflags}"

  rm -rf ${SEA_OUTPUT_PATH}/*
  #go build -mod=vendor -o ${OUT_PUT_PATH}/sealet ./cmd/sealet
  go mod tidy
  go build -o ${SEA_OUTPUT_BIN_PATH}/sealet ./cmd/sealet
  cp -rf conf ${SEA_OUTPUT_PATH}

}

build

# Done!
echo
echo "==> Results:"
ls -hl ${SEA_OUTPUT_PATH}/
