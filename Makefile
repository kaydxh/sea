# example
# make TARGET=seadate
#
# set default target to seadate
TARGET=seadate
DEPS_NAME=deps.yaml
DO_PARSE_DEPS="ON"

MAKEFILE_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
PROJECT_ROOT_DIR := $(realpath ${MAKEFILE_DIR})
PROJECT_CMD_ROOT_DIR := ${PROJECT_ROOT_DIR}/cmd/${TARGET}
PKG_CONFIG_PATH := ${PROJECT_CMD_ROOT_DIR}/pkgconfig
SCRIPT_PATH := ${MAKEFILE_DIR}/../script
BUILD_ENV_VAR :=
OUTPUT_DIR := ${PROJECT_ROOT_DIR}/pack/${TARGET}
OUTPUT_CONF_DIR=${OUTPUT_DIR}/conf
OUTPUT_LIB_DIR=${OUTPUT_DIR}/lib
OUTPUT_BIN_DIR=${OUTPUT_DIR}/bin
OUTPUT_MODEL_DIR=${OUTPUT_DIR}/model
OUTPUT_BIN_PATH=${OUTPUT_DIR}/bin/${TARGET}

.PHONY: all
all:
	@echo "building all ${TARGET}"
	@bash script/build.sh -t "${TARGET}" -e "${BUILD_ENV_VAR}"

.PHONY: generate 
generate:
	@echo "make generate"
	@bash -c "curl -s -L -o ./script/proto-gen.sh https://raw.githubusercontent.com/kaydxh/golang/main/script/go_proto_gen.sh"
	@bash  script/proto-gen.sh -I . --proto_file_path api/protoapi-spec --third_party_path ./third_party --with-go --with-doc

.PHONY: deps
deps:
	@echo "> make deps..."

.PHONY: install
install: all clean
	@echo "> make install..."
	@echo " > install conf..."
	@mkdir -p ${OUTPUT_CONF_DIR}
	@cp -Rv "${PROJECT_ROOT_DIR}/conf/${TARGET}"* ${OUTPUT_CONF_DIR}

	@echo " > install bin..."
	@mkdir -p ${OUTPUT_BIN_DIR}
	@cp -Rv "${PROJECT_CMD_ROOT_DIR}/${TARGET}" ${OUTPUT_BIN_DIR}

	@echo " > install libs..."
	@mkdir -p ${OUTPUT_LIB_DIR}
	@$(eval LINK_THIRD_LIB_PATHS := $(shell find -L ${PROJECT_CMD_ROOT_DIR}/third_path/ -maxdepth 3 -mindepth 2 -type d -iname "lib*" -print0 |xargs -0 -I {} sh -c 'echo {}'|grep -v "stubs"))
	@echo "${LINK_THIRD_LIB_PATHS}"
	@$(eval JOINED_LINK_THIRD_LIB_PATHS := $(call joinwith,:,$(LINK_THIRD_LIB_PATHS)))
	@LD_LIBRARY_PATH="$(JOINED_LINK_THIRD_LIB_PATHS):${LD_LIBRARY_PATH}" ldd "${OUTPUT_BIN_PATH}" | awk '{if (match($$3,"/")){ print $$3  }}' | grep "third_path" | grep -v "^/lib64" | grep -v "^/lib" | xargs -I {} sh -c 'cp -v -L {} ${OUTPUT_LIB_DIR}'
	@LD_LIBRARY_PATH="$(JOINED_LINK_THIRD_LIB_PATHS):${LD_LIBRARY_PATH}" ldd "${OUTPUT_BIN_PATH}" | awk '{if (match($$3,"/")){ print $$3  }}' | grep "libgomp" | xargs -I {} sh -c 'cp -v -L {} ${OUTPUT_LIB_DIR}'
	@LD_LIBRARY_PATH="$(JOINED_LINK_THIRD_LIB_PATHS):${LD_LIBRARY_PATH}" ldd "${OUTPUT_BIN_PATH}" | awk '{if (match($$3,"/")){ print $$3  }}' | grep "libjbig" | xargs -I {} sh -c 'cp -v -L {} ${OUTPUT_LIB_DIR}'
	@LD_LIBRARY_PATH="$(JOINED_LINK_THIRD_LIB_PATHS):${LD_LIBRARY_PATH}" ldd "${OUTPUT_BIN_PATH}" | awk '{if (match($$3,"/")){ print $$3  }}' | grep "libstdc++.so" | xargs -I {} sh -c 'cp -v -L {} ${OUTPUT_LIB_DIR}'

	@echo " > install model..."
	@cd ${PROJECT_CMD_ROOT_DIR}; find -L ./third_path/ -maxdepth 3 -type d \( -iname "model" -o -iname "sdk_data" -o -iname "config" \) -print0 | xargs -0 -I {} sh -c 'mkdir -p ${OUTPUT_MODEL_DIR}; cp -r -v -d {}/* ${OUTPUT_MODEL_DIR}'
	@echo " > install script..."
	@cp -Rv "${PROJECT_ROOT_DIR}/script/${TARGET}/"* ${OUTPUT_DIR}

	@echo -e "\n==> Results:"
	@ls -Rhl ${OUTPUT_DIR}

.PHONY: pkg-config
pkg-config: deps
	@echo "> make pkg-config..."
	@$(eval THIRD_LIB_PATHS := $(shell find -L ${PROJECT_CMD_ROOT_DIR}/third_path/ -maxdepth 3 -mindepth 2 -type d -iname "lib*" -print0 |xargs -0 -I {} bash -c 'echo {}'|grep -v "stubs"))
	@$(eval JOINED_THIRD_LIB_PATHS := $(call joinwith,:,${THIRD_LIB_PATHS}))
	@$(eval BUILD_ENV_VAR = PKG_CONFIG_PATH=${PKG_CONFIG_PATH} LD_LIBRARY_PATH=${JOINED_THIRD_LIB_PATHS}:${LD_LIBRARY_PATH} LIBRARY_PATH=${JOINED_THIRD_LIB_PATHS}:${LIBRARY_PATH})
	@echo "copy pkg-config..."
	@mkdir -p ${PKG_CONFIG_PATH}
	@if [[ (( -d "${PROJECT_CMD_ROOT_DIR}"/third_path/yt-sdk-go/pkgconfig )) ]]; then cp -df "${PROJECT_CMD_ROOT_DIR}"/third_path/yt-sdk-go/pkgconfig/* "${PKG_CONFIG_PATH}/"; fi
	@${BUILD_ENV_VAR} go mod tidy

.PHONY: copyright 
copyright:
	@echo "make copyright"
	@bash -c "curl -s -L -o ./script/copyright.sh https://raw.githubusercontent.com/kaydxh/golang/main/script/copyright.sh"
	@bash -c "curl -s -L -o ./script/copyright.txt https://raw.githubusercontent.com/kaydxh/golang/main/script/copyright.txt"
	@bash script/copyright.sh

.PHONY: clean
clean:
	@echo "clean"
	@bash script/clean.sh ${OUTPUT_DIR}
