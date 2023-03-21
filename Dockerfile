FROM golang:1.20

ADD . /go/src/axobase

RUN go env -w GO111MODULE=auto \
    && go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel

ENTRYPOINT revel run axobase

WORKDIR /go/src

EXPOSE 9100
