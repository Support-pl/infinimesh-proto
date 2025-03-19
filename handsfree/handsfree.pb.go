//
//Copyright © 2021-2023 Infinite Devices GmbH, Nikita Ivanovski info@slnt-opp.xyz
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
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: handsfree/handsfree.proto

package handsfree

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Code int32

const (
	Code_UNKNOWN Code = 0
	Code_AUTH    Code = 1
	Code_DATA    Code = 2
	Code_SUCCESS Code = 3
	Code_ERROR   Code = 4
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "AUTH",
		2: "DATA",
		3: "SUCCESS",
		4: "ERROR",
	}
	Code_value = map[string]int32{
		"UNKNOWN": 0,
		"AUTH":    1,
		"DATA":    2,
		"SUCCESS": 3,
		"ERROR":   4,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_handsfree_handsfree_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_handsfree_handsfree_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_handsfree_handsfree_proto_rawDescGZIP(), []int{0}
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId   string   `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Payload []string `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty"` // anything you want to share with the Console
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handsfree_handsfree_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handsfree_handsfree_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_handsfree_handsfree_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ConnectionRequest) GetPayload() []string {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ControlPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    Code     `protobuf:"varint,1,opt,name=code,proto3,enum=infinimesh.handsfree.Code" json:"code,omitempty"`
	Payload []string `protobuf:"bytes,2,rep,name=payload,proto3" json:"payload,omitempty"` // payload simply depends on returned code
	AppId   *string  `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3,oneof" json:"app_id,omitempty"`
}

func (x *ControlPacket) Reset() {
	*x = ControlPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handsfree_handsfree_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlPacket) ProtoMessage() {}

func (x *ControlPacket) ProtoReflect() protoreflect.Message {
	mi := &file_handsfree_handsfree_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlPacket.ProtoReflect.Descriptor instead.
func (*ControlPacket) Descriptor() ([]byte, []int) {
	return file_handsfree_handsfree_proto_rawDescGZIP(), []int{1}
}

func (x *ControlPacket) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_UNKNOWN
}

func (x *ControlPacket) GetPayload() []string {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ControlPacket) GetAppId() string {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return ""
}

var File_handsfree_handsfree_proto protoreflect.FileDescriptor

var file_handsfree_handsfree_proto_rawDesc = []byte{
	0x0a, 0x19, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x2f, 0x68, 0x61, 0x6e, 0x64,
	0x73, 0x66, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65,
	0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1a, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x2a, 0x3f, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x41, 0x55, 0x54, 0x48, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10,
	0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32, 0xf5, 0x01, 0x0a, 0x10, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67,
	0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d,
	0x65, 0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72,
	0x65, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x12, 0x78, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x27, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72,
	0x65, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x73,
	0x66, 0x72, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x30,
	0x01, 0x42, 0xc2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66, 0x69, 0x6e, 0x69,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x42, 0x0e,
	0x48, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0xa2, 0x02, 0x03, 0x49, 0x48, 0x58, 0xaa, 0x02, 0x14,
	0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73,
	0x66, 0x72, 0x65, 0x65, 0xca, 0x02, 0x14, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73,
	0x68, 0x5c, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0xe2, 0x02, 0x20, 0x49, 0x6e,
	0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x5c, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x66, 0x72,
	0x65, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x15, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x3a, 0x48, 0x61, 0x6e,
	0x64, 0x73, 0x66, 0x72, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_handsfree_handsfree_proto_rawDescOnce sync.Once
	file_handsfree_handsfree_proto_rawDescData = file_handsfree_handsfree_proto_rawDesc
)

func file_handsfree_handsfree_proto_rawDescGZIP() []byte {
	file_handsfree_handsfree_proto_rawDescOnce.Do(func() {
		file_handsfree_handsfree_proto_rawDescData = protoimpl.X.CompressGZIP(file_handsfree_handsfree_proto_rawDescData)
	})
	return file_handsfree_handsfree_proto_rawDescData
}

var file_handsfree_handsfree_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_handsfree_handsfree_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_handsfree_handsfree_proto_goTypes = []interface{}{
	(Code)(0),                 // 0: infinimesh.handsfree.Code
	(*ConnectionRequest)(nil), // 1: infinimesh.handsfree.ConnectionRequest
	(*ControlPacket)(nil),     // 2: infinimesh.handsfree.ControlPacket
}
var file_handsfree_handsfree_proto_depIdxs = []int32{
	0, // 0: infinimesh.handsfree.ControlPacket.code:type_name -> infinimesh.handsfree.Code
	2, // 1: infinimesh.handsfree.HandsfreeService.Send:input_type -> infinimesh.handsfree.ControlPacket
	1, // 2: infinimesh.handsfree.HandsfreeService.Connect:input_type -> infinimesh.handsfree.ConnectionRequest
	2, // 3: infinimesh.handsfree.HandsfreeService.Send:output_type -> infinimesh.handsfree.ControlPacket
	2, // 4: infinimesh.handsfree.HandsfreeService.Connect:output_type -> infinimesh.handsfree.ControlPacket
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_handsfree_handsfree_proto_init() }
func file_handsfree_handsfree_proto_init() {
	if File_handsfree_handsfree_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handsfree_handsfree_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_handsfree_handsfree_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_handsfree_handsfree_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_handsfree_handsfree_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handsfree_handsfree_proto_goTypes,
		DependencyIndexes: file_handsfree_handsfree_proto_depIdxs,
		EnumInfos:         file_handsfree_handsfree_proto_enumTypes,
		MessageInfos:      file_handsfree_handsfree_proto_msgTypes,
	}.Build()
	File_handsfree_handsfree_proto = out.File
	file_handsfree_handsfree_proto_rawDesc = nil
	file_handsfree_handsfree_proto_goTypes = nil
	file_handsfree_handsfree_proto_depIdxs = nil
}
