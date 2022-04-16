package endpoint

import (
	"GolandProjects/base_study/etcd/demo1/service"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

type StringEndpoints struct {
	StringEndpoint      endpoint.Endpoint
	HealthCheckEndpoint endpoint.Endpoint
	SayEndpoint         endpoint.Endpoint
}

var (
	ErrInvalidRequestType = errors.New("RequestType has only one type: Concat,")
)

type StringRequest struct {
	RequestType string `json:"request_type"`
	A           string `json:"a"`
	B           string `json:"b"`
}

type StringResponse struct {
	Request string `json:"request"`
	Error   error  `json:"error"`
}

func MakeStringEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(StringRequest)

		var (
			res, a, b string
			opError   error
		)
		a = req.A
		b = req.B
		//根据请求操作执行不同方法
		if strings.EqualFold(req.RequestType, "Concat") {
			res, _ = svc.Concat(a, b)
		} else {
			return nil, ErrInvalidRequestType
		}
		return StringResponse{Request: res, Error: opError}, nil
	}
}

// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()

		return HealthResponse{status}, nil
	}
}

// 说话
type SayRequest struct {
	A string `json:"a"`
}

// 说话
type SayResponse struct {
	A string `json:"a"`
}

// 说话
func MakeSayEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SayRequest)

		var (
			res, a string
		)
		a = req.A
		res = svc.Say(a)
		fmt.Println("make say endpoint")

		return SayResponse{res}, nil
	}
}
