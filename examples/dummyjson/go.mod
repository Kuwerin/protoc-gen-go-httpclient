module dummyjson

go 1.23.0

toolchain go1.23.11

replace github.com/Kuwerin/protoc-gen-go-httpclient => ../../../protoc-gen-go-httpclient

require (
	github.com/Kuwerin/protoc-gen-go-httpclient v0.3.0
	github.com/go-kit/log v0.2.1
	github.com/golang/protobuf v1.5.4
	google.golang.org/genproto/googleapis/api v0.0.0-20250721164621-a45f3dfb1074
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250715232539-7130f93afb79 // indirect
)
