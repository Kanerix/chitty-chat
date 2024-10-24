# CHITTY-CHAT

## 1 Streaming

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

## 2 System architecture

This project uses a client-server architecture for communication. The client connects to the chat server and messages
is streamed between the server and the client. The clients are never directly exposed to eachother.

## 3 RPC methods

- [ ] Describe what  RPC methods are implemented, of what type, and what messages types are used for communication

## 4  Lamport timestamps

- [ ] Describe how you have implemented the calculation of the Lamport timestamps

## 5 Diagram of lamport

- [ ] Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joins, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.

## Github repository

- [ ] Provide a link to a Git repo with your source code in the report

## System logs

- [ ] Include system logs, that document the requirements are met, in the appendix of your report

## README.md

- [ ] Include a readme.md file that describes how to run your program.
