FROM golang:1.8

WORKDIR /go/src/app

ADD ./go .

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8000

CMD ["app"]

