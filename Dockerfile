FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

COPY storage-napp /go/bin/app

RUN chmod +x /go/bin/storage-napp

CMD ["/go/bin/storage-napp"]
