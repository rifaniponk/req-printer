FROM golang:1.14

RUN go get github.com/codegangsta/gin

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN rm -rf /app/vendor
RUN go mod vendor

RUN go build -o main .

CMD ["/app/main"]