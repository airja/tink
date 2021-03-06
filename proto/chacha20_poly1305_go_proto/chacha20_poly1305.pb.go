// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: third_party/tink/proto/chacha20_poly1305.proto

package chacha20_poly1305_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// key_type: type.googleapis.com/google.crypto.tink.ChaCha20Poly1305.
// This key type actually implements ChaCha20Poly1305 as described
// at https://tools.ietf.org/html/rfc7539#section-2.8.
type ChaCha20Poly1305Key struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	KeyValue             []byte   `protobuf:"bytes,2,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChaCha20Poly1305Key) Reset()         { *m = ChaCha20Poly1305Key{} }
func (m *ChaCha20Poly1305Key) String() string { return proto.CompactTextString(m) }
func (*ChaCha20Poly1305Key) ProtoMessage()    {}
func (*ChaCha20Poly1305Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_ddbcaacfeec7ca2e, []int{0}
}

func (m *ChaCha20Poly1305Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaCha20Poly1305Key.Unmarshal(m, b)
}
func (m *ChaCha20Poly1305Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaCha20Poly1305Key.Marshal(b, m, deterministic)
}
func (m *ChaCha20Poly1305Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaCha20Poly1305Key.Merge(m, src)
}
func (m *ChaCha20Poly1305Key) XXX_Size() int {
	return xxx_messageInfo_ChaCha20Poly1305Key.Size(m)
}
func (m *ChaCha20Poly1305Key) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaCha20Poly1305Key.DiscardUnknown(m)
}

var xxx_messageInfo_ChaCha20Poly1305Key proto.InternalMessageInfo

func (m *ChaCha20Poly1305Key) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ChaCha20Poly1305Key) GetKeyValue() []byte {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaCha20Poly1305Key)(nil), "google.crypto.tink.ChaCha20Poly1305Key")
}

func init() {
	proto.RegisterFile("proto/chacha20_poly1305.proto", fileDescriptor_ddbcaacfeec7ca2e)
}

var fileDescriptor_ddbcaacfeec7ca2e = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2b, 0xc9, 0xc8, 0x2c,
	0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4, 0x2f, 0xc9, 0xcc, 0xcb, 0xd6, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0x48, 0x4c, 0xce, 0x48, 0x34, 0x32, 0x88, 0x2f, 0xc8, 0xcf, 0xa9,
	0x34, 0x34, 0x36, 0x30, 0xd5, 0x03, 0x8b, 0x0b, 0x09, 0xa5, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea,
	0x25, 0x17, 0x55, 0x16, 0x94, 0xe4, 0xeb, 0x81, 0x74, 0x28, 0xf9, 0x70, 0x09, 0x3b, 0x67, 0x24,
	0x3a, 0x83, 0x94, 0x07, 0x40, 0x55, 0x7b, 0xa7, 0x56, 0x0a, 0x49, 0x70, 0xb1, 0x97, 0xa5, 0x16,
	0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0xc1, 0xb8, 0x42, 0xd2, 0x5c,
	0x9c, 0xd9, 0xa9, 0x95, 0xf1, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x1c, 0xd9, 0xa9, 0x95, 0x61, 0x20, 0xbe, 0x53, 0x12, 0x97, 0x4c, 0x72, 0x7e, 0xae, 0x1e,
	0xa6, 0x3d, 0x10, 0x17, 0x04, 0x30, 0x46, 0x99, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x43, 0x94, 0xe1, 0x75, 0x79, 0x7c, 0x7a, 0x7e, 0x3c, 0x58, 0x6a, 0x11, 0x13,
	0x5b, 0x88, 0xa7, 0x9f, 0x77, 0x80, 0x53, 0x12, 0x1b, 0x98, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xa8, 0x37, 0x79, 0xfe, 0x00, 0x00, 0x00,
}
