package descriptor

import (
	"errors"
	"net/http"
	"regexp"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type MethodSemantics int

const (
	MethodSemanticsUnspecified MethodSemantics = iota

	MethodSemanticsListEntities

	MethodSemanticsBatchGetEntities

	MethodSemanticsGetOneEntity
)

type HTTPRule struct {
	RequestType string

	URL string

	Semantics MethodSemantics
}

// Extracts HTTP rule from protogen method.
func HTTPRuleFromMethod(method *protogen.Method) *HTTPRule {
	annotation, err := extractHTTPRuleAnnotation(method)
	if err != nil {
		return nil
	}

	var rule HTTPRule

	rule.RequestType = requestType(annotation.GetPattern())
	rule.URL = url(annotation.GetPattern())
	rule.Semantics = semantics(method.GoName)

	return &rule
}

// Extracts request type from google.api.Http annotation.
func requestType(pattern any) string {
	var requestType string

	switch pattern.(type) {
	case *annotations.HttpRule_Get:
		requestType = http.MethodGet
	case *annotations.HttpRule_Post:
		requestType = http.MethodPost
	case *annotations.HttpRule_Patch:
		requestType = http.MethodPatch
	case *annotations.HttpRule_Put:
		requestType = http.MethodPut
	case *annotations.HttpRule_Delete:
		requestType = http.MethodDelete
	case *annotations.HttpRule_Custom:
		requestType = pattern.(*annotations.HttpRule_Custom).Custom.Kind
	}

	return requestType
}

// Extracts URL from google.api.Http annotation.
func url(pattern any) string {
	var url string

	switch pattern.(type) {
	case *annotations.HttpRule_Get:
		url = pattern.(*annotations.HttpRule_Get).Get
	case *annotations.HttpRule_Post:
		url = pattern.(*annotations.HttpRule_Post).Post
	case *annotations.HttpRule_Patch:
		url = pattern.(*annotations.HttpRule_Patch).Patch
	case *annotations.HttpRule_Put:
		url = pattern.(*annotations.HttpRule_Put).Put
	case *annotations.HttpRule_Delete:
		url = pattern.(*annotations.HttpRule_Delete).Delete
	case *annotations.HttpRule_Custom:
		url = pattern.(*annotations.HttpRule_Custom).Custom.GetPath()
	}

	return url
}

var reBatchGet = regexp.MustCompile(`Batch\w*$`)
var reList = regexp.MustCompile(`List\w*$`)

// Resolves method semantics based on method name which can be:
// `ListEntities`, `BatchGetEntities`, `GetOneEntity`.
func semantics(selector string) MethodSemantics {
	if reList.MatchString(selector) {
		return MethodSemanticsListEntities
	}

	if reBatchGet.MatchString(selector) {
		return MethodSemanticsBatchGetEntities
	}

	return MethodSemanticsGetOneEntity
}

// Extracts method's google.api.Http annotation.
func extractHTTPRuleAnnotation(method *protogen.Method) (*annotations.HttpRule, error) {
	opts, ok := method.Desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, errors.New("method options was not found")
	}

	annotation, ok := proto.GetExtension(opts, annotations.E_Http).(*annotations.HttpRule)
	if !ok {
		return nil, errors.New("http annotation was not found")
	}

	return annotation, nil
}
