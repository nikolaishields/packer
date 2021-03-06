// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/compute/v1/disk.proto

package compute

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Disk_Status int32

const (
	Disk_STATUS_UNSPECIFIED Disk_Status = 0
	// Disk is being created.
	Disk_CREATING Disk_Status = 1
	// Disk is ready to use.
	Disk_READY Disk_Status = 2
	// Disk encountered a problem and cannot operate.
	Disk_ERROR Disk_Status = 3
	// Disk is being deleted.
	Disk_DELETING Disk_Status = 4
)

var Disk_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
	3: "ERROR",
	4: "DELETING",
}

var Disk_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"READY":              2,
	"ERROR":              3,
	"DELETING":           4,
}

func (x Disk_Status) String() string {
	return proto.EnumName(Disk_Status_name, int32(x))
}

func (Disk_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6ed05fced9465d11, []int{0, 0}
}

// A Disk resource. For more information, see [Disks](/docs/compute/concepts/disk).
type Disk struct {
	// ID of the disk.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the disk belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the disk. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the disk. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the disk type.
	TypeId string `protobuf:"bytes,7,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// ID of the availability zone where the disk resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Size of the disk, specified in bytes.
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// Block size of the disk, specifiedin bytes.
	BlockSize int64 `protobuf:"varint,15,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	// License IDs that indicate which licenses are attached to this resource.
	// License IDs are used to calculate additional charges for the use of the virtual machine.
	//
	// The correct license ID is generated by Yandex.Cloud. IDs are inherited by new resources created from this resource.
	//
	// If you know the license IDs, specify them when you create the image.
	// For example, if you create a disk image using a third-party utility and load it into Yandex Object Storage, the license IDs will be lost.
	// You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.
	ProductIds []string `protobuf:"bytes,10,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	// Current status of the disk.
	Status Disk_Status `protobuf:"varint,11,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Disk_Status" json:"status,omitempty"`
	// Types that are valid to be assigned to Source:
	//	*Disk_SourceImageId
	//	*Disk_SourceSnapshotId
	Source isDisk_Source `protobuf_oneof:"source"`
	// Array of instances to which the disk is attached.
	InstanceIds          []string `protobuf:"bytes,14,rep,name=instance_ids,json=instanceIds,proto3" json:"instance_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disk) Reset()         { *m = Disk{} }
func (m *Disk) String() string { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()    {}
func (*Disk) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ed05fced9465d11, []int{0}
}

func (m *Disk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disk.Unmarshal(m, b)
}
func (m *Disk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disk.Marshal(b, m, deterministic)
}
func (m *Disk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disk.Merge(m, src)
}
func (m *Disk) XXX_Size() int {
	return xxx_messageInfo_Disk.Size(m)
}
func (m *Disk) XXX_DiscardUnknown() {
	xxx_messageInfo_Disk.DiscardUnknown(m)
}

var xxx_messageInfo_Disk proto.InternalMessageInfo

func (m *Disk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Disk) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Disk) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Disk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Disk) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Disk) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Disk) GetTypeId() string {
	if m != nil {
		return m.TypeId
	}
	return ""
}

func (m *Disk) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *Disk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Disk) GetBlockSize() int64 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *Disk) GetProductIds() []string {
	if m != nil {
		return m.ProductIds
	}
	return nil
}

func (m *Disk) GetStatus() Disk_Status {
	if m != nil {
		return m.Status
	}
	return Disk_STATUS_UNSPECIFIED
}

type isDisk_Source interface {
	isDisk_Source()
}

type Disk_SourceImageId struct {
	SourceImageId string `protobuf:"bytes,12,opt,name=source_image_id,json=sourceImageId,proto3,oneof"`
}

type Disk_SourceSnapshotId struct {
	SourceSnapshotId string `protobuf:"bytes,13,opt,name=source_snapshot_id,json=sourceSnapshotId,proto3,oneof"`
}

func (*Disk_SourceImageId) isDisk_Source() {}

func (*Disk_SourceSnapshotId) isDisk_Source() {}

func (m *Disk) GetSource() isDisk_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *Disk) GetSourceImageId() string {
	if x, ok := m.GetSource().(*Disk_SourceImageId); ok {
		return x.SourceImageId
	}
	return ""
}

