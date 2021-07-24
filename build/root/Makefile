DBG_MAKEFILE ?=
ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
else
    # If we're not debugging the Makefile, don't echo recipes.
    MAKEFLAGS += -s
endif

#objects := $(wildcard *.go)
objects := $(patsubst %.go,%.o,$(wildcard *.go))
   #$(warning ***** starting Makefile for goal(s) "$(objects)")

.PHONY: all
all:
	@echo "building all"
	bash script/build.sh

.PHONY: generate 
generate:
	@echo "make generate"
	bash api/scripts/go_proto_gen.sh api/openapi-spec/v1.0

.PHONY: clean
clean:
	@echo "clean"
	bash script/clean.sh
