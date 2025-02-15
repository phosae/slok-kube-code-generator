#!/usr/bin/env bash

set -eufo pipefail

## Project specific data
# PROJECT_PACKAGE=github.com/phosae/kube-code-generator/example
# APIS_ROOT=${PROJECT_PACKAGE}/apis
#
## GENERATION_TARGETS="helpers,client,openapi"
## helpers ENVs
#    none
#
## clent ENVs
#   CLIENT_GENERATOR_OUT           
#   CLIENTSET_NAME                 An optional override for the leaf name of the generated "clientset" directory. 
#   VERSIONED_NAME                 An optional override for the leaf name of the generated "<clientset>/versioned" directory.
#
#   WITH_APPLYCONFIG               Enables generation of applyconfiguration files
#     APPLYCONFIG_NAME             An optional override for the leaf name of the generated "applyconfiguration" directory
# 
#   WITH_WATCH                     Enables generation of listers and informers for APIs which support WATCH
#     LISTERS_NAME <listers>       An optional override for the leaf name of the generated "listers" directory.
#     INFORMERS_NAME <informers>   An optional override for the leaf name of the generated "informers" directory
#
## openapi ENVs
#    OPENAPI_GENERATOR_OUT         The root package under which generated directories and files
#    OPENAPI_VIOLATION_EXCEPTIONS  OpenAPI violation_exceptions.list, like https://github.com/kubernetes/kubernetes/tree/master/api/api-rules, optional if project APIs are non-violation    

PROJECT_PACKAGE="${PROJECT_PACKAGE:-""}"
CLIENT_GENERATOR_OUT="${CLIENT_GENERATOR_OUT:-""}"
APIS_ROOT="${APIS_ROOT:-""}"

[ -z "$PROJECT_PACKAGE" ] && echo "PROJECT_PACKAGE env var is required" && exit 1
[ -z "$APIS_ROOT" ] && echo "APIS_ROOT env var is required" && exit 1

GENERATION_TARGETS="${GENERATION_TARGETS:-helpers,client}"

# From >=1.16 we use gomod, so we need to execute from the project directory.
cd "${GOPATH}/src/${PROJECT_PACKAGE}"

# Ugly but needs to be relative if we want to use k8s.io/code-generator
# as it is without touching/sed-ing the code/scripts
RELATIVE_ROOT_PATH=$(realpath --relative-to="${PWD}" /)
CODEGEN_PKG=${RELATIVE_ROOT_PATH}${GOPATH}/src/k8s.io/code-generator

BOILERPLATE_PATH=/tmp/fake-boilerplate.txt
[[ -f ${BOILERPLATE_PATH} ]] || touch ${BOILERPLATE_PATH}

source "${CODEGEN_PKG}/kube_codegen.sh"

if grep -qw "helpers" <<<"${GENERATION_TARGETS}"; then
  kube::codegen::gen_helpers \
    --input-pkg-root ${APIS_ROOT} \
    --output-base "${GOPATH:-"/go"}/src" \
    --boilerplate "${BOILERPLATE_PATH}"
fi

CLIENTSET_NAME="${CLIENTSET_NAME:-"clientset"}"
VERSIONED_NAME="${VERSIONED_NAME:-"versioned"}"

# used by client with-applyconfig
APPLYCONFIG_NAME="${APPLYCONFIG_NAME:-"applyconfiguration"}"
# used by client with watch
LISTERS_NAME="${LISTER_NAME:-"listers"}"
INFORMERS_NAME="${INFORMERS_NAME:-"informers"}"

if grep -qw "client" <<<"${GENERATION_TARGETS}"; then
  [ -z "$CLIENT_GENERATOR_OUT" ] && echo "CLIENT_GENERATOR_OUT env var is required" && exit 1

  applyconfig_opts=""
  watch_opts=""
  [[ -n "${WITH_APPLYCONFIG+set}" ]] && applyconfig_opts="--with-applyconfig --applyconfig-name ${APPLYCONFIG_NAME}"
  [[ -n "${WITH_WATCH+set}" ]] && watch_opts="--with-watch --listers-name ${LISTERS_NAME} --informers-name ${INFORMERS_NAME}"
  
  kube::codegen::gen_client \
    --clientset-name ${CLIENTSET_NAME} \
    --versioned-name ${VERSIONED_NAME} \
    $applyconfig_opts $watch_opts \
    --input-pkg-root ${APIS_ROOT} \
    --output-pkg-root ${CLIENT_GENERATOR_OUT} \
    --output-base "${GOPATH:-"/go"}/src" \
    --boilerplate "${BOILERPLATE_PATH}"
fi

if grep -qw "openapi" <<<"${GENERATION_TARGETS}"; then  
  [ -z "$OPENAPI_GENERATOR_OUT" ] && echo "OPENAPI_GENERATOR_OUT env var is required" && exit 1

  kube::codegen::gen_openapi \
    --input-pkg-root ${APIS_ROOT} \
    --output-pkg-root ${OPENAPI_GENERATOR_OUT} \
    --report-filename ${OPENAPI_VIOLATION_EXCEPTIONS:-"/go/k8s-validation_exceptions.list"} \
    --output-base "${GOPATH:-"/go"}/src" \
    --boilerplate "${BOILERPLATE_PATH}"
fi