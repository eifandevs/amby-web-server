FROM golang:1.8

WORKDIR /go/src

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]