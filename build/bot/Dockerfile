FROM golang:1.19-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -v -o /usr/local/bin/bot github.com/patriarch11/go-tg-task-bot/cmd/bot

FROM alpine
COPY --from=builder /usr/local/bin/bot /usr/local/bin/bot
ENTRYPOINT ["/usr/local/bin/bot"]