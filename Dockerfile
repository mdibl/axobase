FROM golang:1.15

ADD . /go/src/axobase

RUN go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel \
    && go get axobase/...

ENTRYPOINT revel run axobase

WORKDIR /go/src

EXPOSE 9100
