//
//Copyright © 2021-2023 Infinite Devices GmbH
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: shadow/shadow.proto

package shadow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type StateKey int32

const (
	StateKey_REPORTED   StateKey = 0
	StateKey_DESIRED    StateKey = 1
	StateKey_CONNECTION StateKey = 2
)

// Enum value maps for StateKey.
var (
	StateKey_name = map[int32]string{
		0: "REPORTED",
		1: "DESIRED",
		2: "CONNECTION",
	}
	StateKey_value = map[string]int32{
		"REPORTED":   0,
		"DESIRED":    1,
		"CONNECTION": 2,
	}
)

func (x StateKey) Enum() *StateKey {
	p := new(StateKey)
	*p = x
	return p
}

func (x StateKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StateKey) Descriptor() protoreflect.EnumDescriptor {
	return file_shadow_shadow_proto_enumTypes[0].Descriptor()
}

func (StateKey) Type() protoreflect.EnumType {
	return &file_shadow_shadow_proto_enumTypes[0]
}

func (x StateKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StateKey.Descriptor instead.
func (StateKey) EnumDescriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{0}
}

type State struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data          *structpb.Struct       `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *State) Reset() {
	*x = State{}
	mi := &file_shadow_shadow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{0}
}

func (x *State) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *State) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConnectionState struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Connected     bool                   `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	ClientId      string                 `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // MQTT Client ID used for connection
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectionState) Reset() {
	*x = ConnectionState{}
	mi := &file_shadow_shadow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionState) ProtoMessage() {}

func (x *ConnectionState) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionState.ProtoReflect.Descriptor instead.
func (*ConnectionState) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionState) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ConnectionState) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

func (x *ConnectionState) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Shadow struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Device        string                 `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Reported      *State                 `protobuf:"bytes,2,opt,name=reported,proto3" json:"reported,omitempty"`
	Desired       *State                 `protobuf:"bytes,3,opt,name=desired,proto3" json:"desired,omitempty"`
	Connection    *ConnectionState       `protobuf:"bytes,4,opt,name=connection,proto3" json:"connection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Shadow) Reset() {
	*x = Shadow{}
	mi := &file_shadow_shadow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shadow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shadow) ProtoMessage() {}

func (x *Shadow) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shadow.ProtoReflect.Descriptor instead.
func (*Shadow) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{2}
}

func (x *Shadow) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Shadow) GetReported() *State {
	if x != nil {
		return x.Reported
	}
	return nil
}

func (x *Shadow) GetDesired() *State {
	if x != nil {
		return x.Desired
	}
	return nil
}

func (x *Shadow) GetConnection() *ConnectionState {
	if x != nil {
		return x.Connection
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pool          []string               `protobuf:"bytes,1,rep,name=pool,proto3" json:"pool,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_shadow_shadow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{3}
}

func (x *GetRequest) GetPool() []string {
	if x != nil {
		return x.Pool
	}
	return nil
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shadows       []*Shadow              `protobuf:"bytes,1,rep,name=shadows,proto3" json:"shadows,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_shadow_shadow_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{4}
}

func (x *GetResponse) GetShadows() []*Shadow {
	if x != nil {
		return x.Shadows
	}
	return nil
}

type RemoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Device        string                 `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	StateKey      StateKey               `protobuf:"varint,2,opt,name=state_key,json=stateKey,proto3,enum=infinimesh.shadow.StateKey" json:"state_key,omitempty"`
	Key           string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	mi := &file_shadow_shadow_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *RemoveRequest) GetStateKey() StateKey {
	if x != nil {
		return x.StateKey
	}
	return StateKey_REPORTED
}

func (x *RemoveRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type StreamShadowRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Devices       []string               `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
	OnlyDelta     bool                   `protobuf:"varint,2,opt,name=only_delta,json=onlyDelta,proto3" json:"only_delta,omitempty"`
	Sync          bool                   `protobuf:"varint,3,opt,name=sync,proto3" json:"sync,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamShadowRequest) Reset() {
	*x = StreamShadowRequest{}
	mi := &file_shadow_shadow_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamShadowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamShadowRequest) ProtoMessage() {}

