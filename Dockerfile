FROM alpine

COPY gopath/bin/spinlab /go/bin/spinlab

ENRYPOINT /go/bin/spinlab