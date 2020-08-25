build:
	protoc --go_out=plugins=grpc:proto proto/notifications.proto