func (m *Disk) GetSourceSnapshotId() string {
	if x, ok := m.GetSource().(*Disk_SourceSnapshotId); ok {
		return x.SourceSnapshotId
	}
	return ""
}

func (m *Disk) GetInstanceIds() []string {
	if m != nil {
		return m.InstanceIds
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Disk) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Disk_SourceImageId)(nil),
		(*Disk_SourceSnapshotId)(nil),
	}
}

func init() {
	proto.RegisterEnum("yandex.cloud.compute.v1.Disk_Status", Disk_Status_name, Disk_Status_value)
	proto.RegisterType((*Disk)(nil), "yandex.cloud.compute.v1.Disk")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.compute.v1.Disk.LabelsEntry")
}

func init() {
	proto.RegisterFile("yandex/cloud/compute/v1/disk.proto", fileDescriptor_6ed05fced9465d11)
}

var fileDescriptor_6ed05fced9465d11 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4d, 0x4f, 0xdc, 0x3c,
	0x14, 0x85, 0xc9, 0x7c, 0x84, 0xc9, 0x0d, 0x1f, 0x23, 0xeb, 0xd5, 0x4b, 0x04, 0xaa, 0x48, 0x51,
	0x17, 0xe9, 0x82, 0x44, 0xd0, 0x4d, 0x69, 0xbb, 0x19, 0x98, 0xb4, 0x8d, 0x84, 0x68, 0xe5, 0x19,
	0x16, 0xed, 0x66, 0xe4, 0x89, 0x4d, 0xb0, 0x26, 0x13, 0x47, 0xb1, 0x83, 0x3a, 0xfc, 0xbb, 0xfe,
	0xb3, 0xca, 0x76, 0x90, 0xe8, 0x82, 0xee, 0xae, 0xcf, 0x79, 0xec, 0x7b, 0x8f, 0x65, 0xc3, 0xc9,
	0x86, 0x54, 0x94, 0xfd, 0x4a, 0xf2, 0x52, 0xb4, 0x34, 0xc9, 0xc5, 0xba, 0x6e, 0x15, 0x4b, 0x1e,
	0xce, 0x12, 0xca, 0xe5, 0x2a, 0xae, 0x1b, 0xa1, 0x04, 0x3a, 0xb0, 0x4c, 0x6c, 0x98, 0xb8, 0x63,
	0xe2, 0x87, 0xb3, 0xc3, 0xe3, 0x42, 0x88, 0xa2, 0x64, 0x89, 0xc1, 0x96, 0xed, 0x5d, 0xa2, 0xf8,
	0x9a, 0x49, 0x45, 0xd6, 0xb5, 0xdd, 0x79, 0xf2, 0x7b, 0x08, 0x83, 0x29, 0x97, 0x2b, 0xb4, 0x07,
	0x3d, 0x4e, 0x03, 0x27, 0x74, 0x22, 0x0f, 0xf7, 0x38, 0x45, 0x47, 0xe0, 0xdd, 0x89, 0x92, 0xb2,
	0x66, 0xc1, 0x69, 0xd0, 0x33, 0xf2, 0xc8, 0x0a, 0x19, 0x45, 0x17, 0x00, 0x79, 0xc3, 0x88, 0x62,
	0x74, 0x41, 0x54, 0xd0, 0x0f, 0x9d, 0xc8, 0x3f, 0x3f, 0x8c, 0x6d, 0xaf, 0xf8, 0xa9, 0x57, 0x3c,
	0x7f, 0xea, 0x85, 0xbd, 0x8e, 0x9e, 0x28, 0x84, 0x60, 0x50, 0x91, 0x35, 0x0b, 0x06, 0xe6, 0x48,
	0x53, 0xa3, 0x10, 0x7c, 0xca, 0x64, 0xde, 0xf0, 0x5a, 0x71, 0x51, 0x05, 0x43, 0x63, 0x3d, 0x97,
	0xd0, 0x04, 0xdc, 0x92, 0x2c, 0x59, 0x29, 0x03, 0x37, 0xec, 0x47, 0xfe, 0xf9, 0xdb, 0xf8, 0x85,
	0xc4, 0xb1, 0x0e, 0x13, 0x5f, 0x1b, 0x36, 0xad, 0x54, 0xb3, 0xc1, 0xdd, 0x46, 0x74, 0x00, 0xdb,
	0x6a, 0x53, 0x33, 0x1d, 0x67, 0xdb, 0x34, 0x70, 0xf5, 0x32, 0xa3, 0xda, 0x78, 0x14, 0x95, 0x31,
	0x46, 0xd6, 0xd0, 0xcb, 0x8c, 0xea, 0x51, 0x25, 0x7f, 0x64, 0x81, 0x17, 0x3a, 0x51, 0x1f, 0x9b,
	0x1a, 0xbd, 0x02, 0x58, 0x96, 0x22, 0x5f, 0x2d, 0x8c, 0xb3, 0x6f, 0x1c, 0xcf, 0x28, 0x33, 0x6d,
	0x1f, 0x83, 0x5f, 0x37, 0x82, 0xb6, 0xb9, 0x5a, 0x70, 0x2a, 0x03, 0x08, 0xfb, 0x91, 0x87, 0xa1,
	0x93, 0x32, 0x2a, 0xd1, 0x27, 0x70, 0xa5, 0x22, 0xaa, 0x95, 0x81, 0x1f, 0x3a, 0xd1, 0xde, 0xf9,
	0x9b, 0x7f, 0x07, 0x99, 0x19, 0x16, 0x77, 0x7b, 0x50, 0x04, 0xfb, 0x52, 0xb4, 0x4d, 0xce, 0x16,
	0x7c, 0x4d, 0x0a, 0x33, 0xf2, 0x8e, 0x1e, 0xf9, 0xeb, 0x16, 0xde, 0xb5, 0x46, 0xa6, 0xf5, 0x8c,
	0xa2, 0x18, 0x50, 0x47, 0xca, 0x8a, 0xd4, 0xf2, 0x5e, 0xe8, 0x81, 0x82, 0xdd, 0x0e, 0x1e, 0x5b,
	0x6f, 0xd6, 0x59, 0x19, 0x45, 0xaf, 0x61, 0x87, 0x57, 0x52, 0x91, 0x4a, 0x9f, 0x4d, 0x65, 0xb0,
	0x67, 0x26, 0xf7, 0x9f, 0xb4, 0x8c, 0xca, 0xc3, 0x0b, 0xf0, 0x9f, 0xdd, 0x2b, 0x1a, 0x43, 0x7f,
	0xc5, 0x36, 0xdd, 0x8b, 0xd1, 0x25, 0xfa, 0x0f, 0x86, 0x0f, 0xa4, 0x6c, 0x59, 0xf7, 0x5c, 0xec,
	0xe2, 0x43, 0xef, 0xbd, 0x73, 0x82, 0xc1, 0xb5, 0x49, 0xd0, 0xff, 0x80, 0x66, 0xf3, 0xc9, 0xfc,
	0x76, 0xb6, 0xb8, 0xbd, 0x99, 0x7d, 0x4f, 0xaf, 0xb2, 0xcf, 0x59, 0x3a, 0x1d, 0x6f, 0xa1, 0x1d,
	0x18, 0x5d, 0xe1, 0x74, 0x32, 0xcf, 0x6e, 0xbe, 0x8c, 0x1d, 0xe4, 0xc1, 0x10, 0xa7, 0x93, 0xe9,
	0x8f, 0x71, 0x4f, 0x97, 0x29, 0xc6, 0xdf, 0xf0, 0xb8, 0xaf, 0x99, 0x69, 0x7a, 0x9d, 0x1a, 0x66,
	0x70, 0x39, 0x02, 0xd7, 0xa6, 0xb8, 0x5c, 0xc2, 0xd1, 0x5f, 0x97, 0x48, 0x6a, 0xfe, 0xec, 0x22,
	0x7f, 0x5e, 0x15, 0x5c, 0xdd, 0xb7, 0x4b, 0x2d, 0x25, 0x96, 0x3b, 0xb5, 0x7f, 0xa9, 0x10, 0xa7,
	0x05, 0xab, 0xcc, 0x73, 0x4d, 0x5e, 0xf8, 0x64, 0x1f, 0xbb, 0x72, 0xe9, 0x1a, 0xec, 0xdd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xfe, 0x79, 0xb6, 0x8e, 0x03, 0x00, 0x00,
}
