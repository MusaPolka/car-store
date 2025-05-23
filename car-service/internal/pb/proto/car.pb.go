// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/car.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Car struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	BrandId       string                 `protobuf:"bytes,4,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Price         float64                `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty"`
	Year          int32                  `protobuf:"varint,7,opt,name=year,proto3" json:"year,omitempty"`
	Color         string                 `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Car) Reset() {
	*x = Car{}
	mi := &file_proto_car_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Car) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Car) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *Car) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Car) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Car) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Car) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type ListCarsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCarsRequest) Reset() {
	*x = ListCarsRequest{}
	mi := &file_proto_car_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCarsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarsRequest) ProtoMessage() {}

func (x *ListCarsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarsRequest.ProtoReflect.Descriptor instead.
func (*ListCarsRequest) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{1}
}

type ListCarsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cars          []*Car                 `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCarsResponse) Reset() {
	*x = ListCarsResponse{}
	mi := &file_proto_car_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCarsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarsResponse) ProtoMessage() {}

func (x *ListCarsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarsResponse.ProtoReflect.Descriptor instead.
func (*ListCarsResponse) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{2}
}

func (x *ListCarsResponse) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

type GetCarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCarRequest) Reset() {
	*x = GetCarRequest{}
	mi := &file_proto_car_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarRequest) ProtoMessage() {}

func (x *GetCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarRequest.ProtoReflect.Descriptor instead.
func (*GetCarRequest) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{3}
}

func (x *GetCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateCarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	BrandId       string                 `protobuf:"bytes,3,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	Year          int32                  `protobuf:"varint,6,opt,name=year,proto3" json:"year,omitempty"`
	Color         string                 `protobuf:"bytes,7,opt,name=color,proto3" json:"color,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	mi := &file_proto_car_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCarRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCarRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCarRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *CreateCarRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateCarRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CreateCarRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CreateCarRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type DeleteCarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCarRequest) Reset() {
	*x = DeleteCarRequest{}
	mi := &file_proto_car_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarRequest) ProtoMessage() {}

func (x *DeleteCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCarResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCarResponse) Reset() {
	*x = DeleteCarResponse{}
	mi := &file_proto_car_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarResponse) ProtoMessage() {}

func (x *DeleteCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_car_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarResponse.ProtoReflect.Descriptor instead.
func (*DeleteCarResponse) Descriptor() ([]byte, []int) {
	return file_proto_car_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCarResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_car_proto protoreflect.FileDescriptor

const file_proto_car_proto_rawDesc = "" +
	"\n" +
	"\x0fproto/car.proto\x12\x03car\"\xbc\x01\n" +
	"\x03Car\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x19\n" +
	"\bbrand_id\x18\x04 \x01(\tR\abrandId\x12\x14\n" +
	"\x05price\x18\x05 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x06 \x01(\x05R\x05stock\x12\x12\n" +
	"\x04year\x18\a \x01(\x05R\x04year\x12\x14\n" +
	"\x05color\x18\b \x01(\tR\x05color\"\x11\n" +
	"\x0fListCarsRequest\"0\n" +
	"\x10ListCarsResponse\x12\x1c\n" +
	"\x04cars\x18\x01 \x03(\v2\b.car.CarR\x04cars\"\x1f\n" +
	"\rGetCarRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\xb9\x01\n" +
	"\x10CreateCarRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x19\n" +
	"\bbrand_id\x18\x03 \x01(\tR\abrandId\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x12\x14\n" +
	"\x05stock\x18\x05 \x01(\x05R\x05stock\x12\x12\n" +
	"\x04year\x18\x06 \x01(\x05R\x04year\x12\x14\n" +
	"\x05color\x18\a \x01(\tR\x05color\"\"\n" +
	"\x10DeleteCarRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"-\n" +
	"\x11DeleteCarResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xd7\x01\n" +
	"\n" +
	"CarService\x127\n" +
	"\bListCars\x12\x14.car.ListCarsRequest\x1a\x15.car.ListCarsResponse\x12&\n" +
	"\x06GetCar\x12\x12.car.GetCarRequest\x1a\b.car.Car\x12,\n" +
	"\tCreateCar\x12\x15.car.CreateCarRequest\x1a\b.car.Car\x12:\n" +
	"\tDeleteCar\x12\x15.car.DeleteCarRequest\x1a\x16.car.DeleteCarResponseB&Z$ecommerce/car-service/internal/pb;pbb\x06proto3"

var (
	file_proto_car_proto_rawDescOnce sync.Once
	file_proto_car_proto_rawDescData []byte
)

func file_proto_car_proto_rawDescGZIP() []byte {
	file_proto_car_proto_rawDescOnce.Do(func() {
		file_proto_car_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_car_proto_rawDesc), len(file_proto_car_proto_rawDesc)))
	})
	return file_proto_car_proto_rawDescData
}

var file_proto_car_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_car_proto_goTypes = []any{
	(*Car)(nil),               // 0: car.Car
	(*ListCarsRequest)(nil),   // 1: car.ListCarsRequest
	(*ListCarsResponse)(nil),  // 2: car.ListCarsResponse
	(*GetCarRequest)(nil),     // 3: car.GetCarRequest
	(*CreateCarRequest)(nil),  // 4: car.CreateCarRequest
	(*DeleteCarRequest)(nil),  // 5: car.DeleteCarRequest
	(*DeleteCarResponse)(nil), // 6: car.DeleteCarResponse
}
var file_proto_car_proto_depIdxs = []int32{
	0, // 0: car.ListCarsResponse.cars:type_name -> car.Car
	1, // 1: car.CarService.ListCars:input_type -> car.ListCarsRequest
	3, // 2: car.CarService.GetCar:input_type -> car.GetCarRequest
	4, // 3: car.CarService.CreateCar:input_type -> car.CreateCarRequest
	5, // 4: car.CarService.DeleteCar:input_type -> car.DeleteCarRequest
	2, // 5: car.CarService.ListCars:output_type -> car.ListCarsResponse
	0, // 6: car.CarService.GetCar:output_type -> car.Car
	0, // 7: car.CarService.CreateCar:output_type -> car.Car
	6, // 8: car.CarService.DeleteCar:output_type -> car.DeleteCarResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_car_proto_init() }
func file_proto_car_proto_init() {
	if File_proto_car_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_car_proto_rawDesc), len(file_proto_car_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_car_proto_goTypes,
		DependencyIndexes: file_proto_car_proto_depIdxs,
		MessageInfos:      file_proto_car_proto_msgTypes,
	}.Build()
	File_proto_car_proto = out.File
	file_proto_car_proto_goTypes = nil
	file_proto_car_proto_depIdxs = nil
}
