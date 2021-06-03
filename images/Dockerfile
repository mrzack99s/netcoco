# Build stage
FROM node:14-alpine3.13 as build

# Install
RUN apk add git yarn
RUN git clone https://github.com/mrzack99s/netcoco-ui
WORKDIR /netcoco-ui
RUN yarn install 
RUN sed -i "s,<Domain Name or Hostname>,localhost,g" .env.production
RUN yarn build --mode production

# Production stage
FROM ubuntu:20.04 as production

ARG version
ENV NETCOCO_VERSION=$version

RUN apt-get update && \
    apt-get install -y curl
RUN mkdir -p /netcoco
RUN curl -LO https://github.com/mrzack99s/netcoco/releases/download/v${NETCOCO_VERSION}/templates.tar.gz && \
    tar -zxf templates.tar.gz && mv templates /netcoco 
RUN sed -i "s,<Hostname or Domain name>:<port>,localhost:8080,g" /netcoco/templates/js/main.js 
RUN mkdir -p /netcoco/templates/dist
COPY --from=build /netcoco-ui/dist /netcoco/templates/dist
RUN curl -L https://github.com/mrzack99s/netcoco/releases/download/v${NETCOCO_VERSION}/netcoco-linux-amd64 -o /netcoco/netcoco-linux-amd64

## Execution
WORKDIR /netcoco
COPY entrypoint.sh .
RUN chmod +x netcoco-linux-amd64
RUN chmod +x entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]

