go_dir=gen/go

user_files = proto/typicode/user/*.proto

proto_files=\
			$(user_files)

lint:
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm bufbuild/buf lint proto

compile: compile-go compile-redoc

compile-go:
	docker build -t doc-compiler -f Dockerfile.gen .
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm doc-compiler make compile-go-in-docker

compile-go-in-docker:
	cd $(go_dir); rm -rf *

	protoc -Iproto \
		--go_out=$(go_dir) \
		--go_opt paths=source_relative \
		--go-grpc_out=$(go_dir) \
		--go-grpc_opt paths=source_relative \
		$(proto_files)

	# User
	protoc -Iproto \
		--openapiv2_out=gen/openapi/user \
		--openapiv2_opt allow_merge=true \
		--openapiv2_opt json_names_for_fields=false \
		--openapiv2_opt disable_default_errors=true \
		--openapiv2_opt omit_enum_default_value=true \
		--openapiv2_opt openapi_configuration=config/openapi/user.yaml \
		--openapiv2_opt openapi_naming_strategy=fqn \
		$(user_files)

	protoc -Iproto \
		--go-httpclient_out=$(go_dir) \
		--go-httpclient_opt logging_middleware=true \
		--go-httpclient_opt paths=source_relative \
		$(proto_files)

	cd $(go_dir); mv -f typicode/* .; rm -rf typicode

compile-redoc:
	docker build -t redoc-cli -f Dockerfile.redoc .

	docker run --volume "$(shell pwd):/workspace" --workdir /workspace --rm redoc-cli build \
		--options.hideSchemaTitles --options.pathInMiddlePanel --options.hideHostname --options.jsonSampleExpandLevel=all \
		-o gen/openapi/user/apidocs.swagger.html gen/openapi/user/apidocs.swagger.json
