## gRPC implementation in Golang

This project demonstrate use of grpc.
gRPC (gRPC Remote Procedure Call) is a high-performance, open-source, universal RPC framework initially developed by Google.
It uses HTTP/2 for transport, Protocol Buffers (protobuf) as the interface description language.
It also provides features such as authentication, load balancing, and more.

### Included in this project 
- A unary API service which follows single request response model.
- Two stream models, one with client stream and other with server stream.
- One Bi-stream model where both client and server can stream data asynchronously and simultaniously.

### Packages used
- proto - for implementing protocol buffel compiler
- context - to create context which can run in background and helps in prevention of memory leaks
- other packages include io, time, log etc.


For additional notes refer "https://github.com/kushagraag/grpc-golang/blob/main/Notes.txt"