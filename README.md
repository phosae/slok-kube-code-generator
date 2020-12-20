# Kube code generator [![Docker Repository on Quay](https://quay.io/repository/slok/kube-code-generator/status "Docker Repository on Quay")](https://quay.io/repository/slok/kube-code-generator)

A kubernetes code generator ready container to create your CRD autogenerated code like clients or openapi specs.

This uses the [official util](https://github.com/kubernetes/code-generator) created by Kubernetes to autogenerate the code required by Kubernetes resources.

## What can generate

- CRD based Go codes like Kubernetes clients.
- CRD manifest YAMLs to register your CRs on the cluster.
- OpenAPI spec.

## Docker image versions

|                  | Docker image                                            |
| ---------------- | ------------------------------------------------------- |
| Kubernetes v1.20 | `docker pull quay.io/slok/kube-code-generator:v1.20.1`  |
| Kubernetes v1.19 | `docker pull quay.io/slok/kube-code-generator:v1.19.2`  |
| Kubernetes v1.18 | `docker pull quay.io/slok/kube-code-generator:v1.18.0`  |
| Kubernetes v1.17 | `docker pull quay.io/slok/kube-code-generator:v1.17.3`  |
| Kubernetes v1.16 | `docker pull quay.io/slok/kube-code-generator:v1.16.7`  |
| Kubernetes v1.15 | `docker pull quay.io/slok/kube-code-generator:v1.15.10` |
| Kubernetes v1.14 | `docker pull quay.io/slok/kube-code-generator:v1.14.2`  |
| Kubernetes v1.13 | `docker pull quay.io/slok/kube-code-generator:v1.13.5`  |
| Kubernetes v1.12 | `docker pull quay.io/slok/kube-code-generator:v1.12.4`  |
| Kubernetes v1.11 | `docker pull quay.io/slok/kube-code-generator:v1.11.3`  |
| Kubernetes v1.10 | `docker pull quay.io/slok/kube-code-generator:v1.10.0`  |
| Kubernetes v1.9  | `docker pull quay.io/slok/kube-code-generator:v1.9.1`   |

You can use `docker pull quay.io/slok/kube-code-generator:latest` for `master` branch.

## Getting started

The best way to start is by checking an [example](example/) of how
to generate Go code for a CRD.

For a detailed explanation of how to use it continue reading.

## Kubernetes type code generation

### Env vars required

- `PROJECT_PACKAGE`: The project package path.
- `CLIENT_GENERATOR_OUT`: The client generated out path.
- `APIS_ROOT`: The path where our API/resources are.
- `GROUPS_VERSION`: The groups of the resources.
- `GENERATION_TARGETS`: The wanted generated code. [(deepcopy,defaulter,client,lister,informer) or "all"].

### Process

Having a project that would be on `github.com/someone/myproject`, wants to generate the client on `github.com/someone/myproject/client`, has its resources on `github.com/someone/myproject/apis`, has the resources `github.com/someone/myproject/apis/test/v1alpha1` and `github.com/someone/myproject/apis/test2/v1` and wants only `deepcopy` and `client` autogenerated code this would be:

```bash
PROJECT_PACKAGE=github.com/someone/myproject && \
docker run -it --rm \
    -v ${PWD}:/go/src/${PROJECT_PACKAGE}\
    -e PROJECT_PACKAGE=${PROJECT_PACKAGE} \
    -e CLIENT_GENERATOR_OUT=${PROJECT_PACKAGE}/client \
    -e APIS_ROOT=${PROJECT_PACKAGE}/apis \
    -e GROUPS_VERSION="test:v1alpha1 test2:v1" \
    -e GENERATION_TARGETS="deepcopy,client" \
    quay.io/slok/kube-code-generator
```

## Kubernetes CRD manifests

### Env vars required

- `GO_PROJECT_ROOT`: The project root (does not depend on GOPATH).
- `CRD_TYPES_PATH`: The path to the CRDs (searches in the path).
- `CRD_OUT_PATH`: The output path for the CRD manifest.

### Process

```bash
docker run -it --rm \
    -v ${PWD}:/src \
    -e GO_PROJECT_ROOT=/src \
    -e CRD_TYPES_PATH=/src/apis \
    -e CRD_OUT_PATH=/src/manifests \
    quay.io/slok/kube-code-generator update-crd.sh
```

## Kubernetes type openapi spec generation

### Env vars required

- `CRD_PACKAGES`: The packages where all the CRD types are (comma separated).
- `OPENAPI_OUTPUT_PACKAGE`: The output package where the open api spec will be created.

### Process

Having a project that would be on `github.com/someone/myproject`, wants to generate the openapi spec for its types on `github.com/someone/myproject/openapi`, has its resources on `github.com/someone/myproject/apis`, has the resources `github.com/someone/myproject/apis/test/v1alpha1` and `github.com/someone/myproject/apis/test2/v1` autogenerated spec this would be:

```bash
PROJECT_PACKAGE=github.com/someone/myproject && \
docker run -it --rm \
    -v ${PWD}:/go/src/${PROJECT_PACKAGE}\
    -e CRD_PACKAGES=${PROJECT_PACKAGE}/apis/test/v1alpha1,${PROJECT_PACKAGE}/apis/test2/v1 \
    -e OPENAPI_OUTPUT_PACKAGE=${PROJECT_PACKAGE}/openapi \
    quay.io/slok/kube-code-generator update-openapi.sh
```
