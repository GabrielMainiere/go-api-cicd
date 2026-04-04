FROM golang:1.25-alpine3.22 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM alpine:3.22
WORKDIR /app
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/api /app/api
EXPOSE 3000
USER appuser
ENTRYPOINT ["/app/api"]
