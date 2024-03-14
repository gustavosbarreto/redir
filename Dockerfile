FROM golang:1.22-alpine3.19

WORKDIR /app

COPY . /app

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
