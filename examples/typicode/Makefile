go_dir=gen/go

user_files = proto/typicode/user/*.proto

proto_files=\
			$(user_files)

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm bufbuild/buf lint proto

compile: compile-go

compile-go:
	docker build -t doc-compiler -f Dockerfile.gen .
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm doc-compiler make compile-go-in-docker

compile-go-in-docker:
	cd $(go_dir); rm -rf *;

	protoc -Iproto \
		--go_out=$(go_dir) \
		--go_opt paths=source_relative \
		--go-grpc_out=$(go_dir) \
		--go-grpc_opt paths=source_relative \
		$(proto_files)

	# User
	protoc -Iproto \
		--go-httpclient_out=$(go_dir) \
		--go-httpclient_opt logging_middleware=true \
		--go-httpclient_opt paths=source_relative \
		$(proto_files)

	cd $(go_dir); mv -f typicode/* .; rm -rf typicode
