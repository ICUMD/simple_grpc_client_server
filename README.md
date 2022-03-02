# simple_grpc_client_server
A grpc-client server

1) Define the protocol in the .proto file
2) Run to download protoc-gen-go with go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
3) In a shell, cd to the root directory of your project, and run the following command: protoc -I=. --go_out=. --go-grpc_out=. ./api.proto 
