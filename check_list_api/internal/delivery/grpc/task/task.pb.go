// protoc -I api/grpc/proto api/grpc/proto/task.proto --go_out=./internal/delivery/grpc/task --go_opt=paths=source_relative --go-grpc_out=./internal/delivery/grpc/task --go-grpc_opt=paths=source_relative

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.27.3
// source: task.proto

package task

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

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        int64                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Done          bool                   `protobuf:"varint,4,opt,name=Done,proto3" json:"Done,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTaskRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type GetListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        int64                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListRequest) Reset() {
	*x = GetListRequest{}
	mi := &file_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRequest) ProtoMessage() {}

func (x *GetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRequest.ProtoReflect.Descriptor instead.
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *GetListRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        int64                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TaskID        int64                  `protobuf:"varint,2,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTaskRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DeleteTaskRequest) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type DoneTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserID        int64                  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	TaskID        int64                  `protobuf:"varint,2,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	Done          bool                   `protobuf:"varint,3,opt,name=Done,proto3" json:"Done,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoneTaskRequest) Reset() {
	*x = DoneTaskRequest{}
	mi := &file_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoneTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoneTaskRequest) ProtoMessage() {}

func (x *DoneTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoneTaskRequest.ProtoReflect.Descriptor instead.
func (*DoneTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *DoneTaskRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *DoneTaskRequest) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *DoneTaskRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

type GetListResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Tasks         []*GetListResponse_Task `protobuf:"bytes,1,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListResponse) Reset() {
	*x = GetListResponse{}
	mi := &file_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse) ProtoMessage() {}

func (x *GetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse.ProtoReflect.Descriptor instead.
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *GetListResponse) GetTasks() []*GetListResponse_Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type TaskIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskID        int64                  `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskIDResponse) Reset() {
	*x = TaskIDResponse{}
	mi := &file_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskIDResponse) ProtoMessage() {}

func (x *TaskIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskIDResponse.ProtoReflect.Descriptor instead.
func (*TaskIDResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *TaskIDResponse) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

type GetListResponse_Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskID        int64                  `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Done          bool                   `protobuf:"varint,4,opt,name=Done,proto3" json:"Done,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListResponse_Task) Reset() {
	*x = GetListResponse_Task{}
	mi := &file_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListResponse_Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListResponse_Task) ProtoMessage() {}

func (x *GetListResponse_Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListResponse_Task.ProtoReflect.Descriptor instead.
func (*GetListResponse_Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetListResponse_Task) GetTaskID() int64 {
	if x != nil {
		return x.TaskID
	}
	return 0
}

func (x *GetListResponse_Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetListResponse_Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetListResponse_Task) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_task_proto protoreflect.FileDescriptor

var file_task_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x22, 0x77, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x0f, 0x44, 0x6f,
	0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x6f, 0x6e,
	0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x6a, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44,
	0x6f, 0x6e, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x32, 0xf9, 0x01,
	0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x08, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x44, 0x6f, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData = file_task_proto_rawDesc
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_proto_rawDescData)
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),    // 0: task.CreateTaskRequest
	(*GetListRequest)(nil),       // 1: task.GetListRequest
	(*DeleteTaskRequest)(nil),    // 2: task.DeleteTaskRequest
	(*DoneTaskRequest)(nil),      // 3: task.DoneTaskRequest
	(*GetListResponse)(nil),      // 4: task.GetListResponse
	(*TaskIDResponse)(nil),       // 5: task.TaskIDResponse
	(*GetListResponse_Task)(nil), // 6: task.GetListResponse.Task
}
var file_task_proto_depIdxs = []int32{
	6, // 0: task.GetListResponse.Tasks:type_name -> task.GetListResponse.Task
	0, // 1: task.Task.CreateTask:input_type -> task.CreateTaskRequest
	1, // 2: task.Task.GetList:input_type -> task.GetListRequest
	2, // 3: task.Task.DeleteTask:input_type -> task.DeleteTaskRequest
	3, // 4: task.Task.DoneTask:input_type -> task.DoneTaskRequest
	5, // 5: task.Task.CreateTask:output_type -> task.TaskIDResponse
	4, // 6: task.Task.GetList:output_type -> task.GetListResponse
	5, // 7: task.Task.DeleteTask:output_type -> task.TaskIDResponse
	5, // 8: task.Task.DoneTask:output_type -> task.TaskIDResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_rawDesc = nil
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}
