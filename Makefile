# example
# make TARGET=sealet
#
# set default target to sealet
TARGET=sealet

.PHONY: all
all:
	@echo "building all ${TARGET}"
	@bash script/build.sh -t ${TARGET}

.PHONY: generate 
generate:
	@echo "make generate"
	@bash -c "curl -s -L -o ./script/proto-gen.sh https://raw.githubusercontent.com/kaydxh/golang/main/script/go_proto_gen.sh"
	@bash  script/proto-gen.sh -I . --proto_file_path api/protoapi-spec --third_party_path ./third_party --with-go --with-doc
.PHONY: copyright 
copyright:
	@echo "make copyright"
	@bash -c "curl -s -L -o ./script/copyright.sh https://raw.githubusercontent.com/kaydxh/golang/main/script/copyright.sh"
	@bash -c "curl -s -L -o ./script/copyright.txt https://raw.githubusercontent.com/kaydxh/golang/main/script/copyright.txt"
	@bash script/copyright.sh

.PHONY: clean
clean:
	@echo "clean"
	@bash script/clean.sh
