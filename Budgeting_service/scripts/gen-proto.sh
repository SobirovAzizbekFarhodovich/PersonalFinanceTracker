#!/bin/sh

CURRENT_DIR=$(pwd)

rm -rf ${CURRENT_DIR}/genprotos
mkdir -p ${CURRENT_DIR}/genprotos


for x in $(find ${CURRENT_DIR}/protos* -type d); do
  protoc -I=${x} -I/usr/local/include \
    --go_out=${CURRENT_DIR}/genprotos --go_opt=paths=source_relative \
    --go-grpc_out=${CURRENT_DIR}/genprotos --go-grpc_opt=paths=source_relative ${x}/*.proto
done
