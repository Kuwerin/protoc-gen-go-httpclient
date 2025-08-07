go_out_dir=out/go

user_files=proto/dummyjson/user/v1/*.proto

proto_files=\
			$(user_files)

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm bufbuild/buf lint proto

compile: compile-go-in-docker

compile-go-in-docker:
	docker build -t doc-compiler -f Dockerfile.gen .
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm doc-compiler make compile-go

compile-go:
	mkdir -p $(go_out_dir)
	cd $(go_out_dir) && rm -rf *;

  # User
	protoc -Iproto \
		--go_out=$(go_out_dir) \
		--go_opt paths=source_relative \
		--go-grpc_out=$(go_out_dir) \
		--go-grpc_opt paths=source_relative \
		--go-httpclient_out=$(go_out_dir) \
		--go-httpclient_opt logging_middleware=true \
		--go-httpclient_opt paths=source_relative \
		$(proto_files)

	cd $(go_out_dir); mv -f dummyjson/* .; rm -rf dummyjson
