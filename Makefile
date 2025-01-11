default:
	go run build.go
build_proto:
	protoc --go_out=paths=source_relative:. proto/plugin.proto
	protoc-go-inject-tag -input=./proto/driver_plugin/driver_plugin.pb.go
