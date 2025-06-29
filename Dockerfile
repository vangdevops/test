FROM golang:alpine AS builder
COPY . . 
RUN go build -o /app

FROM alpine
COPY --from=builder /app /app
ENTRYPOINT ["/app"]
