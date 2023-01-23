FROM golang:latest

WORKDIR /app

COPY . .

RUN go build main.go

EXPOSE 3000

CMD ["./main"]