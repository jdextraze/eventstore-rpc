// Code generated by protoc-gen-go.
// source: event_store.proto
// DO NOT EDIT!

/*
Package eventstore is a generated protocol buffer package.

It is generated from these files:
	event_store.proto

It has these top-level messages:
	AppendToStreamRequest
	AppendToStreamResponse
	SubscribeToStreamFromRequest
	SubscribeToStreamFromResponse
	EventData
	UserCredentials
	Position
	Error
	RecordedEvent
	ResolvedEvent
*/
package eventstore

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

type SubscribeToStreamFromResponse_DropReason int32

const (
	SubscribeToStreamFromResponse_UserInitiated                 SubscribeToStreamFromResponse_DropReason = 0
	SubscribeToStreamFromResponse_NotAuthenticated              SubscribeToStreamFromResponse_DropReason = 1
	SubscribeToStreamFromResponse_AccessDenied                  SubscribeToStreamFromResponse_DropReason = 2
	SubscribeToStreamFromResponse_SubscribingError              SubscribeToStreamFromResponse_DropReason = 3
	SubscribeToStreamFromResponse_ServerError                   SubscribeToStreamFromResponse_DropReason = 4
	SubscribeToStreamFromResponse_ConnectionClosed              SubscribeToStreamFromResponse_DropReason = 5
	SubscribeToStreamFromResponse_CatchUpError                  SubscribeToStreamFromResponse_DropReason = 6
	SubscribeToStreamFromResponse_ProcessingQueueOverflow       SubscribeToStreamFromResponse_DropReason = 7
	SubscribeToStreamFromResponse_EventHandlerException         SubscribeToStreamFromResponse_DropReason = 8
	SubscribeToStreamFromResponse_MaxSubscribersReached         SubscribeToStreamFromResponse_DropReason = 9
	SubscribeToStreamFromResponse_PersistentSubscriptionDeleted SubscribeToStreamFromResponse_DropReason = 10
	SubscribeToStreamFromResponse_Unknown                       SubscribeToStreamFromResponse_DropReason = 100
	SubscribeToStreamFromResponse_NotFound                      SubscribeToStreamFromResponse_DropReason = 101
)

var SubscribeToStreamFromResponse_DropReason_name = map[int32]string{
	0:   "UserInitiated",
	1:   "NotAuthenticated",
	2:   "AccessDenied",
	3:   "SubscribingError",
	4:   "ServerError",
	5:   "ConnectionClosed",
	6:   "CatchUpError",
	7:   "ProcessingQueueOverflow",
	8:   "EventHandlerException",
	9:   "MaxSubscribersReached",
	10:  "PersistentSubscriptionDeleted",
	100: "Unknown",
	101: "NotFound",
}
var SubscribeToStreamFromResponse_DropReason_value = map[string]int32{
	"UserInitiated":                 0,
	"NotAuthenticated":              1,
	"AccessDenied":                  2,
	"SubscribingError":              3,
	"ServerError":                   4,
	"ConnectionClosed":              5,
	"CatchUpError":                  6,
	"ProcessingQueueOverflow":       7,
	"EventHandlerException":         8,
	"MaxSubscribersReached":         9,
	"PersistentSubscriptionDeleted": 10,
	"Unknown":                       100,
	"NotFound":                      101,
}

