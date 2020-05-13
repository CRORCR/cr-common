// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/his_transfer/his_transfer.proto

package his_transfer

import (
	context "context"
	fmt "fmt"
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
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Direction            string   `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Page                 string   `protobuf:"bytes,6,opt,name=page,proto3" json:"page,omitempty"`
	Size                 string   `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
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

func (m *GetHisTransferReq) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *GetHisTransferReq) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

type GetHisTransferResp struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Direction            string   `protobuf:"bytes,5,opt,name=direction,proto3" json:"direction,omitempty"`
	Amount               string   `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	Reason               string   `protobuf:"bytes,8,opt,name=reason,proto3" json:"reason,omitempty"`
	Status               string   `protobuf:"bytes,9,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt            string   `protobuf:"bytes,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *GetHisTransferResp) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GetHisTransferResp) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *GetHisTransferResp) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *GetHisTransferResp) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetHisTransferResp) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *GetHisTransferResp) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *GetHisTransferResp) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *GetHisTransferResp) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *GetHisTransferResp) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetHisTransferResp) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHisTransferReq)(nil), "his_transfer.GetHisTransferReq")
	proto.RegisterType((*GetHisTransferResp)(nil), "his_transfer.GetHisTransferResp")
}

func init() {
	proto.RegisterFile("proto/his_transfer/his_transfer.proto", fileDescriptor_491afa8a17bc335d)
}

var fileDescriptor_491afa8a17bc335d = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xbf, 0xf6, 0xeb, 0xdf, 0x28, 0x82, 0x59, 0x48, 0x28, 0x82, 0xa5, 0x20, 0xb8, 0x71,
	0x46, 0xf4, 0x09, 0xb4, 0x0b, 0x5d, 0x29, 0x0c, 0x75, 0xe3, 0x46, 0xd2, 0x4c, 0xda, 0x06, 0x9b,
	0x64, 0x4c, 0x6e, 0x16, 0xf5, 0x29, 0x7c, 0x25, 0xdf, 0x4c, 0x92, 0xcc, 0x8c, 0x9d, 0x16, 0x04,
	0x77, 0xf7, 0xfc, 0xce, 0x30, 0x39, 0xe7, 0x72, 0xd1, 0x79, 0x61, 0x34, 0xe8, 0x74, 0x25, 0xec,
	0x2b, 0x18, 0xaa, 0xec, 0x82, 0x9b, 0x86, 0x48, 0x82, 0x8f, 0x0f, 0xb7, 0xd9, 0xe4, 0xab, 0x85,
	0x8e, 0xef, 0x39, 0x3c, 0x08, 0x3b, 0x2b, 0x51, 0xc6, 0xdf, 0x31, 0x46, 0x1d, 0xe7, 0x44, 0x4e,
	0x5a, 0xe3, 0xd6, 0xc5, 0x30, 0x0b, 0x33, 0x1e, 0xa1, 0xc1, 0xc2, 0xad, 0xd7, 0x8f, 0x54, 0x72,
	0xd2, 0x0e, 0xbc, 0xd6, 0xf8, 0x14, 0x0d, 0x73, 0x61, 0x38, 0x03, 0xa1, 0x15, 0xf9, 0x1f, 0xcc,
	0x1f, 0xe0, 0x5d, 0x0b, 0xd4, 0xc0, 0x4c, 0x48, 0x4e, 0x3a, 0xd1, 0xad, 0x01, 0x26, 0xa8, 0xcf,
	0x55, 0x1e, 0xbc, 0x6e, 0xf0, 0x2a, 0xe9, 0x53, 0x14, 0x74, 0xc9, 0x49, 0x2f, 0xa6, 0xf0, 0xb3,
	0x67, 0x56, 0x7c, 0x70, 0xd2, 0x8f, 0xcc, 0xcf, 0x93, 0xcf, 0x36, 0xc2, 0xbb, 0x1d, 0x6c, 0xf1,
	0xe7, 0x12, 0x23, 0x34, 0x60, 0xce, 0x18, 0xae, 0xd8, 0xa6, 0xec, 0x50, 0x6b, 0xff, 0x2f, 0xd8,
	0x14, 0x55, 0xfa, 0x30, 0x37, 0x4b, 0x77, 0x77, 0x4b, 0x9f, 0xa0, 0x1e, 0x95, 0xda, 0x29, 0x28,
	0xe3, 0x97, 0xca, 0x73, 0xc3, 0x25, 0x35, 0x6f, 0x65, 0x85, 0x52, 0x45, 0x4e, 0xad, 0x56, 0x64,
	0x50, 0x71, 0xaf, 0x3c, 0xb7, 0x40, 0xc1, 0x59, 0x32, 0x8c, 0x3c, 0x2a, 0xff, 0x3a, 0x33, 0x9c,
	0x02, 0xcf, 0x6f, 0x81, 0xa0, 0xf8, 0x7a, 0x0d, 0xae, 0x73, 0x74, 0xb0, 0xb5, 0x0e, 0xfc, 0x8c,
	0x8e, 0x9a, 0x0b, 0xc2, 0x67, 0x49, 0xe3, 0x34, 0xf6, 0x4e, 0x60, 0x34, 0xfe, 0xfd, 0x03, 0x5b,
	0x4c, 0xfe, 0xdd, 0x5d, 0xbd, 0x24, 0x4b, 0x01, 0x2b, 0x37, 0x4f, 0x98, 0x96, 0xe9, 0x34, 0x7b,
	0xca, 0xa6, 0x59, 0xca, 0xcc, 0x25, 0xd3, 0x52, 0x6a, 0x95, 0xee, 0xdf, 0xe3, 0xbc, 0x17, 0xd8,
	0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0x30, 0xac, 0xa9, 0xac, 0x02, 0x00, 0x00,
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
