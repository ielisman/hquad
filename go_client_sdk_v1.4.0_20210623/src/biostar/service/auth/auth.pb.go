// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import (
	err "biostar/service/err"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type AuthMode int32

const (
	AuthMode_AUTH_MODE_BIOMETRIC_ONLY        AuthMode = 0
	AuthMode_AUTH_MODE_BIOMETRIC_PIN         AuthMode = 1
	AuthMode_AUTH_MODE_CARD_ONLY             AuthMode = 2
	AuthMode_AUTH_MODE_CARD_BIOMETRIC        AuthMode = 3
	AuthMode_AUTH_MODE_CARD_PIN              AuthMode = 4
	AuthMode_AUTH_MODE_CARD_BIOMETRIC_OR_PIN AuthMode = 5
	AuthMode_AUTH_MODE_CARD_BIOMETRIC_PIN    AuthMode = 6
	AuthMode_AUTH_MODE_ID_BIOMETRIC          AuthMode = 7
	AuthMode_AUTH_MODE_ID_PIN                AuthMode = 8
	AuthMode_AUTH_MODE_ID_BIOMETRIC_OR_PIN   AuthMode = 9
	AuthMode_AUTH_MODE_ID_BIOMETRIC_PIN      AuthMode = 10
	AuthMode_AUTH_MODE_NONE                  AuthMode = 255
	AuthMode_AUTH_MODE_PROHIBITED            AuthMode = 254
	// for F2 only
	AuthMode_AUTH_EXT_MODE_FACE_ONLY                       AuthMode = 11
	AuthMode_AUTH_EXT_MODE_FACE_FINGERPRINT                AuthMode = 12
	AuthMode_AUTH_EXT_MODE_FACE_PIN                        AuthMode = 13
	AuthMode_AUTH_EXT_MODE_FACE_FINGERPRINT_OR_PIN         AuthMode = 14
	AuthMode_AUTH_EXT_MODE_FACE_FINGERPRINT_PIN            AuthMode = 15
	AuthMode_AUTH_EXT_MODE_FINGERPRINT_ONLY                AuthMode = 16
	AuthMode_AUTH_EXT_MODE_FINGERPRINT_FACE                AuthMode = 17
	AuthMode_AUTH_EXT_MODE_FINGERPRINT_PIN                 AuthMode = 18
	AuthMode_AUTH_EXT_MODE_FINGERPRINT_FACE_OR_PIN         AuthMode = 19
	AuthMode_AUTH_EXT_MODE_FINGERPRINT_FACE_PIN            AuthMode = 20
	AuthMode_AUTH_EXT_MODE_CARD_ONLY                       AuthMode = 21
	AuthMode_AUTH_EXT_MODE_CARD_FACE                       AuthMode = 22
	AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT                AuthMode = 23
	AuthMode_AUTH_EXT_MODE_CARD_PIN                        AuthMode = 24
	AuthMode_AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT        AuthMode = 25
	AuthMode_AUTH_EXT_MODE_CARD_FACE_OR_PIN                AuthMode = 26
	AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT_OR_PIN         AuthMode = 27
	AuthMode_AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_OR_PIN AuthMode = 28
	AuthMode_AUTH_EXT_MODE_CARD_FACE_FINGERPRINT           AuthMode = 29
	AuthMode_AUTH_EXT_MODE_CARD_FACE_PIN                   AuthMode = 30
	AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT_FACE           AuthMode = 31
	AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT_PIN            AuthMode = 32
	AuthMode_AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_PIN    AuthMode = 33
	AuthMode_AUTH_EXT_MODE_CARD_FACE_FINGERPRINT_OR_PIN    AuthMode = 34
	AuthMode_AUTH_EXT_MODE_CARD_FINGERPRINT_FACE_OR_PIN    AuthMode = 35
	AuthMode_AUTH_EXT_MODE_ID_FACE                         AuthMode = 36
	AuthMode_AUTH_EXT_MODE_ID_FINGERPRINT                  AuthMode = 37
	AuthMode_AUTH_EXT_MODE_ID_PIN                          AuthMode = 38
	AuthMode_AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT          AuthMode = 39
	AuthMode_AUTH_EXT_MODE_ID_FACE_OR_PIN                  AuthMode = 40
	AuthMode_AUTH_EXT_MODE_ID_FINGERPRINT_OR_PIN           AuthMode = 41
	AuthMode_AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_OR_PIN   AuthMode = 42
	AuthMode_AUTH_EXT_MODE_ID_FACE_FINGERPRINT             AuthMode = 43
	AuthMode_AUTH_EXT_MODE_ID_FACE_PIN                     AuthMode = 44
	AuthMode_AUTH_EXT_MODE_ID_FINGERPRINT_FACE             AuthMode = 45
	AuthMode_AUTH_EXT_MODE_ID_FINGERPRINT_PIN              AuthMode = 46
	AuthMode_AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_PIN      AuthMode = 47
	AuthMode_AUTH_EXT_MODE_ID_FACE_FINGERPRINT_OR_PIN      AuthMode = 48
	AuthMode_AUTH_EXT_MODE_ID_FINGERPRINT_FACE_OR_PIN      AuthMode = 49
)

