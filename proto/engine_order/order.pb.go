// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/engine_order/order.proto

package his_order

import (
	context "context"
	fmt "fmt"
	base "github.com/CRORCR/cr-common/proto/base"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetOrderReq struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"`
	MakerOrder           string   `protobuf:"bytes,2,opt,name=makerOrder,proto3" json:"makerOrder"`
	TakerOrder           string   `protobuf:"bytes,3,opt,name=takerOrder,proto3" json:"takerOrder"`
	Side                 string   `protobuf:"bytes,4,opt,name=side,proto3" json:"side"`
	StartTime            string   `protobuf:"bytes,5,opt,name=startTime,proto3" json:"startTime"`
	EndTime              string   `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime"`
	Symbol               string   `protobuf:"bytes,7,opt,name=symbol,proto3" json:"symbol"`
	Page                 int32    `protobuf:"varint,8,opt,name=Page,proto3" json:"Page"`
	Size                 int32    `protobuf:"varint,9,opt,name=Size,proto3" json:"Size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderReq) Reset()         { *m = GetOrderReq{} }
func (m *GetOrderReq) String() string { return proto.CompactTextString(m) }
func (*GetOrderReq) ProtoMessage()    {}
func (*GetOrderReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2c507e7584a5c5, []int{0}
}

func (m *GetOrderReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderReq.Unmarshal(m, b)
}
func (m *GetOrderReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderReq.Marshal(b, m, deterministic)
}
func (m *GetOrderReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderReq.Merge(m, src)
}
func (m *GetOrderReq) XXX_Size() int {
	return xxx_messageInfo_GetOrderReq.Size(m)
}
func (m *GetOrderReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderReq proto.InternalMessageInfo

func (m *GetOrderReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GetOrderReq) GetMakerOrder() string {
	if m != nil {
		return m.MakerOrder
	}
	return ""
}

func (m *GetOrderReq) GetTakerOrder() string {
	if m != nil {
		return m.TakerOrder
	}
	return ""
}

func (m *GetOrderReq) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *GetOrderReq) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *GetOrderReq) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *GetOrderReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetOrderReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetOrderReq) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type GetOrderResp struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Timestamp            string           `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp"`
	MakerOrderID         string           `protobuf:"bytes,3,opt,name=makerOrderID,proto3" json:"makerOrderID"`
	TakerOrderID         string           `protobuf:"bytes,4,opt,name=takerOrderID,proto3" json:"takerOrderID"`
	Side                 string           `protobuf:"bytes,5,opt,name=side,proto3" json:"side"`
	Size                 float32          `protobuf:"fixed32,6,opt,name=size,proto3" json:"size"`
	Price                float32          `protobuf:"fixed32,7,opt,name=price,proto3" json:"price"`
	Symbol               string           `protobuf:"bytes,8,opt,name=symbol,proto3" json:"symbol"`
	MakerFeeTotal        float32          `protobuf:"fixed32,9,opt,name=makerFeeTotal,proto3" json:"makerFeeTotal"`
	TakerFeeTotal        float32          `protobuf:"fixed32,10,opt,name=takerFeeTotal,proto3" json:"takerFeeTotal"`
	MakerFeeCurrency     string           `protobuf:"bytes,11,opt,name=makerFeeCurrency,proto3" json:"makerFeeCurrency"`
	TakerFeeCurrency     string           `protobuf:"bytes,12,opt,name=takerFeeCurrency,proto3" json:"takerFeeCurrency"`
	MakerAccount         string           `protobuf:"bytes,13,opt,name=makerAccount,proto3" json:"makerAccount"`
	TakerAccount         string           `protobuf:"bytes,14,opt,name=takerAccount,proto3" json:"takerAccount"`
	Pag                  *base.Pagination `protobuf:"bytes,15,opt,name=pag,proto3" json:"pag"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetOrderResp) Reset()         { *m = GetOrderResp{} }
func (m *GetOrderResp) String() string { return proto.CompactTextString(m) }
func (*GetOrderResp) ProtoMessage()    {}
func (*GetOrderResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d2c507e7584a5c5, []int{1}
}

func (m *GetOrderResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderResp.Unmarshal(m, b)
}
func (m *GetOrderResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderResp.Marshal(b, m, deterministic)
}
func (m *GetOrderResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderResp.Merge(m, src)
}
func (m *GetOrderResp) XXX_Size() int {
	return xxx_messageInfo_GetOrderResp.Size(m)
}
func (m *GetOrderResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderResp proto.InternalMessageInfo

func (m *GetOrderResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetOrderResp) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *GetOrderResp) GetMakerOrderID() string {
	if m != nil {
		return m.MakerOrderID
	}
	return ""
}

func (m *GetOrderResp) GetTakerOrderID() string {
	if m != nil {
		return m.TakerOrderID
	}
	return ""
}

func (m *GetOrderResp) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *GetOrderResp) GetSize() float32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetOrderResp) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *GetOrderResp) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetOrderResp) GetMakerFeeTotal() float32 {
	if m != nil {
		return m.MakerFeeTotal
	}
	return 0
}

func (m *GetOrderResp) GetTakerFeeTotal() float32 {
	if m != nil {
		return m.TakerFeeTotal
	}
	return 0
}

func (m *GetOrderResp) GetMakerFeeCurrency() string {
	if m != nil {
		return m.MakerFeeCurrency
	}
	return ""
}

func (m *GetOrderResp) GetTakerFeeCurrency() string {
	if m != nil {
		return m.TakerFeeCurrency
	}
	return ""
}

func (m *GetOrderResp) GetMakerAccount() string {
	if m != nil {
		return m.MakerAccount
	}
	return ""
}

func (m *GetOrderResp) GetTakerAccount() string {
	if m != nil {
		return m.TakerAccount
	}
	return ""
}

func (m *GetOrderResp) GetPag() *base.Pagination {
	if m != nil {
		return m.Pag
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrderReq)(nil), "his_order.GetOrderReq")
	proto.RegisterType((*GetOrderResp)(nil), "his_order.GetOrderResp")
}

func init() { proto.RegisterFile("proto/engine_order/order.proto", fileDescriptor_5d2c507e7584a5c5) }

var fileDescriptor_5d2c507e7584a5c5 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x49, 0xb6, 0x76, 0xeb, 0x6d, 0x37, 0x26, 0x0b, 0x86, 0x35, 0xa1, 0xa9, 0xaa, 0x78,
	0xa8, 0x40, 0x4b, 0xa4, 0xf1, 0x05, 0x80, 0x22, 0xfe, 0x48, 0x88, 0x4e, 0x66, 0x4f, 0xbc, 0x20,
	0x37, 0xb9, 0xca, 0x2c, 0x96, 0x38, 0xd8, 0xce, 0xc3, 0xf6, 0x91, 0x91, 0xf8, 0x0e, 0xc8, 0xd7,
	0x6d, 0xfe, 0x0c, 0x5e, 0xa2, 0xeb, 0xdf, 0x39, 0x49, 0x7c, 0x8e, 0x6c, 0x38, 0xaf, 0x8d, 0x76,
	0x3a, 0xc5, 0xaa, 0x50, 0x15, 0xfe, 0xd0, 0x26, 0x47, 0x93, 0xd2, 0x33, 0x21, 0x81, 0x4d, 0x6e,
	0x94, 0x0d, 0xf8, 0xec, 0x69, 0xb0, 0x6e, 0xa4, 0x45, 0x7a, 0x04, 0xc7, 0xe2, 0x4f, 0x04, 0xd3,
	0x8f, 0xe8, 0xd6, 0xde, 0x23, 0xf0, 0x17, 0x63, 0xb0, 0xdf, 0x34, 0x2a, 0xe7, 0xd1, 0x3c, 0x5a,
	0x4e, 0x04, 0xcd, 0xec, 0x1c, 0xa0, 0x94, 0x3f, 0xd1, 0x90, 0x89, 0xc7, 0xa4, 0xf4, 0x88, 0xd7,
	0x5d, 0xa7, 0xef, 0x05, 0xbd, 0x23, 0xfe, 0x9b, 0x56, 0xe5, 0xc8, 0xf7, 0xc3, 0x37, 0xfd, 0xcc,
	0x9e, 0xc3, 0xc4, 0x3a, 0x69, 0xdc, 0xb5, 0x2a, 0x91, 0x8f, 0x48, 0xe8, 0x00, 0xe3, 0x70, 0x80,
	0x55, 0x4e, 0xda, 0x98, 0xb4, 0xdd, 0x92, 0x9d, 0xc2, 0xd8, 0xde, 0x95, 0x1b, 0x7d, 0xcb, 0x0f,
	0x48, 0xd8, 0xae, 0xfc, 0x3f, 0xae, 0x64, 0x81, 0xfc, 0x70, 0x1e, 0x2d, 0x47, 0x82, 0x66, 0xcf,
	0xbe, 0xa9, 0x7b, 0xe4, 0x93, 0xc0, 0xfc, 0xbc, 0xf8, 0xbd, 0x07, 0xb3, 0x2e, 0xaf, 0xad, 0xd9,
	0x31, 0xc4, 0x6d, 0xdc, 0x58, 0xe5, 0x7e, 0x63, 0x4e, 0x95, 0x68, 0x9d, 0x2c, 0xeb, 0x6d, 0xd6,
	0x0e, 0xb0, 0x05, 0xcc, 0xba, 0xe0, 0x9f, 0xdf, 0x6f, 0xc3, 0x0e, 0x98, 0xf7, 0xb8, 0xbe, 0x27,
	0xc4, 0x1e, 0xb0, 0xb6, 0x92, 0x51, 0xaf, 0x12, 0x62, 0xf7, 0x21, 0x71, 0x2c, 0x68, 0x66, 0x4f,
	0x60, 0x54, 0x1b, 0x95, 0x21, 0xa5, 0x8d, 0x45, 0x58, 0xf4, 0x4a, 0x38, 0x1c, 0x94, 0xf0, 0x02,
	0x8e, 0x68, 0x27, 0x1f, 0x10, 0xaf, 0xb5, 0x93, 0xb7, 0x94, 0x3c, 0x16, 0x43, 0xe8, 0x5d, 0x6e,
	0xe0, 0x82, 0xe0, 0x1a, 0x40, 0xf6, 0x12, 0x4e, 0x76, 0xaf, 0xad, 0x1a, 0x63, 0xb0, 0xca, 0xee,
	0xf8, 0x94, 0xfe, 0xf6, 0x0f, 0xf7, 0x5e, 0xf7, 0xd0, 0x3b, 0x0b, 0xde, 0x87, 0xbc, 0x6d, 0xf0,
	0x6d, 0x96, 0xe9, 0xa6, 0x72, 0xfc, 0xa8, 0xd7, 0xe0, 0x96, 0xb5, 0x0d, 0xee, 0x3c, 0xc7, 0xbd,
	0x06, 0x3b, 0xcf, 0x5e, 0x2d, 0x0b, 0xfe, 0x78, 0x1e, 0x2d, 0xa7, 0x97, 0x27, 0x09, 0x1d, 0xe9,
	0x2b, 0x59, 0xa8, 0x4a, 0x3a, 0xa5, 0x2b, 0xe1, 0xc5, 0xcb, 0xaf, 0x00, 0x9f, 0x94, 0x5d, 0xbb,
	0xec, 0x8b, 0xb2, 0x8e, 0xbd, 0xa1, 0x93, 0xee, 0x01, 0x9d, 0xca, 0xd3, 0xa4, 0xbd, 0x1c, 0x49,
	0xef, 0x06, 0x9c, 0x3d, 0xfb, 0x2f, 0xb7, 0xf5, 0xe2, 0xd1, 0xbb, 0x8b, 0xef, 0xaf, 0x0a, 0xe5,
	0x6e, 0x9a, 0x4d, 0x92, 0xe9, 0x32, 0x5d, 0x89, 0xb5, 0x58, 0x89, 0x34, 0x33, 0x17, 0x99, 0x2e,
	0x4b, 0x5d, 0xa5, 0xe1, 0x86, 0xb5, 0x6f, 0x6f, 0xc6, 0x04, 0x5e, 0xff, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xce, 0xf8, 0x94, 0x6c, 0xa6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HisOtcListClient is the client API for HisOtcList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HisOtcListClient interface {
	GetHisOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error)
}

type hisOtcListClient struct {
	cc *grpc.ClientConn
}

func NewHisOtcListClient(cc *grpc.ClientConn) HisOtcListClient {
	return &hisOtcListClient{cc}
}

func (c *hisOtcListClient) GetHisOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderResp, error) {
	out := new(GetOrderResp)
	err := c.cc.Invoke(ctx, "/his_order.HisOtcList/GetHisOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HisOtcListServer is the server API for HisOtcList service.
type HisOtcListServer interface {
	GetHisOrder(context.Context, *GetOrderReq) (*GetOrderResp, error)
}

// UnimplementedHisOtcListServer can be embedded to have forward compatible implementations.
type UnimplementedHisOtcListServer struct {
}

func (*UnimplementedHisOtcListServer) GetHisOrder(ctx context.Context, req *GetOrderReq) (*GetOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHisOrder not implemented")
}

func RegisterHisOtcListServer(s *grpc.Server, srv HisOtcListServer) {
	s.RegisterService(&_HisOtcList_serviceDesc, srv)
}

func _HisOtcList_GetHisOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HisOtcListServer).GetHisOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/his_order.HisOtcList/GetHisOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HisOtcListServer).GetHisOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HisOtcList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "his_order.HisOtcList",
	HandlerType: (*HisOtcListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHisOrder",
			Handler:    _HisOtcList_GetHisOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/engine_order/order.proto",
}
