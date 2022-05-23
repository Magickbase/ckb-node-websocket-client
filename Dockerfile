##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD . .

RUN go build -o /ckb-node-websocket-client

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

ADD configs configs

COPY --from=build /ckb-node-websocket-client /ckb-node-websocket-client

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ckb-node-websocket-client"]
