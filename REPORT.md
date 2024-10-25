# CHITTY-CHAT

## 1. Streaming

When deciding between server-side streaming, client-side streaming, or bidirectional streaming,
it's important to understand the differences between each.

### Server-Side Streaming

The client sends a single request to the server, and the server responds with a stream of data.
This is useful when the client needs to receive a continuous flow of data once a single request
is made, such as receiving live updates, notifications, or real-time data feeds.

### Client-Side Streaming

The client sends a stream of data to the server with a single request, and the server responds once when
all data is received. This is appropriate when the server needs to process or analyze a large stream of data
sent by the client before responsding, such as file uploads or real-time data collection.

### Bidirectional Streaming

Both client and server send streams of data to each other simultaneously in a single request. Each side
works independently. Ideal for scenarios requiring real-time, two-way communication, such as
chat applications or collaborative tools.

### This project

Since this application is a chat server, bidirectional streaming is best for the purpose of the project.

## 2. System architecture

This project uses a client-server architecture for communication. The client connects to the chat server and messages
is streamed between the server and the client. The clients are never directly exposed to eachother.

## 3. RPC methods

Only 1 RPC method is used called broadcast. It takes a stream of chat events and returns a stream of chat messages.

```proto
service Chat {
    rpc Broadcast(stream ChatEvent) returns (stream ChatMessage);
}
```

A chat event is either a join, leave or message event.

```proto
message ChatEvent {
    oneof event {
        UserJoin join = 2;
        UserLeave leave = 3;
        ChatMessage message = 4;
    }

    message UserJoin {
        string username = 1;
    }

    message UserLeave {
        string username = 1;
    }

    message ChatMessage {
        string username = 1;
        string message = 2;
    }
}
```

The server then turns the event into a message and streams it back to the client.

```proto
message ChatMessage {
    uint64 timestamp = 1;
    string username = 2;
    string message = 3;
}
```

## 4. Lamport timestamps

- [ ] Describe how you have implemented the calculation of the Lamport timestamps

## 5. Diagram of lamport

- [ ] Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joins, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.

## 6. Github repository

<https://github.com/Kanerix/chitty-chat>

## 7. System logs

The [systems logs](https://github.com/Kanerix/chitty-chat/blob/main/example.log) can be found in the GitHub repo.

## 8. README.md

The [README.md](https://github.com/Kanerix/chitty-chat/blob/main/readme.log) can be found in the GitHub repo.
