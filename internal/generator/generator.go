package generator

import (
	"regexp"
	"sort"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/Kuwerin/protoc-gen-go-httpclient/internal/descriptor"
	"github.com/Kuwerin/protoc-gen-go-httpclient/internal/generator/templates/logging"
	"github.com/Kuwerin/protoc-gen-go-httpclient/internal/generator/templates/server"
)

// Generates HTTP-client from input .proto files.
func Generate(file *protogen.GeneratedFile, registry *descriptor.ServiceRegistry) error {
	var headerData = struct {
		Registry *descriptor.ServiceRegistry
		Modules  string
	}{
		Registry: registry,
		Modules:  modules(registry),
	}

	server.Header.Execute(
		file,
		headerData,
	)
	for _, method := range registry.Methods {

		rule, exists := registry.GetHTTPRuleForMethod(method)
		if !exists {
			continue
		}

		var methodTemplate *template.Template

		switch rule.Semantics {
		case descriptor.MethodSemanticsGetOneEntity:
			methodTemplate = server.GetOne
		case descriptor.MethodSemanticsBatchGetEntities:
			methodTemplate = server.BatchGet
		case descriptor.MethodSemanticsListEntities:
			methodTemplate = server.List
		default:
			continue
		}

		var methodData = struct {
			Comment string
			Method  *protogen.Method
			Rule    *descriptor.HTTPRule
		}{
			Comment: method.Comments.Leading.String(),
			Method:  method,
			Rule:    rule,
		}

		methodTemplate.Execute(file, methodData)
	}

	return nil
}

// Generates logging middleware for HTTP-client from input .proto files.
func GenerateLoggingMiddleware(file *protogen.GeneratedFile, registry *descriptor.ServiceRegistry) error {
	var headerData = struct {
		Registry *descriptor.ServiceRegistry
	}{registry}

	logging.Header.Execute(
		file,
		headerData,
	)

	for _, method := range registry.Methods {
		rule, exists := registry.GetHTTPRuleForMethod(method)
		if !exists {
			continue
		}

		var methodTemplate *template.Template

		switch rule.Semantics {
		case descriptor.MethodSemanticsGetOneEntity:
			methodTemplate = logging.GetOne
		case descriptor.MethodSemanticsBatchGetEntities:
			methodTemplate = logging.BatchGet
		case descriptor.MethodSemanticsListEntities:
			methodTemplate = logging.List
		default:
			continue
		}

		var methodData = struct {
			Comment              string
			Package              string
			UnderscoreMethodname string
			Method               *protogen.Method
		}{
			Comment:              method.Comments.Leading.String(),
			Package:              toUndescoreCase(string(method.Desc.FullName().Parent())),
			UnderscoreMethodname: toUndescoreCase(method.GoName),
			Method:               method,
		}

		methodTemplate.Execute(file, methodData)
	}

	return nil
}

// Resolves registry dependencies.
func modules(registry *descriptor.ServiceRegistry) string {
	var modules = []string{
		`"context"`,
		`"net/http"`,
		`"net/url"`,
		`"regexp"`,
		`"strings"`,
		`"github.com/golang/protobuf/jsonpb"`,
	}

	var batchGetPresent bool
	var getByPagePresent bool
	for _, method := range registry.Methods {
		rule, exists := registry.GetHTTPRuleForMethod(method)
		if !exists {
			continue
		}

		switch rule.Semantics {
		case descriptor.MethodSemanticsBatchGetEntities:
			batchGetPresent = true
		case descriptor.MethodSemanticsListEntities:
			getByPagePresent = true
		}
	}

	if batchGetPresent {
		modules = append(modules, []string{`"bytes"`, `"encoding/json"`}...)
	}
	if getByPagePresent {
		modules = append(modules, `"strconv"`)
	}

	sort.Strings(modules)

	return strings.Join(modules, "\n")
}

var reCamelCase = regexp.MustCompile("([a-z])([A-Z])")

// Converts string to underscore case.
func toUndescoreCase(str string) string {
	return strings.ToLower(reCamelCase.ReplaceAllString(str, "${1}_${2}"))
}
