FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /binary
EXPOSE 8080

ENTRYPOINT ["/app/binary"]

FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder binary .
COPY .env .
EXPOSE 8080
CMD ["./binary"]