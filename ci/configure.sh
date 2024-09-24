#!/usr/bin/env bash

fly -t "${CONCOURSE_TARGET:-bosh}" set-pipeline \
    -p bosh-cpi-go \
    -c ci/pipeline.yml