var AuthMode_name = map[int32]string{
	0:   "AUTH_MODE_BIOMETRIC_ONLY",
	1:   "AUTH_MODE_BIOMETRIC_PIN",
	2:   "AUTH_MODE_CARD_ONLY",
	3:   "AUTH_MODE_CARD_BIOMETRIC",
	4:   "AUTH_MODE_CARD_PIN",
	5:   "AUTH_MODE_CARD_BIOMETRIC_OR_PIN",
	6:   "AUTH_MODE_CARD_BIOMETRIC_PIN",
	7:   "AUTH_MODE_ID_BIOMETRIC",
	8:   "AUTH_MODE_ID_PIN",
	9:   "AUTH_MODE_ID_BIOMETRIC_OR_PIN",
	10:  "AUTH_MODE_ID_BIOMETRIC_PIN",
	255: "AUTH_MODE_NONE",
	254: "AUTH_MODE_PROHIBITED",
	11:  "AUTH_EXT_MODE_FACE_ONLY",
	12:  "AUTH_EXT_MODE_FACE_FINGERPRINT",
	13:  "AUTH_EXT_MODE_FACE_PIN",
	14:  "AUTH_EXT_MODE_FACE_FINGERPRINT_OR_PIN",
	15:  "AUTH_EXT_MODE_FACE_FINGERPRINT_PIN",
	16:  "AUTH_EXT_MODE_FINGERPRINT_ONLY",
	17:  "AUTH_EXT_MODE_FINGERPRINT_FACE",
	18:  "AUTH_EXT_MODE_FINGERPRINT_PIN",
	19:  "AUTH_EXT_MODE_FINGERPRINT_FACE_OR_PIN",
	20:  "AUTH_EXT_MODE_FINGERPRINT_FACE_PIN",
	21:  "AUTH_EXT_MODE_CARD_ONLY",
	22:  "AUTH_EXT_MODE_CARD_FACE",
	23:  "AUTH_EXT_MODE_CARD_FINGERPRINT",
	24:  "AUTH_EXT_MODE_CARD_PIN",
	25:  "AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT",
	26:  "AUTH_EXT_MODE_CARD_FACE_OR_PIN",
	27:  "AUTH_EXT_MODE_CARD_FINGERPRINT_OR_PIN",
	28:  "AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_OR_PIN",
	29:  "AUTH_EXT_MODE_CARD_FACE_FINGERPRINT",
	30:  "AUTH_EXT_MODE_CARD_FACE_PIN",
	31:  "AUTH_EXT_MODE_CARD_FINGERPRINT_FACE",
	32:  "AUTH_EXT_MODE_CARD_FINGERPRINT_PIN",
	33:  "AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_PIN",
	34:  "AUTH_EXT_MODE_CARD_FACE_FINGERPRINT_OR_PIN",
	35:  "AUTH_EXT_MODE_CARD_FINGERPRINT_FACE_OR_PIN",
	36:  "AUTH_EXT_MODE_ID_FACE",
	37:  "AUTH_EXT_MODE_ID_FINGERPRINT",
	38:  "AUTH_EXT_MODE_ID_PIN",
	39:  "AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT",
	40:  "AUTH_EXT_MODE_ID_FACE_OR_PIN",
	41:  "AUTH_EXT_MODE_ID_FINGERPRINT_OR_PIN",
	42:  "AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_OR_PIN",
	43:  "AUTH_EXT_MODE_ID_FACE_FINGERPRINT",
	44:  "AUTH_EXT_MODE_ID_FACE_PIN",
	45:  "AUTH_EXT_MODE_ID_FINGERPRINT_FACE",
	46:  "AUTH_EXT_MODE_ID_FINGERPRINT_PIN",
	47:  "AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_PIN",
	48:  "AUTH_EXT_MODE_ID_FACE_FINGERPRINT_OR_PIN",
	49:  "AUTH_EXT_MODE_ID_FINGERPRINT_FACE_OR_PIN",
}

var AuthMode_value = map[string]int32{
	"AUTH_MODE_BIOMETRIC_ONLY":                      0,
	"AUTH_MODE_BIOMETRIC_PIN":                       1,
	"AUTH_MODE_CARD_ONLY":                           2,
	"AUTH_MODE_CARD_BIOMETRIC":                      3,
	"AUTH_MODE_CARD_PIN":                            4,
	"AUTH_MODE_CARD_BIOMETRIC_OR_PIN":               5,
	"AUTH_MODE_CARD_BIOMETRIC_PIN":                  6,
	"AUTH_MODE_ID_BIOMETRIC":                        7,
	"AUTH_MODE_ID_PIN":                              8,
	"AUTH_MODE_ID_BIOMETRIC_OR_PIN":                 9,
	"AUTH_MODE_ID_BIOMETRIC_PIN":                    10,
	"AUTH_MODE_NONE":                                255,
	"AUTH_MODE_PROHIBITED":                          254,
	"AUTH_EXT_MODE_FACE_ONLY":                       11,
	"AUTH_EXT_MODE_FACE_FINGERPRINT":                12,
	"AUTH_EXT_MODE_FACE_PIN":                        13,
	"AUTH_EXT_MODE_FACE_FINGERPRINT_OR_PIN":         14,
	"AUTH_EXT_MODE_FACE_FINGERPRINT_PIN":            15,
	"AUTH_EXT_MODE_FINGERPRINT_ONLY":                16,
	"AUTH_EXT_MODE_FINGERPRINT_FACE":                17,
	"AUTH_EXT_MODE_FINGERPRINT_PIN":                 18,
	"AUTH_EXT_MODE_FINGERPRINT_FACE_OR_PIN":         19,
	"AUTH_EXT_MODE_FINGERPRINT_FACE_PIN":            20,
	"AUTH_EXT_MODE_CARD_ONLY":                       21,
	"AUTH_EXT_MODE_CARD_FACE":                       22,
	"AUTH_EXT_MODE_CARD_FINGERPRINT":                23,
	"AUTH_EXT_MODE_CARD_PIN":                        24,
	"AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT":        25,
	"AUTH_EXT_MODE_CARD_FACE_OR_PIN":                26,
	"AUTH_EXT_MODE_CARD_FINGERPRINT_OR_PIN":         27,
	"AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_OR_PIN": 28,
	"AUTH_EXT_MODE_CARD_FACE_FINGERPRINT":           29,
	"AUTH_EXT_MODE_CARD_FACE_PIN":                   30,
	"AUTH_EXT_MODE_CARD_FINGERPRINT_FACE":           31,
	"AUTH_EXT_MODE_CARD_FINGERPRINT_PIN":            32,
	"AUTH_EXT_MODE_CARD_FACE_OR_FINGERPRINT_PIN":    33,
	"AUTH_EXT_MODE_CARD_FACE_FINGERPRINT_OR_PIN":    34,
	"AUTH_EXT_MODE_CARD_FINGERPRINT_FACE_OR_PIN":    35,
	"AUTH_EXT_MODE_ID_FACE":                         36,
	"AUTH_EXT_MODE_ID_FINGERPRINT":                  37,
	"AUTH_EXT_MODE_ID_PIN":                          38,
	"AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT":          39,
	"AUTH_EXT_MODE_ID_FACE_OR_PIN":                  40,
	"AUTH_EXT_MODE_ID_FINGERPRINT_OR_PIN":           41,
	"AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_OR_PIN":   42,
	"AUTH_EXT_MODE_ID_FACE_FINGERPRINT":             43,
	"AUTH_EXT_MODE_ID_FACE_PIN":                     44,
	"AUTH_EXT_MODE_ID_FINGERPRINT_FACE":             45,
	"AUTH_EXT_MODE_ID_FINGERPRINT_PIN":              46,
	"AUTH_EXT_MODE_ID_FACE_OR_FINGERPRINT_PIN":      47,
	"AUTH_EXT_MODE_ID_FACE_FINGERPRINT_OR_PIN":      48,
	"AUTH_EXT_MODE_ID_FINGERPRINT_FACE_OR_PIN":      49,
}

