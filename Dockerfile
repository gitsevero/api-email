FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main .
FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY config.json /app/config.json
CMD ["/app/main"]
