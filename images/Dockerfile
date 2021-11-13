# Build stage
FROM node:14-alpine3.13 as build_node

# Install
RUN apk add git yarn
RUN git clone https://github.com/mrzack99s/netcoco-ui
WORKDIR /netcoco-ui
RUN git checkout v0.2.0-rc1
RUN yarn install 
RUN sed -i "s,<Domain Name or Hostname>,localhost,g" .env.production
RUN yarn build --mode production

FROM golang:1.16.4-alpine3.13 as build_go

ARG COCO_VERSION
# Install
RUN apk add git make
RUN git clone https://github.com/mrzack99s/netcoco /netcoco

WORKDIR /netcoco
RUN ls
RUN git checkout ${COCO_VERSION}
RUN make build

# Production stage
FROM alpine:3 as production

RUN apk add --no-cache bash

RUN mkdir -p /netcoco/templates
RUN mkdir -p /netcoco/templates/dist

COPY --from=build_go /netcoco/templates /netcoco/templates
COPY --from=build_node /netcoco-ui/dist /netcoco/templates/dist
COPY --from=build_go /netcoco/netcoco-linux-amd64 /netcoco

## Execution
WORKDIR /netcoco
COPY entrypoint.sh .
RUN chmod +x netcoco-linux-amd64
RUN chmod +x entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]
