FROM golang:1.10.0

ENV GOROOT /usr/local/go
ENV GOPATH /go/vendor

RUN go get github.com/labstack/echo

WORKDIR /go/src

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]