// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thermal.proto

package thermal

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

type CheckMode int32

const (
	CheckMode_OFF  CheckMode = 0
	CheckMode_HARD CheckMode = 1
	CheckMode_SOFT CheckMode = 2
)

var CheckMode_name = map[int32]string{
	0: "OFF",
	1: "HARD",
	2: "SOFT",
}

var CheckMode_value = map[string]int32{
	"OFF":  0,
	"HARD": 1,
	"SOFT": 2,
}

func (x CheckMode) String() string {
	return proto.EnumName(CheckMode_name, int32(x))
}

func (CheckMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{0}
}

type CheckOrder int32

const (
	CheckOrder_AFTER_AUTH   CheckOrder = 0
	CheckOrder_BEFORE_AUTH  CheckOrder = 1
	CheckOrder_WITHOUT_AUTH CheckOrder = 2
)

var CheckOrder_name = map[int32]string{
	0: "AFTER_AUTH",
	1: "BEFORE_AUTH",
	2: "WITHOUT_AUTH",
}

var CheckOrder_value = map[string]int32{
	"AFTER_AUTH":   0,
	"BEFORE_AUTH":  1,
	"WITHOUT_AUTH": 2,
}

func (x CheckOrder) String() string {
	return proto.EnumName(CheckOrder_name, int32(x))
}

func (CheckOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{1}
}

type TemperatureFormat int32

const (
	TemperatureFormat_FAHRENHEIT TemperatureFormat = 0
	TemperatureFormat_CELSIUS    TemperatureFormat = 1
)

var TemperatureFormat_name = map[int32]string{
	0: "FAHRENHEIT",
	1: "CELSIUS",
}

var TemperatureFormat_value = map[string]int32{
	"FAHRENHEIT": 0,
	"CELSIUS":    1,
}

func (x TemperatureFormat) String() string {
	return proto.EnumName(TemperatureFormat_name, int32(x))
}

func (TemperatureFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{2}
}

type MaskDetectionLevel int32

const (
	MaskDetectionLevel_NOT_USE   MaskDetectionLevel = 0
	MaskDetectionLevel_NORMAL    MaskDetectionLevel = 1
	MaskDetectionLevel_HIGH      MaskDetectionLevel = 2
	MaskDetectionLevel_VERY_HIGH MaskDetectionLevel = 3
)

var MaskDetectionLevel_name = map[int32]string{
	0: "NOT_USE",
	1: "NORMAL",
	2: "HIGH",
	3: "VERY_HIGH",
}

var MaskDetectionLevel_value = map[string]int32{
	"NOT_USE":   0,
	"NORMAL":    1,
	"HIGH":      2,
	"VERY_HIGH": 3,
}

func (x MaskDetectionLevel) String() string {
	return proto.EnumName(MaskDetectionLevel_name, int32(x))
}

func (MaskDetectionLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{3}
}

type ThermalCameraROI struct {
	X                    uint32   `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    uint32   `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width                uint32   `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32   `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThermalCameraROI) Reset()         { *m = ThermalCameraROI{} }
func (m *ThermalCameraROI) String() string { return proto.CompactTextString(m) }
func (*ThermalCameraROI) ProtoMessage()    {}
func (*ThermalCameraROI) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{0}
}

func (m *ThermalCameraROI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermalCameraROI.Unmarshal(m, b)
}
func (m *ThermalCameraROI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermalCameraROI.Marshal(b, m, deterministic)
}
func (m *ThermalCameraROI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermalCameraROI.Merge(m, src)
}
func (m *ThermalCameraROI) XXX_Size() int {
	return xxx_messageInfo_ThermalCameraROI.Size(m)
}
func (m *ThermalCameraROI) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermalCameraROI.DiscardUnknown(m)
}

var xxx_messageInfo_ThermalCameraROI proto.InternalMessageInfo

