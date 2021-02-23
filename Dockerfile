FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY ./src/main.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o /go/bin/server

FROM scratch
EXPOSE 80
ENV GIN_MODE=release 
COPY ./src/database/IP2LOCATION-LITE-DB3.IPV6.BIN /go/bin/IP2LOCATION-LITE-DB3.IPV6.BIN
COPY ./src/templates /go/templates
COPY --from=builder /go/bin/server /go/bin/server
ENTRYPOINT ["/go/bin/server"]
