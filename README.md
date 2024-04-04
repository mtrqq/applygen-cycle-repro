# applyconfig-gen reproduction

Minimal reproduction in order to demonstrate import cycle issue

## Version information

applyconfig-gen binary is generated based on for 670586590c4574bdb99793d8373d8e4431ee30db of code-generator staging repo

## How to reproduce

Launch `run.sh` script which will remove previous generated configuration and create new on top

## Where to look for an issue

Look inside `applyconfig/resource.io/v1/priority.go` file, it contains cycle import to the same module it's located in (line 6)
