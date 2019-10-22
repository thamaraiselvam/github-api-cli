#! /usr/bin/env bash

set -ex

mkdir -p ../internal/mock
mockgen -destination ../internal/mock/service/service.go -package mock -source ../cli/service/service.go