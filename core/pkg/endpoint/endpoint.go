package endpoint

import (
	service "University-course-selection/http_cli/pkg/service"
	"context"

	"github.com/FFFcomewhere/University-course-selection/http_cli/pkg/model"
	endpoint "github.com/go-kit/kit/endpoint"
)

// GetNameRequest collects the request parameters for the GetName method.
type GetNameRequest struct {
	Id string `json:"id"`
}

// GetNameResponse collects the response parameters for the GetName method.
type GetNameResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakeGetNameEndpoint returns an endpoint that invokes GetName on the service.
func MakeGetNameEndpoint(s service.HttpCliService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNameRequest)
		status, errinfo, data := s.GetName(ctx, req.Id)
		return GetNameResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// PostRequest collects the request parameters for the Post method.
type PostRequest struct {
	User *model.User `json:"user"`
}

// PostResponse collects the response parameters for the Post method.
type PostResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakePostEndpoint returns an endpoint that invokes Post on the service.
func MakePostEndpoint(s service.HttpCliService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostRequest)
		status, errinfo, data := s.Post(ctx, req.User)
		return PostResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// DeleteRequest collects the request parameters for the Delete method.
type DeleteRequest struct {
	Id string `json:"id"`
}

// DeleteResponse collects the response parameters for the Delete method.
type DeleteResponse struct {
	Status  bool        `json:"status"`
	Errinfo string      `json:"errinfo"`
	Data    *model.User `json:"data"`
}

// MakeDeleteEndpoint returns an endpoint that invokes Delete on the service.
func MakeDeleteEndpoint(s service.HttpCliService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		status, errinfo, data := s.Delete(ctx, req.Id)
		return DeleteResponse{
			Data:    data,
			Errinfo: errinfo,
			Status:  status,
		}, nil
	}
}

// GetName implements Service. Primarily useful in a client.
func (e Endpoints) GetName(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	request := GetNameRequest{Id: id}
	response, err := e.GetNameEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetNameResponse).Status, response.(GetNameResponse).Errinfo, response.(GetNameResponse).Data
}

// Post implements Service. Primarily useful in a client.
func (e Endpoints) Post(ctx context.Context, user *model.User) (status bool, errinfo string, data *model.User) {
	request := PostRequest{User: user}
	response, err := e.PostEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(PostResponse).Status, response.(PostResponse).Errinfo, response.(PostResponse).Data
}

// Delete implements Service. Primarily useful in a client.
func (e Endpoints) Delete(ctx context.Context, id string) (status bool, errinfo string, data *model.User) {
	request := DeleteRequest{Id: id}
	response, err := e.DeleteEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteResponse).Status, response.(DeleteResponse).Errinfo, response.(DeleteResponse).Data
}
