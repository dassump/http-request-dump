# http-request-dump

HTTP server that reply the complete dump of the request. Useful for debugging and troubleshooting.

## Getting started

1. Download a pre-compiled binary from the release [page](https://github.com/dassump/http-request-dump/releases).
2. Run `http-request-dump --help`

```
$ http-request-dump --help
NAME:
   http-request-dump - HTTP request dump server

USAGE:
   http-request-dump [global options] [arguments...]

VERSION:
   devel

DESCRIPTION:
   https://github.com/dassump/http-request-dump

AUTHOR:
   Daniel Dias de Assumpção <dassump@gmail.com>

GLOBAL OPTIONS:
   --listen value, -l value  server address and port (default: "0.0.0.0:8888") [$HTTPREQUESTDUMP_LISTEN]
   --body, -b                dump request body (default: true) [$HTTPREQUESTDUMP_BODY]
   --help, -h                show help
   --version, -v             print the version

COPYRIGHT:
   http://www.apache.org/licenses/LICENSE-2.0
```

## Usage

HTTP server listens at 0.0.0.0:8888 and request body dump is enabled by default.

To change the listening address use the `--listen` flag and for body dump use the `--body` flag.

```
$ http-request-dump
>>> 2023/08/18 09:59:14 HTTP server listening on 0.0.0.0:8888
```

```
$ http-request-dump --listen 0.0.0.0:8080 --body=false
>>> 2023/08/18 09:59:54 HTTP server listening on 0.0.0.0:8080
```

### Container

A precompiled version is available as a container on [dockerhub](https://hub.docker.com/r/dassump/http-request-dump).


```
$ docker run --rm -p 8888:8888 dassump/http-request-dump
2023/08/18 13:38:15 Listening on 0.0.0.0:8888
```

## Examples

```
$ curl -X GET "127.0.0.1:8888"
GET / HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
User-Agent: curl/7.88.1
```

```
$ curl -X GET "localhost:8888/page?parameter=value&also=another"
GET /page?parameter=value&also=another HTTP/1.1
Host: localhost:8888
Accept: */*
User-Agent: curl/7.88.1
```

```
$ curl -X POST 127.0.0.1:8888/form -d username=test -d password=test
POST /form HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 27
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.88.1

username=test&password=test
```

```
$ curl -X POST 127.0.0.1:8888/api -H "Content-Type: application/json" -d "{\"key\":\"value\"}"
POST /api HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 15
Content-Type: application/json
User-Agent: curl/7.88.1

{"key":"value"}
```

```
$ curl -X POST 127.0.0.1:8888/form -F "file=@empty.png" --output -
POST /form HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 328
Content-Type: multipart/form-data; boundary=------------------------ee98747b7b09e4ce
User-Agent: curl/7.88.1

--------------------------ee98747b7b09e4ce
Content-Disposition: form-data; name="file"; filename="empty.png"
Content-Type: image/png

?PNG

IHDR??&PLTE?????tRNS@??f9IDATx^??1? ????X??????YIEND?B`?
--------------------------ee98747b7b09e4ce--
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/http-request-dump.

## License
See [LISENSE](https://github.com/dassump/http-request-dump/blob/main/LICENSE) file.
