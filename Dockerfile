FROM golang:1.12.9

ENV GO111MODULE=off
RUN go get github.com/labstack/echo/...

WORKDIR /go/src/github.com/eifandevs/amby

# RUN go get github.com/labstack/echo \
        #    github.com/labstack/echo/middleware

# CMD ["go", "run", "main.go"]
CMD ["tail", "-f", "/dev/null"]