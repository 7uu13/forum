FROM golang:1.20

ADD . /go/src/app
WORKDIR /go/src/app

RUN go get github.com/7uu13/forum
RUN go build -o app
RUN go install

EXPOSE 8080

CMD ["./app"]