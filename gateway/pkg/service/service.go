package service

import (
	"context"

	"github.com/FFFcomewhere/University-course-selection/http_cli/pkg/model"
)

// HttpCliService describes the service.
type HttpCliService interface {
	GetName(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
	Post(ctx context.Context, user *model.User) (status bool, errinfo string, data *model.User)
	Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
}

type basicHttpCliService struct{}

func (b *basicHttpCliService) GetName(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of GetName
	return status, errinfo, data
}
func (b *basicHttpCliService) Post(ctx context.Context, user *model.User) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of Post
	return status, errinfo, data
}
func (b *basicHttpCliService) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	// TODO implement the business logic of Delete
	return status, errinfo, data
}

// NewBasicHttpCliService returns a naive, stateless implementation of HttpCliService.
func NewBasicHttpCliService() HttpCliService {
	return &basicHttpCliService{}
}

// New returns a HttpCliService with all of the expected middleware wired in.
func New(middleware []Middleware) HttpCliService {
	var svc HttpCliService = NewBasicHttpCliService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
