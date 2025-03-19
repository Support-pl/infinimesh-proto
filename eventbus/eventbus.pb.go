// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: eventbus/eventbus.proto

package eventbus

import (
	node "github.com/support-pl/infinimesh-proto/node"
	accounts "github.com/support-pl/infinimesh-proto/node/accounts"
	devices "github.com/support-pl/infinimesh-proto/node/devices"
	namespaces "github.com/support-pl/infinimesh-proto/node/namespaces"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventKind int32

const (
	EventKind_NONE             EventKind = 0
	EventKind_ACCOUNT_CREATE   EventKind = 1
	EventKind_ACCOUNT_UPDATE   EventKind = 2
	EventKind_ACCOUNT_DELETE   EventKind = 3
	EventKind_ACCOUNT_MOVE     EventKind = 4
	EventKind_NAMESPACE_CREATE EventKind = 5
	EventKind_NAMESPACE_UPDATE EventKind = 6
	EventKind_NAMESPACE_DELETE EventKind = 7
	EventKind_DEVICE_CREATE    EventKind = 8
	EventKind_DEVICE_UPDATE    EventKind = 9
	EventKind_DEVICE_DELETE    EventKind = 10
	EventKind_DEVICE_MOVE      EventKind = 11
)

// Enum value maps for EventKind.
var (
	EventKind_name = map[int32]string{
		0:  "NONE",
		1:  "ACCOUNT_CREATE",
		2:  "ACCOUNT_UPDATE",
		3:  "ACCOUNT_DELETE",
		4:  "ACCOUNT_MOVE",
		5:  "NAMESPACE_CREATE",
		6:  "NAMESPACE_UPDATE",
		7:  "NAMESPACE_DELETE",
		8:  "DEVICE_CREATE",
		9:  "DEVICE_UPDATE",
		10: "DEVICE_DELETE",
		11: "DEVICE_MOVE",
	}
	EventKind_value = map[string]int32{
		"NONE":             0,
		"ACCOUNT_CREATE":   1,
		"ACCOUNT_UPDATE":   2,
		"ACCOUNT_DELETE":   3,
		"ACCOUNT_MOVE":     4,
		"NAMESPACE_CREATE": 5,
		"NAMESPACE_UPDATE": 6,
		"NAMESPACE_DELETE": 7,
		"DEVICE_CREATE":    8,
		"DEVICE_UPDATE":    9,
		"DEVICE_DELETE":    10,
		"DEVICE_MOVE":      11,
	}
)

func (x EventKind) Enum() *EventKind {
	p := new(EventKind)
	*p = x
	return p
}

func (x EventKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventKind) Descriptor() protoreflect.EnumDescriptor {
	return file_eventbus_eventbus_proto_enumTypes[0].Descriptor()
}

func (EventKind) Type() protoreflect.EnumType {
	return &file_eventbus_eventbus_proto_enumTypes[0]
}

func (x EventKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventKind.Descriptor instead.
func (EventKind) EnumDescriptor() ([]byte, []int) {
	return file_eventbus_eventbus_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	EventKind EventKind              `protobuf:"varint,1,opt,name=event_kind,json=eventKind,proto3,enum=infinimesh.eventbus.EventKind" json:"event_kind,omitempty"`
	// Types that are valid to be assigned to Entity:
	//
	//	*Event_Account
	//	*Event_Namespace
	//	*Event_Device
	Entity        isEvent_Entity   `protobuf_oneof:"entity"`
	Meta          *structpb.Struct `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Event) Reset() {
	*x = Event{}
	mi := &file_eventbus_eventbus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_eventbus_eventbus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_eventbus_eventbus_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventKind() EventKind {
	if x != nil {
		return x.EventKind
	}
	return EventKind_NONE
}

func (x *Event) GetEntity() isEvent_Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Event) GetAccount() *accounts.Account {
	if x != nil {
		if x, ok := x.Entity.(*Event_Account); ok {
			return x.Account
		}
	}
	return nil
}

func (x *Event) GetNamespace() *namespaces.Namespace {
	if x != nil {
		if x, ok := x.Entity.(*Event_Namespace); ok {
			return x.Namespace
		}
	}
	return nil
}

func (x *Event) GetDevice() *devices.Device {
	if x != nil {
		if x, ok := x.Entity.(*Event_Device); ok {
			return x.Device
		}
	}
	return nil
}

func (x *Event) GetMeta() *structpb.Struct {
	if x != nil {
		return x.Meta
	}
	return nil
}

type isEvent_Entity interface {
	isEvent_Entity()
}

type Event_Account struct {
	Account *accounts.Account `protobuf:"bytes,2,opt,name=account,proto3,oneof"`
}

type Event_Namespace struct {
	Namespace *namespaces.Namespace `protobuf:"bytes,3,opt,name=namespace,proto3,oneof"`
}

type Event_Device struct {
	Device *devices.Device `protobuf:"bytes,4,opt,name=device,proto3,oneof"`
}

func (*Event_Account) isEvent_Entity() {}

func (*Event_Namespace) isEvent_Entity() {}

func (*Event_Device) isEvent_Entity() {}

var File_eventbus_eventbus_proto protoreflect.FileDescriptor

var file_eventbus_eventbus_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x1a, 0x1c,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x3d, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x45, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x42, 0x08,
	0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2a, 0xef, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x04, 0x12, 0x14,
	0x0a, 0x10, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43,
	0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x41,
	0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x07,
	0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x56,
	0x49, 0x43, 0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x0b, 0x32, 0x59, 0x0a, 0x0d, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x30, 0x01, 0x42, 0xc6, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75,
	0x73, 0x42, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x70, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x62, 0x75, 0x73, 0xa2, 0x02, 0x03, 0x49, 0x45, 0x58, 0xaa, 0x02, 0x13, 0x49, 0x6e, 0x66, 0x69,
	0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0xca,
	0x02, 0x13, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x62, 0x75, 0x73, 0xe2, 0x02, 0x1f, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_eventbus_eventbus_proto_rawDescOnce sync.Once
	file_eventbus_eventbus_proto_rawDescData []byte
)

func file_eventbus_eventbus_proto_rawDescGZIP() []byte {
	file_eventbus_eventbus_proto_rawDescOnce.Do(func() {
		file_eventbus_eventbus_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_eventbus_eventbus_proto_rawDesc), len(file_eventbus_eventbus_proto_rawDesc)))
	})
	return file_eventbus_eventbus_proto_rawDescData
}

var file_eventbus_eventbus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eventbus_eventbus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eventbus_eventbus_proto_goTypes = []any{
	(EventKind)(0),               // 0: infinimesh.eventbus.EventKind
	(*Event)(nil),                // 1: infinimesh.eventbus.Event
	(*accounts.Account)(nil),     // 2: infinimesh.node.accounts.Account
	(*namespaces.Namespace)(nil), // 3: infinimesh.node.namespaces.Namespace
	(*devices.Device)(nil),       // 4: infinimesh.node.devices.Device
	(*structpb.Struct)(nil),      // 5: google.protobuf.Struct
	(*node.EmptyMessage)(nil),    // 6: infinimesh.node.EmptyMessage
}
var file_eventbus_eventbus_proto_depIdxs = []int32{
	0, // 0: infinimesh.eventbus.Event.event_kind:type_name -> infinimesh.eventbus.EventKind
	2, // 1: infinimesh.eventbus.Event.account:type_name -> infinimesh.node.accounts.Account
	3, // 2: infinimesh.eventbus.Event.namespace:type_name -> infinimesh.node.namespaces.Namespace
	4, // 3: infinimesh.eventbus.Event.device:type_name -> infinimesh.node.devices.Device
	5, // 4: infinimesh.eventbus.Event.meta:type_name -> google.protobuf.Struct
	6, // 5: infinimesh.eventbus.EventsService.Subscribe:input_type -> infinimesh.node.EmptyMessage
	1, // 6: infinimesh.eventbus.EventsService.Subscribe:output_type -> infinimesh.eventbus.Event
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_eventbus_eventbus_proto_init() }
func file_eventbus_eventbus_proto_init() {
	if File_eventbus_eventbus_proto != nil {
		return
	}
	file_eventbus_eventbus_proto_msgTypes[0].OneofWrappers = []any{
		(*Event_Account)(nil),
		(*Event_Namespace)(nil),
		(*Event_Device)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_eventbus_eventbus_proto_rawDesc), len(file_eventbus_eventbus_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventbus_eventbus_proto_goTypes,
		DependencyIndexes: file_eventbus_eventbus_proto_depIdxs,
		EnumInfos:         file_eventbus_eventbus_proto_enumTypes,
		MessageInfos:      file_eventbus_eventbus_proto_msgTypes,
	}.Build()
	File_eventbus_eventbus_proto = out.File
	file_eventbus_eventbus_proto_goTypes = nil
	file_eventbus_eventbus_proto_depIdxs = nil
}
