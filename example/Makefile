IMAGE := zengxu/kube-code-generator

CUR_DIR := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))
DIRECTORY := $(realpath $(CUR_DIR)/..)

PROJECT_PACKAGE := github.com/phosae/kube-code-generator/example
DEPS_CMD := go mod tidy

default: generate

.PHONY: generate
generate: generate-client generate-crd generate-openapi generate-proto

generate-helpers:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/github.com/phosae/kube-code-generator \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e GENERATION_TARGETS="helpers" \
	$(IMAGE)


# client/clientset/example/... seems bits of werid, but kube_codegen.sh can only do this now
# writing custom scripts using client-gen, lister-gen, informer-gen, applyconfiguration-gen directly  
#  can output client like client-go/kubernetes (without clientset in middle)
generate-client: generate-helpers
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/github.com/phosae/kube-code-generator \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e CLIENT_GENERATOR_OUT=$(PROJECT_PACKAGE)/client \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e GENERATION_TARGETS="client" \
	-e CLIENTSET_NAME="clientset" \
	-e VERSIONED_NAME="example" \
	-e WITH_APPLYCONFIG="true" \
	-e APPLYCONFIG_NAME="applyconfigurations" \
	-e WITH_WATCH="" \
	$(IMAGE)

generate-openapi:
	docker run -it --rm \
	-v $(DIRECTORY):/go/src/github.com/phosae/kube-code-generator \
	-v $(DIRECTORY)/example/openapi-validation-exceptions.list:/tmp/validation-exceptions.list \
	-e PROJECT_PACKAGE=$(PROJECT_PACKAGE) \
	-e APIS_ROOT=$(PROJECT_PACKAGE)/apis \
	-e OPENAPI_VIOLATION_EXCEPTIONS="/tmp/validation-exceptions.list" \
	-e OPENAPI_GENERATOR_OUT=$(PROJECT_PACKAGE)/generated \
	-e GENERATION_TARGETS="openapi" \
	$(IMAGE)

generate-proto:
	docker run -it --rm \
    -v $(DIRECTORY):/go/src/github.com/phosae/kube-code-generator \
    -w /go/src/${PROJECT_PACKAGE} \
    $(IMAGE) \
    update-generated-protobuf.sh $(PROJECT_PACKAGE)/apis/comic/v1

generate-crd:
	docker run -it --rm \
	-v $(DIRECTORY)/example:/src \
	-e GO_PROJECT_ROOT=/src \
	-e CRD_TYPES_PATH=/src/apis \
	-e CRD_OUT_PATH=/src/manifests \
	$(IMAGE) update-crd.sh

.PHONY: deps
deps:
	$(DEPS_CMD)

.PHONY: clean
clean:
	echo "Cleaning generated files..."
	rm -rf ./manifests
	rm -rf ./client
	rm -rf ./apis/comic/v1/zz_generated.deepcopy.go
	rm -rf ./apis/comic/v1/generated.pb.go
	rm -rf ./apis/comic/v1/generated.proto
	rm -rf ./generated