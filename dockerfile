# syntax=docker/dockerfile:1
FROM --platform=linux/amd64 scratch
COPY build/http-request-dump-linux-amd64 /
CMD ["/http-request-dump-linux-amd64"]
EXPOSE 8888