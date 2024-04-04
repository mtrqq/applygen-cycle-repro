#/usr/bin/env bash

rm -rf applyconfig
./applyconfiguration-gen ./resource.io/v1 \
    --output-dir ./applyconfig \
    --output-pkg github.io/applygen-cycle-repro/applyconfig \
    -v 5