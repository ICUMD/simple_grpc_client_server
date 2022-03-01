# simple_grpc_client_server
A grpc-client server

1) Define the protocol in the .proto file
2) In a shell, cd to the root directory of your project, and run the following command: protoc -I=. --go_out=. ./sample_server.proto