func (x SubscribeToStreamFromResponse_DropReason) String() string {
	return proto.EnumName(SubscribeToStreamFromResponse_DropReason_name, int32(x))
}
func (SubscribeToStreamFromResponse_DropReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

type AppendToStreamRequest struct {
	StreamId        string           `protobuf:"bytes,1,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	ExpectedVersion int32            `protobuf:"varint,2,opt,name=expected_version,json=expectedVersion" json:"expected_version,omitempty"`
	Events          []*EventData     `protobuf:"bytes,3,rep,name=events" json:"events,omitempty"`
	UserCredentials *UserCredentials `protobuf:"bytes,4,opt,name=user_credentials,json=userCredentials" json:"user_credentials,omitempty"`
}

func (m *AppendToStreamRequest) Reset()                    { *m = AppendToStreamRequest{} }
func (m *AppendToStreamRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendToStreamRequest) ProtoMessage()               {}
func (*AppendToStreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AppendToStreamRequest) GetStreamId() string {
	if m != nil {
		return m.StreamId
	}
	return ""
}

func (m *AppendToStreamRequest) GetExpectedVersion() int32 {
	if m != nil {
		return m.ExpectedVersion
	}
	return 0
}

func (m *AppendToStreamRequest) GetEvents() []*EventData {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *AppendToStreamRequest) GetUserCredentials() *UserCredentials {
	if m != nil {
		return m.UserCredentials
	}
	return nil
}

type AppendToStreamResponse struct {
	NextExpectedVersion int32     `protobuf:"varint,1,opt,name=next_expected_version,json=nextExpectedVersion" json:"next_expected_version,omitempty"`
	Position            *Position `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	Error               *Error    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *AppendToStreamResponse) Reset()                    { *m = AppendToStreamResponse{} }
func (m *AppendToStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendToStreamResponse) ProtoMessage()               {}
func (*AppendToStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppendToStreamResponse) GetNextExpectedVersion() int32 {
	if m != nil {
		return m.NextExpectedVersion
	}
	return 0
}

func (m *AppendToStreamResponse) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *AppendToStreamResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type SubscribeToStreamFromRequest struct {
	StreamId        string           `protobuf:"bytes,1,opt,name=stream_id,json=streamId" json:"stream_id,omitempty"`
	LastCheckpoint  int32            `protobuf:"varint,2,opt,name=last_checkpoint,json=lastCheckpoint" json:"last_checkpoint,omitempty"`
	UserCredentials *UserCredentials `protobuf:"bytes,3,opt,name=user_credentials,json=userCredentials" json:"user_credentials,omitempty"`
}

func (m *SubscribeToStreamFromRequest) Reset()                    { *m = SubscribeToStreamFromRequest{} }
func (m *SubscribeToStreamFromRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeToStreamFromRequest) ProtoMessage()               {}
func (*SubscribeToStreamFromRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubscribeToStreamFromRequest) GetStreamId() string {
	if m != nil {
		return m.StreamId
	}
	return ""
}

func (m *SubscribeToStreamFromRequest) GetLastCheckpoint() int32 {
	if m != nil {
		return m.LastCheckpoint
	}
	return 0
}

func (m *SubscribeToStreamFromRequest) GetUserCredentials() *UserCredentials {
	if m != nil {
		return m.UserCredentials
	}
	return nil
}

type SubscribeToStreamFromResponse struct {
	Event      *ResolvedEvent                           `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	DropReason SubscribeToStreamFromResponse_DropReason `protobuf:"varint,2,opt,name=drop_reason,json=dropReason,enum=eventstore.SubscribeToStreamFromResponse_DropReason" json:"drop_reason,omitempty"`
	Error      *Error                                   `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *SubscribeToStreamFromResponse) Reset()                    { *m = SubscribeToStreamFromResponse{} }
func (m *SubscribeToStreamFromResponse) String() string            { return proto.CompactTextString(m) }
func (*SubscribeToStreamFromResponse) ProtoMessage()               {}
func (*SubscribeToStreamFromResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SubscribeToStreamFromResponse) GetEvent() *ResolvedEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SubscribeToStreamFromResponse) GetDropReason() SubscribeToStreamFromResponse_DropReason {
	if m != nil {
		return m.DropReason
	}
	return SubscribeToStreamFromResponse_UserInitiated
}

