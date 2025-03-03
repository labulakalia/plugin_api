// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: plugin/plugin.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileEntry_FileType int32

const (
	FileEntry_FileTypeUNSPECIFIED FileEntry_FileType = 0
	FileEntry_FileTypeDir         FileEntry_FileType = 1
	FileEntry_FileTypeFile        FileEntry_FileType = 2
	FileEntry_FileTypeLink        FileEntry_FileType = 3
)

// Enum value maps for FileEntry_FileType.
var (
	FileEntry_FileType_name = map[int32]string{
		0: "FileTypeUNSPECIFIED",
		1: "FileTypeDir",
		2: "FileTypeFile",
		3: "FileTypeLink",
	}
	FileEntry_FileType_value = map[string]int32{
		"FileTypeUNSPECIFIED": 0,
		"FileTypeDir":         1,
		"FileTypeFile":        2,
		"FileTypeLink":        3,
	}
)

func (x FileEntry_FileType) Enum() *FileEntry_FileType {
	p := new(FileEntry_FileType)
	*p = x
	return p
}

func (x FileEntry_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileEntry_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_plugin_proto_enumTypes[0].Descriptor()
}

func (FileEntry_FileType) Type() protoreflect.EnumType {
	return &file_plugin_plugin_proto_enumTypes[0]
}

func (x FileEntry_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileEntry_FileType.Descriptor instead.
func (FileEntry_FileType) EnumDescriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{1, 0}
}

// video display VideoResolution
type FileResource_Resolution int32

const (
	FileResource_ResolutionUNSPECIFIED FileResource_Resolution = 0
	FileResource_Original              FileResource_Resolution = 1 // Origin
	FileResource_SD                    FileResource_Resolution = 2 // 480P high defintion
	FileResource_HD                    FileResource_Resolution = 3 // 720P high defintion
	FileResource_FHD                   FileResource_Resolution = 4 // 1080P full high defintion
	FileResource_QHD                   FileResource_Resolution = 5 // 2K  quad high defintion
	FileResource_UHD                   FileResource_Resolution = 6 // 4K ultra high defintion
)

// Enum value maps for FileResource_Resolution.
var (
	FileResource_Resolution_name = map[int32]string{
		0: "ResolutionUNSPECIFIED",
		1: "Original",
		2: "SD",
		3: "HD",
		4: "FHD",
		5: "QHD",
		6: "UHD",
	}
	FileResource_Resolution_value = map[string]int32{
		"ResolutionUNSPECIFIED": 0,
		"Original":              1,
		"SD":                    2,
		"HD":                    3,
		"FHD":                   4,
		"QHD":                   5,
		"UHD":                   6,
	}
)

func (x FileResource_Resolution) Enum() *FileResource_Resolution {
	p := new(FileResource_Resolution)
	*p = x
	return p
}

func (x FileResource_Resolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileResource_Resolution) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_plugin_proto_enumTypes[1].Descriptor()
}

func (FileResource_Resolution) Type() protoreflect.EnumType {
	return &file_plugin_plugin_proto_enumTypes[1]
}

func (x FileResource_Resolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileResource_Resolution.Descriptor instead.
func (FileResource_Resolution) EnumDescriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{3, 0}
}

type FileResource_ResourceType int32

const (
	FileResource_ResourceTypeUNSPECIFIED FileResource_ResourceType = 0
	FileResource_Video                   FileResource_ResourceType = 1
	FileResource_Subtitle                FileResource_ResourceType = 2
	FileResource_Audio                   FileResource_ResourceType = 3
)

// Enum value maps for FileResource_ResourceType.
var (
	FileResource_ResourceType_name = map[int32]string{
		0: "ResourceTypeUNSPECIFIED",
		1: "Video",
		2: "Subtitle",
		3: "Audio",
	}
	FileResource_ResourceType_value = map[string]int32{
		"ResourceTypeUNSPECIFIED": 0,
		"Video":                   1,
		"Subtitle":                2,
		"Audio":                   3,
	}
)

func (x FileResource_ResourceType) Enum() *FileResource_ResourceType {
	p := new(FileResource_ResourceType)
	*p = x
	return p
}

func (x FileResource_ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileResource_ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_plugin_proto_enumTypes[2].Descriptor()
}

func (FileResource_ResourceType) Type() protoreflect.EnumType {
	return &file_plugin_plugin_proto_enumTypes[2]
}

