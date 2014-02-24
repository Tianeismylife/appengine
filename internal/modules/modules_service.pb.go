// Code generated by protoc-gen-go.
// source: google.golang.org/appengine/internal/modules/modules_service.proto
// DO NOT EDIT!

/*
Package modules is a generated protocol buffer package.

It is generated from these files:
	google.golang.org/appengine/internal/modules/modules_service.proto

It has these top-level messages:
	ModulesServiceError
	GetModulesRequest
	GetModulesResponse
	GetVersionsRequest
	GetVersionsResponse
	GetDefaultVersionRequest
	GetDefaultVersionResponse
	GetNumInstancesRequest
	GetNumInstancesResponse
	SetNumInstancesRequest
	SetNumInstancesResponse
	StartModuleRequest
	StartModuleResponse
	StopModuleRequest
	StopModuleResponse
	GetHostnameRequest
	GetHostnameResponse
*/
package modules

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type ModulesServiceError_ErrorCode int32

const (
	ModulesServiceError_OK                ModulesServiceError_ErrorCode = 0
	ModulesServiceError_INVALID_MODULE    ModulesServiceError_ErrorCode = 1
	ModulesServiceError_INVALID_VERSION   ModulesServiceError_ErrorCode = 2
	ModulesServiceError_INVALID_INSTANCES ModulesServiceError_ErrorCode = 3
	ModulesServiceError_TRANSIENT_ERROR   ModulesServiceError_ErrorCode = 4
	ModulesServiceError_UNEXPECTED_STATE  ModulesServiceError_ErrorCode = 5
)

var ModulesServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INVALID_MODULE",
	2: "INVALID_VERSION",
	3: "INVALID_INSTANCES",
	4: "TRANSIENT_ERROR",
	5: "UNEXPECTED_STATE",
}
var ModulesServiceError_ErrorCode_value = map[string]int32{
	"OK":                0,
	"INVALID_MODULE":    1,
	"INVALID_VERSION":   2,
	"INVALID_INSTANCES": 3,
	"TRANSIENT_ERROR":   4,
	"UNEXPECTED_STATE":  5,
}

func (x ModulesServiceError_ErrorCode) Enum() *ModulesServiceError_ErrorCode {
	p := new(ModulesServiceError_ErrorCode)
	*p = x
	return p
}
func (x ModulesServiceError_ErrorCode) String() string {
	return proto.EnumName(ModulesServiceError_ErrorCode_name, int32(x))
}
func (x *ModulesServiceError_ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ModulesServiceError_ErrorCode_value, data, "ModulesServiceError_ErrorCode")
	if err != nil {
		return err
	}
	*x = ModulesServiceError_ErrorCode(value)
	return nil
}

type ModulesServiceError struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ModulesServiceError) Reset()         { *m = ModulesServiceError{} }
func (m *ModulesServiceError) String() string { return proto.CompactTextString(m) }
func (*ModulesServiceError) ProtoMessage()    {}

type GetModulesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetModulesRequest) Reset()         { *m = GetModulesRequest{} }
func (m *GetModulesRequest) String() string { return proto.CompactTextString(m) }
func (*GetModulesRequest) ProtoMessage()    {}

type GetModulesResponse struct {
	Module           []string `protobuf:"bytes,1,rep,name=module" json:"module,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetModulesResponse) Reset()         { *m = GetModulesResponse{} }
func (m *GetModulesResponse) String() string { return proto.CompactTextString(m) }
func (*GetModulesResponse) ProtoMessage()    {}

func (m *GetModulesResponse) GetModule() []string {
	if m != nil {
		return m.Module
	}
	return nil
}

type GetVersionsRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetVersionsRequest) Reset()         { *m = GetVersionsRequest{} }
func (m *GetVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionsRequest) ProtoMessage()    {}

func (m *GetVersionsRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

type GetVersionsResponse struct {
	Version          []string `protobuf:"bytes,1,rep,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetVersionsResponse) Reset()         { *m = GetVersionsResponse{} }
func (m *GetVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionsResponse) ProtoMessage()    {}

func (m *GetVersionsResponse) GetVersion() []string {
	if m != nil {
		return m.Version
	}
	return nil
}

type GetDefaultVersionRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDefaultVersionRequest) Reset()         { *m = GetDefaultVersionRequest{} }
func (m *GetDefaultVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetDefaultVersionRequest) ProtoMessage()    {}

