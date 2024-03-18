
FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go build -o main .
FROM alpine
COPY --from=builder /app/main /app/main
CMD ["/app/main"]
