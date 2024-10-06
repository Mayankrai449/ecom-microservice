// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: inventory.proto

package inventory

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

type StockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductIds []uint32 `protobuf:"varint,1,rep,packed,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	Quantities []uint32 `protobuf:"varint,2,rep,packed,name=quantities,proto3" json:"quantities,omitempty"`
}

func (x *StockRequest) Reset() {
	*x = StockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockRequest) ProtoMessage() {}

func (x *StockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockRequest.ProtoReflect.Descriptor instead.
func (*StockRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *StockRequest) GetProductIds() []uint32 {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *StockRequest) GetQuantities() []uint32 {
	if x != nil {
		return x.Quantities
	}
	return nil
}

type StockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InStock bool   `protobuf:"varint,1,opt,name=in_stock,json=inStock,proto3" json:"in_stock,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StockResponse) Reset() {
	*x = StockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockResponse) ProtoMessage() {}

func (x *StockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockResponse.ProtoReflect.Descriptor instead.
func (*StockResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *StockResponse) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

func (x *StockResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

var file_inventory_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x4f, 0x0a, 0x0c,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0a, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x44, 0x0a,
	0x0d, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x53, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData = file_inventory_proto_rawDesc
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_proto_rawDescData)
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_proto_goTypes = []any{
	(*StockRequest)(nil),  // 0: inventory.StockRequest
	(*StockResponse)(nil), // 1: inventory.StockResponse
}
var file_inventory_proto_depIdxs = []int32{
	0, // 0: inventory.InventoryService.CheckStock:input_type -> inventory.StockRequest
	1, // 1: inventory.InventoryService.CheckStock:output_type -> inventory.StockResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StockRequest); i {
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
		file_inventory_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*StockResponse); i {
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
			RawDescriptor: file_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_rawDesc = nil
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}