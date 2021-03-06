// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/his_transfer/his_transfer.proto

package his_transfer

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

type GetHisTransferReq struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName"`
	Direction            string   `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction"`
	StartTime            string   `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime"`
	EndTime              string   `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime"`
	Page                 int32    `protobuf:"varint,6,opt,name=page,proto3" json:"page"`
	Size                 int32    `protobuf:"varint,7,opt,name=size,proto3" json:"size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHisTransferReq) Reset()         { *m = GetHisTransferReq{} }
func (m *GetHisTransferReq) String() string { return proto.CompactTextString(m) }
func (*GetHisTransferReq) ProtoMessage()    {}
func (*GetHisTransferReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_491afa8a17bc335d, []int{0}
}

func (m *GetHisTransferReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHisTransferReq.Unmarshal(m, b)
}
func (m *GetHisTransferReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHisTransferReq.Marshal(b, m, deterministic)
}
func (m *GetHisTransferReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHisTransferReq.Merge(m, src)
}
func (m *GetHisTransferReq) XXX_Size() int {
	return xxx_messageInfo_GetHisTransferReq.Size(m)
}
func (m *GetHisTransferReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHisTransferReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetHisTransferReq proto.InternalMessageInfo

func (m *GetHisTransferReq) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GetHisTransferReq) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *GetHisTransferReq) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *GetHisTransferReq) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *GetHisTransferReq) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *GetHisTransferReq) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetHisTransferReq) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type GetHisTransferResp struct {
	Data                 []*GetHisTransferResp_Data `protobuf:"bytes,1,rep,name=data,proto3" json:"data"`
	Pagination           *base.Pagination           `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *GetHisTransferResp) Reset()         { *m = GetHisTransferResp{} }
func (m *GetHisTransferResp) String() string { return proto.CompactTextString(m) }
func (*GetHisTransferResp) ProtoMessage()    {}
func (*GetHisTransferResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_491afa8a17bc335d, []int{1}
}

