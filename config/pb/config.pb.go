// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: config/pb/config.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Logger_Level int32

const (
	Logger_UNSPECIFIED Logger_Level = 0
	Logger_DEBUG       Logger_Level = 1
	Logger_INFO        Logger_Level = 2
	Logger_WARN        Logger_Level = 3
	Logger_ERROR       Logger_Level = 4
	Logger_PANIC       Logger_Level = 5
	Logger_FATAL       Logger_Level = 6
)

// Enum value maps for Logger_Level.
var (
	Logger_Level_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DEBUG",
		2: "INFO",
		3: "WARN",
		4: "ERROR",
		5: "PANIC",
		6: "FATAL",
	}
	Logger_Level_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DEBUG":       1,
		"INFO":        2,
		"WARN":        3,
		"ERROR":       4,
		"PANIC":       5,
		"FATAL":       6,
	}
)

func (x Logger_Level) Enum() *Logger_Level {
	p := new(Logger_Level)
	*p = x
	return p
}

func (x Logger_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Logger_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_config_pb_config_proto_enumTypes[0].Descriptor()
}

func (Logger_Level) Type() protoreflect.EnumType {
	return &file_config_pb_config_proto_enumTypes[0]
}

func (x Logger_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Logger_Level.Descriptor instead.
func (Logger_Level) EnumDescriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{2, 0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listener *Listener `protobuf:"bytes,1,opt,name=listener,proto3" json:"listener,omitempty"`
	Logger   *Logger   `protobuf:"bytes,2,opt,name=logger,proto3" json:"logger,omitempty"`
	Database *Database `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	// tenants
	Tenants []*TenantInfo `protobuf:"bytes,4,rep,name=tenants,proto3" json:"tenants,omitempty"`
	// auth
	Auth *Auth `protobuf:"bytes,7,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetListener() *Listener {
	if x != nil {
		return x.Listener
	}
	return nil
}

func (x *Config) GetLogger() *Logger {
	if x != nil {
		return x.Logger
	}
	return nil
}

func (x *Config) GetDatabase() *Database {
	if x != nil {
		return x.Database
	}
	return nil
}

func (x *Config) GetTenants() []*TenantInfo {
	if x != nil {
		return x.Tenants
	}
	return nil
}

func (x *Config) GetAuth() *Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

type TenantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenantId     string `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	TenantApiKey string `protobuf:"bytes,2,opt,name=tenant_api_key,json=tenantApiKey,proto3" json:"tenant_api_key,omitempty"`
	ListGameUrl  string `protobuf:"bytes,3,opt,name=list_game_url,json=listGameUrl,proto3" json:"list_game_url,omitempty"`
}

func (x *TenantInfo) Reset() {
	*x = TenantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TenantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantInfo) ProtoMessage() {}

func (x *TenantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantInfo.ProtoReflect.Descriptor instead.
func (*TenantInfo) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{1}
}

func (x *TenantInfo) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *TenantInfo) GetTenantApiKey() string {
	if x != nil {
		return x.TenantApiKey
	}
	return ""
}

func (x *TenantInfo) GetListGameUrl() string {
	if x != nil {
		return x.ListGameUrl
	}
	return ""
}

type Logger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level Logger_Level `protobuf:"varint,1,opt,name=level,proto3,enum=pb.Logger_Level" json:"level,omitempty"`
	// Types that are assignable to Format:
	//
	//	*Logger_Pretty
	Format isLogger_Format `protobuf_oneof:"format"`
}

func (x *Logger) Reset() {
	*x = Logger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logger) ProtoMessage() {}

func (x *Logger) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logger.ProtoReflect.Descriptor instead.
func (*Logger) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{2}
}

func (x *Logger) GetLevel() Logger_Level {
	if x != nil {
		return x.Level
	}
	return Logger_UNSPECIFIED
}

func (m *Logger) GetFormat() isLogger_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (x *Logger) GetPretty() bool {
	if x, ok := x.GetFormat().(*Logger_Pretty); ok {
		return x.Pretty
	}
	return false
}

type isLogger_Format interface {
	isLogger_Format()
}

type Logger_Pretty struct {
	Pretty bool `protobuf:"varint,2,opt,name=pretty,proto3,oneof"`
}

func (*Logger_Pretty) isLogger_Format() {}

type TCPSocket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Secure  bool   `protobuf:"varint,3,opt,name=secure,proto3" json:"secure,omitempty"`
}

func (x *TCPSocket) Reset() {
	*x = TCPSocket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCPSocket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCPSocket) ProtoMessage() {}

func (x *TCPSocket) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCPSocket.ProtoReflect.Descriptor instead.
func (*TCPSocket) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{3}
}

func (x *TCPSocket) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TCPSocket) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *TCPSocket) GetSecure() bool {
	if x != nil {
		return x.Secure
	}
	return false
}

type UnixSocket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *UnixSocket) Reset() {
	*x = UnixSocket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnixSocket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnixSocket) ProtoMessage() {}

func (x *UnixSocket) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnixSocket.ProtoReflect.Descriptor instead.
func (*UnixSocket) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{4}
}

func (x *UnixSocket) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Listener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Socket:
	//
	//	*Listener_Tcp
	//	*Listener_Unix
	Socket isListener_Socket `protobuf_oneof:"socket"`
}

func (x *Listener) Reset() {
	*x = Listener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listener) ProtoMessage() {}

func (x *Listener) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listener.ProtoReflect.Descriptor instead.
func (*Listener) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{5}
}

func (m *Listener) GetSocket() isListener_Socket {
	if m != nil {
		return m.Socket
	}
	return nil
}

func (x *Listener) GetTcp() *TCPSocket {
	if x, ok := x.GetSocket().(*Listener_Tcp); ok {
		return x.Tcp
	}
	return nil
}

func (x *Listener) GetUnix() *UnixSocket {
	if x, ok := x.GetSocket().(*Listener_Unix); ok {
		return x.Unix
	}
	return nil
}

type isListener_Socket interface {
	isListener_Socket()
}

type Listener_Tcp struct {
	Tcp *TCPSocket `protobuf:"bytes,1,opt,name=tcp,proto3,oneof"`
}

type Listener_Unix struct {
	Unix *UnixSocket `protobuf:"bytes,2,opt,name=unix,proto3,oneof"`
}

func (*Listener_Tcp) isListener_Socket() {}

func (*Listener_Unix) isListener_Socket() {}

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"` // Keep it simple first
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{6}
}

