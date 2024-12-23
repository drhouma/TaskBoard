FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -o app cmd/main.go

EXPOSE 8080

CMD ["./app"]
