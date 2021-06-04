# Docker Compose

> Docker compose file **docker-compose.yaml** 

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

### Requirements
*   Docker verison >= 19.03.xx
*   Docker-compose

### How to use

1.  Create file docker-compose.yaml
2.  Run docker compose
    ```bash
        docker-compose up -d
    ```