FROM golang:1.17

ADD . /go/src/axobase

ADD /data/antibodies.csv /tmp/
ADD /papers/Axolotl_White_Paper_Final.pdf /tmp/

RUN CGO_ENABLED=0 go env -w GO111MODULE=auto \
    && git config --global --add safe.directory '*' \ 
    && go get github.com/revel/revel \
    && go get github.com/revel/cmd/revel

ENTRYPOINT revel run axobase

WORKDIR /go/src

EXPOSE 9100
