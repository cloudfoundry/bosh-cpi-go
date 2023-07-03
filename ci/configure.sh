#!/usr/bin/env bash

fly -t bosh-ecosystem set-pipeline \
    -p bosh-cpi-go \
    -c ci/pipeline.yml
