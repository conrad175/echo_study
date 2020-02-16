FROM golang:1.12.0

RUN go get github.com/labstack/echo/...
RUN go get github.com/gomodule/redigo/redis

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]
