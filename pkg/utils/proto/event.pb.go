// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: event.proto

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

type SearchEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueName string `protobuf:"bytes,1,opt,name=LeagueName,proto3" json:"LeagueName,omitempty"`
	SportType  string `protobuf:"bytes,2,opt,name=SportType,proto3" json:"SportType,omitempty"`
	StartDate  string `protobuf:"bytes,3,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate    string `protobuf:"bytes,4,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
}

func (x *SearchEventsRequest) Reset() {
	*x = SearchEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchEventsRequest) ProtoMessage() {}

func (x *SearchEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchEventsRequest.ProtoReflect.Descriptor instead.
func (*SearchEventsRequest) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{0}
}

func (x *SearchEventsRequest) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *SearchEventsRequest) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

func (x *SearchEventsRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SearchEventsRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type EventInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeagueName string `protobuf:"bytes,1,opt,name=leagueName,proto3" json:"leagueName,omitempty"`
	RaceTime   string `protobuf:"bytes,2,opt,name=raceTime,proto3" json:"raceTime,omitempty"`
	HomeName   string `protobuf:"bytes,3,opt,name=homeName,proto3" json:"homeName,omitempty"`
	Score      string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	AwayName   string `protobuf:"bytes,5,opt,name=awayName,proto3" json:"awayName,omitempty"`
	HomeOdds   string `protobuf:"bytes,6,opt,name=homeOdds,proto3" json:"homeOdds,omitempty"`
	AwayOdds   string `protobuf:"bytes,7,opt,name=awayOdds,proto3" json:"awayOdds,omitempty"`
	Date       string `protobuf:"bytes,8,opt,name=date,proto3" json:"date,omitempty"`
	Timestamp  int64  `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *EventInfo) Reset() {
	*x = EventInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInfo) ProtoMessage() {}

func (x *EventInfo) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInfo.ProtoReflect.Descriptor instead.
func (*EventInfo) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{1}
}

func (x *EventInfo) GetLeagueName() string {
	if x != nil {
		return x.LeagueName
	}
	return ""
}

func (x *EventInfo) GetRaceTime() string {
	if x != nil {
		return x.RaceTime
	}
	return ""
}

func (x *EventInfo) GetHomeName() string {
	if x != nil {
		return x.HomeName
	}
	return ""
}

func (x *EventInfo) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *EventInfo) GetAwayName() string {
	if x != nil {
		return x.AwayName
	}
	return ""
}

func (x *EventInfo) GetHomeOdds() string {
	if x != nil {
		return x.HomeOdds
	}
	return ""
}

func (x *EventInfo) GetAwayOdds() string {
	if x != nil {
		return x.AwayOdds
	}
	return ""
}

func (x *EventInfo) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *EventInfo) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type EventsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventInfo `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	Count  int32        `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *EventsReply) Reset() {
	*x = EventsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsReply) ProtoMessage() {}

func (x *EventsReply) ProtoReflect() protoreflect.Message {
	mi := &file_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsReply.ProtoReflect.Descriptor instead.
func (*EventsReply) Descriptor() ([]byte, []int) {
	return file_event_proto_rawDescGZIP(), []int{2}
}

func (x *EventsReply) GetEvents() []*EventInfo {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *EventsReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_event_proto protoreflect.FileDescriptor

var file_event_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x6d, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x6d, 0x65, 0x4f, 0x64, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4f, 0x64,
	0x64, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x77, 0x61, 0x79, 0x4f, 0x64,
	0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x4d, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x32, 0x50, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x68,
	0x6f, 0x74, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_proto_rawDescOnce sync.Once
	file_event_proto_rawDescData = file_event_proto_rawDesc
)

func file_event_proto_rawDescGZIP() []byte {
	file_event_proto_rawDescOnce.Do(func() {
		file_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_proto_rawDescData)
	})
	return file_event_proto_rawDescData
}

var file_event_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_proto_goTypes = []interface{}{
	(*SearchEventsRequest)(nil), // 0: event.SearchEventsRequest
	(*EventInfo)(nil),           // 1: event.EventInfo
	(*EventsReply)(nil),         // 2: event.EventsReply
}
var file_event_proto_depIdxs = []int32{
	1, // 0: event.EventsReply.events:type_name -> event.EventInfo
	0, // 1: event.EventService.SearchEvents:input_type -> event.SearchEventsRequest
	2, // 2: event.EventService.SearchEvents:output_type -> event.EventsReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_proto_init() }
func file_event_proto_init() {
	if File_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchEventsRequest); i {
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
		file_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventInfo); i {
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
		file_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsReply); i {
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
			RawDescriptor: file_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_event_proto_goTypes,
		DependencyIndexes: file_event_proto_depIdxs,
		MessageInfos:      file_event_proto_msgTypes,
	}.Build()
	File_event_proto = out.File
	file_event_proto_rawDesc = nil
	file_event_proto_goTypes = nil
	file_event_proto_depIdxs = nil
}