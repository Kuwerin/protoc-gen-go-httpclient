package descriptor

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// Registry is a registry of information extracted from protogen.File.
type ServiceRegistry struct {
	PackageName string

	ServiceName string

	Methods []*protogen.Method

	HTTPRules map[string]*HTTPRule
}

func NewServiceRegistry(file *protogen.File) *ServiceRegistry {
	var r ServiceRegistry

	service := file.Services[0]

	r.PackageName = string(file.GoPackageName)
	r.ServiceName = service.GoName
	r.Methods = make([]*protogen.Method, 0)
	r.HTTPRules = make(map[string]*HTTPRule, 0)

	for _, method := range service.Methods {
		r.HTTPRules[method.GoName] = HTTPRuleFromMethod(method)
		r.Methods = append(r.Methods, method)
	}

	return &r
}

// Gets HTTPRule for protogen.Method.
func (r *ServiceRegistry) GetHTTPRuleForMethod(method *protogen.Method) (*HTTPRule, bool) {
	rule, exists := r.HTTPRules[method.GoName]
	return rule, exists
}
