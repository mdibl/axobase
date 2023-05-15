FROM golang:1.17

ADD . /go/src/axobase

RUN go env -w GO111MODULE=auto \
    && git config --global --add safe.directory '*' \ 
    && go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel

ENTRYPOINT revel run axobase

WORKDIR /go/src

EXPOSE 9100
