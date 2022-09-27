.PHONY: all
all:
	@echo "building all"
	@bash script/build.sh

.PHONY: generate 
generate:
	@echo "make generate"
	@bash  script/go_proto_gen.sh -I . --proto_file_path api/openapi-spec --with-go --with-doc
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
