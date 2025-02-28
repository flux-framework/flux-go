HERE ?= $(shell pwd)
LOCALBIN ?= $(shell pwd)/bin
FLUX_INSTALL_ROOT ?= /usr
COMMONENVVAR=GO111MODULE="on" GOOS=$(shell uname -s | tr A-Z a-z)
CPP = $(shell which cpp)

REGISTRY=ghcr.io/flux-framework
IMAGE=flux-go:latest
RELEASE_VERSION?=v$(shell date +%Y%m%d)-$(shell git describe --tags --match "v*")

.PHONY: all
all: build

.PHONY: $(LOCALBIN)
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

BUILDENVVAR=CGO_CFLAGS="-I$(FLUX_INSTALL_ROOT)/include" CGO_LDFLAGS="-L$(FLUX_INSTALL_ROOT)/lib -L$(FLUX_INSTALL_ROOT)/lib/flux -L$(FLUX_INSTALL_ROOT)/lib/flux/resource -lflux-core -lflux-idset -lflux-hostlist -lstdc++ -ljansson -lczmq -lzmq"

# This should be run outside of the devcontainer environment
.PHONY: build
build:
	docker build --build-arg ARCH="amd64" --build-arg RELEASE_VERSION="$(RELEASE_VERSION)" -t $(REGISTRY)/$(IMAGE) .

# Build the API server
.PHONY: server
server:
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o bin/flux-grpc-server cmd/api/api.go

.PHONY: examples
examples: submit-example keygen-example list-jobs-example

.PHONY: test-examples
test-examples: examples
	$(LOCALBIN)/flux-list-jobs
	$(LOCALBIN)/flux-keygen
	$(LOCALBIN)/flux-submit

.PHONY: submit-example
submit-example: $(LOCALBIN)
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o $(LOCALBIN)/flux-submit example/submit/submit.go

.PHONY: list-jobs-example
list-jobs-example: $(LOCALBIN)
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o $(LOCALBIN)/flux-list-jobs example/list-jobs/list-jobs.go

.PHONY: keygen-example
keygen-example: $(LOCALBIN)
	$(COMMONENVVAR) $(BUILDENVVAR) go build -ldflags '-w' -o $(LOCALBIN)/flux-keygen example/keygen/keygen.go


.PHONY: test
test: test-examples
#	$(COMMONENVVAR) $(BUILDENVVAR) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) go test -count 1 -run TestCancel -ldflags '-w' ./pkg/fluxcli ./pkg/types
	$(COMMONENVVAR) $(BUILDENVVAR) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) go test -ldflags '-w' ./pkg/flux

.PHONY: test-v
test-v: test-examples
#	$(COMMONENVVAR) $(BUILDENVVAR) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) go test -count 1 -run TestFind -v -ldflags '-w' ./pkg/fluxcli ./pkg/types
	$(COMMONENVVAR) $(BUILDENVVAR) LD_LIBRARY_PATH=$(LD_LIBRARY_PATH) go test -v -ldflags '-w' ./pkg/flux

.PHONY: protoc
protoc: $(LOCALBIN)
	GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# You can use make protoc to download proto
.PHONY: proto
proto: protoc
	PATH=$(LOCALBIN):${PATH} protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pkg/flux-grpc/flux.proto

.PHONY: python
python: python ## Generate python proto files in python
	# pip install grpcio-tools
	# pip freeze | grep grpcio-tools
	mkdir -p python/v1/fluxgrpc/protos
	cd python/v1/fluxgrpc/protos
	python3 -m grpc_tools.protoc -I./pkg/flux-grpc --python_out=./python/v1/fluxgrpc/protos --pyi_out=./python/v1/fluxgrpc/protos --grpc_python_out=./python/v1/fluxgrpc/protos ./pkg/flux-grpc/flux.proto
	sed -i 's/import flux_pb2 as flux__pb2/from . import flux_pb2 as flux__pb2/' ./python/v1/fluxgrpc/protos/flux_pb2_grpc.py
