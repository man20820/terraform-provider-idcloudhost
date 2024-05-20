FROM golang:1.16

RUN echo $GOPATH
RUN mkdir -p $GOPATH/src/github.com/bapung
WORKDIR $GOPATH/src/github.com/bapung
RUN git clone https://github.com/man20820/terraform-provider-idcloudhost
WORKDIR $GOPATH/src/github.com/man20820/terraform-provider-idcloudhost
RUN make install
