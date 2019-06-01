# Errol

Errol is a networking application demonstrating message delivery, written in Go.

It contains the following components:

### Server

The Server relays incoming message bodies to receivers based on client ID(s) defined in the message. The server does not perform authentication, rather arbitrarily assigns an integer client id to the client once its connected.

- id - unsigned 64 bit integer (monotonically increasing)
- Connection to the server must be done using pure TCP. Protocol doesnt require multiplexing.

### Clients

Clients are clients who are connected to the server. Client may send three types of messages which are described below.

### WhoAmI message

Client can send a identity message which the server will answer with the client_id of the connected client.

### GetConnectedClients (List) message

Client can send a list message which the server will answer with the list of all connected client client_ids (excluding the requesting client).

### Relay message

Client can send a relay messages which body is relayed to receivers marked in the message. The data format used is JSON. Message body can be relayed to one or multiple receivers.

- max 255 receivers (client_id:s) per message
- message body - byte array (text, JSON, binary, or anything), max length 1024 kilobytes

*Relay example: receivers: 2 and 3, body: foobar*

## Known issues:

- server should detect client connection closures and remove those clients
- implement request IDs and request queueing on the client so that requesting methods can return the result
- server, client networking code is not well-tested with unit tests
- serialization suppresses errors; should return them instead
