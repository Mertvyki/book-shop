FROM golang:1.26-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/bookshop/

FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -h /app appuser
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/migrations ./migrations
USER appuser
EXPOSE 5050
CMD ["./server"]