func (x *Database) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Database) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Database) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerSigning       *JwtSigning `protobuf:"bytes,1,opt,name=player_signing,json=playerSigning,proto3" json:"player_signing,omitempty"`
	TenantPlayerSigning *JwtSigning `protobuf:"bytes,2,opt,name=tenant_player_signing,json=tenantPlayerSigning,proto3" json:"tenant_player_signing,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{7}
}

func (x *Auth) GetPlayerSigning() *JwtSigning {
	if x != nil {
		return x.PlayerSigning
	}
	return nil
}

func (x *Auth) GetTenantPlayerSigning() *JwtSigning {
	if x != nil {
		return x.TenantPlayerSigning
	}
	return nil
}

type JwtSigning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SigningKey string `protobuf:"bytes,1,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	Audience   string `protobuf:"bytes,2,opt,name=audience,proto3" json:"audience,omitempty"`
	Issuer     string `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
}

func (x *JwtSigning) Reset() {
	*x = JwtSigning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_pb_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JwtSigning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JwtSigning) ProtoMessage() {}

func (x *JwtSigning) ProtoReflect() protoreflect.Message {
	mi := &file_config_pb_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JwtSigning.ProtoReflect.Descriptor instead.
func (*JwtSigning) Descriptor() ([]byte, []int) {
	return file_config_pb_config_proto_rawDescGZIP(), []int{8}
}

func (x *JwtSigning) GetSigningKey() string {
	if x != nil {
		return x.SigningKey
	}
	return ""
}

func (x *JwtSigning) GetAudience() string {
	if x != nil {
		return x.Audience
	}
	return ""
}

func (x *JwtSigning) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

var File_config_pb_config_proto protoreflect.FileDescriptor

var file_config_pb_config_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x28, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x12, 0x1c, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x73, 0x0a, 0x0a, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x22, 0x58, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10,
	0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x06, 0x42, 0x08, 0x0a, 0x06,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x65, 0x0a, 0x09, 0x54, 0x43, 0x50, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0xff, 0xff, 0x03, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x22, 0x29, 0x0a,
	0x0a, 0x55, 0x6e, 0x69, 0x78, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x20, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x62, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x43, 0x50, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x03, 0x74, 0x63, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x78, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x48, 0x00, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x78, 0x42, 0x0d, 0x0a,
	0x06, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0xa4, 0x01, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0xff, 0xff, 0x03, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x0e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x77, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x69, 0x6e, 0x67, 0x52, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x15, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x77, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x52, 0x13, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x7c, 0x0a, 0x0a, 0x4a, 0x77, 0x74, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x20, 0x01, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12,
	0x23, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x75, 0x74, 0x69, 0x6c, 0x75, 0x73, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x62,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_pb_config_proto_rawDescOnce sync.Once
	file_config_pb_config_proto_rawDescData = file_config_pb_config_proto_rawDesc
)

func file_config_pb_config_proto_rawDescGZIP() []byte {
	file_config_pb_config_proto_rawDescOnce.Do(func() {
		file_config_pb_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_pb_config_proto_rawDescData)
	})
	return file_config_pb_config_proto_rawDescData
}

var file_config_pb_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_pb_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_config_pb_config_proto_goTypes = []interface{}{
	(Logger_Level)(0),  // 0: pb.Logger.Level
	(*Config)(nil),     // 1: pb.Config
	(*TenantInfo)(nil), // 2: pb.TenantInfo
	(*Logger)(nil),     // 3: pb.Logger
	(*TCPSocket)(nil),  // 4: pb.TCPSocket
	(*UnixSocket)(nil), // 5: pb.UnixSocket
	(*Listener)(nil),   // 6: pb.Listener
	(*Database)(nil),   // 7: pb.Database
	(*Auth)(nil),       // 8: pb.Auth
	(*JwtSigning)(nil), // 9: pb.JwtSigning
}
var file_config_pb_config_proto_depIdxs = []int32{
	6,  // 0: pb.Config.listener:type_name -> pb.Listener
	3,  // 1: pb.Config.logger:type_name -> pb.Logger
	7,  // 2: pb.Config.database:type_name -> pb.Database
	2,  // 3: pb.Config.tenants:type_name -> pb.TenantInfo
	8,  // 4: pb.Config.auth:type_name -> pb.Auth
	0,  // 5: pb.Logger.level:type_name -> pb.Logger.Level
	4,  // 6: pb.Listener.tcp:type_name -> pb.TCPSocket
	5,  // 7: pb.Listener.unix:type_name -> pb.UnixSocket
	9,  // 8: pb.Auth.player_signing:type_name -> pb.JwtSigning
	9,  // 9: pb.Auth.tenant_player_signing:type_name -> pb.JwtSigning
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_config_pb_config_proto_init() }
func file_config_pb_config_proto_init() {
	if File_config_pb_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_pb_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TenantInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logger); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCPSocket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnixSocket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listener); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_pb_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JwtSigning); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_config_pb_config_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Logger_Pretty)(nil),
	}
	file_config_pb_config_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Listener_Tcp)(nil),
		(*Listener_Unix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_pb_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_pb_config_proto_goTypes,
		DependencyIndexes: file_config_pb_config_proto_depIdxs,
		EnumInfos:         file_config_pb_config_proto_enumTypes,
		MessageInfos:      file_config_pb_config_proto_msgTypes,
	}.Build()
	File_config_pb_config_proto = out.File
	file_config_pb_config_proto_rawDesc = nil
	file_config_pb_config_proto_goTypes = nil
	file_config_pb_config_proto_depIdxs = nil
}
