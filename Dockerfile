FROM golang:1.23.2-alpine3.20 AS chef

WORKDIR /app

RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache protoc make


FROM chef AS proto-builder

WORKDIR /build

RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache protoc make

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY proto proto

RUN --mount=type=bind,source=Makefile,target=Makefile \
    make proto


FROM chef AS grpc-builder

WORKDIR /build

COPY --from=proto-builder /build/proto proto

ADD . .

RUN make grpc-build


FROM alpine:3.20 AS runner

WORKDIR /var/app

COPY --from=grpc-builder /build/bin/grpc .

RUN addgroup -S app && \
    adduser -S chitty -G app && \
    chown -R chitty:app /var/app

USER chitty

EXPOSE 8080

CMD ["./grpc"]
