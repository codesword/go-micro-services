// Code generated by protoc-gen-go.
// source: pb/profile/profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	pb/profile/profile.proto

It has these top-level messages:
	Request
	Result
	Hotel
	Address
	Image
*/
package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	Locale   string   `protobuf:"bytes,2,opt,name=locale" json:"locale,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Hotels []*Hotel `protobuf:"bytes,1,rep,name=hotels" json:"hotels,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber string `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Profile service

type ProfileClient interface {
	GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetProfiles(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/profile.Profile/GetProfiles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	GetProfiles(context.Context, *Request) (*Result, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_GetProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProfileServer).GetProfiles(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProfiles",
			Handler:    _Profile_GetProfiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0xd5, 0x9f, 0x24, 0xed, 0xa4, 0xb7, 0xad, 0x7c, 0x59, 0x58, 0x2c, 0x50, 0x89, 0x84,
	0x54, 0x09, 0xa9, 0xa0, 0xb2, 0x64, 0x85, 0x58, 0x40, 0x37, 0x08, 0xf5, 0x0d, 0xd2, 0x78, 0x4a,
	0x23, 0xb9, 0x71, 0xb0, 0xc7, 0x8b, 0xbe, 0x04, 0xcf, 0x8c, 0xe3, 0x38, 0x14, 0xb1, 0x72, 0xe6,
	0xe8, 0xcc, 0x9c, 0x2f, 0x33, 0xc0, 0xeb, 0xdd, 0x5d, 0xad, 0xd5, 0xbe, 0x94, 0xd8, 0xbd, 0x2b,
	0xf7, 0x92, 0x62, 0x49, 0x28, 0xb3, 0x5b, 0x48, 0xb6, 0xf8, 0x69, 0xd1, 0x10, 0x9b, 0xc3, 0xe8,
	0xa0, 0x08, 0xe5, 0x46, 0x18, 0xde, 0x5b, 0x0c, 0x96, 0x63, 0x36, 0x85, 0x58, 0xaa, 0x22, 0x97,
	0xc8, 0xfb, 0x8b, 0xde, 0x72, 0x9c, 0x2d, 0x21, 0xde, 0xa2, 0xb1, 0x92, 0xd8, 0x15, 0xc4, 0xde,
	0xdb, 0x3a, 0xd3, 0xf5, 0x74, 0xd5, 0xcd, 0x7f, 0x6d, 0xe4, 0xec, 0xab, 0x07, 0x91, 0xff, 0x62,
	0x00, 0xfd, 0x52, 0x38, 0x97, 0xeb, 0x67, 0x13, 0x18, 0x56, 0xf9, 0x31, 0x4c, 0x63, 0xff, 0x21,
	0xad, 0x0f, 0xaa, 0xc2, 0x37, 0x7b, 0xdc, 0xa1, 0xe6, 0x83, 0x4e, 0x14, 0x68, 0x0a, 0x5d, 0xd6,
	0x54, 0xaa, 0x8a, 0x0f, 0xbd, 0x78, 0x0d, 0x49, 0x2e, 0x84, 0x46, 0x63, 0x78, 0xe4, 0x84, 0x74,
	0x3d, 0xff, 0x89, 0x7b, 0x6a, 0xf5, 0x06, 0xa8, 0x3c, 0xe6, 0x1f, 0x68, 0x78, 0xfc, 0x07, 0x68,
	0xd3, 0xc8, 0x99, 0x85, 0xa4, 0xb3, 0x5e, 0xc0, 0xc4, 0x90, 0x46, 0xa4, 0x10, 0xdc, 0xb2, 0x39,
	0xd0, 0xa0, 0x9e, 0x09, 0x1d, 0x6f, 0x51, 0xd2, 0x29, 0xa0, 0xfd, 0x83, 0xc8, 0x50, 0x4e, 0x18,
	0xa0, 0x66, 0x90, 0x14, 0xca, 0x56, 0xa4, 0x4f, 0x1e, 0xca, 0x4f, 0xa8, 0x95, 0x73, 0xc8, 0x67,
	0x25, 0xd0, 0x61, 0x34, 0x1b, 0xbb, 0x81, 0xc8, 0xe7, 0xb3, 0x14, 0x06, 0x56, 0xcb, 0x90, 0xe5,
	0x5a, 0x05, 0xee, 0x73, 0xb7, 0x48, 0x1f, 0x34, 0x5a, 0x3f, 0x42, 0xf2, 0xde, 0xe2, 0xb2, 0x7b,
	0x48, 0x5f, 0x90, 0x42, 0x65, 0xd8, 0xf9, 0x4f, 0xc3, 0x99, 0x2e, 0x67, 0xbf, 0x94, 0xe6, 0x16,
	0xbb, 0xd8, 0x9f, 0xf4, 0xe1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x6b, 0x2a, 0xa1, 0xee, 0x01,
	0x00, 0x00,
}
