CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -ldflags "-s -w" -o http-request-dump-linux-amd64       .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o http-request-dump-windows-amd64.exe .
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -ldflags "-s -w" -o http-request-dump-darwin-amd64      .
CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -ldflags "-s -w" -o http-request-dump-darwin-arm64      .

lipo -create -output http-request-dump-darwin-universal http-request-dump-darwin-amd64 http-request-dump-darwin-arm64