func (m *ThermalCameraROI) GetX() uint32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ThermalCameraROI) GetY() uint32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *ThermalCameraROI) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ThermalCameraROI) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type ThermalCamera struct {
	Distance                uint32            `protobuf:"varint,1,opt,name=distance,proto3" json:"distance,omitempty"`
	EmissionRate            uint32            `protobuf:"varint,2,opt,name=emissionRate,proto3" json:"emissionRate,omitempty"`
	ROI                     *ThermalCameraROI `protobuf:"bytes,3,opt,name=ROI,proto3" json:"ROI,omitempty"`
	UseBodyCompensation     bool              `protobuf:"varint,4,opt,name=useBodyCompensation,proto3" json:"useBodyCompensation,omitempty"`
	CompensationTemperature int32             `protobuf:"varint,5,opt,name=compensationTemperature,proto3" json:"compensationTemperature,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}          `json:"-"`
	XXX_unrecognized        []byte            `json:"-"`
	XXX_sizecache           int32             `json:"-"`
}

func (m *ThermalCamera) Reset()         { *m = ThermalCamera{} }
func (m *ThermalCamera) String() string { return proto.CompactTextString(m) }
func (*ThermalCamera) ProtoMessage()    {}
func (*ThermalCamera) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{1}
}

func (m *ThermalCamera) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermalCamera.Unmarshal(m, b)
}
func (m *ThermalCamera) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermalCamera.Marshal(b, m, deterministic)
}
func (m *ThermalCamera) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermalCamera.Merge(m, src)
}
func (m *ThermalCamera) XXX_Size() int {
	return xxx_messageInfo_ThermalCamera.Size(m)
}
func (m *ThermalCamera) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermalCamera.DiscardUnknown(m)
}

var xxx_messageInfo_ThermalCamera proto.InternalMessageInfo

func (m *ThermalCamera) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *ThermalCamera) GetEmissionRate() uint32 {
	if m != nil {
		return m.EmissionRate
	}
	return 0
}

func (m *ThermalCamera) GetROI() *ThermalCameraROI {
	if m != nil {
		return m.ROI
	}
	return nil
}

func (m *ThermalCamera) GetUseBodyCompensation() bool {
	if m != nil {
		return m.UseBodyCompensation
	}
	return false
}

func (m *ThermalCamera) GetCompensationTemperature() int32 {
	if m != nil {
		return m.CompensationTemperature
	}
	return 0
}

type ThermalConfig struct {
	CheckMode            CheckMode         `protobuf:"varint,1,opt,name=checkMode,proto3,enum=thermal.CheckMode" json:"checkMode,omitempty"`
	CheckOrder           CheckOrder        `protobuf:"varint,2,opt,name=checkOrder,proto3,enum=thermal.CheckOrder" json:"checkOrder,omitempty"`
	TemperatureFormat    TemperatureFormat `protobuf:"varint,3,opt,name=temperatureFormat,proto3,enum=thermal.TemperatureFormat" json:"temperatureFormat,omitempty"`
	TemperatureThreshold uint32            `protobuf:"varint,4,opt,name=temperatureThreshold,proto3" json:"temperatureThreshold,omitempty"`
	AuditTemperature     bool              `protobuf:"varint,5,opt,name=auditTemperature,proto3" json:"auditTemperature,omitempty"`
	UseRejectSound       bool              `protobuf:"varint,6,opt,name=useRejectSound,proto3" json:"useRejectSound,omitempty"`
	UseOverlapThermal    bool              `protobuf:"varint,7,opt,name=useOverlapThermal,proto3" json:"useOverlapThermal,omitempty"`
	Camera               *ThermalCamera    `protobuf:"bytes,8,opt,name=camera,proto3" json:"camera,omitempty"`
	// only for FaceStation F2
	MaskCheckMode        CheckMode          `protobuf:"varint,9,opt,name=maskCheckMode,proto3,enum=thermal.CheckMode" json:"maskCheckMode,omitempty"`
	MaskDetectionLevel   MaskDetectionLevel `protobuf:"varint,10,opt,name=maskDetectionLevel,proto3,enum=thermal.MaskDetectionLevel" json:"maskDetectionLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ThermalConfig) Reset()         { *m = ThermalConfig{} }
func (m *ThermalConfig) String() string { return proto.CompactTextString(m) }
func (*ThermalConfig) ProtoMessage()    {}
func (*ThermalConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{2}
}

func (m *ThermalConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermalConfig.Unmarshal(m, b)
}
func (m *ThermalConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermalConfig.Marshal(b, m, deterministic)
}
func (m *ThermalConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermalConfig.Merge(m, src)
}
func (m *ThermalConfig) XXX_Size() int {
	return xxx_messageInfo_ThermalConfig.Size(m)
}
func (m *ThermalConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermalConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ThermalConfig proto.InternalMessageInfo

func (m *ThermalConfig) GetCheckMode() CheckMode {
	if m != nil {
		return m.CheckMode
	}
	return CheckMode_OFF
}

func (m *ThermalConfig) GetCheckOrder() CheckOrder {
	if m != nil {
		return m.CheckOrder
	}
	return CheckOrder_AFTER_AUTH
}

func (m *ThermalConfig) GetTemperatureFormat() TemperatureFormat {
	if m != nil {
		return m.TemperatureFormat
	}
	return TemperatureFormat_FAHRENHEIT
}

func (m *ThermalConfig) GetTemperatureThreshold() uint32 {
	if m != nil {
		return m.TemperatureThreshold
	}
	return 0
}

func (m *ThermalConfig) GetAuditTemperature() bool {
	if m != nil {
		return m.AuditTemperature
	}
	return false
}

func (m *ThermalConfig) GetUseRejectSound() bool {
	if m != nil {
		return m.UseRejectSound
	}
	return false
}

func (m *ThermalConfig) GetUseOverlapThermal() bool {
	if m != nil {
		return m.UseOverlapThermal
	}
	return false
}

func (m *ThermalConfig) GetCamera() *ThermalCamera {
	if m != nil {
		return m.Camera
	}
	return nil
}

func (m *ThermalConfig) GetMaskCheckMode() CheckMode {
	if m != nil {
		return m.MaskCheckMode
	}
	return CheckMode_OFF
}

func (m *ThermalConfig) GetMaskDetectionLevel() MaskDetectionLevel {
	if m != nil {
		return m.MaskDetectionLevel
	}
	return MaskDetectionLevel_NOT_USE
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
	return fileDescriptor_21313f5eb7527e98, []int{3}
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
	Config               *ThermalConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetConfigResponse) Reset()         { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()    {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{4}
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

func (m *GetConfigResponse) GetConfig() *ThermalConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type SetConfigRequest struct {
	DeviceID             uint32         `protobuf:"varint,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	Config               *ThermalConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{5}
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

func (m *SetConfigRequest) GetConfig() *ThermalConfig {
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
	return fileDescriptor_21313f5eb7527e98, []int{6}
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
	DeviceIDs            []uint32       `protobuf:"varint,1,rep,packed,name=deviceIDs,proto3" json:"deviceIDs,omitempty"`
	Config               *ThermalConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetConfigMultiRequest) Reset()         { *m = SetConfigMultiRequest{} }
func (m *SetConfigMultiRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigMultiRequest) ProtoMessage()    {}
func (*SetConfigMultiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{7}
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

func (m *SetConfigMultiRequest) GetConfig() *ThermalConfig {
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
	return fileDescriptor_21313f5eb7527e98, []int{8}
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

type TemperatureLog struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp            uint32   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DeviceID             uint32   `protobuf:"varint,3,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	UserID               string   `protobuf:"bytes,4,opt,name=userID,proto3" json:"userID,omitempty"`
	EventCode            uint32   `protobuf:"varint,5,opt,name=eventCode,proto3" json:"eventCode,omitempty"`
	SubCode              uint32   `protobuf:"varint,6,opt,name=subCode,proto3" json:"subCode,omitempty"`
	Temperature          uint32   `protobuf:"varint,7,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemperatureLog) Reset()         { *m = TemperatureLog{} }
func (m *TemperatureLog) String() string { return proto.CompactTextString(m) }
func (*TemperatureLog) ProtoMessage()    {}
func (*TemperatureLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{9}
}

func (m *TemperatureLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemperatureLog.Unmarshal(m, b)
}
func (m *TemperatureLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemperatureLog.Marshal(b, m, deterministic)
}
func (m *TemperatureLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureLog.Merge(m, src)
}
func (m *TemperatureLog) XXX_Size() int {
	return xxx_messageInfo_TemperatureLog.Size(m)
}
func (m *TemperatureLog) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureLog.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureLog proto.InternalMessageInfo

func (m *TemperatureLog) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TemperatureLog) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TemperatureLog) GetDeviceID() uint32 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *TemperatureLog) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TemperatureLog) GetEventCode() uint32 {
	if m != nil {
		return m.EventCode
	}
	return 0
}

func (m *TemperatureLog) GetSubCode() uint32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func (m *TemperatureLog) GetTemperature() uint32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

type GetTemperatureLogRequest struct {
	DeviceID             uint32   `protobuf:"varint,1,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	StartEventID         uint32   `protobuf:"varint,2,opt,name=startEventID,proto3" json:"startEventID,omitempty"`
	MaxNumOfLog          uint32   `protobuf:"varint,3,opt,name=maxNumOfLog,proto3" json:"maxNumOfLog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTemperatureLogRequest) Reset()         { *m = GetTemperatureLogRequest{} }
func (m *GetTemperatureLogRequest) String() string { return proto.CompactTextString(m) }
func (*GetTemperatureLogRequest) ProtoMessage()    {}
func (*GetTemperatureLogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{10}
}

func (m *GetTemperatureLogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemperatureLogRequest.Unmarshal(m, b)
}
func (m *GetTemperatureLogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemperatureLogRequest.Marshal(b, m, deterministic)
}
func (m *GetTemperatureLogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemperatureLogRequest.Merge(m, src)
}
func (m *GetTemperatureLogRequest) XXX_Size() int {
	return xxx_messageInfo_GetTemperatureLogRequest.Size(m)
}
func (m *GetTemperatureLogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemperatureLogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemperatureLogRequest proto.InternalMessageInfo

func (m *GetTemperatureLogRequest) GetDeviceID() uint32 {
	if m != nil {
		return m.DeviceID
	}
	return 0
}

func (m *GetTemperatureLogRequest) GetStartEventID() uint32 {
	if m != nil {
		return m.StartEventID
	}
	return 0
}

func (m *GetTemperatureLogRequest) GetMaxNumOfLog() uint32 {
	if m != nil {
		return m.MaxNumOfLog
	}
	return 0
}

type GetTemperatureLogResponse struct {
	TemperatureEvents    []*TemperatureLog `protobuf:"bytes,1,rep,name=temperatureEvents,proto3" json:"temperatureEvents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetTemperatureLogResponse) Reset()         { *m = GetTemperatureLogResponse{} }
func (m *GetTemperatureLogResponse) String() string { return proto.CompactTextString(m) }
func (*GetTemperatureLogResponse) ProtoMessage()    {}
func (*GetTemperatureLogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21313f5eb7527e98, []int{11}
}

func (m *GetTemperatureLogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTemperatureLogResponse.Unmarshal(m, b)
}
func (m *GetTemperatureLogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTemperatureLogResponse.Marshal(b, m, deterministic)
}
func (m *GetTemperatureLogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTemperatureLogResponse.Merge(m, src)
}
func (m *GetTemperatureLogResponse) XXX_Size() int {
	return xxx_messageInfo_GetTemperatureLogResponse.Size(m)
}
func (m *GetTemperatureLogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTemperatureLogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTemperatureLogResponse proto.InternalMessageInfo

func (m *GetTemperatureLogResponse) GetTemperatureEvents() []*TemperatureLog {
	if m != nil {
		return m.TemperatureEvents
	}
	return nil
}

func init() {
	proto.RegisterEnum("thermal.CheckMode", CheckMode_name, CheckMode_value)
	proto.RegisterEnum("thermal.CheckOrder", CheckOrder_name, CheckOrder_value)
	proto.RegisterEnum("thermal.TemperatureFormat", TemperatureFormat_name, TemperatureFormat_value)
	proto.RegisterEnum("thermal.MaskDetectionLevel", MaskDetectionLevel_name, MaskDetectionLevel_value)
	proto.RegisterType((*ThermalCameraROI)(nil), "thermal.ThermalCameraROI")
	proto.RegisterType((*ThermalCamera)(nil), "thermal.ThermalCamera")
	proto.RegisterType((*ThermalConfig)(nil), "thermal.ThermalConfig")
	proto.RegisterType((*GetConfigRequest)(nil), "thermal.GetConfigRequest")
	proto.RegisterType((*GetConfigResponse)(nil), "thermal.GetConfigResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "thermal.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "thermal.SetConfigResponse")
	proto.RegisterType((*SetConfigMultiRequest)(nil), "thermal.SetConfigMultiRequest")
	proto.RegisterType((*SetConfigMultiResponse)(nil), "thermal.SetConfigMultiResponse")
	proto.RegisterType((*TemperatureLog)(nil), "thermal.TemperatureLog")
	proto.RegisterType((*GetTemperatureLogRequest)(nil), "thermal.GetTemperatureLogRequest")
	proto.RegisterType((*GetTemperatureLogResponse)(nil), "thermal.GetTemperatureLogResponse")
}

func init() { proto.RegisterFile("thermal.proto", fileDescriptor_21313f5eb7527e98) }

var fileDescriptor_21313f5eb7527e98 = []byte{
	// 959 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0xae, 0x93, 0x6d, 0xd2, 0x4c, 0x93, 0xe0, 0x9c, 0x2e, 0xad, 0x37, 0x20, 0x28, 0xbe, 0x40,
	0x55, 0x40, 0xd9, 0x2a, 0x2b, 0xc1, 0xde, 0xa1, 0xfc, 0x38, 0x8d, 0x45, 0x5a, 0x2f, 0x27, 0x2e,
	0x88, 0x15, 0xa2, 0x72, 0x9d, 0xd9, 0xc6, 0x34, 0x8e, 0xc3, 0x39, 0xc7, 0xa5, 0x95, 0x78, 0x25,
	0xde, 0x81, 0x07, 0xe0, 0x41, 0x78, 0x0c, 0xe4, 0x63, 0x3b, 0x71, 0xfe, 0xd0, 0xee, 0x5d, 0xe6,
	0x9b, 0x6f, 0x66, 0xce, 0xfc, 0xc6, 0x50, 0x11, 0x13, 0x64, 0xbe, 0x33, 0x6d, 0xce, 0x59, 0x20,
	0x02, 0x52, 0x4c, 0xc4, 0x7a, 0x09, 0x19, 0x8b, 0x31, 0xfd, 0x17, 0x50, 0xed, 0x18, 0xed, 0x3a,
	0x3e, 0x32, 0x87, 0x5a, 0x26, 0x29, 0x83, 0xf2, 0xa8, 0x29, 0xa7, 0xca, 0x59, 0x85, 0x2a, 0x8f,
	0x91, 0xf4, 0xa4, 0xe5, 0x62, 0xe9, 0x89, 0x3c, 0x87, 0xfd, 0x3f, 0xbc, 0xb1, 0x98, 0x68, 0x79,
	0x89, 0xc4, 0x02, 0x39, 0x86, 0xc2, 0x04, 0xbd, 0xbb, 0x89, 0xd0, 0x9e, 0x49, 0x38, 0x91, 0xf4,
	0x7f, 0x15, 0xa8, 0xac, 0xb8, 0x27, 0x75, 0x38, 0x18, 0x7b, 0x5c, 0x38, 0x33, 0x17, 0x93, 0x10,
	0x0b, 0x99, 0xe8, 0x50, 0x46, 0xdf, 0xe3, 0xdc, 0x0b, 0x66, 0xd4, 0x11, 0x98, 0x04, 0x5d, 0xc1,
	0xc8, 0x57, 0x90, 0xa7, 0x96, 0x29, 0xa3, 0x1f, 0xb6, 0x5e, 0x34, 0xd3, 0x04, 0xd7, 0x73, 0xa0,
	0x11, 0x8b, 0x9c, 0xc3, 0x51, 0xc8, 0xb1, 0x13, 0x8c, 0x9f, 0xba, 0x81, 0x3f, 0xc7, 0x19, 0x77,
	0x84, 0x17, 0xcc, 0xe4, 0x1b, 0x0f, 0xe8, 0x36, 0x15, 0x79, 0x0d, 0x27, 0x6e, 0x46, 0xb6, 0xd1,
	0x9f, 0x23, 0x73, 0x44, 0xc8, 0x50, 0xdb, 0x3f, 0x55, 0xce, 0xf6, 0xe9, 0x2e, 0xb5, 0xfe, 0xd7,
	0xb3, 0x65, 0xaa, 0xc1, 0xec, 0x9d, 0x77, 0x47, 0xce, 0xa1, 0xe4, 0x4e, 0xd0, 0xbd, 0xbf, 0x0c,
	0xc6, 0x71, 0xae, 0xd5, 0x16, 0x59, 0x3c, 0xb8, 0x9b, 0x6a, 0xe8, 0x92, 0x44, 0x5e, 0x01, 0x48,
	0xc1, 0x62, 0x63, 0x64, 0x32, 0xfd, 0x6a, 0xeb, 0x68, 0xd5, 0x44, 0xaa, 0x68, 0x86, 0x46, 0x06,
	0x50, 0x13, 0xcb, 0x77, 0xf4, 0x03, 0xe6, 0x3b, 0x42, 0xd6, 0xa7, 0xda, 0xaa, 0x2f, 0xeb, 0xb3,
	0xce, 0xa0, 0x9b, 0x46, 0xa4, 0x05, 0xcf, 0x33, 0xa0, 0x3d, 0x61, 0xc8, 0x27, 0xc1, 0x74, 0x9c,
	0xf4, 0x74, 0xab, 0x8e, 0x34, 0x40, 0x75, 0xc2, 0xb1, 0x27, 0xd6, 0x2b, 0x75, 0x40, 0x37, 0x70,
	0xf2, 0x25, 0x54, 0x43, 0x8e, 0x14, 0x7f, 0x43, 0x57, 0x8c, 0x82, 0x70, 0x36, 0xd6, 0x0a, 0x92,
	0xb9, 0x86, 0x92, 0xaf, 0xa1, 0x16, 0x72, 0xb4, 0x1e, 0x90, 0x4d, 0x9d, 0x79, 0x52, 0x53, 0xad,
	0x28, 0xa9, 0x9b, 0x0a, 0xd2, 0x84, 0x82, 0x2b, 0xdb, 0xae, 0x1d, 0xc8, 0xa1, 0x38, 0xde, 0x31,
	0x14, 0x09, 0x8b, 0xbc, 0x86, 0x8a, 0xef, 0xf0, 0xfb, 0x45, 0x03, 0xb4, 0xd2, 0xce, 0xd6, 0xac,
	0x12, 0xc9, 0xf7, 0x40, 0x22, 0xa0, 0x87, 0x02, 0xdd, 0xa8, 0xfd, 0x43, 0x7c, 0xc0, 0xa9, 0x06,
	0xd2, 0xfc, 0x93, 0x85, 0xf9, 0xe5, 0x06, 0x85, 0x6e, 0x31, 0xd3, 0x9b, 0xa0, 0x5e, 0xa0, 0x88,
	0x47, 0x85, 0xe2, 0xef, 0x21, 0x72, 0x21, 0x97, 0x03, 0x1f, 0x3c, 0x17, 0xcd, 0xde, 0x62, 0x39,
	0x12, 0x59, 0xef, 0x42, 0x2d, 0xc3, 0xe7, 0xf3, 0x60, 0xc6, 0x51, 0xe6, 0x2e, 0x11, 0x49, 0xdf,
	0x96, 0x7b, 0xcc, 0x4f, 0x58, 0xfa, 0xaf, 0xa0, 0x8e, 0x3e, 0x20, 0x68, 0xc6, 0x7f, 0xee, 0xbd,
	0xfc, 0x1f, 0x41, 0x6d, 0xb4, 0xfe, 0x48, 0x1d, 0xe1, 0xe3, 0x05, 0x78, 0x19, 0x4e, 0x85, 0x97,
	0x46, 0xfe, 0x14, 0x4a, 0x69, 0x24, 0xae, 0x29, 0xa7, 0xf9, 0xb3, 0x0a, 0x5d, 0x02, 0x1f, 0x1c,
	0xfb, 0x0d, 0x1c, 0xaf, 0x87, 0x49, 0xaa, 0xf4, 0x0d, 0x94, 0x63, 0xb7, 0x06, 0x63, 0x01, 0x8b,
	0x43, 0x1d, 0xb6, 0x48, 0x33, 0xba, 0x82, 0x12, 0x4a, 0x99, 0x74, 0x85, 0xa7, 0xff, 0xa3, 0x40,
	0x35, 0x33, 0xbf, 0xc3, 0xe0, 0x8e, 0x54, 0x21, 0xb7, 0x28, 0x53, 0xce, 0xec, 0x45, 0x29, 0x08,
	0xcf, 0x47, 0x2e, 0x1c, 0x7f, 0x9e, 0xdc, 0xab, 0x25, 0xb0, 0x52, 0xda, 0xfc, 0x5a, 0x69, 0x8f,
	0xa1, 0x10, 0x72, 0x64, 0x66, 0x4f, 0xae, 0x57, 0x89, 0x26, 0x52, 0xe4, 0x11, 0x1f, 0x70, 0x26,
	0xba, 0xd1, 0x68, 0xee, 0xc7, 0x1e, 0x17, 0x00, 0xd1, 0xa0, 0xc8, 0xc3, 0x5b, 0xa9, 0x2b, 0x48,
	0x5d, 0x2a, 0x92, 0x53, 0x38, 0xcc, 0x2c, 0xa8, 0x5c, 0x97, 0x0a, 0xcd, 0x42, 0xfa, 0x9f, 0xa0,
	0x5d, 0xa0, 0x58, 0x4d, 0xe8, 0x7d, 0x86, 0x40, 0x87, 0x32, 0x17, 0x0e, 0x13, 0x46, 0xf4, 0x0a,
	0xb3, 0x97, 0x9e, 0xe5, 0x2c, 0x16, 0x45, 0xf7, 0x9d, 0xc7, 0xab, 0xd0, 0xb7, 0xde, 0x0d, 0x83,
	0xbb, 0x24, 0xd9, 0x2c, 0xa4, 0xdf, 0xc2, 0x8b, 0x2d, 0xd1, 0x93, 0x0e, 0x19, 0x2b, 0x37, 0x4c,
	0x3a, 0x4d, 0xdb, 0x74, 0xb2, 0xed, 0x86, 0x45, 0xb6, 0x9b, 0x16, 0x8d, 0x33, 0x28, 0x2d, 0xb7,
	0xb5, 0x08, 0x79, 0xab, 0xdf, 0x57, 0xf7, 0xc8, 0x01, 0x3c, 0x1b, 0xb4, 0x69, 0x4f, 0x55, 0xa2,
	0x5f, 0x23, 0xab, 0x6f, 0xab, 0xb9, 0xc6, 0x77, 0x00, 0xcb, 0x73, 0x4a, 0xaa, 0x00, 0xed, 0xbe,
	0x6d, 0xd0, 0x9b, 0xf6, 0xb5, 0x3d, 0x50, 0xf7, 0xc8, 0x47, 0x70, 0xd8, 0x31, 0xfa, 0x16, 0x35,
	0x62, 0x40, 0x21, 0x2a, 0x94, 0x7f, 0x32, 0xed, 0x81, 0x75, 0x6d, 0xc7, 0x48, 0xae, 0x71, 0x0e,
	0xb5, 0x8d, 0x9b, 0x1a, 0xf9, 0xe9, 0xb7, 0x07, 0xd4, 0xb8, 0x1a, 0x18, 0xa6, 0xad, 0xee, 0x91,
	0x43, 0x28, 0x76, 0x8d, 0xe1, 0xc8, 0xbc, 0x1e, 0xa9, 0x4a, 0xa3, 0x0f, 0x64, 0xf3, 0x34, 0x44,
	0x94, 0x2b, 0xcb, 0xbe, 0xb9, 0x1e, 0x19, 0xea, 0x1e, 0x01, 0x28, 0x5c, 0x59, 0xf4, 0xb2, 0x3d,
	0x8c, 0xdf, 0x3a, 0x30, 0x2f, 0x06, 0x6a, 0x8e, 0x54, 0xa0, 0xf4, 0xa3, 0x41, 0x7f, 0xbe, 0x91,
	0x62, 0xbe, 0xf5, 0x77, 0x0e, 0x8a, 0xe9, 0xed, 0xeb, 0x40, 0x69, 0x71, 0x14, 0xc8, 0xf2, 0xdf,
	0x70, 0xfd, 0xb0, 0xd4, 0xeb, 0xdb, 0x54, 0x49, 0xed, 0x3b, 0x50, 0x1a, 0x6d, 0xf1, 0x31, 0xda,
	0xed, 0x63, 0x63, 0xc5, 0xc9, 0x0f, 0x50, 0x5d, 0xdd, 0x3d, 0xf2, 0xd9, 0x26, 0x3b, 0xbb, 0xfb,
	0xf5, 0xcf, 0x77, 0xea, 0x13, 0x97, 0x6f, 0xe5, 0xbd, 0x5b, 0x5b, 0xbf, 0x2f, 0xb2, 0x79, 0x6c,
	0x9d, 0xe4, 0xba, 0xfe, 0x7f, 0x94, 0xd8, 0x77, 0xe7, 0x5b, 0xa8, 0xbb, 0x81, 0xdf, 0xe4, 0xe1,
	0x9c, 0xa1, 0xef, 0x78, 0x33, 0xb7, 0xc9, 0xc7, 0xf7, 0xa9, 0xdd, 0x1b, 0xe5, 0xed, 0xc9, 0xad,
	0x17, 0x44, 0xe3, 0xfd, 0x92, 0x23, 0x8b, 0x96, 0xe0, 0x65, 0xa2, 0xba, 0x2d, 0xc8, 0x8f, 0xa6,
	0x57, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x60, 0x26, 0x66, 0x47, 0x59, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThermalClient is the client API for Thermal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThermalClient interface {
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	SetConfigMulti(ctx context.Context, in *SetConfigMultiRequest, opts ...grpc.CallOption) (*SetConfigMultiResponse, error)
	GetTemperatureLog(ctx context.Context, in *GetTemperatureLogRequest, opts ...grpc.CallOption) (*GetTemperatureLogResponse, error)
}

type thermalClient struct {
	cc *grpc.ClientConn
}

func NewThermalClient(cc *grpc.ClientConn) ThermalClient {
	return &thermalClient{cc}
}

func (c *thermalClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/thermal.Thermal/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermalClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/thermal.Thermal/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermalClient) SetConfigMulti(ctx context.Context, in *SetConfigMultiRequest, opts ...grpc.CallOption) (*SetConfigMultiResponse, error) {
	out := new(SetConfigMultiResponse)
	err := c.cc.Invoke(ctx, "/thermal.Thermal/SetConfigMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thermalClient) GetTemperatureLog(ctx context.Context, in *GetTemperatureLogRequest, opts ...grpc.CallOption) (*GetTemperatureLogResponse, error) {
	out := new(GetTemperatureLogResponse)
	err := c.cc.Invoke(ctx, "/thermal.Thermal/GetTemperatureLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThermalServer is the server API for Thermal service.
type ThermalServer interface {
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	SetConfigMulti(context.Context, *SetConfigMultiRequest) (*SetConfigMultiResponse, error)
	GetTemperatureLog(context.Context, *GetTemperatureLogRequest) (*GetTemperatureLogResponse, error)
}

func RegisterThermalServer(s *grpc.Server, srv ThermalServer) {
	s.RegisterService(&_Thermal_serviceDesc, srv)
}

func _Thermal_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermalServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thermal.Thermal/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermalServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thermal_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermalServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thermal.Thermal/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermalServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thermal_SetConfigMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigMultiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermalServer).SetConfigMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thermal.Thermal/SetConfigMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermalServer).SetConfigMulti(ctx, req.(*SetConfigMultiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thermal_GetTemperatureLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemperatureLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThermalServer).GetTemperatureLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thermal.Thermal/GetTemperatureLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThermalServer).GetTemperatureLog(ctx, req.(*GetTemperatureLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Thermal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thermal.Thermal",
	HandlerType: (*ThermalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _Thermal_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Thermal_SetConfig_Handler,
		},
		{
			MethodName: "SetConfigMulti",
			Handler:    _Thermal_SetConfigMulti_Handler,
		},
		{
			MethodName: "GetTemperatureLog",
			Handler:    _Thermal_GetTemperatureLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thermal.proto",
}
