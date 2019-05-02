// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/image.proto

package containerregistry // import "github.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// An Image resource. For more information, see [Docker image](/docs/cloud/container-registry/docker-image).
type Image struct {
	// Output only. ID of the Docker image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the Docker image.
	// The name is unique within the registry.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Content-addressable identifier of the Docker image.
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Compressed size of the Docker image, specified in bytes.
	CompressedSize int64 `protobuf:"varint,4,opt,name=compressed_size,json=compressedSize,proto3" json:"compressed_size,omitempty"`
	// Configuration of the Docker image.
	Config *Blob `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	// Layers of the Docker image.
	Layers []*Blob `protobuf:"bytes,6,rep,name=layers,proto3" json:"layers,omitempty"`
	// Tags of the Docker image.
	//
	// Each tag is unique within the repository.
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// Output only. Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_image_cab1ec775e9195d2, []int{0}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Image) GetCompressedSize() int64 {
	if m != nil {
		return m.CompressedSize
	}
	return 0
}

func (m *Image) GetConfig() *Blob {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Image) GetLayers() []*Blob {
	if m != nil {
		return m.Layers
	}
	return nil
}

func (m *Image) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Image) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Image)(nil), "yandex.cloud.containerregistry.v1.Image")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/image.proto", fileDescriptor_image_cab1ec775e9195d2)
}

var fileDescriptor_image_cab1ec775e9195d2 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x41, 0x4b, 0x33, 0x31,
	0x14, 0x64, 0x77, 0xdb, 0xfd, 0xbe, 0x46, 0xa8, 0x90, 0x83, 0x84, 0x5e, 0x5c, 0xbd, 0xb4, 0x07,
	0x9b, 0x50, 0x3d, 0x89, 0x07, 0xb1, 0x37, 0xaf, 0xab, 0x20, 0x7a, 0x29, 0xd9, 0xcd, 0x6b, 0x0c,
	0xec, 0x26, 0x25, 0x49, 0x8b, 0xed, 0x7f, 0xf3, 0xbf, 0xc9, 0x26, 0x2d, 0x1e, 0x7a, 0x28, 0xde,
	0xde, 0x1b, 0x66, 0x86, 0x37, 0xf3, 0xd0, 0x74, 0xcb, 0xb5, 0x80, 0x2f, 0x56, 0x37, 0x66, 0x2d,
	0x58, 0x6d, 0xb4, 0xe7, 0x4a, 0x83, 0xb5, 0x20, 0x95, 0xf3, 0x76, 0xcb, 0x36, 0x33, 0xa6, 0x5a,
	0x2e, 0x81, 0xae, 0xac, 0xf1, 0x06, 0x5f, 0x45, 0x3a, 0x0d, 0x74, 0x7a, 0x44, 0xa7, 0x9b, 0xd9,
	0xe8, 0xe6, 0xb4, 0x63, 0xd5, 0x98, 0x2a, 0x1a, 0x8e, 0x2e, 0xa5, 0x31, 0xb2, 0x01, 0x16, 0xb6,
	0x6a, 0xbd, 0x64, 0x5e, 0xb5, 0xe0, 0x3c, 0x6f, 0x57, 0x91, 0x70, 0xfd, 0x9d, 0xa2, 0xfe, 0x73,
	0x77, 0x01, 0x1e, 0xa2, 0x54, 0x09, 0x92, 0x14, 0xc9, 0x64, 0x50, 0xa6, 0x4a, 0x60, 0x8c, 0x7a,
	0x9a, 0xb7, 0x40, 0xd2, 0x80, 0x84, 0x19, 0x5f, 0xa0, 0x5c, 0x28, 0x09, 0xce, 0x93, 0x2c, 0xa0,
	0xfb, 0x0d, 0x8f, 0xd1, 0x79, 0x6d, 0xda, 0x95, 0x05, 0xe7, 0x40, 0x2c, 0x9c, 0xda, 0x01, 0xe9,
	0x15, 0xc9, 0x24, 0x2b, 0x87, 0xbf, 0xf0, 0x8b, 0xda, 0x01, 0x7e, 0x44, 0x79, 0x6d, 0xf4, 0x52,
	0x49, 0xd2, 0x2f, 0x92, 0xc9, 0xd9, 0xed, 0x98, 0x9e, 0x4c, 0x4c, 0xe7, 0x8d, 0xa9, 0xca, 0xbd,
	0xac, 0x33, 0x68, 0xf8, 0x16, 0xac, 0x23, 0x79, 0x91, 0xfd, 0xc9, 0x20, 0xca, 0xba, 0x58, 0x9e,
	0x4b, 0x47, 0xfe, 0x15, 0x59, 0x17, 0xab, 0x9b, 0xf1, 0x3d, 0x42, 0xb5, 0x05, 0xee, 0x41, 0x2c,
	0xb8, 0x27, 0xff, 0xc3, 0x65, 0x23, 0x1a, 0xab, 0xa3, 0x87, 0xea, 0xe8, 0xeb, 0xa1, 0xba, 0x72,
	0xb0, 0x67, 0x3f, 0xf9, 0xf9, 0xfb, 0xc7, 0x9b, 0x54, 0xfe, 0x73, 0x5d, 0xd1, 0xda, 0xb4, 0x2c,
	0xde, 0x32, 0x8d, 0xbf, 0x91, 0x66, 0x2a, 0x41, 0x07, 0x39, 0x3b, 0xf9, 0xb4, 0x87, 0x23, 0xb0,
	0xca, 0x83, 0xf4, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xfc, 0xfe, 0xdb, 0x44, 0x02, 0x00,
	0x00,
}