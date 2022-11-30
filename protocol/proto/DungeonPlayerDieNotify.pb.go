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
// source: DungeonPlayerDieNotify.proto

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

// CmdId: 931
// EnetChannelId: 0
// EnetIsReliable: true
type DungeonPlayerDieNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StrengthenPointDataMap map[uint32]*StrengthenPointData `protobuf:"bytes,15,rep,name=strengthen_point_data_map,json=strengthenPointDataMap,proto3" json:"strengthen_point_data_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WaitTime               uint32                          `protobuf:"varint,1,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	DungeonId              uint32                          `protobuf:"varint,9,opt,name=dungeon_id,json=dungeonId,proto3" json:"dungeon_id,omitempty"`
	MurdererEntityId       uint32                          `protobuf:"varint,13,opt,name=murderer_entity_id,json=murdererEntityId,proto3" json:"murderer_entity_id,omitempty"`
	DieType                PlayerDieType                   `protobuf:"varint,3,opt,name=die_type,json=dieType,proto3,enum=proto.PlayerDieType" json:"die_type,omitempty"`
	ReviveCount            uint32                          `protobuf:"varint,6,opt,name=revive_count,json=reviveCount,proto3" json:"revive_count,omitempty"`
	// Types that are assignable to Entity:
	//
	//	*DungeonPlayerDieNotify_MonsterId
	//	*DungeonPlayerDieNotify_GadgetId
	Entity isDungeonPlayerDieNotify_Entity `protobuf_oneof:"entity"`
}

func (x *DungeonPlayerDieNotify) Reset() {
	*x = DungeonPlayerDieNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DungeonPlayerDieNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DungeonPlayerDieNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DungeonPlayerDieNotify) ProtoMessage() {}

func (x *DungeonPlayerDieNotify) ProtoReflect() protoreflect.Message {
	mi := &file_DungeonPlayerDieNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DungeonPlayerDieNotify.ProtoReflect.Descriptor instead.
func (*DungeonPlayerDieNotify) Descriptor() ([]byte, []int) {
	return file_DungeonPlayerDieNotify_proto_rawDescGZIP(), []int{0}
}

func (x *DungeonPlayerDieNotify) GetStrengthenPointDataMap() map[uint32]*StrengthenPointData {
	if x != nil {
		return x.StrengthenPointDataMap
	}
	return nil
}

func (x *DungeonPlayerDieNotify) GetWaitTime() uint32 {
	if x != nil {
		return x.WaitTime
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetDungeonId() uint32 {
	if x != nil {
		return x.DungeonId
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetMurdererEntityId() uint32 {
	if x != nil {
		return x.MurdererEntityId
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetDieType() PlayerDieType {
	if x != nil {
		return x.DieType
	}
	return PlayerDieType_PLAYER_DIE_TYPE_NONE
}

func (x *DungeonPlayerDieNotify) GetReviveCount() uint32 {
	if x != nil {
		return x.ReviveCount
	}
	return 0
}

func (m *DungeonPlayerDieNotify) GetEntity() isDungeonPlayerDieNotify_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (x *DungeonPlayerDieNotify) GetMonsterId() uint32 {
	if x, ok := x.GetEntity().(*DungeonPlayerDieNotify_MonsterId); ok {
		return x.MonsterId
	}
	return 0
}

func (x *DungeonPlayerDieNotify) GetGadgetId() uint32 {
	if x, ok := x.GetEntity().(*DungeonPlayerDieNotify_GadgetId); ok {
		return x.GadgetId
	}
	return 0
}

type isDungeonPlayerDieNotify_Entity interface {
	isDungeonPlayerDieNotify_Entity()
}

type DungeonPlayerDieNotify_MonsterId struct {
	MonsterId uint32 `protobuf:"varint,4,opt,name=monster_id,json=monsterId,proto3,oneof"`
}

type DungeonPlayerDieNotify_GadgetId struct {
	GadgetId uint32 `protobuf:"varint,8,opt,name=gadget_id,json=gadgetId,proto3,oneof"`
}

func (*DungeonPlayerDieNotify_MonsterId) isDungeonPlayerDieNotify_Entity() {}

func (*DungeonPlayerDieNotify_GadgetId) isDungeonPlayerDieNotify_Entity() {}

var File_DungeonPlayerDieNotify_proto protoreflect.FileDescriptor

var file_DungeonPlayerDieNotify_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44,
	0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x53, 0x74, 0x72, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x03, 0x0a, 0x16, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x74, 0x0a, 0x19, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x5f, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x0f, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x6e, 0x67,
	0x65, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x69, 0x65, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16,
	0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x77, 0x61, 0x69, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x75, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x6d, 0x75, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x2f, 0x0a, 0x08, 0x64, 0x69, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x44, 0x69, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x64, 0x69, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x76, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x67, 0x61, 0x64, 0x67, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x67, 0x61, 0x64, 0x67,
	0x65, 0x74, 0x49, 0x64, 0x1a, 0x65, 0x0a, 0x1b, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x70, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x65, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DungeonPlayerDieNotify_proto_rawDescOnce sync.Once
	file_DungeonPlayerDieNotify_proto_rawDescData = file_DungeonPlayerDieNotify_proto_rawDesc
)

func file_DungeonPlayerDieNotify_proto_rawDescGZIP() []byte {
	file_DungeonPlayerDieNotify_proto_rawDescOnce.Do(func() {
		file_DungeonPlayerDieNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_DungeonPlayerDieNotify_proto_rawDescData)
	})
	return file_DungeonPlayerDieNotify_proto_rawDescData
}

var file_DungeonPlayerDieNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_DungeonPlayerDieNotify_proto_goTypes = []interface{}{
	(*DungeonPlayerDieNotify)(nil), // 0: proto.DungeonPlayerDieNotify
	nil,                            // 1: proto.DungeonPlayerDieNotify.StrengthenPointDataMapEntry
	(PlayerDieType)(0),             // 2: proto.PlayerDieType
	(*StrengthenPointData)(nil),    // 3: proto.StrengthenPointData
}
var file_DungeonPlayerDieNotify_proto_depIdxs = []int32{
	1, // 0: proto.DungeonPlayerDieNotify.strengthen_point_data_map:type_name -> proto.DungeonPlayerDieNotify.StrengthenPointDataMapEntry
	2, // 1: proto.DungeonPlayerDieNotify.die_type:type_name -> proto.PlayerDieType
	3, // 2: proto.DungeonPlayerDieNotify.StrengthenPointDataMapEntry.value:type_name -> proto.StrengthenPointData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_DungeonPlayerDieNotify_proto_init() }
func file_DungeonPlayerDieNotify_proto_init() {
	if File_DungeonPlayerDieNotify_proto != nil {
		return
	}
	file_PlayerDieType_proto_init()
	file_StrengthenPointData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DungeonPlayerDieNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DungeonPlayerDieNotify); i {
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
	file_DungeonPlayerDieNotify_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DungeonPlayerDieNotify_MonsterId)(nil),
		(*DungeonPlayerDieNotify_GadgetId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_DungeonPlayerDieNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DungeonPlayerDieNotify_proto_goTypes,
		DependencyIndexes: file_DungeonPlayerDieNotify_proto_depIdxs,
		MessageInfos:      file_DungeonPlayerDieNotify_proto_msgTypes,
	}.Build()
	File_DungeonPlayerDieNotify_proto = out.File
	file_DungeonPlayerDieNotify_proto_rawDesc = nil
	file_DungeonPlayerDieNotify_proto_goTypes = nil
	file_DungeonPlayerDieNotify_proto_depIdxs = nil
}