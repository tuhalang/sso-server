dev:
	go run cmd/main.go -conf=config.yml
prod:
	go run cmd/main.go -conf=prod-config.yml
proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./api/grpc/proto/*.proto
proto-clear:
	rm -rf ./domain/proto

.PHONY: dev prod proto-gen