func (x AuthMode) String() string {
	return proto.EnumName(AuthMode_name, int32(x))
}

func (AuthMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

type OperatorLevel int32

const (
	OperatorLevel_OPERATOR_LEVEL_NONE   OperatorLevel = 0
	OperatorLevel_OPERATOR_LEVEL_ADMIN  OperatorLevel = 1
	OperatorLevel_OPERATOR_LEVEL_CONFIG OperatorLevel = 2
	OperatorLevel_OPERATOR_LEVEL_USER   OperatorLevel = 3
)

var OperatorLevel_name = map[int32]string{
	0: "OPERATOR_LEVEL_NONE",
	1: "OPERATOR_LEVEL_ADMIN",
	2: "OPERATOR_LEVEL_CONFIG",
	3: "OPERATOR_LEVEL_USER",
}

var OperatorLevel_value = map[string]int32{
	"OPERATOR_LEVEL_NONE":   0,
	"OPERATOR_LEVEL_ADMIN":  1,
	"OPERATOR_LEVEL_CONFIG": 2,
	"OPERATOR_LEVEL_USER":   3,
}

func (x OperatorLevel) String() string {
	return proto.EnumName(OperatorLevel_name, int32(x))
}

func (OperatorLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

type FaceDetectionLevel int32

const (
	FaceDetectionLevel_FACE_DETECTION_NONE   FaceDetectionLevel = 0
	FaceDetectionLevel_FACE_DETECTION_NORMAL FaceDetectionLevel = 1
	FaceDetectionLevel_FACE_DETECTION_STRICT FaceDetectionLevel = 2
)

var FaceDetectionLevel_name = map[int32]string{
	0: "FACE_DETECTION_NONE",
	1: "FACE_DETECTION_NORMAL",
	2: "FACE_DETECTION_STRICT",
}

var FaceDetectionLevel_value = map[string]int32{
	"FACE_DETECTION_NONE":   0,
	"FACE_DETECTION_NORMAL": 1,
	"FACE_DETECTION_STRICT": 2,
}

func (x FaceDetectionLevel) String() string {
	return proto.EnumName(FaceDetectionLevel_name, int32(x))
}

func (FaceDetectionLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

type GlobalAPBFailActionType int32

const (
	GlobalAPBFailActionType_GLOBAL_APB_FAIL_ACTION_NONE GlobalAPBFailActionType = 0
	GlobalAPBFailActionType_GLOBAL_APB_FAIL_ACTION_SOFT GlobalAPBFailActionType = 1
	GlobalAPBFailActionType_GLOBAL_APB_FAIL_ACTION_HARD GlobalAPBFailActionType = 2
)

var GlobalAPBFailActionType_name = map[int32]string{
	0: "GLOBAL_APB_FAIL_ACTION_NONE",
	1: "GLOBAL_APB_FAIL_ACTION_SOFT",
	2: "GLOBAL_APB_FAIL_ACTION_HARD",
}

var GlobalAPBFailActionType_value = map[string]int32{
	"GLOBAL_APB_FAIL_ACTION_NONE": 0,
	"GLOBAL_APB_FAIL_ACTION_SOFT": 1,
	"GLOBAL_APB_FAIL_ACTION_HARD": 2,
}

func (x GlobalAPBFailActionType) String() string {
	return proto.EnumName(GlobalAPBFailActionType_name, int32(x))
}

func (GlobalAPBFailActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

type Operator struct {
	UserID               string        `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Level                OperatorLevel `protobuf:"varint,2,opt,name=level,proto3,enum=auth.OperatorLevel" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Operator) Reset()         { *m = Operator{} }
func (m *Operator) String() string { return proto.CompactTextString(m) }
func (*Operator) ProtoMessage()    {}
func (*Operator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Operator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operator.Unmarshal(m, b)
}
func (m *Operator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operator.Marshal(b, m, deterministic)
}
func (m *Operator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operator.Merge(m, src)
}
func (m *Operator) XXX_Size() int {
	return xxx_messageInfo_Operator.Size(m)
}
func (m *Operator) XXX_DiscardUnknown() {
	xxx_messageInfo_Operator.DiscardUnknown(m)
}

var xxx_messageInfo_Operator proto.InternalMessageInfo

func (m *Operator) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *Operator) GetLevel() OperatorLevel {
	if m != nil {
		return m.Level
	}
	return OperatorLevel_OPERATOR_LEVEL_NONE
}

type AuthSchedule struct {
	Mode                 AuthMode `protobuf:"varint,1,opt,name=mode,proto3,enum=auth.AuthMode" json:"mode,omitempty"`
	ScheduleID           uint32   `protobuf:"varint,2,opt,name=scheduleID,proto3" json:"scheduleID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSchedule) Reset()         { *m = AuthSchedule{} }
func (m *AuthSchedule) String() string { return proto.CompactTextString(m) }
func (*AuthSchedule) ProtoMessage()    {}
func (*AuthSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AuthSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSchedule.Unmarshal(m, b)
}
func (m *AuthSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSchedule.Marshal(b, m, deterministic)
}
func (m *AuthSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSchedule.Merge(m, src)
}
func (m *AuthSchedule) XXX_Size() int {
	return xxx_messageInfo_AuthSchedule.Size(m)
}
func (m *AuthSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSchedule proto.InternalMessageInfo

func (m *AuthSchedule) GetMode() AuthMode {
	if m != nil {
		return m.Mode
	}
	return AuthMode_AUTH_MODE_BIOMETRIC_ONLY
}

func (m *AuthSchedule) GetScheduleID() uint32 {
	if m != nil {
		return m.ScheduleID
	}
	return 0
}

type AuthConfig struct {
	AuthSchedules        []*AuthSchedule         `protobuf:"bytes,1,rep,name=authSchedules,proto3" json:"authSchedules,omitempty"`
	UseGlobalAPB         bool                    `protobuf:"varint,2,opt,name=useGlobalAPB,proto3" json:"useGlobalAPB,omitempty"`
	GlobalAPBFailAction  GlobalAPBFailActionType `protobuf:"varint,3,opt,name=globalAPBFailAction,proto3,enum=auth.GlobalAPBFailActionType" json:"globalAPBFailAction,omitempty"`
	UseGroupMatching     bool                    `protobuf:"varint,4,opt,name=useGroupMatching,proto3" json:"useGroupMatching,omitempty"`
	UsePrivateAuth       bool                    `protobuf:"varint,5,opt,name=usePrivateAuth,proto3" json:"usePrivateAuth,omitempty"`
	FaceDetectionLevel   FaceDetectionLevel      `protobuf:"varint,6,opt,name=faceDetectionLevel,proto3,enum=auth.FaceDetectionLevel" json:"faceDetectionLevel,omitempty"`
	UseServerMatching    bool                    `protobuf:"varint,7,opt,name=useServerMatching,proto3" json:"useServerMatching,omitempty"`
	UseFullAccess        bool                    `protobuf:"varint,8,opt,name=useFullAccess,proto3" json:"useFullAccess,omitempty"`
	MatchTimeout         uint32                  `protobuf:"varint,9,opt,name=matchTimeout,proto3" json:"matchTimeout,omitempty"`
	AuthTimeout          uint32                  `protobuf:"varint,10,opt,name=authTimeout,proto3" json:"authTimeout,omitempty"`
	Operators            []*Operator             `protobuf:"bytes,11,rep,name=operators,proto3" json:"operators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AuthConfig) Reset()         { *m = AuthConfig{} }
func (m *AuthConfig) String() string { return proto.CompactTextString(m) }
func (*AuthConfig) ProtoMessage()    {}
func (*AuthConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *AuthConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthConfig.Unmarshal(m, b)
}
func (m *AuthConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthConfig.Marshal(b, m, deterministic)
}
func (m *AuthConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthConfig.Merge(m, src)
}
func (m *AuthConfig) XXX_Size() int {
	return xxx_messageInfo_AuthConfig.Size(m)
}
func (m *AuthConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AuthConfig proto.InternalMessageInfo

func (m *AuthConfig) GetAuthSchedules() []*AuthSchedule {
	if m != nil {
		return m.AuthSchedules
	}
	return nil
}

func (m *AuthConfig) GetUseGlobalAPB() bool {
	if m != nil {
		return m.UseGlobalAPB
	}
	return false
}

func (m *AuthConfig) GetGlobalAPBFailAction() GlobalAPBFailActionType {
	if m != nil {
		return m.GlobalAPBFailAction
	}
	return GlobalAPBFailActionType_GLOBAL_APB_FAIL_ACTION_NONE
}

func (m *AuthConfig) GetUseGroupMatching() bool {
	if m != nil {
		return m.UseGroupMatching
	}
	return false
}

func (m *AuthConfig) GetUsePrivateAuth() bool {
	if m != nil {
		return m.UsePrivateAuth
	}
	return false
}

func (m *AuthConfig) GetFaceDetectionLevel() FaceDetectionLevel {
	if m != nil {
		return m.FaceDetectionLevel
	}
	return FaceDetectionLevel_FACE_DETECTION_NONE
}

func (m *AuthConfig) GetUseServerMatching() bool {
	if m != nil {
		return m.UseServerMatching
	}
	return false
}

func (m *AuthConfig) GetUseFullAccess() bool {
	if m != nil {
		return m.UseFullAccess
	}
	return false
}

func (m *AuthConfig) GetMatchTimeout() uint32 {
	if m != nil {
		return m.MatchTimeout
	}
	return 0
}

func (m *AuthConfig) GetAuthTimeout() uint32 {
	if m != nil {
		return m.AuthTimeout
	}
	return 0
}

func (m *AuthConfig) GetOperators() []*Operator {
	if m != nil {
		return m.Operators
	}
	return nil
}

type GetConfigRequest struct {
	DeviceID             uint32   `protobuf:"varint,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{3}
}

func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(m, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

func (m *GetConfigRequest) GetDeviceID() uint32 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

type GetConfigResponse struct {
	Config               *AuthConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetConfigResponse) Reset()         { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()    {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{4}
}

func (m *GetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigResponse.Unmarshal(m, b)
}
func (m *GetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigResponse.Marshal(b, m, deterministic)
}
func (m *GetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigResponse.Merge(m, src)
}
func (m *GetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetConfigResponse.Size(m)
}
func (m *GetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigResponse proto.InternalMessageInfo

func (m *GetConfigResponse) GetConfig() *AuthConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type SetConfigRequest struct {
	DeviceID             uint32      `protobuf:"varint,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Config               *AuthConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{5}
}

func (m *SetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigRequest.Unmarshal(m, b)
}
func (m *SetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigRequest.Marshal(b, m, deterministic)
}
func (m *SetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigRequest.Merge(m, src)
}
func (m *SetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigRequest.Size(m)
}
func (m *SetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigRequest proto.InternalMessageInfo

func (m *SetConfigRequest) GetDeviceID() uint32 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *SetConfigRequest) GetConfig() *AuthConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type SetConfigResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigResponse) Reset()         { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()    {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{6}
}

func (m *SetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigResponse.Unmarshal(m, b)
}
func (m *SetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigResponse.Marshal(b, m, deterministic)
}
func (m *SetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigResponse.Merge(m, src)
}
func (m *SetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigResponse.Size(m)
}
func (m *SetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigResponse proto.InternalMessageInfo

type SetConfigMultiRequest struct {
	DeviceIDs            []uint32    `protobuf:"varint,1,rep,packed,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	Config               *AuthConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SetConfigMultiRequest) Reset()         { *m = SetConfigMultiRequest{} }
func (m *SetConfigMultiRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigMultiRequest) ProtoMessage()    {}
func (*SetConfigMultiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{7}
}

func (m *SetConfigMultiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigMultiRequest.Unmarshal(m, b)
}
func (m *SetConfigMultiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigMultiRequest.Marshal(b, m, deterministic)
}
func (m *SetConfigMultiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigMultiRequest.Merge(m, src)
}
func (m *SetConfigMultiRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigMultiRequest.Size(m)
}
func (m *SetConfigMultiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigMultiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigMultiRequest proto.InternalMessageInfo

func (m *SetConfigMultiRequest) GetDeviceIDs() []uint32 {
	if m != nil {
		return m.DeviceIDs
	}
	return nil
}

func (m *SetConfigMultiRequest) GetConfig() *AuthConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type SetConfigMultiResponse struct {
	DeviceErrors         []*err.ErrorResponse `protobuf:"bytes,1,rep,name=deviceErrors,proto3" json:"deviceErrors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SetConfigMultiResponse) Reset()         { *m = SetConfigMultiResponse{} }
func (m *SetConfigMultiResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigMultiResponse) ProtoMessage()    {}
func (*SetConfigMultiResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{8}
}

func (m *SetConfigMultiResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigMultiResponse.Unmarshal(m, b)
}
func (m *SetConfigMultiResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigMultiResponse.Marshal(b, m, deterministic)
}
func (m *SetConfigMultiResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigMultiResponse.Merge(m, src)
}
func (m *SetConfigMultiResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigMultiResponse.Size(m)
}
func (m *SetConfigMultiResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigMultiResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigMultiResponse proto.InternalMessageInfo

func (m *SetConfigMultiResponse) GetDeviceErrors() []*err.ErrorResponse {
	if m != nil {
		return m.DeviceErrors
	}
	return nil
}

func init() {
	proto.RegisterEnum("auth.AuthMode", AuthMode_name, AuthMode_value)
	proto.RegisterEnum("auth.OperatorLevel", OperatorLevel_name, OperatorLevel_value)
	proto.RegisterEnum("auth.FaceDetectionLevel", FaceDetectionLevel_name, FaceDetectionLevel_value)
	proto.RegisterEnum("auth.GlobalAPBFailActionType", GlobalAPBFailActionType_name, GlobalAPBFailActionType_value)
	proto.RegisterType((*Operator)(nil), "auth.Operator")
	proto.RegisterType((*AuthSchedule)(nil), "auth.AuthSchedule")
	proto.RegisterType((*AuthConfig)(nil), "auth.AuthConfig")
	proto.RegisterType((*GetConfigRequest)(nil), "auth.GetConfigRequest")
	proto.RegisterType((*GetConfigResponse)(nil), "auth.GetConfigResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "auth.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "auth.SetConfigResponse")
	proto.RegisterType((*SetConfigMultiRequest)(nil), "auth.SetConfigMultiRequest")
	proto.RegisterType((*SetConfigMultiResponse)(nil), "auth.SetConfigMultiResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 1182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdd, 0x72, 0xda, 0x56,
	0x10, 0x0e, 0x36, 0x71, 0x60, 0x6d, 0xa8, 0xbc, 0xd8, 0x20, 0xf0, 0x4f, 0x88, 0xe2, 0x38, 0x84,
	0x38, 0xb8, 0x71, 0x67, 0x3a, 0xbd, 0x48, 0x2f, 0x04, 0x08, 0xac, 0x29, 0x20, 0xe6, 0xa0, 0x74,
	0xd2, 0xde, 0x30, 0x04, 0x9f, 0xc4, 0x4c, 0xb1, 0x45, 0xf5, 0xe3, 0x99, 0xce, 0xf4, 0x65, 0xfa,
	0x44, 0x7d, 0x80, 0xbe, 0x4b, 0xdb, 0x39, 0x47, 0x08, 0xf4, 0xeb, 0x38, 0x77, 0x66, 0xbf, 0x6f,
	0xbf, 0x6f, 0xf7, 0x9c, 0x95, 0xb4, 0x06, 0x98, 0x38, 0xf6, 0x75, 0x63, 0x61, 0x1a, 0xb6, 0x81,
	0x69, 0xf6, 0x77, 0x25, 0x4b, 0x4d, 0xd3, 0x0d, 0x48, 0x7d, 0xc8, 0x68, 0x0b, 0x6a, 0x4e, 0x6c,
	0xc3, 0xc4, 0x22, 0x6c, 0x39, 0x16, 0x35, 0xd5, 0xb6, 0x98, 0xaa, 0xa6, 0x6a, 0x59, 0xb2, 0xfc,
	0x85, 0xaf, 0xe0, 0xf1, 0x9c, 0xde, 0xd1, 0xb9, 0xb8, 0x51, 0x4d, 0xd5, 0xf2, 0x17, 0x85, 0x06,
	0x17, 0xf4, 0xd2, 0x7a, 0x0c, 0x22, 0x2e, 0x43, 0x22, 0xb0, 0x23, 0x3b, 0xf6, 0xf5, 0x68, 0x7a,
	0x4d, 0xaf, 0x9c, 0x39, 0x45, 0x09, 0xd2, 0x37, 0xc6, 0x15, 0xe5, 0x82, 0xf9, 0x8b, 0xbc, 0x9b,
	0xc9, 0x18, 0x7d, 0xe3, 0x8a, 0x12, 0x8e, 0xe1, 0x31, 0x80, 0xb5, 0xe4, 0xab, 0x6d, 0xee, 0x91,
	0x23, 0xbe, 0x88, 0xf4, 0x57, 0x1a, 0x80, 0xa5, 0xb4, 0x8c, 0xdb, 0x4f, 0xb3, 0xcf, 0xf8, 0x03,
	0xe4, 0x26, 0x3e, 0x0b, 0x4b, 0x4c, 0x55, 0x37, 0x6b, 0xdb, 0x17, 0xb8, 0xd6, 0xf6, 0x20, 0x12,
	0x24, 0xa2, 0x04, 0x3b, 0x8e, 0x45, 0xbb, 0x73, 0xe3, 0xe3, 0x64, 0x2e, 0x0f, 0x9b, 0xdc, 0x2a,
	0x43, 0x02, 0x31, 0xd4, 0xa0, 0xf0, 0xd9, 0xfb, 0xd1, 0x99, 0xcc, 0xe6, 0xf2, 0xd4, 0x9e, 0x19,
	0xb7, 0xe2, 0x26, 0xaf, 0xff, 0xc8, 0xf5, 0xe8, 0x46, 0x09, 0xfa, 0x1f, 0x0b, 0x4a, 0xe2, 0x32,
	0xb1, 0x0e, 0x02, 0x33, 0x30, 0x0d, 0x67, 0xd1, 0x9f, 0xd8, 0xd3, 0xeb, 0xd9, 0xed, 0x67, 0x31,
	0xcd, 0x8d, 0x23, 0x71, 0x3c, 0x85, 0xbc, 0x63, 0xd1, 0xa1, 0x39, 0xbb, 0x9b, 0xd8, 0x94, 0x75,
	0x22, 0x3e, 0xe6, 0xcc, 0x50, 0x14, 0x2f, 0x01, 0x3f, 0x4d, 0xa6, 0xb4, 0x4d, 0x6d, 0xca, 0x4d,
	0xf8, 0x15, 0x88, 0x5b, 0xbc, 0x46, 0xd1, 0xad, 0xb1, 0x13, 0xc1, 0x49, 0x4c, 0x0e, 0x9e, 0xc1,
	0xae, 0x63, 0xd1, 0x11, 0x35, 0xef, 0xa8, 0xb9, 0x2a, 0xef, 0x09, 0x37, 0x8d, 0x02, 0x78, 0x02,
	0x39, 0xc7, 0xa2, 0x1d, 0x67, 0x3e, 0x97, 0xa7, 0x53, 0x6a, 0x59, 0x62, 0x86, 0x33, 0x83, 0x41,
	0x76, 0xcc, 0x37, 0x2c, 0x43, 0x9f, 0xdd, 0x50, 0xc3, 0xb1, 0xc5, 0x2c, 0xbf, 0xd1, 0x40, 0x0c,
	0xab, 0xb0, 0xcd, 0xca, 0xf4, 0x28, 0xc0, 0x29, 0xfe, 0x10, 0x9e, 0x41, 0xd6, 0x58, 0x4e, 0x98,
	0x25, 0x6e, 0xf3, 0x2b, 0xce, 0x07, 0x07, 0x8f, 0xac, 0x09, 0x52, 0x03, 0x84, 0x2e, 0xb5, 0xdd,
	0x09, 0x21, 0xf4, 0x77, 0x87, 0x5a, 0x36, 0x56, 0x20, 0x73, 0x45, 0xef, 0x66, 0x53, 0xba, 0x1c,
	0xe8, 0x1c, 0x59, 0xfd, 0x96, 0x7e, 0x84, 0x5d, 0x1f, 0xdf, 0x5a, 0x18, 0xb7, 0x16, 0xc5, 0x1a,
	0x6c, 0x4d, 0x79, 0x84, 0xd3, 0xb7, 0x2f, 0x84, 0xf5, 0x48, 0x2d, 0x99, 0x4b, 0x5c, 0xfa, 0x00,
	0xc2, 0xe8, 0x2b, 0xec, 0x7c, 0xca, 0x1b, 0x5f, 0x50, 0x2e, 0xc0, 0xee, 0x28, 0x5c, 0x98, 0x34,
	0x86, 0xfd, 0x55, 0xb0, 0xef, 0xcc, 0xed, 0x99, 0xe7, 0x79, 0x08, 0x59, 0xcf, 0xc3, 0x7d, 0x0e,
	0x72, 0x64, 0x1d, 0xf8, 0x0a, 0xd7, 0x21, 0x14, 0xc3, 0x06, 0xcb, 0x33, 0xf9, 0x1e, 0x76, 0x5c,
	0x41, 0xc5, 0x34, 0xd9, 0x4d, 0x78, 0x0f, 0x1b, 0x7b, 0x83, 0xf0, 0x90, 0xc7, 0x24, 0x01, 0x5e,
	0xfd, 0x9f, 0x3c, 0x64, 0xbc, 0xe7, 0x1c, 0x0f, 0x41, 0x94, 0xdf, 0xeb, 0x97, 0xe3, 0xbe, 0xd6,
	0x56, 0xc6, 0x4d, 0x55, 0xeb, 0x2b, 0x3a, 0x51, 0x5b, 0x63, 0x6d, 0xd0, 0xfb, 0x45, 0x78, 0x84,
	0x07, 0x50, 0x8a, 0x43, 0x87, 0xea, 0x40, 0x48, 0x61, 0x09, 0x0a, 0x6b, 0xb0, 0x25, 0x93, 0xb6,
	0x9b, 0xb5, 0x11, 0xd4, 0xe4, 0xc0, 0x2a, 0x55, 0xd8, 0xc4, 0x22, 0x60, 0x08, 0x65, 0x72, 0x69,
	0x7c, 0x0e, 0x4f, 0x93, 0xb2, 0xc6, 0x1a, 0xe1, 0xa4, 0xc7, 0x58, 0x85, 0xc3, 0x44, 0x12, 0x63,
	0x6c, 0x61, 0x05, 0x8a, 0x6b, 0x86, 0xea, 0xb7, 0x7e, 0x82, 0x7b, 0x20, 0x04, 0x30, 0x96, 0x91,
	0xc1, 0x67, 0x70, 0x14, 0x9f, 0xe1, 0xd9, 0x66, 0xf1, 0x18, 0x2a, 0x09, 0x14, 0x86, 0x03, 0x16,
	0x20, 0xbf, 0xc6, 0x07, 0xda, 0x40, 0x11, 0xfe, 0x4b, 0x61, 0x19, 0xf6, 0xd6, 0xc1, 0x21, 0xd1,
	0x2e, 0xd5, 0xa6, 0xaa, 0x2b, 0x6d, 0xe1, 0xdf, 0xd4, 0xea, 0x5c, 0x95, 0x0f, 0xba, 0x0b, 0x77,
	0xe4, 0x96, 0xe2, 0x1e, 0xdf, 0x36, 0x4a, 0x70, 0x1c, 0x03, 0x76, 0xd4, 0x41, 0x57, 0x21, 0x43,
	0xa2, 0x0e, 0x74, 0x61, 0x67, 0xd5, 0x65, 0x90, 0xc3, 0x8a, 0xc9, 0xe1, 0x2b, 0x78, 0x71, 0x7f,
	0xbe, 0xd7, 0x57, 0x1e, 0x4f, 0x41, 0xfa, 0x02, 0x95, 0xf1, 0xbe, 0x89, 0x29, 0xc9, 0xaf, 0xc6,
	0xca, 0x16, 0xee, 0xe7, 0x30, 0x5d, 0x61, 0x77, 0x75, 0xd4, 0xb1, 0x1c, 0x66, 0x85, 0x31, 0xd5,
	0x87, 0x64, 0xbc, 0xea, 0x0b, 0x31, 0xd5, 0x87, 0xa9, 0x8c, 0xb7, 0x17, 0x3d, 0xed, 0xf5, 0xb0,
	0xee, 0x27, 0x80, 0xbc, 0xde, 0x62, 0xb4, 0x27, 0x17, 0xf4, 0x5d, 0x45, 0x29, 0x7a, 0x15, 0xab,
	0x99, 0x16, 0xb1, 0x0e, 0xa7, 0x09, 0xe2, 0xac, 0x0b, 0xbf, 0x4e, 0x39, 0xc9, 0xcb, 0xd7, 0x71,
	0x25, 0x7a, 0x38, 0xe1, 0x7a, 0x3c, 0xea, 0x01, 0xbe, 0x85, 0x37, 0x0f, 0xb3, 0xf6, 0x52, 0x0e,
	0xf1, 0x25, 0x3c, 0x4f, 0x4a, 0xf1, 0x97, 0x7a, 0x84, 0x4f, 0xe1, 0x20, 0x89, 0xc8, 0x94, 0x8e,
	0x93, 0x94, 0xc2, 0x03, 0xf1, 0x34, 0x7a, 0x85, 0x11, 0x22, 0x13, 0xac, 0x62, 0x03, 0xea, 0x0f,
	0xec, 0x86, 0xf1, 0x9f, 0xdd, 0xc7, 0x8f, 0x69, 0x5d, 0x4a, 0xe2, 0x27, 0x8c, 0xde, 0x73, 0x2c,
	0xc3, 0x7e, 0x90, 0xaf, 0x2e, 0x67, 0xe6, 0x64, 0xf5, 0x8a, 0x0a, 0x40, 0xbe, 0xe3, 0x7b, 0x81,
	0xe2, 0xf2, 0xc5, 0xe0, 0x67, 0x30, 0xd9, 0x53, 0xac, 0xc1, 0x49, 0xac, 0x6c, 0x78, 0x5a, 0x5e,
	0xc6, 0xbb, 0xf8, 0x4a, 0xac, 0x45, 0xef, 0x40, 0x8d, 0x9d, 0x94, 0x57, 0x78, 0x0e, 0xaf, 0x1f,
	0x62, 0xea, 0x25, 0xd4, 0xf1, 0x05, 0x3c, 0x8b, 0x4f, 0xf0, 0x97, 0xf8, 0x1a, 0x8f, 0xa0, 0x1c,
	0x4f, 0x63, 0x2a, 0x67, 0xf1, 0x2a, 0xe1, 0x09, 0x79, 0x83, 0x27, 0x50, 0xbd, 0x97, 0xc6, 0xc4,
	0x1a, 0x78, 0x06, 0xb5, 0x07, 0xf5, 0xc0, 0xd8, 0xe7, 0xc9, 0xec, 0x98, 0x76, 0xbf, 0x8d, 0x67,
	0x27, 0x4c, 0xc6, 0xdb, 0xfa, 0x1d, 0xe4, 0x02, 0xeb, 0x37, 0xfb, 0x4c, 0x6a, 0x43, 0x85, 0xc8,
	0xba, 0x46, 0xc6, 0x3d, 0xe5, 0x67, 0xa5, 0xe7, 0x7e, 0x20, 0x1e, 0xb1, 0x31, 0x08, 0x01, 0x72,
	0xbb, 0xcf, 0xbf, 0xac, 0x65, 0xd8, 0x0f, 0x21, 0x2d, 0x6d, 0xd0, 0x51, 0xbb, 0xc2, 0x46, 0x8c,
	0xda, 0xfb, 0x91, 0x42, 0x84, 0xcd, 0xfa, 0x14, 0x30, 0xba, 0x58, 0x32, 0x3a, 0x2f, 0xaf, 0xad,
	0xe8, 0x4a, 0x4b, 0x57, 0xb5, 0x81, 0x67, 0x5e, 0x86, 0xfd, 0x08, 0x40, 0xfa, 0x72, 0xcf, 0x75,
	0x0f, 0x41, 0x23, 0xf6, 0xa9, 0xd3, 0x85, 0x8d, 0xfa, 0x9f, 0x50, 0x4a, 0xd8, 0xb0, 0xd9, 0x3b,
	0xa1, 0xdb, 0xd3, 0x9a, 0x72, 0x6f, 0x2c, 0x0f, 0x9b, 0xe3, 0x8e, 0xac, 0xf6, 0xc6, 0x72, 0xc0,
	0x31, 0x99, 0x30, 0xd2, 0x3a, 0xba, 0x90, 0xba, 0x87, 0x70, 0x29, 0x93, 0xb6, 0xb0, 0x71, 0xf1,
	0x77, 0x0a, 0xd2, 0x7c, 0xc9, 0x7e, 0x07, 0xd9, 0xd5, 0x8a, 0x88, 0xc5, 0xe5, 0xe6, 0x1f, 0x5a,
	0xfa, 0x2a, 0xa5, 0x48, 0x7c, 0xb9, 0x37, 0xbd, 0x83, 0xec, 0x28, 0x9c, 0x3d, 0x4a, 0xc8, 0x8e,
	0x2c, 0x7c, 0xf8, 0x13, 0xe4, 0x83, 0xfb, 0x18, 0x1e, 0x84, 0xa8, 0xfe, 0x35, 0xb0, 0x72, 0x18,
	0x0f, 0xba, 0x62, 0xcd, 0xb7, 0x50, 0x9a, 0x1a, 0x37, 0x0d, 0xcb, 0x59, 0x98, 0xf4, 0x66, 0x32,
	0xbb, 0x9d, 0x36, 0xac, 0xab, 0xdf, 0x78, 0xc6, 0x30, 0xf5, 0xeb, 0xde, 0xc7, 0x99, 0x61, 0xd9,
	0x13, 0xf3, 0xdc, 0xa2, 0x26, 0x5b, 0xdf, 0xce, 0x59, 0xfc, 0xe3, 0x16, 0xff, 0xe7, 0xf0, 0xbb,
	0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xdb, 0xb1, 0x53, 0x3b, 0x0e, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	SetConfigMulti(ctx context.Context, in *SetConfigMultiRequest, opts ...grpc.CallOption) (*SetConfigMultiResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetConfigMulti(ctx context.Context, in *SetConfigMultiRequest, opts ...grpc.CallOption) (*SetConfigMultiResponse, error) {
	out := new(SetConfigMultiResponse)
	err := c.cc.Invoke(ctx, "/auth.Auth/SetConfigMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	SetConfigMulti(context.Context, *SetConfigMultiRequest) (*SetConfigMultiResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetConfigMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetConfigMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/SetConfigMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetConfigMulti(ctx, req.(*SetConfigMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Auth_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Auth_SetConfig_Handler,
		},
		{
			MethodName: "SetConfigMulti",
			Handler:    _Auth_SetConfigMulti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
