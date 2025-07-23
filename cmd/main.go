package main

import (
	"flag"
	"os"

	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/Kuwerin/protoc-gen-go-httpclient/internal/descriptor"
	"github.com/Kuwerin/protoc-gen-go-httpclient/internal/generator"
)

var (
	loggingMiddleware = flag.Bool("logging_middleware", false, "specifies whether to generate logging middleware")
)

func main() {
	flag.Parse()

	var opts protogen.Options

	opts.ParamFunc = flag.CommandLine.Set

	opts.Run(run)

}

func run(gen *protogen.Plugin) error {
	// Add proto3 syntax optional support.
	gen.SupportedFeatures = supportedCodeGeneratorFeatures()

	// Finds all files that contains services description.
	var serviceFiles = make([]*protogen.File, 0)
	{
		for _, file := range gen.Files {
			if len(file.Services) == 0 {
				continue
			}

			serviceFiles = append(serviceFiles, file)
		}
	}

	for _, serviceFile := range serviceFiles {
		registry := descriptor.NewServiceRegistry(serviceFile)

		file := gen.NewGeneratedFile(serviceFile.GeneratedFilenamePrefix+"_httpclient.pb.go", serviceFile.GoImportPath)

		if err := generator.Generate(file, registry); err != nil {
			emitError(err)
			return nil
		}

		if *loggingMiddleware {
			file := gen.NewGeneratedFile(serviceFile.GeneratedFilenamePrefix+"_httpclient_logging.pb.go", serviceFile.GoImportPath)

			if err := generator.GenerateLoggingMiddleware(file, registry); err != nil {
				emitError(err)
				return nil
			}
		}
	}

	return nil
}

// Add protobuf3 features for optional features suport.
func supportedCodeGeneratorFeatures() uint64 {
	return uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
}

func emitError(err error) {
	resp := &pluginpb.CodeGeneratorResponse{Error: proto.String(err.Error())}

	buf, err := proto.Marshal(resp)
	if err != nil {
		grpclog.Fatal(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		grpclog.Fatal(err)
	}
}
