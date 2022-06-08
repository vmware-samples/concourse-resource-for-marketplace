# Copyright 2022 VMware, Inc.
# SPDX-License-Identifier: BSD-2-Clause

ARG MKPCLI_VERSION=latest
FROM harbor-repo.vmware.com/dockerhub-proxy-cache/library/golang:1.18 as builder
ARG VERSION

COPY . /concourse-resource-for-marketplace/
ENV PATH="${PATH}:/root/go/bin"
WORKDIR /concourse-resource-for-marketplace/
RUN make build

FROM projects.registry.vmware.com/tanzu_isv_engineering/mkpcli:$MKPCLI_VERSION
LABEL description="Concourse Resource for VMware Marketplace"
LABEL maintainer="tanzu-isv-engineering@groups.vmware.com"

COPY --from=builder /concourse-resource-for-marketplace/build/check /opt/resource/check
COPY --from=builder /concourse-resource-for-marketplace/build/in    /opt/resource/in
COPY --from=builder /concourse-resource-for-marketplace/build/out   /opt/resource/out
WORKDIR /opt/resource
