FROM golang:1.23.2-alpine3.20 AS proto-builder

WORKDIR /app

RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache protoc

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ADD ./grpc /app/grpc

RUN protoc --go_out=. --go-grpc_out=. ./grpc/*.proto


FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app

COPY --from=proto-builder /app/pb /app/pb

ADD . .

RUN go build -o bin/server server/main.go


FROM alpine:3.20 AS runner

WORKDIR /var/app

COPY --from=builder /app/bin/server .

RUN addgroup -S app && \
    adduser -S chitty-chat -G app && \
    chown -R chitty-chat:app /app

USER chitty-chat

EXPOSE 8080

CMD ["./server"]
