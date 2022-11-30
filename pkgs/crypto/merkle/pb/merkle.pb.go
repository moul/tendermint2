// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: merkle.proto

package pb

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

// messages
type ProofOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Key  []byte `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ProofOp) Reset() {
	*x = ProofOp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merkle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofOp) ProtoMessage() {}

func (x *ProofOp) ProtoReflect() protoreflect.Message {
	mi := &file_merkle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofOp.ProtoReflect.Descriptor instead.
func (*ProofOp) Descriptor() ([]byte, []int) {
	return file_merkle_proto_rawDescGZIP(), []int{0}
}

func (x *ProofOp) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ProofOp) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *ProofOp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Proof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ops []*ProofOp `protobuf:"bytes,1,rep,name=Ops,proto3" json:"Ops,omitempty"`
}

func (x *Proof) Reset() {
	*x = Proof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merkle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proof) ProtoMessage() {}

func (x *Proof) ProtoReflect() protoreflect.Message {
	mi := &file_merkle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proof.ProtoReflect.Descriptor instead.
func (*Proof) Descriptor() ([]byte, []int) {
	return file_merkle_proto_rawDescGZIP(), []int{1}
}

func (x *Proof) GetOps() []*ProofOp {
	if x != nil {
		return x.Ops
	}
	return nil
}

type SimpleProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64    `protobuf:"zigzag64,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Index    int64    `protobuf:"zigzag64,2,opt,name=Index,proto3" json:"Index,omitempty"`
	LeafHash []byte   `protobuf:"bytes,3,opt,name=LeafHash,proto3" json:"LeafHash,omitempty"`
	Aunts    [][]byte `protobuf:"bytes,4,rep,name=Aunts,proto3" json:"Aunts,omitempty"`
}

func (x *SimpleProof) Reset() {
	*x = SimpleProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merkle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleProof) ProtoMessage() {}

func (x *SimpleProof) ProtoReflect() protoreflect.Message {
	mi := &file_merkle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleProof.ProtoReflect.Descriptor instead.
func (*SimpleProof) Descriptor() ([]byte, []int) {
	return file_merkle_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleProof) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SimpleProof) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SimpleProof) GetLeafHash() []byte {
	if x != nil {
		return x.LeafHash
	}
	return nil
}

func (x *SimpleProof) GetAunts() [][]byte {
	if x != nil {
		return x.Aunts
	}
	return nil
}

type SimpleProofNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte           `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Parent *SimpleProofNode `protobuf:"bytes,2,opt,name=Parent,proto3" json:"Parent,omitempty"`
	Left   *SimpleProofNode `protobuf:"bytes,3,opt,name=Left,proto3" json:"Left,omitempty"`
	Right  *SimpleProofNode `protobuf:"bytes,4,opt,name=Right,proto3" json:"Right,omitempty"`
}

func (x *SimpleProofNode) Reset() {
	*x = SimpleProofNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merkle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleProofNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleProofNode) ProtoMessage() {}

func (x *SimpleProofNode) ProtoReflect() protoreflect.Message {
	mi := &file_merkle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleProofNode.ProtoReflect.Descriptor instead.
func (*SimpleProofNode) Descriptor() ([]byte, []int) {
	return file_merkle_proto_rawDescGZIP(), []int{3}
}

func (x *SimpleProofNode) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *SimpleProofNode) GetParent() *SimpleProofNode {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *SimpleProofNode) GetLeft() *SimpleProofNode {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *SimpleProofNode) GetRight() *SimpleProofNode {
	if x != nil {
		return x.Right
	}
	return nil
}

type MERKLE_BytesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value [][]byte `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
}

func (x *MERKLE_BytesList) Reset() {
	*x = MERKLE_BytesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_merkle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MERKLE_BytesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MERKLE_BytesList) ProtoMessage() {}

func (x *MERKLE_BytesList) ProtoReflect() protoreflect.Message {
	mi := &file_merkle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MERKLE_BytesList.ProtoReflect.Descriptor instead.
func (*MERKLE_BytesList) Descriptor() ([]byte, []int) {
	return file_merkle_proto_rawDescGZIP(), []int{4}
}

func (x *MERKLE_BytesList) GetValue() [][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_merkle_proto protoreflect.FileDescriptor

var file_merkle_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x74, 0x6d, 0x22, 0x43, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4f, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x1d, 0x0a, 0x03, 0x4f, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x74, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4f, 0x70, 0x52, 0x03, 0x4f, 0x70, 0x73, 0x22,
	0x6b, 0x0a, 0x0b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x12, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x65,
	0x61, 0x66, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4c, 0x65,
	0x61, 0x66, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x41, 0x75, 0x6e, 0x74, 0x73, 0x22, 0xa6, 0x01, 0x0a,
	0x0f, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x2b, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x74, 0x6d, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x52, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x6d, 0x2e, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x52, 0x69, 0x67, 0x68, 0x74, 0x22, 0x28, 0x0a, 0x10, 0x4d, 0x45, 0x52, 0x4b, 0x4c, 0x45, 0x5f,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6e,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x73, 0x2f, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_merkle_proto_rawDescOnce sync.Once
	file_merkle_proto_rawDescData = file_merkle_proto_rawDesc
)

func file_merkle_proto_rawDescGZIP() []byte {
	file_merkle_proto_rawDescOnce.Do(func() {
		file_merkle_proto_rawDescData = protoimpl.X.CompressGZIP(file_merkle_proto_rawDescData)
	})
	return file_merkle_proto_rawDescData
}

var file_merkle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_merkle_proto_goTypes = []interface{}{
	(*ProofOp)(nil),          // 0: tm.ProofOp
	(*Proof)(nil),            // 1: tm.Proof
	(*SimpleProof)(nil),      // 2: tm.SimpleProof
	(*SimpleProofNode)(nil),  // 3: tm.SimpleProofNode
	(*MERKLE_BytesList)(nil), // 4: tm.MERKLE_BytesList
}
var file_merkle_proto_depIdxs = []int32{
	0, // 0: tm.Proof.Ops:type_name -> tm.ProofOp
	3, // 1: tm.SimpleProofNode.Parent:type_name -> tm.SimpleProofNode
	3, // 2: tm.SimpleProofNode.Left:type_name -> tm.SimpleProofNode
	3, // 3: tm.SimpleProofNode.Right:type_name -> tm.SimpleProofNode
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_merkle_proto_init() }
func file_merkle_proto_init() {
	if File_merkle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_merkle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofOp); i {
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
		file_merkle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proof); i {
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
		file_merkle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleProof); i {
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
		file_merkle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleProofNode); i {
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
		file_merkle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MERKLE_BytesList); i {
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
			RawDescriptor: file_merkle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_merkle_proto_goTypes,
		DependencyIndexes: file_merkle_proto_depIdxs,
		MessageInfos:      file_merkle_proto_msgTypes,
	}.Build()
	File_merkle_proto = out.File
	file_merkle_proto_rawDesc = nil
	file_merkle_proto_goTypes = nil
	file_merkle_proto_depIdxs = nil
}