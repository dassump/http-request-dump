# syntax=docker/dockerfile:1
FROM --platform=linux/amd64 golang:1.21-alpine as build
RUN apk --no-cache upgrade --available
RUN apk --no-cache add alpine-sdk
WORKDIR /go/src/http-request-dump
COPY go.mod go.sum main.go ./
ARG VERSION=unknown
RUN go mod download -x
RUN CGO_ENABLED=0 go build -tags urfave_cli_no_docs -ldflags "-s -w -X main.version=${VERSION}" -o http-request-dump-linux-amd64 .

FROM --platform=linux/amd64 scratch as runtime
COPY --from=build /go/src/http-request-dump/http-request-dump-linux-amd64 /
CMD ["/http-request-dump-linux-amd64"]
EXPOSE 8888