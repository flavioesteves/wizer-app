// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: routine.proto

package proto

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Step struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	ImageUrl    string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
}

func (x *Step) Reset() {
	*x = Step{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Step) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Step) ProtoMessage() {}

func (x *Step) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Step.ProtoReflect.Descriptor instead.
func (*Step) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{0}
}

func (x *Step) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Step) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

type Exercise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MuscleGroup          string  `protobuf:"bytes,3,opt,name=muscle_group,json=muscleGroup,proto3" json:"muscle_group,omitempty"`
	Description          string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Steps                []*Step `protobuf:"bytes,5,rep,name=steps,proto3" json:"steps,omitempty"`
	VideoUrl             string  `protobuf:"bytes,6,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	VideoDurationSeconds int32   `protobuf:"varint,7,opt,name=video_duration_seconds,json=videoDurationSeconds,proto3" json:"video_duration_seconds,omitempty"`
	CreatedBy            string  `protobuf:"bytes,8,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy            string  `protobuf:"bytes,9,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedAt            string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Exercise) Reset() {
	*x = Exercise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exercise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exercise) ProtoMessage() {}

func (x *Exercise) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exercise.ProtoReflect.Descriptor instead.
func (*Exercise) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{1}
}

func (x *Exercise) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Exercise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Exercise) GetMuscleGroup() string {
	if x != nil {
		return x.MuscleGroup
	}
	return ""
}

func (x *Exercise) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Exercise) GetSteps() []*Step {
	if x != nil {
		return x.Steps
	}
	return nil
}

func (x *Exercise) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *Exercise) GetVideoDurationSeconds() int32 {
	if x != nil {
		return x.VideoDurationSeconds
	}
	return 0
}

