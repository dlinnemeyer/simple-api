FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/kardianos/govendor
RUN govendor sync

# RUN go-wrapper download   # "go get -d -v ./..."
# RUN go-wrapper install    # "go install -v ./..."

# CMD ["go-wrapper", "run"] # ["app"]
CMD ["go", "run", "run.go"]
