module accommodation_service

go 1.20

replace github.com/tamararankovic/microservices_demo/common => ../common

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/tamararankovic/microservices_demo/common v1.0.0
	go.mongodb.org/mongo-driver v1.8.4
<<<<<<< HEAD
	google.golang.org/grpc v1.55.0
=======
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
>>>>>>> 859ba3a (implemented creating of accommodation)
)

require (
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
<<<<<<< HEAD
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
=======
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
>>>>>>> 859ba3a (implemented creating of accommodation)
)
