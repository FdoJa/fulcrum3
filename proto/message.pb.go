// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: message.proto

package proto

import (
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

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector   string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base     string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	Soldados string `protobuf:"bytes,3,opt,name=soldados,proto3" json:"soldados,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *Base) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *Base) GetSoldados() string {
	if x != nil {
		return x.Soldados
	}
	return ""
}

type BaseModificada struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector            string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base              string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	ActualizacionBase string `protobuf:"bytes,3,opt,name=actualizacion_base,json=actualizacionBase,proto3" json:"actualizacion_base,omitempty"`
}

func (x *BaseModificada) Reset() {
	*x = BaseModificada{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseModificada) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseModificada) ProtoMessage() {}

func (x *BaseModificada) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseModificada.ProtoReflect.Descriptor instead.
func (*BaseModificada) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *BaseModificada) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *BaseModificada) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *BaseModificada) GetActualizacionBase() string {
	if x != nil {
		return x.ActualizacionBase
	}
	return ""
}

type ActualizarSoldados struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sector                string `protobuf:"bytes,1,opt,name=sector,proto3" json:"sector,omitempty"`
	Base                  string `protobuf:"bytes,2,opt,name=base,proto3" json:"base,omitempty"`
	ActualizacionSoldados string `protobuf:"bytes,3,opt,name=actualizacion_soldados,json=actualizacionSoldados,proto3" json:"actualizacion_soldados,omitempty"`
}

func (x *ActualizarSoldados) Reset() {
	*x = ActualizarSoldados{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActualizarSoldados) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActualizarSoldados) ProtoMessage() {}

func (x *ActualizarSoldados) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActualizarSoldados.ProtoReflect.Descriptor instead.
func (*ActualizarSoldados) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *ActualizarSoldados) GetSector() string {
	if x != nil {
		return x.Sector
	}
	return ""
}

func (x *ActualizarSoldados) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *ActualizarSoldados) GetActualizacionSoldados() string {
	if x != nil {
		return x.ActualizacionSoldados
	}
	return ""
}

type Recepcion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok string `protobuf:"bytes,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *Recepcion) Reset() {
	*x = Recepcion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recepcion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recepcion) ProtoMessage() {}

func (x *Recepcion) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recepcion.ProtoReflect.Descriptor instead.
func (*Recepcion) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Recepcion) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

type Vector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S1 int32 `protobuf:"varint,1,opt,name=s1,proto3" json:"s1,omitempty"`
	S2 int32 `protobuf:"varint,2,opt,name=s2,proto3" json:"s2,omitempty"`
	S3 int32 `protobuf:"varint,3,opt,name=s3,proto3" json:"s3,omitempty"`
}

func (x *Vector) Reset() {
	*x = Vector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vector) ProtoMessage() {}

func (x *Vector) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vector.ProtoReflect.Descriptor instead.
func (*Vector) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Vector) GetS1() int32 {
	if x != nil {
		return x.S1
	}
	return 0
}

func (x *Vector) GetS2() int32 {
	if x != nil {
		return x.S2
	}
	return 0
}

func (x *Vector) GetS3() int32 {
	if x != nil {
		return x.S3
	}
	return 0
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tiempo         string `protobuf:"bytes,1,opt,name=tiempo,proto3" json:"tiempo,omitempty"`
	Accion         string `protobuf:"bytes,2,opt,name=accion,proto3" json:"accion,omitempty"`
	SectorAfectado string `protobuf:"bytes,3,opt,name=sectorAfectado,proto3" json:"sectorAfectado,omitempty"`
	BaseAfectada   string `protobuf:"bytes,4,opt,name=baseAfectada,proto3" json:"baseAfectada,omitempty"`
	NuevoValor     string `protobuf:"bytes,5,opt,name=nuevoValor,proto3" json:"nuevoValor,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *Log) GetTiempo() string {
	if x != nil {
		return x.Tiempo
	}
	return ""
}

func (x *Log) GetAccion() string {
	if x != nil {
		return x.Accion
	}
	return ""
}

func (x *Log) GetSectorAfectado() string {
	if x != nil {
		return x.SectorAfectado
	}
	return ""
}

func (x *Log) GetBaseAfectada() string {
	if x != nil {
		return x.BaseAfectada
	}
	return ""
}

func (x *Log) GetNuevoValor() string {
	if x != nil {
		return x.NuevoValor
	}
	return ""
}

type LogList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaLogs []*Log  `protobuf:"bytes,1,rep,name=listaLogs,proto3" json:"listaLogs,omitempty"`
	Vector    []int32 `protobuf:"varint,2,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *LogList) Reset() {
	*x = LogList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogList) ProtoMessage() {}