func (m *GetDefaultVersionRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

type GetDefaultVersionResponse struct {
	Version          *string `protobuf:"bytes,1,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetDefaultVersionResponse) Reset()         { *m = GetDefaultVersionResponse{} }
func (m *GetDefaultVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetDefaultVersionResponse) ProtoMessage()    {}

func (m *GetDefaultVersionResponse) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type GetNumInstancesRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version          *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetNumInstancesRequest) Reset()         { *m = GetNumInstancesRequest{} }
func (m *GetNumInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*GetNumInstancesRequest) ProtoMessage()    {}

func (m *GetNumInstancesRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *GetNumInstancesRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type GetNumInstancesResponse struct {
	Instances        *int64 `protobuf:"varint,1,req,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetNumInstancesResponse) Reset()         { *m = GetNumInstancesResponse{} }
func (m *GetNumInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*GetNumInstancesResponse) ProtoMessage()    {}

func (m *GetNumInstancesResponse) GetInstances() int64 {
	if m != nil && m.Instances != nil {
		return *m.Instances
	}
	return 0
}

type SetNumInstancesRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version          *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instances        *int64  `protobuf:"varint,3,req,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SetNumInstancesRequest) Reset()         { *m = SetNumInstancesRequest{} }
func (m *SetNumInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*SetNumInstancesRequest) ProtoMessage()    {}

func (m *SetNumInstancesRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *SetNumInstancesRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *SetNumInstancesRequest) GetInstances() int64 {
	if m != nil && m.Instances != nil {
		return *m.Instances
	}
	return 0
}

type SetNumInstancesResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SetNumInstancesResponse) Reset()         { *m = SetNumInstancesResponse{} }
func (m *SetNumInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*SetNumInstancesResponse) ProtoMessage()    {}

type StartModuleRequest struct {
	Module           *string `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Version          *string `protobuf:"bytes,2,req,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StartModuleRequest) Reset()         { *m = StartModuleRequest{} }
func (m *StartModuleRequest) String() string { return proto.CompactTextString(m) }
func (*StartModuleRequest) ProtoMessage()    {}

func (m *StartModuleRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *StartModuleRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type StartModuleResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StartModuleResponse) Reset()         { *m = StartModuleResponse{} }
func (m *StartModuleResponse) String() string { return proto.CompactTextString(m) }
func (*StartModuleResponse) ProtoMessage()    {}

type StopModuleRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version          *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StopModuleRequest) Reset()         { *m = StopModuleRequest{} }
func (m *StopModuleRequest) String() string { return proto.CompactTextString(m) }
func (*StopModuleRequest) ProtoMessage()    {}

func (m *StopModuleRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *StopModuleRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

type StopModuleResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StopModuleResponse) Reset()         { *m = StopModuleResponse{} }
func (m *StopModuleResponse) String() string { return proto.CompactTextString(m) }
func (*StopModuleResponse) ProtoMessage()    {}

type GetHostnameRequest struct {
	Module           *string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Version          *string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Instance         *string `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetHostnameRequest) Reset()         { *m = GetHostnameRequest{} }
func (m *GetHostnameRequest) String() string { return proto.CompactTextString(m) }
func (*GetHostnameRequest) ProtoMessage()    {}

func (m *GetHostnameRequest) GetModule() string {
	if m != nil && m.Module != nil {
		return *m.Module
	}
	return ""
}

func (m *GetHostnameRequest) GetVersion() string {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return ""
}

func (m *GetHostnameRequest) GetInstance() string {
	if m != nil && m.Instance != nil {
		return *m.Instance
	}
	return ""
}

type GetHostnameResponse struct {
	Hostname         *string `protobuf:"bytes,1,req,name=hostname" json:"hostname,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetHostnameResponse) Reset()         { *m = GetHostnameResponse{} }
func (m *GetHostnameResponse) String() string { return proto.CompactTextString(m) }
func (*GetHostnameResponse) ProtoMessage()    {}

func (m *GetHostnameResponse) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func init() {
	proto.RegisterEnum("appengine.ModulesServiceError_ErrorCode", ModulesServiceError_ErrorCode_name, ModulesServiceError_ErrorCode_value)
}
