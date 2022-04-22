package service

import (
	"context"

	"github.com/FFFcomewhere/University-course-selection/http-cli/pkg/model"
)

// HttpCliService describes the service.
type HttpCliService interface {
	GetName(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
	Post(ctx context.Context, user *model.User) (status bool, errinfo string, data *model.User)
	Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User)
}
