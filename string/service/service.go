package service

import "errors"

// Service constants
const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

type Service interface {
	Concat(a, b string) (string, error)

	HealthCheck() bool

	Say(a string) string
}

type StringService struct {
}

func (s StringService) Concat(a, b string) (string, error) {
	if len(a)+len(b) > StrMaxSize {
		return "", ErrMaxSize
	}
	return a + b, nil
}

// 用于检查服务的健康状态，这里仅仅返回true。
func (s StringService) HealthCheck() bool {
	return true
}

// 对话
func (s StringService) Say(a string) string {
	return a
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(Service) Service
