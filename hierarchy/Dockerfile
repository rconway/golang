FROM golang:latest

WORKDIR /go/src/github.com/rconway/golang/hierarchy
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["hierarchy"]
