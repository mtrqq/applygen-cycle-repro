#/usr/bin/env bash

rm -rf applyconfig
./applyconfiguration-gen ./resource.io/v1 \
    --output-dir ./applyconfig \
    --output-pkg github.com/applygen-cycle-repro/applyconfig \
    -v 5