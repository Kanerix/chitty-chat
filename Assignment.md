# Assignment

You have to implement Chitty-Chat a distributed system, that is providing a chatting service, and keeps track of logical time using Lamport Timestamps.

We call clients of the Chitty-Chat service Participants.

## System Requirements

- [x] R1: Chitty-Chat is a distributed service, that enables its clients to chat. The service is using gRPC for communication. You have to design the API, including gRPC methods and data types.
- [x] R2: Clients in Chitty-Chat can Publish a valid chat message at any time they wish.  A valid message is a string of UTF-8 encoded text with a maximum length of 128 characters. A client publishes a message by making a gRPC call to Chitty-Chat.
- [x] R3: The Chitty-Chat service has to broadcast every published message, together with the current logical timestamp, to all participants in the system, by using gRPC. It is an implementation decision left to the students, whether a Vector Clock or a Lamport timestamp is sent.
- [x] R4: When a client receives a broadcasted message, it has to write the message and the current logical timestamp to the log
- [x] R5: Chat clients can join at any time.
- [x] R6: A "Participant X  joined Chitty-Chat at Lamport time L" message is broadcast to all Participants when client X joins, including the new Participant.
- [x] R7: Chat clients can drop out at any time.
- [x] R8: A "Participant X left Chitty-Chat at Lamport time L" message is broadcast to all remaining Participants when Participant X leaves.

## Technical Requirements

- [x] Use gRPC for all messages passing between nodes
- [x] Use Golang to implement the service and clients
- [x] Every client has to be deployed as a separate process
- [x] Log all service calls (Publish, Broadcast, ...) using the log package
- [x] Demonstrate that the system can be started with at least 3 client nodes
- [x] Demonstrate that a client node can join the system
- [x] Demonstrate that a client node can leave the system
- [x] Optional: All elements of the Chitty-Chat service are deployed as Docker containers

## Hand-in requirements

- [x] Hand in a single report in a pdf file
- [x] Discuss, whether you are going to use server-side streaming, client-side streaming, or bidirectional streaming?
- [x] Describe your system architecture - do you have a server-client architecture, peer-to-peer, or something else?
- [x] Describe what  RPC methods are implemented, of what type, and what messages types are used for communication
- [x] Describe how you have implemented the calculation of the Lamport timestamps
- [x] Provide a diagram, that traces a sequence of RPC calls together with the Lamport timestamps, that corresponds to a chosen sequence of interactions: Client X joins, Client X Publishes, ..., Client X leaves. Include documentation (system logs) in your appendix.
- [x] Provide a link to a Git repo with your source code in the report
- [x] Include system logs, that document the requirements are met, in the appendix of your report
- [x] Include a readme.md file that describes how to run your program.

## Grading notes

Partial implementations may be accepted, if the students can reason what they should have done in the report.
The group must at least *attempt* to solve each of the requirements. The whole assignment is failed, if a requirement is not addressed.
