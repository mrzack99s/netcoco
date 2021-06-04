# Docker Compose

## Environment Vars
*   `COCO_HOSTNAME` is base URL of netcoco
*   `COCO_PORT` is serve port of netcoco
*   `COCO_DB_HOSTNAME` is URL of databese (SQL)
*   `COCO_DB_USERNAME` is username of databese
*   `COCO_DB_PASSWORD` is password of databese
*   `COCO_DB_NAME` is databese name
*   `COCO_SEC_SALT` is an key of encryption (hex in 32 digit)
---
> Docker compose file **docker-compose.yaml** with config file

```yaml
version: '3'

services:
  netcoco:
    image: quay.io/netcoco-io/netcoco
    container_name: netcoco
    restart: always
    sysctl:
        net.ipv4.ping_group_range: "0 2147483647"
    environment:
      COCO_HOSTNAME: Hostname or domain name
    ports:
      - "8080:8080"
    volumes:
      - ./config.yaml:/netcoco/config.yaml
```

> Docker compose file **docker-compose.yaml** with environment

```yaml
version: '3'

services:
  netcoco:
    image: quay.io/netcoco-io/netcoco
    container_name: netcoco
    restart: always
    sysctl:
        net.ipv4.ping_group_range: "0 2147483647"
    environment:
      COCO_HOSTNAME: Hostname or domain name
    ports:
      - "8080:8080"
```
---
## Requirements

* Docker verison &gt;= 19.03.xx
* Docker-compose
---
## How to use

1. Create file docker-compose.yaml
2. Run docker compose

   ```bash
       docker-compose up -d
   ```

