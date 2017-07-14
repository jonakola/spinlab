FROM alpine

COPY gopath/bin/spinlab /go/bin/spinlab

ENTRYPOINT /go/bin/spinlab