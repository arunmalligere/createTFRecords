// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/allocation_description.proto

/*
Package tensorflow is a generated protocol buffer package.

It is generated from these files:
	tensorflow/core/framework/allocation_description.proto
	tensorflow/core/framework/api_def.proto
	tensorflow/core/framework/attr_value.proto
	tensorflow/core/framework/cost_graph.proto
	tensorflow/core/framework/device_attributes.proto
	tensorflow/core/framework/function.proto
	tensorflow/core/framework/graph.proto
	tensorflow/core/framework/graph_transfer_info.proto
	tensorflow/core/framework/kernel_def.proto
	tensorflow/core/framework/log_memory.proto
	tensorflow/core/framework/node_def.proto
	tensorflow/core/framework/op_def.proto
	tensorflow/core/framework/op_gen_overrides.proto
	tensorflow/core/framework/reader_base.proto
	tensorflow/core/framework/remote_fused_graph_execute_info.proto
	tensorflow/core/framework/resource_handle.proto
	tensorflow/core/framework/step_stats.proto
	tensorflow/core/framework/summary.proto
	tensorflow/core/framework/tensor.proto
	tensorflow/core/framework/tensor_description.proto
	tensorflow/core/framework/tensor_shape.proto
	tensorflow/core/framework/tensor_slice.proto
	tensorflow/core/framework/types.proto
	tensorflow/core/framework/variable.proto
	tensorflow/core/framework/versions.proto

It has these top-level messages:
	AllocationDescription
	ApiDef
	ApiDefs
	AttrValue
	NameAttrList
	CostGraphDef
	DeviceLocality
	DeviceAttributes
	FunctionDefLibrary
	FunctionDef
	GradientDef
	GraphDef
	GraphTransferInfo
	KernelDef
	MemoryLogStep
	MemoryLogTensorAllocation
	MemoryLogTensorDeallocation
	MemoryLogTensorOutput
	MemoryLogRawAllocation
	MemoryLogRawDeallocation
	NodeDef
	OpDef
	OpDeprecation
	OpList
	OpGenOverride
	OpGenOverrides
	ReaderBaseState
	RemoteFusedGraphExecuteInfo
	ResourceHandleProto
	AllocationRecord
	AllocatorMemoryUsed
	NodeOutput
	MemoryStats
	NodeExecStats
	DeviceStepStats
	StepStats
	SummaryDescription
	HistogramProto
	SummaryMetadata
	Summary
	TensorProto
	VariantTensorDataProto
	TensorDescription
	TensorShapeProto
	TensorSliceProto
	VariableDef
	SaveSliceInfoDef
	VersionDef
*/
package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr uint64 `protobuf:"varint,6,opt,name=ptr" json:"ptr,omitempty"`
}

func (m *AllocationDescription) Reset()                    { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string            { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()               {}
func (*AllocationDescription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/allocation_description.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x89, 0x9b, 0x43, 0x83, 0x53, 0x09, 0x0a, 0x01, 0x2f, 0x45, 0x11, 0x7b, 0x6a, 0x05,
	0xc1, 0x93, 0x17, 0x8b, 0x17, 0x2f, 0x32, 0xe2, 0x1f, 0x10, 0xb2, 0xf6, 0x75, 0x2b, 0xb6, 0x79,
	0xf5, 0x25, 0x32, 0xfc, 0xcf, 0xf5, 0x26, 0xe9, 0x5c, 0xea, 0xc1, 0xdb, 0xe3, 0x7b, 0x5f, 0x7e,
	0xf0, 0xf1, 0x7b, 0x0f, 0xd6, 0x21, 0xd5, 0x2d, 0x6e, 0xf2, 0x12, 0x09, 0xf2, 0x9a, 0x4c, 0x07,
	0x1b, 0xa4, 0xb7, 0xdc, 0xb4, 0x2d, 0x96, 0xc6, 0x37, 0x68, 0x75, 0x05, 0xae, 0xa4, 0xa6, 0x0f,
	0x73, 0xd6, 0x13, 0x7a, 0x14, 0x7c, 0x3c, 0x77, 0xf9, 0xcd, 0xf8, 0xf9, 0x63, 0x94, 0x9f, 0x46,
	0x57, 0xdc, 0xf0, 0x13, 0x82, 0xf7, 0x0f, 0x70, 0x1e, 0x2a, 0xbd, 0xfc, 0xf4, 0xe0, 0x24, 0x4b,
	0x58, 0x3a, 0x51, 0xc7, 0x11, 0x17, 0x81, 0x06, 0xf1, 0xf7, 0xb9, 0x28, 0xee, 0x6d, 0xc5, 0x88,
	0xb7, 0xe2, 0x35, 0xdf, 0x11, 0x24, 0x6d, 0x4d, 0x07, 0x72, 0x92, 0xb0, 0xf4, 0x50, 0xcd, 0x23,
	0x7d, 0x31, 0x1d, 0x88, 0x2b, 0x3e, 0xff, 0xf3, 0xfd, 0xa6, 0x92, 0xd3, 0xe1, 0xb6, 0xa3, 0x11,
	0x3e, 0x57, 0xe2, 0x96, 0x9f, 0xad, 0x8d, 0xd3, 0xae, 0xb1, 0xab, 0x16, 0x34, 0x41, 0x0d, 0x04,
	0xb6, 0x04, 0xb9, 0x9f, 0xb0, 0xf4, 0x40, 0x89, 0xb5, 0x71, 0xaf, 0xc3, 0x4a, 0xed, 0x36, 0xe2,
	0x94, 0x4f, 0x7a, 0x4f, 0x72, 0x96, 0xb0, 0x74, 0xaa, 0xc2, 0x58, 0x3c, 0x70, 0x89, 0xb4, 0xca,
	0xc6, 0x1a, 0x59, 0x0c, 0x58, 0x5c, 0xfc, 0x1b, 0x65, 0x11, 0xfa, 0xb9, 0x05, 0xfb, 0x62, 0x6c,
	0x39, 0x1b, 0x62, 0xde, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xb7, 0xe1, 0x85, 0x86, 0x01,
	0x00, 0x00,
}
