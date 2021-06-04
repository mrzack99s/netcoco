# Docker Run

## Requirements

* Docker version &gt;= 19.03.xx
---
## Environment Vars
*   `COCO_HOSTNAME` is base URL of netcoco
*   `COCO_PORT` is serve port of netcoco
*   `COCO_DB_HOSTNAME` is URL of databese (SQL)
*   `COCO_DB_USERNAME` is username of databese
*   `COCO_DB_PASSWORD` is password of databese
*   `COCO_DB_NAME` is databese name
*   `COCO_SEC_SALT` is an key of encryption (hex in 32 digit)
---
## Run

* With **config.yaml** file
    ```console
    # docker run -p 8080:8080 -d \ 
    --sysctl net.ipv4.ping_group_range="0 2147483647" \
    -v $(pwd)/config.yaml:/netcoco/config.yaml \
    quay.io/netcoco-io/netcoco
    ```

* With **environments** file
    ```console
    # docker run -p 8080:8080 -d \
    --sysctl net.ipv4.ping_group_range="0 2147483647" \
    -e <env_name>=<value> \
    quay.io/netcoco-io/netcoco
    ```
