# Chitty-Chat

## gRPC usage

Start of by logging in as an user. This will give you a token you can use to join the chat.

```bash
$ grpcurl -plaintext -format text -d 'username:"Bobby"' localhost:8080 chitty_chat.AuthService.Login
session_token: "Bobby:false:9ca1c330-0a60-4776-962a-3670e62459ce"
```

Then join the chat with your new session token.

```bash
$ grpcurl -plaintext -d '@' -H 'authorization: <session_token>' localhost:8080 chitty_chat.ChatService.Chat
"Bobby join the room"
```

Then logout and make your username available for others.

```bash
$ grpcurl -plaintext -format text -d 'session_token: <session_token>' localhost:8080 chitty_chat.AuthService.Logout
"success"
```
