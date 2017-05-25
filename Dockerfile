FROM golang

WORKDIR /go/src/github.com/talbor49/HoneyBee
COPY . .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

EXPOSE 8080

CMD ["go-wrapper", "run"] # ["app"]