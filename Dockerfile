FROM alpine

COPY gopath/bin/test /go/bin/test

ENTRYPOINT /go/bin/test