func (m *SubscribeToStreamFromResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type EventData struct {
	EventId   []byte `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType string `protobuf:"bytes,2,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	IsJson    bool   `protobuf:"varint,3,opt,name=is_json,json=isJson" json:"is_json,omitempty"`
	Data      []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Metadata  []byte `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *EventData) Reset()                    { *m = EventData{} }
func (m *EventData) String() string            { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()               {}
func (*EventData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EventData) GetEventId() []byte {
	if m != nil {
		return m.EventId
	}
	return nil
}

func (m *EventData) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventData) GetIsJson() bool {
	if m != nil {
		return m.IsJson
	}
	return false
}

func (m *EventData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *EventData) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type UserCredentials struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *UserCredentials) Reset()                    { *m = UserCredentials{} }
func (m *UserCredentials) String() string            { return proto.CompactTextString(m) }
func (*UserCredentials) ProtoMessage()               {}
func (*UserCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserCredentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserCredentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Position struct {
	CommitPosition  int64 `protobuf:"varint,1,opt,name=commit_position,json=commitPosition" json:"commit_position,omitempty"`
	PreparePosition int64 `protobuf:"varint,2,opt,name=prepare_position,json=preparePosition" json:"prepare_position,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Position) GetCommitPosition() int64 {
	if m != nil {
		return m.CommitPosition
	}
	return 0
}

func (m *Position) GetPreparePosition() int64 {
	if m != nil {
		return m.PreparePosition
	}
	return 0
}

type Error struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Error) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Error) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type RecordedEvent struct {
	EventStreamId string `protobuf:"bytes,1,opt,name=event_stream_id,json=eventStreamId" json:"event_stream_id,omitempty"`
	EventId       []byte `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventNumber   int32  `protobuf:"varint,3,opt,name=event_number,json=eventNumber" json:"event_number,omitempty"`
	EventType     string `protobuf:"bytes,4,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Data          []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Metadata      []byte `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	IsJson        bool   `protobuf:"varint,7,opt,name=is_json,json=isJson" json:"is_json,omitempty"`
	Created       int64  `protobuf:"varint,8,opt,name=created" json:"created,omitempty"`
	CreatedEpoch  int64  `protobuf:"varint,9,opt,name=created_epoch,json=createdEpoch" json:"created_epoch,omitempty"`
}

func (m *RecordedEvent) Reset()                    { *m = RecordedEvent{} }
func (m *RecordedEvent) String() string            { return proto.CompactTextString(m) }
func (*RecordedEvent) ProtoMessage()               {}
func (*RecordedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RecordedEvent) GetEventStreamId() string {
	if m != nil {
		return m.EventStreamId
	}
	return ""
}

func (m *RecordedEvent) GetEventId() []byte {
	if m != nil {
		return m.EventId
	}
	return nil
}

func (m *RecordedEvent) GetEventNumber() int32 {
	if m != nil {
		return m.EventNumber
	}
	return 0
}

func (m *RecordedEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *RecordedEvent) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RecordedEvent) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RecordedEvent) GetIsJson() bool {
	if m != nil {
		return m.IsJson
	}
	return false
}

func (m *RecordedEvent) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RecordedEvent) GetCreatedEpoch() int64 {
	if m != nil {
		return m.CreatedEpoch
	}
	return 0
}

