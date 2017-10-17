// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event_service.proto

/*
Package event_service is a generated protocol buffer package.

It is generated from these files:
	event_service.proto

It has these top-level messages:
	EventTarget
	EventSource
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

type EventTarget struct {
	TargetType string            `protobuf:"bytes,1,opt,name=target_type,json=targetType" json:"target_type,omitempty"`
	TargetId   string            `protobuf:"bytes,2,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	Params     map[string]string `protobuf:"bytes,3,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EventTarget) Reset()                    { *m = EventTarget{} }
func (m *EventTarget) String() string            { return proto.CompactTextString(m) }
func (*EventTarget) ProtoMessage()               {}
func (*EventTarget) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EventTarget) GetTargetType() string {
	if m != nil {
		return m.TargetType
	}
	return ""
}

func (m *EventTarget) GetTargetId() string {
	if m != nil {
		return m.TargetId
	}
	return ""
}

func (m *EventTarget) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type EventSource struct {
	SourceType string            `protobuf:"bytes,1,opt,name=source_type,json=sourceType" json:"source_type,omitempty"`
	SourceId   string            `protobuf:"bytes,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	Origin     string            `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	Params     map[string]string `protobuf:"bytes,4,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EventSource) Reset()                    { *m = EventSource{} }
func (m *EventSource) String() string            { return proto.CompactTextString(m) }
func (*EventSource) ProtoMessage()               {}
func (*EventSource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EventSource) GetSourceType() string {
	if m != nil {
		return m.SourceType
	}
	return ""
}

func (m *EventSource) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *EventSource) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *EventSource) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

type Event struct {
	Source    *EventSource      `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Target    *EventTarget      `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	EventName string            `protobuf:"bytes,3,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	UserId    string            `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	CreatedAt int32             `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Params    map[string]string `protobuf:"bytes,6,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Event) GetSource() *EventSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Event) GetTarget() *EventTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Event) GetEventName() string {
	if m != nil {
		return m.EventName
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
	RKey      string `protobuf:"bytes,2,opt,name=r_key,json=rKey" json:"r_key,omitempty"`
	Event     *Event `protobuf:"bytes,3,opt,name=event" json:"event,omitempty"`
}

func (m *EventRequest) Reset()                    { *m = EventRequest{} }
func (m *EventRequest) String() string            { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()               {}
func (*EventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EventRequest) GetUserToken() string {
	if m != nil {
		return m.UserToken
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
func (*EventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*EventTarget)(nil), "EventTarget")
	proto.RegisterType((*EventSource)(nil), "EventSource")
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
	Metadata: "event_service.proto",
}

func init() { proto.RegisterFile("event_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x25, 0xbb, 0x49, 0x68, 0x26, 0x2d, 0x42, 0x2e, 0x82, 0xa8, 0x80, 0x58, 0x45, 0x1c, 0x56,
	0x1c, 0x22, 0x14, 0x2e, 0x80, 0xe0, 0x80, 0x50, 0x0f, 0x2b, 0xc4, 0x87, 0xcc, 0xde, 0x83, 0xd9,
	0x8c, 0xaa, 0x28, 0x6c, 0x12, 0x6c, 0x67, 0xa5, 0xfc, 0x0c, 0xfe, 0x0d, 0x3f, 0x83, 0x9f, 0x84,
	0x3c, 0xe3, 0xb6, 0xe9, 0xbd, 0x37, 0xcf, 0x9b, 0xc9, 0x9b, 0xf7, 0x9e, 0x1d, 0x38, 0xc5, 0x03,
	0x76, 0xb6, 0x32, 0xa8, 0x0f, 0xcd, 0x0e, 0x8b, 0x41, 0xf7, 0xb6, 0xcf, 0xff, 0x06, 0x90, 0x9e,
	0x3b, 0x7c, 0xab, 0xf4, 0x05, 0x5a, 0xf1, 0x0c, 0x52, 0x4b, 0xa7, 0xca, 0x4e, 0x03, 0x66, 0xc1,
	0x2a, 0x58, 0x27, 0x12, 0x18, 0xda, 0x4e, 0x03, 0x8a, 0xc7, 0x90, 0xf8, 0x81, 0xa6, 0xce, 0x16,
	0xd4, 0x3e, 0x62, 0x60, 0x53, 0x8b, 0x97, 0x10, 0x0f, 0x4a, 0xab, 0xbd, 0xc9, 0x96, 0xab, 0xe5,
	0x3a, 0x2d, 0xb3, 0x62, 0xc6, 0x5d, 0x7c, 0xa3, 0xd6, 0x79, 0x67, 0xf5, 0x24, 0xfd, 0xdc, 0xd9,
	0x1b, 0x48, 0x67, 0xb0, 0xb8, 0x0f, 0xcb, 0x16, 0x27, 0xbf, 0xd6, 0x1d, 0xc5, 0x03, 0x88, 0x0e,
	0xea, 0xd7, 0x88, 0x7e, 0x17, 0x17, 0x6f, 0x17, 0xaf, 0x83, 0xfc, 0xdf, 0xa5, 0xf4, 0xef, 0xfd,
	0xa8, 0x77, 0xe8, 0xa4, 0x1b, 0x3a, 0xdd, 0x90, 0xce, 0xd0, 0xa5, 0x74, 0x3f, 0x70, 0x2d, 0x9d,
	0x81, 0x4d, 0x2d, 0x1e, 0x42, 0xdc, 0xeb, 0xe6, 0xa2, 0xe9, 0xb2, 0x25, 0x75, 0x7c, 0x35, 0xb3,
	0x14, 0xce, 0x2d, 0xf1, 0xce, 0xdb, 0xb6, 0xf4, 0x67, 0x01, 0x11, 0xd1, 0x8b, 0xe7, 0x10, 0xb3,
	0x34, 0xfa, 0x30, 0x2d, 0x8f, 0xe7, 0x6b, 0xa5, 0xef, 0xb9, 0x29, 0xce, 0x9e, 0xa8, 0xae, 0xa6,
	0x38, 0x6f, 0xe9, 0x7b, 0xe2, 0x29, 0x00, 0x5f, 0x7d, 0xa7, 0xf6, 0xe8, 0xed, 0x25, 0x84, 0x7c,
	0x51, 0x7b, 0x14, 0x8f, 0xe0, 0xee, 0x68, 0x50, 0xbb, 0x50, 0x42, 0xb6, 0xee, 0xca, 0x4d, 0xed,
	0xbe, 0xdb, 0x69, 0x54, 0x16, 0xeb, 0x4a, 0xd9, 0x2c, 0x5a, 0x05, 0xeb, 0x48, 0x26, 0x1e, 0xf9,
	0x60, 0xc5, 0x8b, 0xab, 0x64, 0x62, 0x4a, 0x46, 0xf0, 0xf2, 0xdb, 0xce, 0xe4, 0x07, 0x1c, 0x13,
	0xaf, 0xc4, 0xdf, 0x23, 0x1a, 0x72, 0x43, 0x72, 0x6d, 0xdf, 0x62, 0xe7, 0x29, 0x12, 0x87, 0x6c,
	0x1d, 0x20, 0x4e, 0x21, 0xd2, 0x95, 0x23, 0x67, 0xa2, 0x50, 0x7f, 0xc2, 0x49, 0x3c, 0x81, 0x88,
	0xfc, 0x92, 0xf9, 0xb4, 0x8c, 0x59, 0xa9, 0x64, 0x30, 0x7f, 0x07, 0x27, 0x7e, 0x83, 0x19, 0xfa,
	0xce, 0xa0, 0xe3, 0x68, 0x4c, 0xd5, 0xb7, 0xc4, 0x7e, 0x24, 0xc3, 0xc6, 0x7c, 0x6d, 0xdd, 0x03,
	0x31, 0x56, 0xd9, 0xd1, 0x78, 0x66, 0x5f, 0x95, 0xef, 0xfd, 0x2b, 0xfc, 0xac, 0x86, 0x01, 0xb5,
	0x28, 0x20, 0xfd, 0x48, 0x11, 0xf1, 0x3d, 0x9e, 0x14, 0x73, 0xf1, 0x67, 0xf7, 0x8a, 0x1b, 0x9b,
	0xf2, 0x3b, 0x3f, 0x63, 0xfa, 0x0f, 0x5f, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x97, 0x20, 0xf6,
	0x14, 0x9e, 0x03, 0x00, 0x00,
}