func (x *LogList) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogList.ProtoReflect.Descriptor instead.
func (*LogList) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *LogList) GetListaLogs() []*Log {
	if x != nil {
		return x.ListaLogs
	}
	return nil
}

func (x *LogList) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

type Datos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListaBases []*Base `protobuf:"bytes,1,rep,name=listaBases,proto3" json:"listaBases,omitempty"`
	Vector     []int32 `protobuf:"varint,2,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *Datos) Reset() {
	*x = Datos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Datos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Datos) ProtoMessage() {}

func (x *Datos) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Datos.ProtoReflect.Descriptor instead.
func (*Datos) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *Datos) GetListaBases() []*Base {
	if x != nil {
		return x.ListaBases
	}
	return nil
}

func (x *Datos) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4e, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x6c,
	0x64, 0x61, 0x64, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c,
	0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0x6b, 0x0a, 0x0e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x73, 0x65, 0x22, 0x77, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72,
	0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x16, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x63, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x52,
	0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x38, 0x0a, 0x06, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x73, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x73, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x73, 0x33, 0x22, 0xa1, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69,
	0x65, 0x6d, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6d,
	0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61,
	0x64, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x66, 0x65, 0x63, 0x74, 0x61,
	0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x66,
	0x65, 0x63, 0x74, 0x61, 0x64, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x75, 0x65, 0x76,
	0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x22, 0x4a, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x09, 0x6c, 0x69, 0x73, 0x74, 0x61, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x22, 0x4b, 0x0a, 0x05, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x6c,
	0x69, 0x73, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x0a, 0x6c, 0x69, 0x73,
	0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32,
	0xce, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x67, 0x61, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0a,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x6f,
	0x6d, 0x62, 0x72, 0x61, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x64, 0x61, 0x1a,
	0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x39, 0x0a,
	0x0f, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x6f, 0x72,
	0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x72, 0x53, 0x6f, 0x6c, 0x64, 0x61, 0x64, 0x6f, 0x73, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72,
	0x61, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x32, 0x78, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x69, 0x61,
	0x12, 0x2f, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x67, 0x75, 0x69, 0x72, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69,
	0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x37, 0x0a, 0x17, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x44, 0x61, 0x74, 0x6f, 0x73,
	0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x64, 0x6f, 0x73, 0x12, 0x0b, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x6f, 0x73, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x63, 0x65, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x46, 0x64, 0x6f, 0x4a, 0x61, 0x2f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []interface{}{
	(*Base)(nil),               // 0: grpc.Base
	(*BaseModificada)(nil),     // 1: grpc.BaseModificada
	(*ActualizarSoldados)(nil), // 2: grpc.ActualizarSoldados
	(*Recepcion)(nil),          // 3: grpc.Recepcion
	(*Vector)(nil),             // 4: grpc.Vector
	(*Log)(nil),                // 5: grpc.Log
	(*LogList)(nil),            // 6: grpc.LogList
	(*Datos)(nil),              // 7: grpc.Datos
}
var file_message_proto_depIdxs = []int32{
	5, // 0: grpc.LogList.listaLogs:type_name -> grpc.Log
	0, // 1: grpc.Datos.listaBases:type_name -> grpc.Base
	0, // 2: grpc.Informantes.AgregarBase:input_type -> grpc.Base
	1, // 3: grpc.Informantes.RenombrarBase:input_type -> grpc.BaseModificada
	2, // 4: grpc.Informantes.ActualizarValor:input_type -> grpc.ActualizarSoldados
	0, // 5: grpc.Informantes.BorrarBase:input_type -> grpc.Base
	3, // 6: grpc.Consistencia.ConseguirLogs:input_type -> grpc.Recepcion
	7, // 7: grpc.Consistencia.EnviarDatosActualizados:input_type -> grpc.Datos
	4, // 8: grpc.Informantes.AgregarBase:output_type -> grpc.Vector
	4, // 9: grpc.Informantes.RenombrarBase:output_type -> grpc.Vector
	4, // 10: grpc.Informantes.ActualizarValor:output_type -> grpc.Vector
	4, // 11: grpc.Informantes.BorrarBase:output_type -> grpc.Vector
	6, // 12: grpc.Consistencia.ConseguirLogs:output_type -> grpc.LogList
	3, // 13: grpc.Consistencia.EnviarDatosActualizados:output_type -> grpc.Recepcion
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseModificada); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActualizarSoldados); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recepcion); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vector); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogList); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Datos); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
