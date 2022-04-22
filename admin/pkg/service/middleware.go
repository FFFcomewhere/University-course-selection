package service

import (
	"context"

	"github.com/FFFcomewhere/University-course-selection/http_cli/pkg/model"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HttpCliService) HttpCliService

type loggingMiddleware struct {
	logger log.Logger
	next   HttpCliService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HttpCliService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HttpCliService) HttpCliService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetName(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "GetName", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.GetName(ctx, id)
}
func (l loggingMiddleware) Post(ctx context.Context, user *model.User) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Post", "user", user, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Post(ctx, user)
}
func (l loggingMiddleware) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	defer func() {
		l.logger.Log("method", "Delete", "id", id, "status", status, "errinfo", errinfo, "data", data)
	}()
	return l.next.Delete(ctx, id)
}