func (x FileResource_ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileResource_ResourceType.Descriptor instead.
func (FileResource_ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{3, 1}
}

type Auth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthMethods   []*anypb.Any           `protobuf:"bytes,1,rep,name=auth_methods,json=authMethods,proto3" json:"auth_methods,omitempty"` //
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth) Reset() {
	*x = Auth{}
	mi := &file_plugin_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[0]
	if x != nil {
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
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetAuthMethods() []*anypb.Any {
	if x != nil {
		return x.AuthMethods
	}
	return nil
}

type FileEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FileType      FileEntry_FileType     `protobuf:"varint,2,opt,name=file_type,json=fileType,proto3,enum=plugin.FileEntry_FileType" json:"file_type,omitempty"`
	Size          uint64                 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	CreatedTime   uint64                 `protobuf:"varint,10,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	ModifiedTime  uint64                 `protobuf:"varint,11,opt,name=modified_time,json=modifiedTime,proto3" json:"modified_time,omitempty"`
	AccessedTime  uint64                 `protobuf:"varint,12,opt,name=accessed_time,json=accessedTime,proto3" json:"accessed_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileEntry) Reset() {
	*x = FileEntry{}
	mi := &file_plugin_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileEntry) ProtoMessage() {}

func (x *FileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileEntry.ProtoReflect.Descriptor instead.
func (*FileEntry) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *FileEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileEntry) GetFileType() FileEntry_FileType {
	if x != nil {
		return x.FileType
	}
	return FileEntry_FileTypeUNSPECIFIED
}

func (x *FileEntry) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileEntry) GetCreatedTime() uint64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

func (x *FileEntry) GetModifiedTime() uint64 {
	if x != nil {
		return x.ModifiedTime
	}
	return 0
}

func (x *FileEntry) GetAccessedTime() uint64 {
	if x != nil {
		return x.AccessedTime
	}
	return 0
}

type DirEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileEntries   []*FileEntry           `protobuf:"bytes,1,rep,name=file_entries,json=fileEntries,proto3" json:"file_entries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DirEntry) Reset() {
	*x = DirEntry{}
	mi := &file_plugin_plugin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DirEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirEntry) ProtoMessage() {}

func (x *DirEntry) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirEntry.ProtoReflect.Descriptor instead.
func (*DirEntry) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *DirEntry) GetFileEntries() []*FileEntry {
	if x != nil {
		return x.FileEntries
	}
	return nil
}

type FileResource struct {
	state            protoimpl.MessageState           `protogen:"open.v1"`
	FileResourceData []*FileResource_FileResourceData `protobuf:"bytes,1,rep,name=file_resource_data,json=fileResourceData,proto3" json:"file_resource_data,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *FileResource) Reset() {
	*x = FileResource{}
	mi := &file_plugin_plugin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResource) ProtoMessage() {}

func (x *FileResource) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResource.ProtoReflect.Descriptor instead.
func (*FileResource) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *FileResource) GetFileResourceData() []*FileResource_FileResourceData {
	if x != nil {
		return x.FileResourceData
	}
	return nil
}