func (x *Exercise) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Exercise) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Exercise) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Exercise) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Routine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProfileId string      `protobuf:"bytes,3,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Exercises []*Exercise `protobuf:"bytes,4,rep,name=exercises,proto3" json:"exercises,omitempty"`
	CreatedBy string      `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedBy string      `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	CreatedAt string      `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string      `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Routine) Reset() {
	*x = Routine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Routine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Routine) ProtoMessage() {}

func (x *Routine) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Routine.ProtoReflect.Descriptor instead.
func (*Routine) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{2}
}

func (x *Routine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Routine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Routine) GetProfileId() string {
	if x != nil {
		return x.ProfileId
	}
	return ""
}

func (x *Routine) GetExercises() []*Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

func (x *Routine) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Routine) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Routine) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Routine) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateRoutineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routine *Routine `protobuf:"bytes,1,opt,name=routine,proto3" json:"routine,omitempty"`
}

func (x *CreateRoutineRequest) Reset() {
	*x = CreateRoutineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoutineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoutineRequest) ProtoMessage() {}

func (x *CreateRoutineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoutineRequest.ProtoReflect.Descriptor instead.
func (*CreateRoutineRequest) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRoutineRequest) GetRoutine() *Routine {
	if x != nil {
		return x.Routine
	}
	return nil
}

type GetRoutineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRoutineRequest) Reset() {
	*x = GetRoutineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoutineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoutineRequest) ProtoMessage() {}

func (x *GetRoutineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoutineRequest.ProtoReflect.Descriptor instead.
func (*GetRoutineRequest) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{4}
}

func (x *GetRoutineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateRoutineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routine *Routine `protobuf:"bytes,1,opt,name=routine,proto3" json:"routine,omitempty"`
}

func (x *UpdateRoutineRequest) Reset() {
	*x = UpdateRoutineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoutineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoutineRequest) ProtoMessage() {}

func (x *UpdateRoutineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoutineRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoutineRequest) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRoutineRequest) GetRoutine() *Routine {
	if x != nil {
		return x.Routine
	}
	return nil
}

type DeleteRoutineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRoutineRequest) Reset() {
	*x = DeleteRoutineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoutineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoutineRequest) ProtoMessage() {}

func (x *DeleteRoutineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoutineRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoutineRequest) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRoutineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllRoutinesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRoutinesRequest) Reset() {
	*x = GetAllRoutinesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRoutinesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRoutinesRequest) ProtoMessage() {}

func (x *GetAllRoutinesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRoutinesRequest.ProtoReflect.Descriptor instead.
func (*GetAllRoutinesRequest) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{7}
}

type GetAllRoutinesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routines []*Routine `protobuf:"bytes,1,rep,name=routines,proto3" json:"routines,omitempty"`
}

func (x *GetAllRoutinesResponse) Reset() {
	*x = GetAllRoutinesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRoutinesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRoutinesResponse) ProtoMessage() {}

func (x *GetAllRoutinesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRoutinesResponse.ProtoReflect.Descriptor instead.
func (*GetAllRoutinesResponse) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllRoutinesResponse) GetRoutines() []*Routine {
	if x != nil {
		return x.Routines
	}
	return nil
}

type RoutineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routine *Routine `protobuf:"bytes,1,opt,name=routine,proto3" json:"routine,omitempty"`
}

func (x *RoutineResponse) Reset() {
	*x = RoutineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_routine_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutineResponse) ProtoMessage() {}

func (x *RoutineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routine_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutineResponse.ProtoReflect.Descriptor instead.
func (*RoutineResponse) Descriptor() ([]byte, []int) {
	return file_routine_proto_rawDescGZIP(), []int{9}
}

func (x *RoutineResponse) GetRoutine() *Routine {
	if x != nil {
		return x.Routine
	}
	return nil
}

var File_routine_proto protoreflect.FileDescriptor

var file_routine_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x04, 0x53, 0x74, 0x65, 0x70, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xe7, 0x02, 0x0a,
	0x08, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x6d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x75, 0x73, 0x63, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x74, 0x65, 0x70,
	0x52, 0x05, 0x73, 0x74, 0x65, 0x70, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x55, 0x72, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xf9, 0x01, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x09, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x22,
	0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x46, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x08,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x3d, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x07,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x32, 0x83, 0x03, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x12, 0x1a, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x75,
	0x74, 0x69, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a,
	0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x76,
	0x69, 0x6f, 0x65, 0x73, 0x74, 0x65, 0x76, 0x65, 0x73, 0x2f, 0x77, 0x69, 0x7a, 0x65, 0x72, 0x2d,
	0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routine_proto_rawDescOnce sync.Once
	file_routine_proto_rawDescData = file_routine_proto_rawDesc
)

func file_routine_proto_rawDescGZIP() []byte {
	file_routine_proto_rawDescOnce.Do(func() {
		file_routine_proto_rawDescData = protoimpl.X.CompressGZIP(file_routine_proto_rawDescData)
	})
	return file_routine_proto_rawDescData
}

var file_routine_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_routine_proto_goTypes = []any{
	(*Step)(nil),                   // 0: routine.Step
	(*Exercise)(nil),               // 1: routine.Exercise
	(*Routine)(nil),                // 2: routine.Routine
	(*CreateRoutineRequest)(nil),   // 3: routine.CreateRoutineRequest
	(*GetRoutineRequest)(nil),      // 4: routine.GetRoutineRequest
	(*UpdateRoutineRequest)(nil),   // 5: routine.UpdateRoutineRequest
	(*DeleteRoutineRequest)(nil),   // 6: routine.DeleteRoutineRequest
	(*GetAllRoutinesRequest)(nil),  // 7: routine.GetAllRoutinesRequest
	(*GetAllRoutinesResponse)(nil), // 8: routine.GetAllRoutinesResponse
	(*RoutineResponse)(nil),        // 9: routine.RoutineResponse
	(*empty.Empty)(nil),            // 10: google.protobuf.Empty
}
var file_routine_proto_depIdxs = []int32{
	0,  // 0: routine.Exercise.steps:type_name -> routine.Step
	1,  // 1: routine.Routine.exercises:type_name -> routine.Exercise
	2,  // 2: routine.CreateRoutineRequest.routine:type_name -> routine.Routine
	2,  // 3: routine.UpdateRoutineRequest.routine:type_name -> routine.Routine
	2,  // 4: routine.GetAllRoutinesResponse.routines:type_name -> routine.Routine
	2,  // 5: routine.RoutineResponse.routine:type_name -> routine.Routine
	3,  // 6: routine.RoutineService.CreateRoutine:input_type -> routine.CreateRoutineRequest
	4,  // 7: routine.RoutineService.GetRoutine:input_type -> routine.GetRoutineRequest
	5,  // 8: routine.RoutineService.UpdateRoutine:input_type -> routine.UpdateRoutineRequest
	6,  // 9: routine.RoutineService.DeleteRoutine:input_type -> routine.DeleteRoutineRequest
	7,  // 10: routine.RoutineService.GetAllRoutines:input_type -> routine.GetAllRoutinesRequest
	9,  // 11: routine.RoutineService.CreateRoutine:output_type -> routine.RoutineResponse
	9,  // 12: routine.RoutineService.GetRoutine:output_type -> routine.RoutineResponse
	9,  // 13: routine.RoutineService.UpdateRoutine:output_type -> routine.RoutineResponse
	10, // 14: routine.RoutineService.DeleteRoutine:output_type -> google.protobuf.Empty
	8,  // 15: routine.RoutineService.GetAllRoutines:output_type -> routine.GetAllRoutinesResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_routine_proto_init() }
func file_routine_proto_init() {
	if File_routine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_routine_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Step); i {
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
		file_routine_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Exercise); i {
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
		file_routine_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Routine); i {
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
		file_routine_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRoutineRequest); i {
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
		file_routine_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoutineRequest); i {
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
		file_routine_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRoutineRequest); i {
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
		file_routine_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRoutineRequest); i {
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
		file_routine_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllRoutinesRequest); i {
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
		file_routine_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllRoutinesResponse); i {
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
		file_routine_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RoutineResponse); i {
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
			RawDescriptor: file_routine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routine_proto_goTypes,
		DependencyIndexes: file_routine_proto_depIdxs,
		MessageInfos:      file_routine_proto_msgTypes,
	}.Build()
	File_routine_proto = out.File
	file_routine_proto_rawDesc = nil
	file_routine_proto_goTypes = nil
	file_routine_proto_depIdxs = nil
}
