FROM golang:1.22-alpine3.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build

EXPOSE 8080

CMD ["/app/redir"]
