FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync

CMD ["go", "run", "run.go"]
