#!/usr/bin/env bash

make OS_NAME=ubuntu OS_VERSION=20.04_focal TEMPLATE=base BUILDER=qemu BUILD_OPTS='-var-file=vars_ubuntu-2004-base.json'