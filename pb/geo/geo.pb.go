// Code generated by protoc-gen-go.
// source: pb/geo/geo.proto
// DO NOT EDIT!

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
	pb/geo/geo.proto

It has these top-level messages:
	Request
	Result
*/
package geo

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

// The latitude and longitude of the current location.
type Request struct {
	Lat float32 `protobuf:"fixed32,1,opt,name=lat" json:"lat,omitempty"`
	Lon float32 `protobuf:"fixed32,2,opt,name=lon" json:"lon,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "geo.Request")
	proto.RegisterType((*Result)(nil), "geo.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Geo service

type GeoClient interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type geoClient struct {
	cc *grpc.ClientConn
}

func NewGeoClient(cc *grpc.ClientConn) GeoClient {
	return &geoClient{cc}
}

func (c *geoClient) Nearby(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/geo.Geo/Nearby", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geo service

type GeoServer interface {
	// Finds the hotels contained nearby the current lat/lon.
	Nearby(context.Context, *Request) (*Result, error)
}

func RegisterGeoServer(s *grpc.Server, srv GeoServer) {
	s.RegisterService(&_Geo_serviceDesc, srv)
}

func _Geo_Nearby_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoServer).Nearby(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Geo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geo.Geo",
	HandlerType: (*GeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Nearby",
			Handler:    _Geo_Nearby_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x48, 0xd2, 0x4f,
	0x4f, 0xcd, 0x07, 0x61, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x66, 0x20, 0x53, 0x49, 0x99,
	0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x88, 0x9b, 0x8b, 0x39, 0x27, 0xb1, 0x44,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x09, 0xcc, 0xc9, 0xcf, 0x93, 0x60, 0x02, 0x71, 0x94, 0xa4, 0xb8,
	0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0x04, 0xb8, 0x38, 0x32, 0xf2, 0x4b, 0x52, 0x73,
	0x3c, 0x53, 0x8a, 0x81, 0x0a, 0x99, 0x35, 0x38, 0x8d, 0xb4, 0xb8, 0x98, 0xdd, 0x53, 0xf3, 0x85,
	0x94, 0xb9, 0xd8, 0xfc, 0x52, 0x13, 0x8b, 0x92, 0x2a, 0x85, 0x78, 0xf4, 0x40, 0x56, 0x40, 0x0d,
	0x95, 0xe2, 0x86, 0xf2, 0x40, 0xba, 0x93, 0xd8, 0xc0, 0x16, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x02, 0x2a, 0x8d, 0xcc, 0x8c, 0x00, 0x00, 0x00,
}
