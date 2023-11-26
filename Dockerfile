FROM golang:1.21.4 AS builder

WORKDIR /app

COPY . .

RUN go build -o fetch cmd/main.go

FROM scratch

COPY --from=builder /app/fetch /fetch

ENTRYPOINT ["/fetch"]

