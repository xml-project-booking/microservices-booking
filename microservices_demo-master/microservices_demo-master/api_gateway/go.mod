module github.com/tamararankovic/microservices_demo/api_gateway

go 1.17

replace github.com/tamararankovic/microservices_demo/common => ../common

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0
	github.com/tamararankovic/microservices_demo/common v1.0.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
