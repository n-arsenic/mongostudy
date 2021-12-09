FROM golang:1.16-alpine

WORKDIR /go/src/mongostudy

COPY . .
RUN go mod download

RUN go build -o /go/bin/mongostudy

ENTRYPOINT ["/go/bin/mongostudy"]