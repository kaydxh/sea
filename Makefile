.PHONY: all
all:
	@echo "building all"
	@bash script/build.sh

.PHONY: generate 
generate:
	@echo "make generate"
	@bash  script/go_proto_gen.sh -I . --proto_file_path api/openapi-spec --with-go --with-doc
.PHONY: clean
clean:
	@echo "clean"
	@bash script/clean.sh
