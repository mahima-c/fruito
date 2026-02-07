package health

import "context"

type Service interface {
	Health(ctx context.Context) string
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Health(ctx context.Context) string {
	return "pass"
}
