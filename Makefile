HERE ?= $(shell pwd)
LOCALBIN ?= $(shell pwd)/bin
FLUX_INSTALL_ROOT ?= /usr
COMMONENVVAR=GO111MODULE="on" GOOS=$(shell uname -s | tr A-Z a-z)
CPP = $(shell which cpp)

.PHONY: all
all: examples

.PHONY: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

BUILDENVVAR=CGO_CFLAGS="-I$(FLUX_INSTALL_ROOT)/include" CGO_LDFLAGS="-L$(FLUX_INSTALL_ROOT)/lib -L$(FLUX_INSTALL_ROOT)/lib/flux -L$(FLUX_INSTALL_ROOT)/lib/flux/resource -lflux-core -lflux-idset -lflux-hostlist -lstdc++ -ljansson -lczmq -lzmq"

.PHONY: examples
examples: submit-example keygen-example

.PHONY: submit-example
submit-example: $(LOCALBIN)
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o $(LOCALBIN)/fluxgo-submit example/submit/submit.go

.PHONY: keygen-example
keygen-example: $(LOCALBIN)
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o $(LOCALBIN)/fluxgo-keygen example/keygen/keygen.go

.PHONY: test
test: examples
	bats -t test/bats/cli.bats
