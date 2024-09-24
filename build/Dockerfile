FROM golang:1.23.0-bullseye

WORKDIR /usr/src/app

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download && go mod verify

COPY ./src .
RUN go build -v -o /usr/local/bin/ ./...

CMD ["api"]
