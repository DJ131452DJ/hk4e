// Sorapointa - A server software re-implementation for a certain anime game, and avoid sorapointa.
// Copyright (C) 2022  Sorapointa Team
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: GCGMsgSkillResult.proto

package proto

import (
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

type GCGMsgSkillResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkillId        uint32             `protobuf:"varint,12,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	LastHp         uint32             `protobuf:"varint,14,opt,name=last_hp,json=lastHp,proto3" json:"last_hp,omitempty"`
	DetailList     []*GCGDamageDetail `protobuf:"bytes,2,rep,name=detail_list,json=detailList,proto3" json:"detail_list,omitempty"`
	TargetCardGuid uint32             `protobuf:"varint,7,opt,name=target_card_guid,json=targetCardGuid,proto3" json:"target_card_guid,omitempty"`
	EffectElement  uint32             `protobuf:"varint,5,opt,name=effect_element,json=effectElement,proto3" json:"effect_element,omitempty"`
	FromResultSeq  uint32             `protobuf:"varint,15,opt,name=from_result_seq,json=fromResultSeq,proto3" json:"from_result_seq,omitempty"`
	Damage         uint32             `protobuf:"varint,6,opt,name=damage,proto3" json:"damage,omitempty"`
	ResultSeq      uint32             `protobuf:"varint,4,opt,name=result_seq,json=resultSeq,proto3" json:"result_seq,omitempty"`
	SrcCardGuid    uint32             `protobuf:"varint,8,opt,name=src_card_guid,json=srcCardGuid,proto3" json:"src_card_guid,omitempty"`
}

func (x *GCGMsgSkillResult) Reset() {
	*x = GCGMsgSkillResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GCGMsgSkillResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GCGMsgSkillResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GCGMsgSkillResult) ProtoMessage() {}

func (x *GCGMsgSkillResult) ProtoReflect() protoreflect.Message {
	mi := &file_GCGMsgSkillResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GCGMsgSkillResult.ProtoReflect.Descriptor instead.
func (*GCGMsgSkillResult) Descriptor() ([]byte, []int) {
	return file_GCGMsgSkillResult_proto_rawDescGZIP(), []int{0}
}

func (x *GCGMsgSkillResult) GetSkillId() uint32 {
	if x != nil {
		return x.SkillId
	}
	return 0
}

func (x *GCGMsgSkillResult) GetLastHp() uint32 {
	if x != nil {
		return x.LastHp
	}
	return 0
}

func (x *GCGMsgSkillResult) GetDetailList() []*GCGDamageDetail {
	if x != nil {
		return x.DetailList
	}
	return nil
}

func (x *GCGMsgSkillResult) GetTargetCardGuid() uint32 {
	if x != nil {
		return x.TargetCardGuid
	}
	return 0
}

func (x *GCGMsgSkillResult) GetEffectElement() uint32 {
	if x != nil {
		return x.EffectElement
	}
	return 0
}

func (x *GCGMsgSkillResult) GetFromResultSeq() uint32 {
	if x != nil {
		return x.FromResultSeq
	}
	return 0
}

func (x *GCGMsgSkillResult) GetDamage() uint32 {
	if x != nil {
		return x.Damage
	}
	return 0
}

func (x *GCGMsgSkillResult) GetResultSeq() uint32 {
	if x != nil {
		return x.ResultSeq
	}
	return 0
}

func (x *GCGMsgSkillResult) GetSrcCardGuid() uint32 {
	if x != nil {
		return x.SrcCardGuid
	}
	return 0
}

var File_GCGMsgSkillResult_proto protoreflect.FileDescriptor

var file_GCGMsgSkillResult_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x43, 0x47, 0x4d, 0x73, 0x67, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x47, 0x43, 0x47, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x11, 0x47, 0x43, 0x47, 0x4d,
	0x73, 0x67, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x68, 0x70, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x48,
	0x70, 0x12, 0x37, 0x0a, 0x0b, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x43, 0x47, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x47, 0x75, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x5f, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x53, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x72,
	0x63, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x43, 0x61, 0x72, 0x64, 0x47, 0x75, 0x69, 0x64, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GCGMsgSkillResult_proto_rawDescOnce sync.Once
	file_GCGMsgSkillResult_proto_rawDescData = file_GCGMsgSkillResult_proto_rawDesc
)

func file_GCGMsgSkillResult_proto_rawDescGZIP() []byte {
	file_GCGMsgSkillResult_proto_rawDescOnce.Do(func() {
		file_GCGMsgSkillResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_GCGMsgSkillResult_proto_rawDescData)
	})
	return file_GCGMsgSkillResult_proto_rawDescData
}

var file_GCGMsgSkillResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GCGMsgSkillResult_proto_goTypes = []interface{}{
	(*GCGMsgSkillResult)(nil), // 0: proto.GCGMsgSkillResult
	(*GCGDamageDetail)(nil),   // 1: proto.GCGDamageDetail
}
var file_GCGMsgSkillResult_proto_depIdxs = []int32{
	1, // 0: proto.GCGMsgSkillResult.detail_list:type_name -> proto.GCGDamageDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GCGMsgSkillResult_proto_init() }
func file_GCGMsgSkillResult_proto_init() {
	if File_GCGMsgSkillResult_proto != nil {
		return
	}
	file_GCGDamageDetail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GCGMsgSkillResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GCGMsgSkillResult); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GCGMsgSkillResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GCGMsgSkillResult_proto_goTypes,
		DependencyIndexes: file_GCGMsgSkillResult_proto_depIdxs,
		MessageInfos:      file_GCGMsgSkillResult_proto_msgTypes,
	}.Build()
	File_GCGMsgSkillResult_proto = out.File
	file_GCGMsgSkillResult_proto_rawDesc = nil
	file_GCGMsgSkillResult_proto_goTypes = nil
	file_GCGMsgSkillResult_proto_depIdxs = nil
}