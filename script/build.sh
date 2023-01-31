# Created by kayxhding on 2020-10-11 19:34:08
#!/usr/bin/env bash

# exit by command return non-zero exit code
set -o errexit
# Indicate an error when it encounters an undefined variable
set -o nounset
# Fail on any error.
set -o pipefail
# set -o xtrace

TARGET=sealet
ENV=

help() {
    echo "Usage:"
    echo "getopts.sh [-t target -e env]"
    echo "Description:"
    echo "target,the name of server."
    echo "env,environment variable of server."
    exit -1
}

while getopts 't:e:' option; do
  case ${option} in
    t) TARGET=${OPTARG};;
    e) ENV=${OPTARG};;
    ?) help ;;
  esac
done


SEA_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
SEA_CMD_ROOT=${SEA_ROOT}/cmd
SEA_OUTPUT_SUB_PATH=${SEA_OUTPUT_SUB_PATH:-output/${TARGET}}
SEA_OUTPUT_PATH=${SEA_ROOT}/${SEA_OUTPUT_SUB_PATH}
SEA_OUTPUT_BIN_PATH=${SEA_OUTPUT_PATH}/bin


VERSION_PATH="${SEA_ROOT}/script/version.sh"
if [[ ! -f "${VERSION_PATH}" ]]; then curl -s -L -o ${VERSION_PATH} https://raw.githubusercontent.com/kaydxh/sea/master/script/version.sh; fi
source "${VERSION_PATH}"

export CGO_ENABLED=1
export GOFLAGS="-mod=readonly"

function platform() {
  echo "$(go env GOHOSTOS)/$(go env GOHOSTARCH)"
}

function make_build_args() {
  local goldflags

  goldflags="all=$(ldflags ${TARGET}) ${GOLDFLAGS:-}"
  local -a build_args
   build_args=(
   -ldflags="${goldflags}"
   ) 

   echo "${build_args[*]}"
}

function build() {
  GO_BUILD_ARGS="$(make_build_args)"

  #go build -mod=vendor -o ${OUT_PUT_PATH}/sealet ./cmd/sealet
  go mod tidy
  if [[ ! -z "${ENV}" ]]; then (set -x; export ${ENV}); fi
  (set -x; go build "${GO_BUILD_ARGS}" -o "${SEA_CMD_ROOT}/${TARGET}/${TARGET}" ${SEA_CMD_ROOT}/${TARGET}/*.go)
}

echo "==> Building in $(platform)"
build
