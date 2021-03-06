// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/apps/apps.proto

package go_micro_service_apps

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Apps service

type AppsService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
}

type appsService struct {
	c    client.Client
	name string
}

func NewAppsService(name string, c client.Client) AppsService {
	return &appsService{
		c:    c,
		name: name,
	}
}

func (c *appsService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Apps.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Apps.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Apps.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Apps.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Apps service

type AppsHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
}

func RegisterAppsHandler(s server.Server, hdlr AppsHandler, opts ...server.HandlerOption) error {
	type apps interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
	}
	type Apps struct {
		apps
	}
	h := &appsHandler{hdlr}
	return s.Handle(s.NewHandler(&Apps{h}, opts...))
}

type appsHandler struct {
	AppsHandler
}

func (h *appsHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AppsHandler.Create(ctx, in, out)
}

func (h *appsHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.AppsHandler.List(ctx, in, out)
}

func (h *appsHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AppsHandler.Read(ctx, in, out)
}

func (h *appsHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AppsHandler.Update(ctx, in, out)
}
