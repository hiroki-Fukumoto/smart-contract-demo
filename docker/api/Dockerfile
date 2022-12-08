FROM golang:1.19.3-buster

ENV GO111MODULE=auto

WORKDIR /go/src
COPY ./src .

RUN go get -u github.com/cosmtrek/air && \
  go build -o /go/bin/air github.com/cosmtrek/air && \
  go get -u github.com/gin-gonic/gin && \
  go get -u github.com/go-delve/delve/cmd/dlv

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go get -u github.com/swaggo/swag/cmd/swag

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
