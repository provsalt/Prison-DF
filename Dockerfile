FROM golang:latest

WORKDIR /go/src/prisons
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 19132

CMD ["build -v ."]