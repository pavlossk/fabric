// Code generated by protoc-gen-go.
// source: devops.proto
// DO NOT EDIT!

package protos

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

type BuildResult_StatusCode int32

const (
	BuildResult_UNDEFINED BuildResult_StatusCode = 0
	BuildResult_SUCCESS   BuildResult_StatusCode = 1
	BuildResult_FAILURE   BuildResult_StatusCode = 2
)

var BuildResult_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "SUCCESS",
	2: "FAILURE",
}
var BuildResult_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"SUCCESS":   1,
	"FAILURE":   2,
}

func (x BuildResult_StatusCode) String() string {
	return proto.EnumName(BuildResult_StatusCode_name, int32(x))
}
func (BuildResult_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{4, 0} }

// Secret is a temporary object to establish security with the Devops.
// A better solution using certificate will be introduced later
type Secret struct {
	EnrollId     string `protobuf:"bytes,1,opt,name=enrollId" json:"enrollId,omitempty"`
	EnrollSecret string `protobuf:"bytes,2,opt,name=enrollSecret" json:"enrollSecret,omitempty"`
}

func (m *Secret) Reset()                    { *m = Secret{} }
func (m *Secret) String() string            { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()               {}
func (*Secret) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type SigmaInput struct {
	Secret   *Secret `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
	AppTCert []byte  `protobuf:"bytes,2,opt,name=appTCert,proto3" json:"appTCert,omitempty"`
	Data     []byte  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *SigmaInput) Reset()                    { *m = SigmaInput{} }
func (m *SigmaInput) String() string            { return proto.CompactTextString(m) }
func (*SigmaInput) ProtoMessage()               {}
func (*SigmaInput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SigmaInput) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

type ExecuteWithBinding struct {
	ChaincodeInvocationSpec *ChaincodeInvocationSpec `protobuf:"bytes,1,opt,name=chaincodeInvocationSpec" json:"chaincodeInvocationSpec,omitempty"`
	Binding                 []byte                   `protobuf:"bytes,2,opt,name=binding,proto3" json:"binding,omitempty"`
}

func (m *ExecuteWithBinding) Reset()                    { *m = ExecuteWithBinding{} }
func (m *ExecuteWithBinding) String() string            { return proto.CompactTextString(m) }
func (*ExecuteWithBinding) ProtoMessage()               {}
func (*ExecuteWithBinding) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ExecuteWithBinding) GetChaincodeInvocationSpec() *ChaincodeInvocationSpec {
	if m != nil {
		return m.ChaincodeInvocationSpec
	}
	return nil
}

type SigmaOutput struct {
	Tcert        []byte `protobuf:"bytes,1,opt,name=tcert,proto3" json:"tcert,omitempty"`
	Sigma        []byte `protobuf:"bytes,2,opt,name=sigma,proto3" json:"sigma,omitempty"`
	Asn1Encoding []byte `protobuf:"bytes,3,opt,name=asn1Encoding,proto3" json:"asn1Encoding,omitempty"`
}

func (m *SigmaOutput) Reset()                    { *m = SigmaOutput{} }
func (m *SigmaOutput) String() string            { return proto.CompactTextString(m) }
func (*SigmaOutput) ProtoMessage()               {}
func (*SigmaOutput) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type BuildResult struct {
	Status         BuildResult_StatusCode   `protobuf:"varint,1,opt,name=status,enum=protos.BuildResult_StatusCode" json:"status,omitempty"`
	Msg            string                   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	DeploymentSpec *ChaincodeDeploymentSpec `protobuf:"bytes,3,opt,name=deploymentSpec" json:"deploymentSpec,omitempty"`
}

func (m *BuildResult) Reset()                    { *m = BuildResult{} }
func (m *BuildResult) String() string            { return proto.CompactTextString(m) }
func (*BuildResult) ProtoMessage()               {}
func (*BuildResult) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *BuildResult) GetDeploymentSpec() *ChaincodeDeploymentSpec {
	if m != nil {
		return m.DeploymentSpec
	}
	return nil
}

type TransactionRequest struct {
	TransactionUuid string `protobuf:"bytes,1,opt,name=transactionUuid" json:"transactionUuid,omitempty"`
}

func (m *TransactionRequest) Reset()                    { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()               {}
func (*TransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func init() {
	proto.RegisterType((*Secret)(nil), "protos.Secret")
	proto.RegisterType((*SigmaInput)(nil), "protos.SigmaInput")
	proto.RegisterType((*ExecuteWithBinding)(nil), "protos.ExecuteWithBinding")
	proto.RegisterType((*SigmaOutput)(nil), "protos.SigmaOutput")
	proto.RegisterType((*BuildResult)(nil), "protos.BuildResult")
	proto.RegisterType((*TransactionRequest)(nil), "protos.TransactionRequest")
	proto.RegisterEnum("protos.BuildResult_StatusCode", BuildResult_StatusCode_name, BuildResult_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Devops service

type DevopsClient interface {
	// Log in - passed Secret object and returns Response object, where
	// msg is the security context to be used in subsequent invocations
	Login(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Build the chaincode package.
	Build(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error)
	// Query chaincode.
	Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error)
	// Retrieve a TCert.
	EXP_GetApplicationTCert(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_PrepareForTx(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_ProduceSigma(ctx context.Context, in *SigmaInput, opts ...grpc.CallOption) (*Response, error)
	// Execute a transaction with a specific binding
	EXP_ExecuteWithBinding(ctx context.Context, in *ExecuteWithBinding, opts ...grpc.CallOption) (*Response, error)
}

type devopsClient struct {
	cc *grpc.ClientConn
}

func NewDevopsClient(cc *grpc.ClientConn) DevopsClient {
	return &devopsClient{cc}
}

func (c *devopsClient) Login(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Build(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error) {
	out := new(ChaincodeDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Deploy(ctx context.Context, in *ChaincodeSpec, opts ...grpc.CallOption) (*ChaincodeDeploymentSpec, error) {
	out := new(ChaincodeDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_GetApplicationTCert(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_GetApplicationTCert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_PrepareForTx(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_PrepareForTx", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_ProduceSigma(ctx context.Context, in *SigmaInput, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_ProduceSigma", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) EXP_ExecuteWithBinding(ctx context.Context, in *ExecuteWithBinding, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/protos.Devops/EXP_ExecuteWithBinding", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Devops service

type DevopsServer interface {
	// Log in - passed Secret object and returns Response object, where
	// msg is the security context to be used in subsequent invocations
	Login(context.Context, *Secret) (*Response, error)
	// Build the chaincode package.
	Build(context.Context, *ChaincodeSpec) (*ChaincodeDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(context.Context, *ChaincodeSpec) (*ChaincodeDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(context.Context, *ChaincodeInvocationSpec) (*Response, error)
	// Query chaincode.
	Query(context.Context, *ChaincodeInvocationSpec) (*Response, error)
	// Retrieve a TCert.
	EXP_GetApplicationTCert(context.Context, *Secret) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_PrepareForTx(context.Context, *Secret) (*Response, error)
	// Prepare for performing a TX, which will return a binding that can later be used to sign and then execute a transaction.
	EXP_ProduceSigma(context.Context, *SigmaInput) (*Response, error)
	// Execute a transaction with a specific binding
	EXP_ExecuteWithBinding(context.Context, *ExecuteWithBinding) (*Response, error)
}

func RegisterDevopsServer(s *grpc.Server, srv DevopsServer) {
	s.RegisterService(&_Devops_serviceDesc, srv)
}

func _Devops_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).Login(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).Build(ctx, req.(*ChaincodeSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).Deploy(ctx, req.(*ChaincodeSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).Invoke(ctx, req.(*ChaincodeInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).Query(ctx, req.(*ChaincodeInvocationSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_EXP_GetApplicationTCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).EXP_GetApplicationTCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/EXP_GetApplicationTCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).EXP_GetApplicationTCert(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_EXP_PrepareForTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).EXP_PrepareForTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/EXP_PrepareForTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).EXP_PrepareForTx(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_EXP_ProduceSigma_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigmaInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).EXP_ProduceSigma(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/EXP_ProduceSigma",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).EXP_ProduceSigma(ctx, req.(*SigmaInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devops_EXP_ExecuteWithBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteWithBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevopsServer).EXP_ExecuteWithBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Devops/EXP_ExecuteWithBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevopsServer).EXP_ExecuteWithBinding(ctx, req.(*ExecuteWithBinding))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devops_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Devops",
	HandlerType: (*DevopsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Devops_Login_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Devops_Build_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Devops_Deploy_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Devops_Invoke_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Devops_Query_Handler,
		},
		{
			MethodName: "EXP_GetApplicationTCert",
			Handler:    _Devops_EXP_GetApplicationTCert_Handler,
		},
		{
			MethodName: "EXP_PrepareForTx",
			Handler:    _Devops_EXP_PrepareForTx_Handler,
		},
		{
			MethodName: "EXP_ProduceSigma",
			Handler:    _Devops_EXP_ProduceSigma_Handler,
		},
		{
			MethodName: "EXP_ExecuteWithBinding",
			Handler:    _Devops_EXP_ExecuteWithBinding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor5,
}

func init() { proto.RegisterFile("devops.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0xff, 0x4e, 0x13, 0x41,
	0x10, 0xa6, 0x40, 0x8b, 0x4c, 0x2b, 0x34, 0x13, 0x95, 0xa6, 0x7f, 0xa8, 0xb9, 0x18, 0x42, 0x62,
	0x52, 0x22, 0x46, 0xfe, 0x30, 0x62, 0x42, 0xdb, 0x03, 0x9a, 0x10, 0xc4, 0x3b, 0x1a, 0x7f, 0x24,
	0xc6, 0x5c, 0xef, 0xc6, 0xf6, 0xe2, 0x75, 0x77, 0xdd, 0xdd, 0x23, 0xf0, 0x08, 0xbe, 0x90, 0xaf,
	0xe2, 0xeb, 0x98, 0xdd, 0xed, 0x15, 0x81, 0x12, 0x88, 0xfe, 0xd5, 0x9d, 0x99, 0xef, 0x9b, 0xaf,
	0xdf, 0xdc, 0xec, 0x42, 0x2d, 0xa1, 0x53, 0x2e, 0x54, 0x4b, 0x48, 0xae, 0x39, 0x56, 0xec, 0x8f,
	0x6a, 0xae, 0xc6, 0xa3, 0x28, 0x65, 0x31, 0x4f, 0xc8, 0x15, 0x9a, 0xb5, 0x6f, 0xd1, 0x40, 0xa6,
	0xb1, 0x8b, 0xbc, 0x03, 0xa8, 0x84, 0x14, 0x4b, 0xd2, 0xd8, 0x84, 0x7b, 0xc4, 0x24, 0xcf, 0xb2,
	0x5e, 0xd2, 0x28, 0x3d, 0x2d, 0x6d, 0x2c, 0x07, 0xd3, 0x18, 0x3d, 0xa8, 0xb9, 0xb3, 0xc3, 0x36,
	0xe6, 0x6d, 0xfd, 0x52, 0xce, 0x4b, 0x00, 0xc2, 0x74, 0x38, 0x8e, 0x7a, 0x4c, 0xe4, 0x1a, 0xd7,
	0xa1, 0xa2, 0x1c, 0xd6, 0xf4, 0xaa, 0x6e, 0xad, 0x38, 0x3d, 0xd5, 0x72, 0xe8, 0x60, 0x52, 0x35,
	0xaa, 0x91, 0x10, 0x27, 0x1d, 0x92, 0xae, 0x6b, 0x2d, 0x98, 0xc6, 0x88, 0xb0, 0x98, 0x44, 0x3a,
	0x6a, 0x2c, 0xd8, 0xbc, 0x3d, 0x7b, 0x3f, 0x4b, 0x80, 0xfe, 0x19, 0xc5, 0xb9, 0xa6, 0x0f, 0xa9,
	0x1e, 0xb5, 0x53, 0x96, 0xa4, 0x6c, 0x88, 0x9f, 0x60, 0x6d, 0xea, 0xb3, 0xc7, 0x4e, 0x79, 0x1c,
	0xe9, 0x94, 0xb3, 0x50, 0x50, 0x3c, 0xd1, 0x7f, 0x52, 0xe8, 0x77, 0x66, 0xc3, 0x82, 0x9b, 0xf8,
	0xd8, 0x80, 0xa5, 0x81, 0x53, 0x99, 0xfc, 0xc1, 0x22, 0xf4, 0xbe, 0x40, 0xd5, 0x3a, 0x7e, 0x97,
	0x6b, 0x63, 0xf9, 0x01, 0x94, 0x75, 0x6c, 0x7c, 0x94, 0x2c, 0xcc, 0x05, 0x26, 0xab, 0x0c, 0x68,
	0x42, 0x76, 0x81, 0x19, 0x68, 0xa4, 0xd8, 0x0b, 0xdf, 0x08, 0x9a, 0xce, 0xce, 0xe2, 0xa5, 0x9c,
	0xf7, 0xbb, 0x04, 0xd5, 0x76, 0x9e, 0x66, 0x49, 0x40, 0x2a, 0xcf, 0x34, 0x6e, 0x43, 0x45, 0xe9,
	0x48, 0xe7, 0xca, 0x0a, 0xac, 0x6c, 0x3d, 0x2e, 0x2c, 0xfd, 0x05, 0x6a, 0x85, 0x16, 0xd1, 0xe1,
	0x09, 0x05, 0x13, 0x34, 0xd6, 0x61, 0x61, 0xac, 0x86, 0x93, 0x6f, 0x66, 0x8e, 0xb8, 0x0f, 0x2b,
	0x09, 0x89, 0x8c, 0x9f, 0x8f, 0x89, 0x69, 0x3b, 0xa4, 0x85, 0x1b, 0x86, 0xd4, 0xbd, 0x04, 0x0b,
	0xae, 0xd0, 0xbc, 0x57, 0x00, 0x17, 0x82, 0x78, 0x1f, 0x96, 0xfb, 0x47, 0x5d, 0x7f, 0xaf, 0x77,
	0xe4, 0x77, 0xeb, 0x73, 0x58, 0x85, 0xa5, 0xb0, 0xdf, 0xe9, 0xf8, 0x61, 0x58, 0x2f, 0x99, 0x60,
	0x6f, 0xb7, 0x77, 0xd8, 0x0f, 0xfc, 0xfa, 0xbc, 0xf7, 0x16, 0xf0, 0x44, 0x46, 0x4c, 0x45, 0xb1,
	0x99, 0x72, 0x40, 0x3f, 0x72, 0x52, 0x1a, 0x37, 0x60, 0x55, 0x5f, 0x64, 0xfb, 0x79, 0x5a, 0xec,
	0xe1, 0xd5, 0xf4, 0xd6, 0xaf, 0x45, 0xa8, 0x74, 0xed, 0xb2, 0xe3, 0x73, 0x28, 0x1f, 0xf2, 0x61,
	0xca, 0xf0, 0xca, 0x82, 0x35, 0xeb, 0x45, 0x1c, 0x90, 0x12, 0x9c, 0x29, 0xf2, 0xe6, 0x70, 0x17,
	0xca, 0x76, 0x56, 0xf8, 0xf0, 0x9a, 0x51, 0x63, 0xa7, 0x79, 0x9b, 0x7f, 0x6f, 0x0e, 0xdb, 0x46,
	0xd9, 0xe4, 0xfe, 0xa3, 0xc7, 0x0e, 0x54, 0xcc, 0x8e, 0x7d, 0x27, 0xbc, 0x6d, 0x2b, 0x67, 0xba,
	0x78, 0x03, 0xe5, 0xf7, 0x39, 0xc9, 0xf3, 0x7f, 0x63, 0xef, 0xc0, 0x9a, 0xff, 0xf1, 0xf8, 0xeb,
	0x3e, 0xe9, 0x5d, 0x21, 0xb2, 0xd4, 0xa1, 0xdd, 0x7d, 0xbb, 0xcb, 0x08, 0xb7, 0xa1, 0x6e, 0xe8,
	0xc7, 0x92, 0x44, 0x24, 0x69, 0x8f, 0xcb, 0x93, 0xb3, 0x3b, 0xf1, 0x5e, 0x17, 0x3c, 0x9e, 0xe4,
	0x31, 0xd9, 0x6b, 0x83, 0x38, 0xe5, 0x4d, 0xdf, 0x8d, 0x99, 0xdc, 0x03, 0x78, 0x64, 0xb8, 0x33,
	0xae, 0x7d, 0xb3, 0x40, 0x5f, 0xaf, 0xcd, 0xea, 0xd4, 0x5e, 0xff, 0xfc, 0x6c, 0x98, 0xea, 0x51,
	0x3e, 0x68, 0xc5, 0x7c, 0xbc, 0x39, 0x3a, 0x17, 0x24, 0x33, 0x4a, 0x86, 0x24, 0x37, 0xdd, 0xa3,
	0xb8, 0xe9, 0x28, 0x03, 0xf7, 0x78, 0xbe, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xea, 0xb7, 0x48,
	0x7e, 0x53, 0x05, 0x00, 0x00,
}