type ResolvedEvent struct {
	Event    *RecordedEvent `protobuf:"bytes,1,opt,name=event" json:"event,omitempty"`
	Position *Position      `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
}

func (m *ResolvedEvent) Reset()                    { *m = ResolvedEvent{} }
func (m *ResolvedEvent) String() string            { return proto.CompactTextString(m) }
func (*ResolvedEvent) ProtoMessage()               {}
func (*ResolvedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ResolvedEvent) GetEvent() *RecordedEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ResolvedEvent) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func init() {
	proto.RegisterType((*AppendToStreamRequest)(nil), "eventstore.AppendToStreamRequest")
	proto.RegisterType((*AppendToStreamResponse)(nil), "eventstore.AppendToStreamResponse")
	proto.RegisterType((*SubscribeToStreamFromRequest)(nil), "eventstore.SubscribeToStreamFromRequest")
	proto.RegisterType((*SubscribeToStreamFromResponse)(nil), "eventstore.SubscribeToStreamFromResponse")
	proto.RegisterType((*EventData)(nil), "eventstore.EventData")
	proto.RegisterType((*UserCredentials)(nil), "eventstore.UserCredentials")
	proto.RegisterType((*Position)(nil), "eventstore.Position")
	proto.RegisterType((*Error)(nil), "eventstore.Error")
	proto.RegisterType((*RecordedEvent)(nil), "eventstore.RecordedEvent")
	proto.RegisterType((*ResolvedEvent)(nil), "eventstore.ResolvedEvent")
	proto.RegisterEnum("eventstore.SubscribeToStreamFromResponse_DropReason", SubscribeToStreamFromResponse_DropReason_name, SubscribeToStreamFromResponse_DropReason_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventStore service

type EventStoreClient interface {
	AppendToStream(ctx context.Context, in *AppendToStreamRequest, opts ...grpc.CallOption) (*AppendToStreamResponse, error)
	SubscribeToStreamFrom(ctx context.Context, opts ...grpc.CallOption) (EventStore_SubscribeToStreamFromClient, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) AppendToStream(ctx context.Context, in *AppendToStreamRequest, opts ...grpc.CallOption) (*AppendToStreamResponse, error) {
	out := new(AppendToStreamResponse)
	err := grpc.Invoke(ctx, "/eventstore.EventStore/AppendToStream", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventStoreClient) SubscribeToStreamFrom(ctx context.Context, opts ...grpc.CallOption) (EventStore_SubscribeToStreamFromClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventStore_serviceDesc.Streams[0], c.cc, "/eventstore.EventStore/SubscribeToStreamFrom", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStoreSubscribeToStreamFromClient{stream}
	return x, nil
}

type EventStore_SubscribeToStreamFromClient interface {
	Send(*SubscribeToStreamFromRequest) error
	Recv() (*SubscribeToStreamFromResponse, error)
	grpc.ClientStream
}

type eventStoreSubscribeToStreamFromClient struct {
	grpc.ClientStream
}

func (x *eventStoreSubscribeToStreamFromClient) Send(m *SubscribeToStreamFromRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventStoreSubscribeToStreamFromClient) Recv() (*SubscribeToStreamFromResponse, error) {
	m := new(SubscribeToStreamFromResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EventStore service

type EventStoreServer interface {
	AppendToStream(context.Context, *AppendToStreamRequest) (*AppendToStreamResponse, error)
	SubscribeToStreamFrom(EventStore_SubscribeToStreamFromServer) error
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_AppendToStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendToStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).AppendToStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventstore.EventStore/AppendToStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).AppendToStream(ctx, req.(*AppendToStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventStore_SubscribeToStreamFrom_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventStoreServer).SubscribeToStreamFrom(&eventStoreSubscribeToStreamFromServer{stream})
}

type EventStore_SubscribeToStreamFromServer interface {
	Send(*SubscribeToStreamFromResponse) error
	Recv() (*SubscribeToStreamFromRequest, error)
	grpc.ServerStream
}

type eventStoreSubscribeToStreamFromServer struct {
	grpc.ServerStream
}

func (x *eventStoreSubscribeToStreamFromServer) Send(m *SubscribeToStreamFromResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventStoreSubscribeToStreamFromServer) Recv() (*SubscribeToStreamFromRequest, error) {
	m := new(SubscribeToStreamFromRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventstore.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppendToStream",
			Handler:    _EventStore_AppendToStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToStreamFrom",
			Handler:       _EventStore_SubscribeToStreamFrom_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "event_store.proto",
}

func init() { proto.RegisterFile("event_store.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x72, 0x23, 0x35,
	0x10, 0xce, 0xd8, 0xf1, 0x5f, 0xdb, 0xb1, 0x27, 0x62, 0xcd, 0x3a, 0x1b, 0x52, 0x95, 0x0c, 0x55,
	0xac, 0x73, 0xc0, 0xbb, 0x65, 0xb8, 0x70, 0xcc, 0x3a, 0x0e, 0x84, 0x2a, 0x82, 0x91, 0x37, 0x54,
	0x71, 0x61, 0x6a, 0x32, 0x6a, 0xd6, 0xc3, 0xda, 0xd2, 0x20, 0xc9, 0x8e, 0xf7, 0x25, 0x78, 0x02,
	0x4e, 0x5c, 0xb6, 0x8a, 0xe2, 0x79, 0xb8, 0xf0, 0x26, 0x9c, 0x28, 0x69, 0x66, 0xec, 0xb1, 0xc9,
	0x86, 0xc0, 0x4d, 0xfa, 0xba, 0xd5, 0xea, 0xfe, 0x3e, 0x75, 0xcf, 0xc0, 0x3e, 0x2e, 0x90, 0x6b,
	0x5f, 0x69, 0x21, 0xb1, 0x17, 0x4b, 0xa1, 0x05, 0x01, 0x0b, 0x59, 0xc4, 0xfb, 0xc3, 0x81, 0xf6,
	0x59, 0x1c, 0x23, 0x67, 0x2f, 0xc5, 0x58, 0x4b, 0x0c, 0x66, 0x14, 0x7f, 0x9a, 0xa3, 0xd2, 0xe4,
	0x10, 0x6a, 0xca, 0x02, 0x7e, 0xc4, 0x3a, 0xce, 0xb1, 0xd3, 0xad, 0xd1, 0x6a, 0x02, 0x5c, 0x32,
	0x72, 0x0a, 0x2e, 0x2e, 0x63, 0x0c, 0x35, 0x32, 0x7f, 0x81, 0x52, 0x45, 0x82, 0x77, 0x0a, 0xc7,
	0x4e, 0xb7, 0x44, 0x5b, 0x19, 0xfe, 0x6d, 0x02, 0x93, 0x8f, 0xa1, 0x9c, 0xdc, 0xd7, 0x29, 0x1e,
	0x17, 0xbb, 0xf5, 0x7e, 0xbb, 0xb7, 0xbe, 0xbe, 0x37, 0x34, 0xcb, 0xf3, 0x40, 0x07, 0x34, 0x75,
	0x22, 0x17, 0xe0, 0xce, 0x15, 0x4a, 0x3f, 0x94, 0xc8, 0x90, 0xeb, 0x28, 0x98, 0xaa, 0xce, 0xee,
	0xb1, 0xd3, 0xad, 0xf7, 0x0f, 0xf3, 0x07, 0xaf, 0x15, 0xca, 0xc1, 0xda, 0x85, 0xb6, 0xe6, 0x9b,
	0x80, 0xf7, 0xd6, 0x81, 0xf7, 0xb7, 0x0b, 0x53, 0xb1, 0xe0, 0x0a, 0x49, 0x1f, 0xda, 0x1c, 0x97,
	0xda, 0xff, 0x47, 0x05, 0x8e, 0xad, 0xe0, 0x3d, 0x63, 0x1c, 0x6e, 0x55, 0xf1, 0x1c, 0xaa, 0xb1,
	0x50, 0x91, 0xce, 0x0a, 0xad, 0xf7, 0x1f, 0xe5, 0xd3, 0x19, 0xa5, 0x36, 0xba, 0xf2, 0x22, 0x4f,
	0xa1, 0x84, 0x52, 0x0a, 0xd9, 0x29, 0x5a, 0xf7, 0xfd, 0x8d, 0xb2, 0x8d, 0x81, 0x26, 0x76, 0xef,
	0x77, 0x07, 0x3e, 0x18, 0xcf, 0x6f, 0x54, 0x28, 0xa3, 0x1b, 0xcc, 0x92, 0xbd, 0x90, 0xe2, 0x61,
	0x4a, 0x3c, 0x85, 0xd6, 0x34, 0x50, 0xda, 0x0f, 0x27, 0x18, 0xbe, 0x8e, 0x45, 0xc4, 0x75, 0x2a,
	0x44, 0xd3, 0xc0, 0x83, 0x15, 0x7a, 0x27, 0xb1, 0xc5, 0xff, 0x41, 0xec, 0x5f, 0x45, 0x38, 0x7a,
	0x47, 0xba, 0x29, 0xbf, 0xcf, 0xa0, 0x64, 0x03, 0xda, 0x5c, 0xeb, 0xfd, 0x83, 0x7c, 0x78, 0x8a,
	0x4a, 0x4c, 0x17, 0xc8, 0xac, 0xf0, 0x34, 0xf1, 0x23, 0xd7, 0x50, 0x67, 0x52, 0xc4, 0xbe, 0xc4,
	0x40, 0xa5, 0xfc, 0x36, 0xfb, 0x9f, 0xe6, 0x8f, 0xdd, 0x7b, 0x61, 0xef, 0x5c, 0x8a, 0x98, 0xda,
	0xb3, 0x14, 0xd8, 0x6a, 0xfd, 0x70, 0x05, 0xde, 0x16, 0x00, 0xd6, 0x31, 0xc8, 0x3e, 0xec, 0x19,
	0x16, 0x2e, 0x79, 0xa4, 0xa3, 0x40, 0x23, 0x73, 0x77, 0xc8, 0x23, 0x70, 0xaf, 0x84, 0x3e, 0x9b,
	0xeb, 0x89, 0xa1, 0x21, 0xb4, 0xa8, 0x43, 0x5c, 0x68, 0x9c, 0x85, 0x21, 0x2a, 0x75, 0x8e, 0x3c,
	0x42, 0xe6, 0x16, 0x8c, 0x5f, 0x96, 0x6a, 0xc4, 0x5f, 0xd9, 0x4b, 0xdc, 0x22, 0x69, 0x41, 0x7d,
	0x8c, 0x72, 0x81, 0x32, 0x01, 0x76, 0x8d, 0xdb, 0x40, 0x70, 0x8e, 0xa1, 0x79, 0x29, 0x83, 0xa9,
	0x50, 0xc8, 0xdc, 0x92, 0x09, 0x37, 0x08, 0x74, 0x38, 0xb9, 0x8e, 0x13, 0xbf, 0x32, 0x39, 0x84,
	0xc7, 0x23, 0x29, 0xcc, 0x0d, 0x11, 0x7f, 0xf5, 0xcd, 0x1c, 0xe7, 0xf8, 0xf5, 0x02, 0xe5, 0x0f,
	0x53, 0x71, 0xeb, 0x56, 0xc8, 0x01, 0xb4, 0x2d, 0x8b, 0x5f, 0x04, 0x9c, 0x4d, 0x51, 0x0e, 0x97,
	0x21, 0xc6, 0x26, 0x9e, 0x5b, 0x35, 0xa6, 0xaf, 0x82, 0xe5, 0x8a, 0x34, 0xa9, 0x28, 0x06, 0xe1,
	0x04, 0x99, 0x5b, 0x23, 0x27, 0x70, 0x34, 0x32, 0x6f, 0x5a, 0x69, 0xe4, 0x3a, 0xf5, 0xb0, 0xc7,
	0xce, 0x71, 0x8a, 0xa6, 0x2c, 0x20, 0x75, 0xa8, 0x5c, 0xf3, 0xd7, 0x5c, 0xdc, 0x72, 0x97, 0x91,
	0x06, 0x54, 0xaf, 0x84, 0xbe, 0x10, 0x73, 0xce, 0x5c, 0xf4, 0x7e, 0x76, 0xa0, 0xb6, 0xea, 0x59,
	0x72, 0x00, 0xd5, 0x64, 0xba, 0xa4, 0xef, 0xb2, 0x41, 0x2b, 0x76, 0x7f, 0xc9, 0xc8, 0x11, 0x24,
	0x53, 0xc6, 0xd7, 0x6f, 0x62, 0xb4, 0x8a, 0xd6, 0x68, 0xcd, 0x22, 0x2f, 0xdf, 0xc4, 0x48, 0x1e,
	0x43, 0x25, 0x52, 0xfe, 0x8f, 0x46, 0x6d, 0x23, 0x4e, 0x95, 0x96, 0x23, 0xf5, 0xa5, 0xe1, 0x9e,
	0xc0, 0x2e, 0x0b, 0x74, 0x60, 0x5b, 0xbe, 0x41, 0xed, 0x9a, 0x3c, 0x81, 0xea, 0x0c, 0x75, 0x60,
	0xf1, 0x92, 0xc5, 0x57, 0x7b, 0xef, 0x12, 0x5a, 0x5b, 0x2f, 0xd6, 0xb8, 0x9b, 0x37, 0xcb, 0x83,
	0x19, 0x66, 0xdd, 0x92, 0xed, 0x8d, 0x2d, 0x0e, 0x94, 0xba, 0x15, 0x92, 0xa5, 0x49, 0xad, 0xf6,
	0xde, 0xf7, 0x50, 0x1d, 0xad, 0x9b, 0xb7, 0x15, 0x8a, 0xd9, 0x2c, 0xd2, 0xfe, 0xaa, 0xeb, 0x4d,
	0xa8, 0x22, 0x6d, 0x26, 0xf0, 0xca, 0xf1, 0x14, 0xdc, 0x58, 0x62, 0x1c, 0x48, 0xf4, 0x37, 0xe6,
	0x43, 0x91, 0xb6, 0x52, 0x3c, 0x73, 0xf5, 0x9e, 0x41, 0xc9, 0xea, 0x6a, 0x6a, 0xb4, 0xac, 0x24,
	0xc9, 0xd9, 0xb5, 0xc5, 0x70, 0xa9, 0xd3, 0xa4, 0xec, 0xda, 0xfb, 0xa5, 0x00, 0x7b, 0x14, 0x43,
	0x21, 0x59, 0xda, 0x2f, 0xe4, 0x23, 0x68, 0x65, 0xe3, 0x7c, 0x73, 0x1e, 0xec, 0x59, 0x78, 0x9c,
	0x0d, 0x85, 0xbc, 0x30, 0x85, 0x4d, 0x61, 0x4e, 0xa0, 0x91, 0x98, 0xf8, 0x7c, 0x76, 0x83, 0x49,
	0x6f, 0x94, 0x68, 0xdd, 0x62, 0x57, 0x16, 0xda, 0xd2, 0x6e, 0x77, 0x5b, 0xbb, 0x4c, 0xa2, 0xd2,
	0x3b, 0x24, 0x2a, 0x6f, 0x4a, 0x94, 0xd7, 0xba, 0xb2, 0xa1, 0x75, 0x07, 0x2a, 0xa1, 0x44, 0xd3,
	0x4b, 0x9d, 0xaa, 0xa5, 0x2c, 0xdb, 0x92, 0x0f, 0x61, 0x2f, 0x5d, 0xfa, 0x18, 0x8b, 0x70, 0xd2,
	0xa9, 0x59, 0x7b, 0x23, 0x05, 0x87, 0x06, 0xf3, 0xa4, 0x61, 0x27, 0x37, 0x4d, 0xfe, 0x65, 0xee,
	0xe4, 0x78, 0xcc, 0xe6, 0xce, 0x7f, 0x1e, 0xea, 0xfd, 0x3f, 0x1d, 0x80, 0x61, 0x42, 0xb5, 0x90,
	0x48, 0xbe, 0x83, 0xe6, 0xe6, 0x37, 0x86, 0x9c, 0xe4, 0x03, 0xdc, 0xf9, 0x61, 0x7d, 0xe2, 0xdd,
	0xe7, 0x92, 0x4c, 0x34, 0x6f, 0x87, 0x48, 0x68, 0xdf, 0x39, 0xf4, 0x48, 0xf7, 0x01, 0x73, 0x31,
	0xb9, 0xe8, 0xf4, 0xc1, 0x13, 0xd4, 0xdb, 0xe9, 0x3a, 0xcf, 0x9d, 0x17, 0x9f, 0x01, 0x09, 0xc5,
	0x2c, 0x7f, 0x4a, 0xc6, 0xe1, 0x8b, 0xfa, 0xe7, 0x74, 0x34, 0x18, 0x8e, 0x47, 0xe6, 0xdf, 0x61,
	0xe4, 0xfc, 0x5a, 0x28, 0x0c, 0xc7, 0xbf, 0x15, 0x9a, 0x6b, 0x16, 0x7a, 0x74, 0x34, 0xb8, 0x29,
	0xdb, 0x5f, 0x8b, 0x4f, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x86, 0xae, 0x65, 0x3e, 0x6f, 0x08,
	0x00, 0x00,
}
