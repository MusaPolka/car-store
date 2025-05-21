FROM golang:1.22-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o inventory_service ./cmd/app
EXPOSE 8081
CMD ["./inventory_service"]
