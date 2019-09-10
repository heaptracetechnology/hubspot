FROM golang

RUN go get github.com/gorilla/mux

WORKDIR /go/src/github.com/heaptracetechnology/hubspot

ADD . /go/src/github.com/heaptracetechnology/hubspot

RUN go install github.com/heaptracetechnology/hubspot

ENTRYPOINT hubspot

EXPOSE 3000