func (x *StreamShadowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shadow_shadow_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamShadowRequest.ProtoReflect.Descriptor instead.
func (*StreamShadowRequest) Descriptor() ([]byte, []int) {
	return file_shadow_shadow_proto_rawDescGZIP(), []int{6}
}

func (x *StreamShadowRequest) GetDevices() []string {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *StreamShadowRequest) GetOnlyDelta() bool {
	if x != nil {
		return x.OnlyDelta
	}
	return false
}

func (x *StreamShadowRequest) GetSync() bool {
	if x != nil {
		return x.Sync
	}
	return false
}

var File_shadow_shadow_proto protoreflect.FileDescriptor

var file_shadow_shadow_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x22, 0xce, 0x01, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x64, 0x65, 0x73,
	0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6f, 0x6f, 0x6c, 0x22, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52, 0x07,
	0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x22, 0x73, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x62, 0x0a, 0x13,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x79, 0x6e, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x79, 0x6e, 0x63,
	0x2a, 0x35, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x53, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0xb0, 0x02, 0x0a, 0x0d, 0x53, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x45,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x53, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x30, 0x01, 0x42, 0xb8, 0x01, 0x0a, 0x15, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x42, 0x0b, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x70, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0xa2, 0x02, 0x03, 0x49, 0x53, 0x58, 0xaa, 0x02, 0x11, 0x49, 0x6e, 0x66, 0x69, 0x6e,
	0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0xca, 0x02, 0x11, 0x49,
	0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77,
	0xe2, 0x02, 0x1d, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x53, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x12, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x53,
	0x68, 0x61, 0x64, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_shadow_shadow_proto_rawDescOnce sync.Once
	file_shadow_shadow_proto_rawDescData []byte
)

func file_shadow_shadow_proto_rawDescGZIP() []byte {
	file_shadow_shadow_proto_rawDescOnce.Do(func() {
		file_shadow_shadow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shadow_shadow_proto_rawDesc), len(file_shadow_shadow_proto_rawDesc)))
	})
	return file_shadow_shadow_proto_rawDescData
}

var file_shadow_shadow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shadow_shadow_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_shadow_shadow_proto_goTypes = []any{
	(StateKey)(0),                 // 0: infinimesh.shadow.StateKey
	(*State)(nil),                 // 1: infinimesh.shadow.State
	(*ConnectionState)(nil),       // 2: infinimesh.shadow.ConnectionState
	(*Shadow)(nil),                // 3: infinimesh.shadow.Shadow
	(*GetRequest)(nil),            // 4: infinimesh.shadow.GetRequest
	(*GetResponse)(nil),           // 5: infinimesh.shadow.GetResponse
	(*RemoveRequest)(nil),         // 6: infinimesh.shadow.RemoveRequest
	(*StreamShadowRequest)(nil),   // 7: infinimesh.shadow.StreamShadowRequest
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 9: google.protobuf.Struct
}
var file_shadow_shadow_proto_depIdxs = []int32{
	8,  // 0: infinimesh.shadow.State.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 1: infinimesh.shadow.State.data:type_name -> google.protobuf.Struct
	8,  // 2: infinimesh.shadow.ConnectionState.timestamp:type_name -> google.protobuf.Timestamp
	1,  // 3: infinimesh.shadow.Shadow.reported:type_name -> infinimesh.shadow.State
	1,  // 4: infinimesh.shadow.Shadow.desired:type_name -> infinimesh.shadow.State
	2,  // 5: infinimesh.shadow.Shadow.connection:type_name -> infinimesh.shadow.ConnectionState
	3,  // 6: infinimesh.shadow.GetResponse.shadows:type_name -> infinimesh.shadow.Shadow
	0,  // 7: infinimesh.shadow.RemoveRequest.state_key:type_name -> infinimesh.shadow.StateKey
	4,  // 8: infinimesh.shadow.ShadowService.Get:input_type -> infinimesh.shadow.GetRequest
	3,  // 9: infinimesh.shadow.ShadowService.Patch:input_type -> infinimesh.shadow.Shadow
	6,  // 10: infinimesh.shadow.ShadowService.Remove:input_type -> infinimesh.shadow.RemoveRequest
	7,  // 11: infinimesh.shadow.ShadowService.StreamShadow:input_type -> infinimesh.shadow.StreamShadowRequest
	5,  // 12: infinimesh.shadow.ShadowService.Get:output_type -> infinimesh.shadow.GetResponse
	3,  // 13: infinimesh.shadow.ShadowService.Patch:output_type -> infinimesh.shadow.Shadow
	3,  // 14: infinimesh.shadow.ShadowService.Remove:output_type -> infinimesh.shadow.Shadow
	3,  // 15: infinimesh.shadow.ShadowService.StreamShadow:output_type -> infinimesh.shadow.Shadow
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_shadow_shadow_proto_init() }
func file_shadow_shadow_proto_init() {
	if File_shadow_shadow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shadow_shadow_proto_rawDesc), len(file_shadow_shadow_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shadow_shadow_proto_goTypes,
		DependencyIndexes: file_shadow_shadow_proto_depIdxs,
		EnumInfos:         file_shadow_shadow_proto_enumTypes,
		MessageInfos:      file_shadow_shadow_proto_msgTypes,
	}.Build()
	File_shadow_shadow_proto = out.File
	file_shadow_shadow_proto_goTypes = nil
	file_shadow_shadow_proto_depIdxs = nil
}
