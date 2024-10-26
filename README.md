# Chitty-Chat

## Usage Server

This section will tell you how to start the Chitty-Chat server.

### Go CLI

Use the Go CLI to start the server.

```bash
$ go run ./cmd/grpc
2024/10/26 20:52:41 server is listening on [::]:8080
...
```

You can also use `make` to do the same.

```bash
$ make grpc-serve
2024/10/26 20:52:41 server is listening on [::]:8080
...
```

### Docker

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
...
```

## Usage Client

This section will tell you how to install and use the client.

### Install from repo

Install the `chitty` client from the github repository.

```bash
go install github.com/kanerix/chitty-chat/cmd/chitty
```

You can now use the CLI to connect to the server.

```bash
# chitty chat -u [username] -H [hostname]
chitty chat -u Kanerix -H localhost:8080
```

### Run locally

You can also clone the repository and run it from there.

```bash
git clone https://github.com/kanerix/chitty-chat
```

Then run the CLI.

```bash
# go run ./cmd/chitty -u [username] -H [hostname]
go run ./cmd/chitty -u kanerix -H localhost:8080
```

### Help

If you need any help, you can use the `--help` flag.

```bash
chitty chat -u [username]
```