func (m *GetHisTransferResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHisTransferResp.Unmarshal(m, b)
}
func (m *GetHisTransferResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHisTransferResp.Marshal(b, m, deterministic)
}
func (m *GetHisTransferResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHisTransferResp.Merge(m, src)
}
func (m *GetHisTransferResp) XXX_Size() int {
	return xxx_messageInfo_GetHisTransferResp.Size(m)
}
func (m *GetHisTransferResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHisTransferResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetHisTransferResp proto.InternalMessageInfo

func (m *GetHisTransferResp) GetData() []*GetHisTransferResp_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetHisTransferResp) GetPagination() *base.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type GetHisTransferResp_Data struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	Direction            int32    `protobuf:"varint,5,opt,name=direction,proto3" json:"direction"`
	Amount               float32  `protobuf:"fixed32,6,opt,name=amount,proto3" json:"amount"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark"`
	Reason               string   `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason"`
	Status               int32    `protobuf:"varint,9,opt,name=status,proto3" json:"status"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHisTransferResp_Data) Reset()         { *m = GetHisTransferResp_Data{} }
func (m *GetHisTransferResp_Data) String() string { return proto.CompactTextString(m) }
func (*GetHisTransferResp_Data) ProtoMessage()    {}
func (*GetHisTransferResp_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_491afa8a17bc335d, []int{1, 0}
}

func (m *GetHisTransferResp_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHisTransferResp_Data.Unmarshal(m, b)
}
func (m *GetHisTransferResp_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHisTransferResp_Data.Marshal(b, m, deterministic)
}
func (m *GetHisTransferResp_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHisTransferResp_Data.Merge(m, src)
}
func (m *GetHisTransferResp_Data) XXX_Size() int {
	return xxx_messageInfo_GetHisTransferResp_Data.Size(m)
}
func (m *GetHisTransferResp_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHisTransferResp_Data.DiscardUnknown(m)
}

var xxx_messageInfo_GetHisTransferResp_Data proto.InternalMessageInfo

func (m *GetHisTransferResp_Data) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *GetHisTransferResp_Data) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *GetHisTransferResp_Data) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *GetHisTransferResp_Data) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GetHisTransferResp_Data) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHisTransferReq)(nil), "his_transfer.GetHisTransferReq")
	proto.RegisterType((*GetHisTransferResp)(nil), "his_transfer.GetHisTransferResp")
	proto.RegisterType((*GetHisTransferResp_Data)(nil), "his_transfer.GetHisTransferResp.Data")
}

func init() {
	proto.RegisterFile("proto/his_transfer/his_transfer.proto", fileDescriptor_491afa8a17bc335d)
}

var fileDescriptor_491afa8a17bc335d = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0xed, 0x4c, 0x67, 0x76, 0xfa, 0x46, 0x44, 0x03, 0x4a, 0x28, 0x82, 0x65, 0x60, 0x61,
	0x2e, 0xb6, 0xcb, 0x78, 0xf2, 0xa8, 0x2b, 0xe8, 0x49, 0x25, 0xac, 0x17, 0x2f, 0xf2, 0xa6, 0xcd,
	0xce, 0x06, 0xb7, 0x49, 0x37, 0x49, 0x0f, 0xeb, 0xd1, 0x83, 0xdf, 0xcb, 0x6f, 0x26, 0x79, 0xed,
	0xd4, 0xd6, 0x05, 0xc5, 0xcb, 0xf0, 0xfe, 0xbf, 0xf7, 0x92, 0x4c, 0x7e, 0xa1, 0x70, 0xda, 0x58,
	0xe3, 0x4d, 0x71, 0xa5, 0xdc, 0x17, 0x6f, 0x51, 0xbb, 0x4b, 0x69, 0x27, 0x21, 0xa7, 0x3e, 0xbb,
	0x3f, 0x66, 0xe9, 0xe3, 0x6e, 0xd1, 0x1e, 0x9d, 0xa4, 0x9f, 0x6e, 0x68, 0xf3, 0x33, 0x82, 0x47,
	0x6f, 0xa5, 0x7f, 0xa7, 0xdc, 0x45, 0x3f, 0x29, 0xe4, 0x0d, 0x63, 0x10, 0xb7, 0xad, 0xaa, 0x78,
	0x94, 0x45, 0xdb, 0x44, 0x50, 0xcd, 0x52, 0x58, 0x5d, 0xb6, 0xd7, 0xd7, 0xef, 0xb1, 0x96, 0x7c,
	0x46, 0x7c, 0xc8, 0xec, 0x29, 0x24, 0x95, 0xb2, 0xb2, 0xf4, 0xca, 0x68, 0x3e, 0xa7, 0xe6, 0x6f,
	0x10, 0xba, 0xce, 0xa3, 0xf5, 0x17, 0xaa, 0x96, 0x3c, 0xee, 0xba, 0x03, 0x60, 0x1c, 0x4e, 0xa4,
	0xae, 0xa8, 0xb7, 0xa0, 0xde, 0x31, 0x86, 0x7f, 0xd1, 0xe0, 0x41, 0xf2, 0x65, 0x16, 0x6d, 0x17,
	0x82, 0xea, 0xc0, 0x9c, 0xfa, 0x26, 0xf9, 0x49, 0xc7, 0x42, 0xbd, 0xf9, 0x31, 0x07, 0xf6, 0xe7,
	0x1d, 0x5c, 0xc3, 0x5e, 0x42, 0x5c, 0xa1, 0x47, 0x1e, 0x65, 0xf3, 0xed, 0x7a, 0x77, 0x9a, 0x4f,
	0x14, 0xdd, 0x9d, 0xcf, 0xdf, 0xa0, 0x47, 0x41, 0x4b, 0xd8, 0x19, 0x40, 0x83, 0x07, 0xa5, 0x91,
	0x2e, 0x14, 0x6e, 0xbb, 0xde, 0x3d, 0xcc, 0x49, 0xdb, 0xc7, 0x81, 0x8b, 0xd1, 0x4c, 0xfa, 0x7d,
	0x06, 0x71, 0xd8, 0xe0, 0xbf, 0xd5, 0xa5, 0xb0, 0x2a, 0x5b, 0x6b, 0xa5, 0x2e, 0x6f, 0x7b, 0x73,
	0x43, 0x0e, 0x7b, 0xf9, 0xdb, 0xe6, 0xe8, 0x8c, 0xea, 0xa9, 0xea, 0x05, 0x59, 0x18, 0xa9, 0x7e,
	0x02, 0x4b, 0xac, 0x4d, 0xab, 0x3d, 0x49, 0x9b, 0x89, 0x3e, 0x05, 0x6e, 0x65, 0x8d, 0xf6, 0x2b,
	0x89, 0x4b, 0x44, 0x9f, 0x3a, 0x8e, 0xce, 0x68, 0xbe, 0x3a, 0xf2, 0x90, 0x02, 0x77, 0x1e, 0x7d,
	0xeb, 0x78, 0x42, 0x47, 0xf4, 0x29, 0x9c, 0x5e, 0x5a, 0x89, 0x5e, 0x56, 0xaf, 0x3c, 0x87, 0xee,
	0x29, 0x07, 0xb0, 0xab, 0x60, 0x3d, 0x92, 0xca, 0x3e, 0xc1, 0x83, 0xa9, 0x66, 0xf6, 0xec, 0xef,
	0x8f, 0x70, 0x93, 0x66, 0xff, 0x7a, 0xa5, 0xcd, 0xbd, 0xd7, 0x67, 0x9f, 0xf3, 0x83, 0xf2, 0x57,
	0xed, 0x3e, 0x2f, 0x4d, 0x5d, 0x9c, 0x8b, 0x0f, 0xe2, 0x5c, 0x14, 0xa5, 0x7d, 0x5e, 0x9a, 0xba,
	0x36, 0xba, 0xb8, 0xfb, 0x71, 0xec, 0x97, 0xc4, 0x5e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xbd,
	0xc6, 0xdc, 0x48, 0x39, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HisTransferClient is the client API for HisTransfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HisTransferClient interface {
	GetHisTransfer(ctx context.Context, in *GetHisTransferReq, opts ...grpc.CallOption) (*GetHisTransferResp, error)
}

type hisTransferClient struct {
	cc *grpc.ClientConn
}

func NewHisTransferClient(cc *grpc.ClientConn) HisTransferClient {
	return &hisTransferClient{cc}
}

func (c *hisTransferClient) GetHisTransfer(ctx context.Context, in *GetHisTransferReq, opts ...grpc.CallOption) (*GetHisTransferResp, error) {
	out := new(GetHisTransferResp)
	err := c.cc.Invoke(ctx, "/his_transfer.HisTransfer/GetHisTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HisTransferServer is the server API for HisTransfer service.
type HisTransferServer interface {
	GetHisTransfer(context.Context, *GetHisTransferReq) (*GetHisTransferResp, error)
}

// UnimplementedHisTransferServer can be embedded to have forward compatible implementations.
type UnimplementedHisTransferServer struct {
}

func (*UnimplementedHisTransferServer) GetHisTransfer(ctx context.Context, req *GetHisTransferReq) (*GetHisTransferResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHisTransfer not implemented")
}

func RegisterHisTransferServer(s *grpc.Server, srv HisTransferServer) {
	s.RegisterService(&_HisTransfer_serviceDesc, srv)
}

func _HisTransfer_GetHisTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHisTransferReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HisTransferServer).GetHisTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/his_transfer.HisTransfer/GetHisTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HisTransferServer).GetHisTransfer(ctx, req.(*GetHisTransferReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HisTransfer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "his_transfer.HisTransfer",
	HandlerType: (*HisTransferServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHisTransfer",
			Handler:    _HisTransfer_GetHisTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/his_transfer/his_transfer.proto",
}
