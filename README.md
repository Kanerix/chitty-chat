# Chitty-Chat

Welcome to the Chitty-Chat project. This is developed following the course **Distributed Systems** at ITU.

## Usage server

This section will tell you how to start the Chitty-Chat server.

### Run server using Docker

Build the Chitty-Chat server docker container.

```bash
$ docker build . -t chitty-chat-server
...
[+] Building 10.0s (22/22) FINISHED
...
```

Run the docker container.

```bash
$ docker run chitty-chat-server
2024/10/26 20:52:41 server is listening on [::]:8080
```

### Run server using Go CLI

Clone the repository and run it using the Go CLI.

```bash
git clone https://github.com/kanerix/chitty-chat
```

Use the Go CLI to start the server.

```bash
$ go run ./cmd/grpc
2024/10/26 20:52:41 server is listening on [::]:8080
```

OPTIONAL: You can also use `make` to do the same.

```bash
$ make grpc-serve
2024/10/26 20:52:41 server is listening on [::]:8080
```

## Usage client

This section will tell you how to install and use the client.

### Install client from repo

Install the `chitty` client from the github repository.

```bash
go install github.com/kanerix/chitty-chat/cmd/chitty
```

Use the CLI to connect to the server.

```bash
# chitty chat -u [username] -H [hostname]
chitty chat -u Kanerix -H localhost:8080
```

### Install client using Go CLI

Clone the repository and run it using the Go CLI.

```bash
git clone https://github.com/kanerix/chitty-chat
```

Run the CLI with a username and optional hostname (default localhost:8080).

```bash
# go run ./cmd/chitty -u [username] -H [hostname]
go run ./cmd/chitty -u kanerix -H localhost:8080
```

### Get help using client

If you need any help, you can use the `--help` flag.

```bash
chitty chat --help
```
