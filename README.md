# http-request-dump

Server that listens for any http request and logs the full dump of the request.

## Getting started

1. Download a pre-compiled binary from the release [page](https://github.com/dassump/http-request-dump/releases).
2. Run `http-request-dump --help`

```shell
$ http-request-dump --help
http-request-dump (dev)

HTTP request dump server
https://github.com/dassump/http-request-dump

Usage of http-request-dump:
  -listen string
        Server address and port (default "0.0.0.0:8888")
```

## Usage

Listens at 0.0.0.0:8888 by default.

```shell
$ http-request-dump
>>> 2022/05/17 14:27:57 Listening on 0.0.0.0:8888
```

```shell
$ http-request-dump -listen 127.0.0.1:8080
>>> 2022/05/17 14:28:22 Listening on 127.0.0.1:8080
```

## Examples

### GET
```
$ curl -X GET "127.0.0.1:8888"

>>> 2022/05/17 14:23:37 Request from 127.0.0.1:58609
GET / HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
User-Agent: curl/7.79.1
```

```
$ curl -X GET "localhost:8888/page?parameter=value&also=another"

>>> 2022/05/17 14:33:35 Request from 127.0.0.1:58715
GET /page?parameter=value&also=another HTTP/1.1
Host: localhost:8888
Accept: */*
User-Agent: curl/7.79.1
```

### POST
```
$ curl -X POST 127.0.0.1:8888/form -d username=test -d password=test

>>> 2022/05/17 14:36:39 Request from 127.0.0.1:58734
POST /form HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 27
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.79.1

username=test&password=test
```

```
$ curl -X POST 127.0.0.1:8888/api -H "Content-Type: application/json" -d "{\"key\":\"value\"}"

>>> 2022/05/17 14:48:32 Request from 127.0.0.1:58970
POST /api HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 15
Content-Type: application/json
User-Agent: curl/7.79.1

{"key":"value"}
```

```
$ curl -X POST 127.0.0.1:8888/form -F "file=@empty.png"

>>> 2022/05/17 14:41:15 Request from 127.0.0.1:58814
POST /form HTTP/1.1
Host: 127.0.0.1:8888
Accept: */*
Content-Length: 328
Content-Type: multipart/form-data; boundary=------------------------9fc5f46d87def612
User-Agent: curl/7.79.1

--------------------------9fc5f46d87def612
Content-Disposition: form-data; name="file"; filename="empty.png"
Content-Type: image/png

�PNG

IHDR��&PLTE�����tRNS@��f9IDATx^��1� ����X������YIEND�B`�
--------------------------9fc5f46d87def612--
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/dassump/http-request-dump.
