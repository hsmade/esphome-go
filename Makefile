proto/api.proto:
	curl https://raw.githubusercontent.com/esphome/esphome/dev/esphome/components/api/api.proto > proto/api.proto
proto/api_options.proto:
	curl https://raw.githubusercontent.com/esphome/esphome/dev/esphome/components/api/api_options.proto > proto/api_options.proto

.PHONY=generate
generate: proto/api.proto protobuf
	go generate ./...

.PHONY=build
build: generate
	go build cmd/server/server.go

all: build
clean:
	rm -f proto/*.proto proto/api/api.pb.go proto/api/options/api_options.pb.go