// form data input
type Auth_FormData struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	FormItems     []*Auth_FormData_FormItem `protobuf:"bytes,1,rep,name=form_items,json=formItems,proto3" json:"form_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth_FormData) Reset() {
	*x = Auth_FormData{}
	mi := &file_plugin_plugin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_FormData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_FormData) ProtoMessage() {}

func (x *Auth_FormData) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_FormData.ProtoReflect.Descriptor instead.
func (*Auth_FormData) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Auth_FormData) GetFormItems() []*Auth_FormData_FormItem {
	if x != nil {
		return x.FormItems
	}
	return nil
}

type Auth_ScanQrcode struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	QrcodeImage      []byte                 `protobuf:"bytes,1,opt,name=qrcode_image,json=qrcodeImage,proto3" json:"qrcode_image,omitempty"`                   // qrcode image
	QrcodeImageParam string                 `protobuf:"bytes,2,opt,name=qrcode_image_param,json=qrcodeImageParam,proto3" json:"qrcode_image_param,omitempty"`  // qrcode image some param,like qrcode check key
	QrcodeExpireTime uint64                 `protobuf:"varint,3,opt,name=qrcode_expire_time,json=qrcodeExpireTime,proto3" json:"qrcode_expire_time,omitempty"` // qrcode expire time
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Auth_ScanQrcode) Reset() {
	*x = Auth_ScanQrcode{}
	mi := &file_plugin_plugin_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_ScanQrcode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_ScanQrcode) ProtoMessage() {}

func (x *Auth_ScanQrcode) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_ScanQrcode.ProtoReflect.Descriptor instead.
func (*Auth_ScanQrcode) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Auth_ScanQrcode) GetQrcodeImage() []byte {
	if x != nil {
		return x.QrcodeImage
	}
	return nil
}

func (x *Auth_ScanQrcode) GetQrcodeImageParam() string {
	if x != nil {
		return x.QrcodeImageParam
	}
	return ""
}

func (x *Auth_ScanQrcode) GetQrcodeExpireTime() uint64 {
	if x != nil {
		return x.QrcodeExpireTime
	}
	return 0
}

type Auth_Callback struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	CallbackUrl      string                 `protobuf:"bytes,1,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`                  // callback url
	CallbackUrlParam string                 `protobuf:"bytes,2,opt,name=callback_url_param,json=callbackUrlParam,proto3" json:"callback_url_param,omitempty"` // param for call back url
	CallbackUrlData  string                 `protobuf:"bytes,3,opt,name=callback_url_data,json=callbackUrlData,proto3" json:"callback_url_data,omitempty"`    // url callback data
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Auth_Callback) Reset() {
	*x = Auth_Callback{}
	mi := &file_plugin_plugin_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_Callback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_Callback) ProtoMessage() {}

func (x *Auth_Callback) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_Callback.ProtoReflect.Descriptor instead.
func (*Auth_Callback) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Auth_Callback) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *Auth_Callback) GetCallbackUrlParam() string {
	if x != nil {
		return x.CallbackUrlParam
	}
	return ""
}

func (x *Auth_Callback) GetCallbackUrlData() string {
	if x != nil {
		return x.CallbackUrlData
	}
	return ""
}

