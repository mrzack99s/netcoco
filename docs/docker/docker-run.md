# Docker Run

## Requirements
* Docker version >= 19.03.xx

## Run
```console
# docker run -p 8080:8080 --sysctl net.ipv4.ping_group_range="0 2147483647" -v $(pwd)/config.yaml:/netcoco/config.yaml -d quay.io/netcoco-io/netcoco
```
