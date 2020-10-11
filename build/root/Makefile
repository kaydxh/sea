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
    $(warning ***** starting Makefile for goal(s) "$(objects)")

.PHONY: all
ifeq ($(PRINT_HELP),y)
all:
	@echo "$$ALL_HELP_INFO"
endif