type Auth_FormData_FormItem struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Name          string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         *anypb.Any                `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EnumValues    []*Auth_FormData_FormItem `protobuf:"bytes,3,rep,name=enum_values,json=enumValues,proto3" json:"enum_values,omitempty"` // for dropdown widget
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Auth_FormData_FormItem) Reset() {
	*x = Auth_FormData_FormItem{}
	mi := &file_plugin_plugin_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Auth_FormData_FormItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth_FormData_FormItem) ProtoMessage() {}

func (x *Auth_FormData_FormItem) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth_FormData_FormItem.ProtoReflect.Descriptor instead.
func (*Auth_FormData_FormItem) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Auth_FormData_FormItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Auth_FormData_FormItem) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Auth_FormData_FormItem) GetEnumValues() []*Auth_FormData_FormItem {
	if x != nil {
		return x.EnumValues
	}
	return nil
}

type FileResource_FileResourceData struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	Url           string                    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Resolution    FileResource_Resolution   `protobuf:"varint,2,opt,name=resolution,proto3,enum=plugin.FileResource_Resolution" json:"resolution,omitempty"`
	ExpireTime    uint64                    `protobuf:"varint,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	ResourceType  FileResource_ResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=plugin.FileResource_ResourceType" json:"resource_type,omitempty"`
	Title         string                    `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Header        map[string]string         `protobuf:"bytes,6,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileResource_FileResourceData) Reset() {
	*x = FileResource_FileResourceData{}
	mi := &file_plugin_plugin_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileResource_FileResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResource_FileResourceData) ProtoMessage() {}

func (x *FileResource_FileResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResource_FileResourceData.ProtoReflect.Descriptor instead.
func (*FileResource_FileResourceData) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FileResource_FileResourceData) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FileResource_FileResourceData) GetResolution() FileResource_Resolution {
	if x != nil {
		return x.Resolution
	}
	return FileResource_ResolutionUNSPECIFIED
}

func (x *FileResource_FileResourceData) GetExpireTime() uint64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *FileResource_FileResourceData) GetResourceType() FileResource_ResourceType {
	if x != nil {
		return x.ResourceType
	}
	return FileResource_ResourceTypeUNSPECIFIED
}

func (x *FileResource_FileResourceData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FileResource_FileResourceData) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_plugin_plugin_proto protoreflect.FileDescriptor

var file_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0xd7, 0x01, 0x0a, 0x08, 0x46,
	0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x66, 0x6f, 0x72,
	0x6d, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x8b, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x72, 0x6d, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x1a, 0x8b, 0x01, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x51, 0x72, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x71, 0x72, 0x63, 0x6f, 0x64,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x10, 0x71, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x1a, 0x87, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55,
	0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75,
	0x72, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x2a, 0x0a, 0x11, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x22, 0xb3, 0x02, 0x0a,
	0x09, 0x46, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x44, 0x69, 0x72, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x02,
	0x12, 0x10, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x10, 0x03, 0x22, 0x40, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x34,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x83, 0x05, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0xea, 0x02, 0x0a, 0x10, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x3f, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x49, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x60, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x44, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x44, 0x10, 0x03, 0x12, 0x07,
	0x0a, 0x03, 0x46, 0x48, 0x44, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x51, 0x48, 0x44, 0x10, 0x05,
	0x12, 0x07, 0x0a, 0x03, 0x55, 0x48, 0x44, 0x10, 0x06, 0x22, 0x4f, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x10, 0x03, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x6b,
	0x61, 0x6c, 0x69, 0x61, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_plugin_proto_rawDescOnce sync.Once
	file_plugin_plugin_proto_rawDescData = file_plugin_plugin_proto_rawDesc
)

func file_plugin_plugin_proto_rawDescGZIP() []byte {
	file_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_plugin_proto_rawDescData)
	})
	return file_plugin_plugin_proto_rawDescData
}

var file_plugin_plugin_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_plugin_plugin_proto_goTypes = []any{
	(FileEntry_FileType)(0),               // 0: plugin.FileEntry.FileType
	(FileResource_Resolution)(0),          // 1: plugin.FileResource.Resolution
	(FileResource_ResourceType)(0),        // 2: plugin.FileResource.ResourceType
	(*Auth)(nil),                          // 3: plugin.Auth
	(*FileEntry)(nil),                     // 4: plugin.FileEntry
	(*DirEntry)(nil),                      // 5: plugin.DirEntry
	(*FileResource)(nil),                  // 6: plugin.FileResource
	(*Auth_FormData)(nil),                 // 7: plugin.Auth.FormData
	(*Auth_ScanQrcode)(nil),               // 8: plugin.Auth.ScanQrcode
	(*Auth_Callback)(nil),                 // 9: plugin.Auth.Callback
	(*Auth_FormData_FormItem)(nil),        // 10: plugin.Auth.FormData.FormItem
	(*FileResource_FileResourceData)(nil), // 11: plugin.FileResource.FileResourceData
	nil,                                   // 12: plugin.FileResource.FileResourceData.HeaderEntry
	(*anypb.Any)(nil),                     // 13: google.protobuf.Any
}
var file_plugin_plugin_proto_depIdxs = []int32{
	13, // 0: plugin.Auth.auth_methods:type_name -> google.protobuf.Any
	0,  // 1: plugin.FileEntry.file_type:type_name -> plugin.FileEntry.FileType
	4,  // 2: plugin.DirEntry.file_entries:type_name -> plugin.FileEntry
	11, // 3: plugin.FileResource.file_resource_data:type_name -> plugin.FileResource.FileResourceData
	10, // 4: plugin.Auth.FormData.form_items:type_name -> plugin.Auth.FormData.FormItem
	13, // 5: plugin.Auth.FormData.FormItem.value:type_name -> google.protobuf.Any
	10, // 6: plugin.Auth.FormData.FormItem.enum_values:type_name -> plugin.Auth.FormData.FormItem
	1,  // 7: plugin.FileResource.FileResourceData.resolution:type_name -> plugin.FileResource.Resolution
	2,  // 8: plugin.FileResource.FileResourceData.resource_type:type_name -> plugin.FileResource.ResourceType
	12, // 9: plugin.FileResource.FileResourceData.header:type_name -> plugin.FileResource.FileResourceData.HeaderEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_plugin_plugin_proto_init() }
func file_plugin_plugin_proto_init() {
	if File_plugin_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_plugin_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_plugin_proto_depIdxs,
		EnumInfos:         file_plugin_plugin_proto_enumTypes,
		MessageInfos:      file_plugin_plugin_proto_msgTypes,
	}.Build()
	File_plugin_plugin_proto = out.File
	file_plugin_plugin_proto_rawDesc = nil
	file_plugin_plugin_proto_goTypes = nil
	file_plugin_plugin_proto_depIdxs = nil
}
