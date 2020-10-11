# Created by kayxhding on 2020-10-11 12:40:37
#!/usr/bin/env bash

# Fail on any error.
set -eo pipefail

SCRIPTPATH=$(cd `dirname $0`;pwd)

#:<<EOF
#SCRIPTPATH=$(dirname "$SCRIPT")
#SCRIPT=$(readlink -f "$0")
#SCRIPTPATH=$(dirname "$SCRIPT")
#echo ${SCRIPTPATH}
#EOF

PROTOC_FILE_DIR="$1"

function die() {
  echo 1>&2 "$*"
  exit 1
}

<<'COMMENT'
# This will place three binaries in your $GOBIN
# Make sure that your $GOBIN is in your $PATH
 go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go
COMMENT

GEN_PROTO_TOOLS=(protoc protoc-gen-go protoc-gen-grpc-gateway protoc-gen-govalidators)
for tool in ${GEN_PROTO_TOOLS[@]}; do
   q=$(command -v ${tool}) || die "didn't find ${tool}"
   echo 1>&2 "${tool}: ${q}"
done

grpc_option="--grpc-gateway_out=logtostderr=true"

for f in $(find ${PROTOC_FILE_DIR} -type f -name '*.proto' -print0 | xargs -0); do
  protoc -I . --govalidators_out=paths=source_relative:. --go_out=plugins=grpc,paths=source_relative:. ${grpc_option} ${f}
  echo $f
done
