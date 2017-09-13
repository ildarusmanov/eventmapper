// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/event_service.proto

/*
Package event_service is a generated protocol buffer package.

It is generated from these files:
	pb/event_service.proto

It has these top-level messages:
	Event
	EventRequest
	EventResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Event struct {
	EventName   string            `protobuf:"bytes,1,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	EventTarget string            `protobuf:"bytes,2,opt,name=event_target,json=eventTarget" json:"event_target,omitempty"`
	UserId      string            `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	CreatedAt   int32             `protobuf:"varint,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Params      map[string]string `protobuf:"bytes,5,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Event) GetEventTarget() string {
	if m != nil {
		return m.EventTarget
	}
	return ""
}

func (m *Event) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Event) GetCreatedAt() int32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Event) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type EventRequest struct {
	UserToken string `protobuf:"bytes,1,opt,name=user_token,json=userToken" json:"user_token,omitempty"`
	Signature string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	RKey      string `protobuf:"bytes,3,opt,name=r_key,json=rKey" json:"r_key,omitempty"`
	Event     *Event `protobuf:"bytes,4,opt,name=event" json:"event,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EventRequest) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *EventRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *EventRequest) GetRKey() string {
	if m != nil {
		return m.RKey
	}
	return ""
}

func (m *EventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type EventResponse struct {
	IsOk   bool   `protobuf:"varint,1,opt,name=is_ok,json=isOk" json:"is_ok,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *EventResponse) Reset()                    { *m = EventResponse{} }
func (m *EventResponse) String() string            { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()               {}
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *EventResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*EventRequest)(nil), "EventRequest")
	proto.RegisterType((*EventResponse)(nil), "EventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventMapper service

type EventMapperClient interface {
	CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventMapperClient struct {
	cc *grpc.ClientConn
}

func NewEventMapperClient(cc *grpc.ClientConn) EventMapperClient {
	return &eventMapperClient{cc}
}

func (c *eventMapperClient) CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := grpc.Invoke(ctx, "/EventMapper/CreateEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventMapper service

type EventMapperServer interface {
	CreateEvent(context.Context, *EventRequest) (*EventResponse, error)
}

func RegisterEventMapperServer(s *grpc.Server, srv EventMapperServer) {
	s.RegisterService(&_EventMapper_serviceDesc, srv)
}

func _EventMapper_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventMapperServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EventMapper/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventMapperServer).CreateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventMapper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EventMapper",
	HandlerType: (*EventMapperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventMapper_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/event_service.proto",
}

func init() { proto.RegisterFile("pb/event_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc7, 0x9f, 0xb4, 0x4d, 0x1e, 0x33, 0x69, 0x45, 0x56, 0xa9, 0xa1, 0x54, 0xa8, 0x39, 0x15,
	0x0f, 0x11, 0xea, 0x45, 0x45, 0x0f, 0x22, 0x3d, 0x88, 0xf8, 0x42, 0xe8, 0x3d, 0x6c, 0xdb, 0xa1,
	0x84, 0xd8, 0x64, 0xdd, 0xdd, 0x14, 0x7a, 0xf1, 0xf3, 0xfa, 0x31, 0x64, 0x67, 0xb7, 0x58, 0x6f,
	0x99, 0xdf, 0x4e, 0xfe, 0x2f, 0xbb, 0xd0, 0x17, 0xf3, 0x4b, 0xdc, 0x60, 0xa5, 0x73, 0x85, 0x72,
	0x53, 0x2c, 0x30, 0x15, 0xb2, 0xd6, 0x75, 0xf2, 0xed, 0x81, 0x3f, 0x35, 0x9c, 0x9d, 0x01, 0xd8,
	0x85, 0x8a, 0xaf, 0x31, 0xf6, 0x46, 0xde, 0x38, 0xcc, 0x42, 0x22, 0xaf, 0x7c, 0x8d, 0xec, 0x1c,
	0xba, 0xf6, 0x58, 0x73, 0xb9, 0x42, 0x1d, 0xb7, 0x68, 0x21, 0x22, 0x36, 0x23, 0xc4, 0x4e, 0xe1,
	0x7f, 0xa3, 0x50, 0xe6, 0xc5, 0x32, 0x6e, 0xd3, 0x69, 0x60, 0xc6, 0xa7, 0xa5, 0x91, 0x5e, 0x48,
	0xe4, 0x1a, 0x97, 0x39, 0xd7, 0x71, 0x67, 0xe4, 0x8d, 0xfd, 0x2c, 0x74, 0xe4, 0x41, 0xb3, 0x0b,
	0x08, 0x04, 0x97, 0x7c, 0xad, 0x62, 0x7f, 0xd4, 0x1e, 0x47, 0x13, 0x96, 0x52, 0xa2, 0xf4, 0x9d,
	0xe0, 0xb4, 0xd2, 0x72, 0x9b, 0xb9, 0x8d, 0xc1, 0x0d, 0x44, 0x7b, 0x98, 0x1d, 0x41, 0xbb, 0xc4,
	0xad, 0x4b, 0x6b, 0x3e, 0xd9, 0x09, 0xf8, 0x1b, 0xfe, 0xd1, 0xa0, 0x0b, 0x68, 0x87, 0xdb, 0xd6,
	0xb5, 0x97, 0x7c, 0x41, 0x97, 0x74, 0x33, 0xfc, 0x6c, 0x50, 0x51, 0x61, 0x8a, 0xab, 0xeb, 0x12,
	0xab, 0x5d, 0x61, 0x43, 0x66, 0x06, 0xb0, 0x21, 0x84, 0xaa, 0x58, 0x55, 0x5c, 0x37, 0x72, 0x27,
	0xf6, 0x0b, 0xd8, 0x31, 0xf8, 0x32, 0x37, 0xd6, 0xb6, 0x69, 0x47, 0x3e, 0xe3, 0x96, 0x0d, 0xc1,
	0xa7, 0xfb, 0xa0, 0x8a, 0xd1, 0x24, 0xb0, 0x3d, 0x32, 0x0b, 0x93, 0x3b, 0xe8, 0x39, 0x7f, 0x25,
	0xea, 0x4a, 0x91, 0x46, 0xa1, 0xf2, 0xba, 0x24, 0xef, 0x83, 0xac, 0x53, 0xa8, 0xb7, 0x92, 0xf5,
	0x21, 0x50, 0x9a, 0xeb, 0x46, 0x39, 0x4f, 0x37, 0x4d, 0xee, 0x21, 0xa2, 0xbf, 0x5f, 0xb8, 0x10,
	0x28, 0x59, 0x0a, 0xd1, 0x23, 0x5d, 0xa0, 0x7d, 0xbc, 0x5e, 0xba, 0x5f, 0x6d, 0x70, 0x98, 0xfe,
	0x71, 0x4a, 0xfe, 0xcd, 0x03, 0x7a, 0xee, 0xab, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x27,
	0x6f, 0x94, 0x08, 0x02, 0x00, 0x00,
}
