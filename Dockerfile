FROM golang:1.21-alpine AS builder

WORKDIR /cmd

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN make build]

FROM apline:latest

WORKDIR /root/

COPY --from=builder /bin/chat-backend-app .

EXPOSE 8080

CMD ["./chat-backend-app"]
