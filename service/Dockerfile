FROM golang:1.20-alpine3.16

ARG WORK_DIR=/service

RUN apk add gcc musl-dev

WORKDIR $WORK_DIR

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN cd ./src/di && go run github.com/google/wire/cmd/wire && cd -

RUN go build ./src/server.go

CMD ./server
