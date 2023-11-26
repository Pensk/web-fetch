FROM golang:1.21.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o /out/fetch cmd/main.go

FROM alpine

COPY --from=builder /out/fetch /fetch

ENTRYPOINT ["/fetch"]

