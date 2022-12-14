rpc_server:
	#set up a gRPC service
	protoc --go_out=. --go_opt=paths=source_relative \
		--micro_out=. --micro_opt=paths=source_relative \
		consignment/consignment.proto
docker:
	#Go build generate an executable file
	GOOS=linux GOARCH=amd64 go build
	#create an docker image by Dockerfile
	docker build -t kxxx/consignment-service-gomicro .
	#run a container using consignment-service-gomicro image.
	docker run -d -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 --name consignment-service-gomicro kxxx/consignment-service-gomicro

build:
