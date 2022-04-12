FROM golang:alpine

COPY . /app
WORKDIR /app

EXPOSE 8080

RUN CGO_ENABLE=0 GOOS=linux go build -o main

CMD ["./main"]