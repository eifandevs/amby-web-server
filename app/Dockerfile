FROM golang:1.12.9

ENV GO111MODULE=off

RUN apt-get update \
 && apt-get -y install default-mysql-client \
 && go get -u github.com/labstack/echo/... \
 && go get -u github.com/jinzhu/gorm \
 && go get -u github.com/jinzhu/gorm/dialects/mysql

WORKDIR /go/src/github.com/eifandevs/amby

# RUN go get github.com/labstack/echo \
        #    github.com/labstack/echo/middleware

# CMD ["go", "run", "main.go"]
# CMD ["tail", "-f", "/dev/null"]
