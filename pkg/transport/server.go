package transport

import (
	"errors"

	"google.golang.org/grpc"
)

type RegisterServerFunc[T any] func(grpc.ServiceRegistrar, T)

type Server[T any] struct {
	*grpc.Server
	// The server API of the service.
	ContextAPI *T

	// Server registrer function.
	RegisterServerFunc[T]
}

func (s *Server[T]) Register() error {
	if s.ContextAPI == nil {
		return errors.New("service must be provided")
	}

	s.RegisterServerFunc(s.Server, *s.ContextAPI)

	return nil
}
