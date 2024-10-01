// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: exercise.proto

package exercise

import (
	model "github.com/flavioesteves/wizer-app/broker/proto/model"
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

type CreateExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *model.Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *CreateExerciseRequest) Reset() {
	*x = CreateExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExerciseRequest) ProtoMessage() {}

func (x *CreateExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExerciseRequest.ProtoReflect.Descriptor instead.
func (*CreateExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExerciseRequest) GetExercise() *model.Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type GetExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetExerciseRequest) Reset() {
	*x = GetExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExerciseRequest) ProtoMessage() {}

func (x *GetExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExerciseRequest.ProtoReflect.Descriptor instead.
func (*GetExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{1}
}

func (x *GetExerciseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *model.Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *UpdateExerciseRequest) Reset() {
	*x = UpdateExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateExerciseRequest) ProtoMessage() {}

func (x *UpdateExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateExerciseRequest.ProtoReflect.Descriptor instead.
func (*UpdateExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateExerciseRequest) GetExercise() *model.Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type DeleteExerciseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteExerciseRequest) Reset() {
	*x = DeleteExerciseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteExerciseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteExerciseRequest) ProtoMessage() {}

func (x *DeleteExerciseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteExerciseRequest.ProtoReflect.Descriptor instead.
func (*DeleteExerciseRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteExerciseRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllExercisesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllExercisesRequest) Reset() {
	*x = GetAllExercisesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllExercisesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExercisesRequest) ProtoMessage() {}

func (x *GetAllExercisesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExercisesRequest.ProtoReflect.Descriptor instead.
func (*GetAllExercisesRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{4}
}

type GetAllExercisesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercises []*model.Exercise `protobuf:"bytes,1,rep,name=exercises,proto3" json:"exercises,omitempty"`
}

func (x *GetAllExercisesResponse) Reset() {
	*x = GetAllExercisesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllExercisesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllExercisesResponse) ProtoMessage() {}

func (x *GetAllExercisesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllExercisesResponse.ProtoReflect.Descriptor instead.
func (*GetAllExercisesResponse) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllExercisesResponse) GetExercises() []*model.Exercise {
	if x != nil {
		return x.Exercises
	}
	return nil
}

type ExerciseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exercise *model.Exercise `protobuf:"bytes,1,opt,name=exercise,proto3" json:"exercise,omitempty"`
}

func (x *ExerciseResponse) Reset() {
	*x = ExerciseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExerciseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExerciseResponse) ProtoMessage() {}

func (x *ExerciseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExerciseResponse.ProtoReflect.Descriptor instead.
func (*ExerciseResponse) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{6}
}

func (x *ExerciseResponse) GetExercise() *model.Exercise {
	if x != nil {
		return x.Exercise
	}
	return nil
}

type GetRoutinesByExerciseIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExerciseId string `protobuf:"bytes,1,opt,name=exercise_id,json=exerciseId,proto3" json:"exercise_id,omitempty"`
}

func (x *GetRoutinesByExerciseIdRequest) Reset() {
	*x = GetRoutinesByExerciseIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoutinesByExerciseIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoutinesByExerciseIdRequest) ProtoMessage() {}

func (x *GetRoutinesByExerciseIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoutinesByExerciseIdRequest.ProtoReflect.Descriptor instead.
func (*GetRoutinesByExerciseIdRequest) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{7}
}

func (x *GetRoutinesByExerciseIdRequest) GetExerciseId() string {
	if x != nil {
		return x.ExerciseId
	}
	return ""
}

type GetRoutinesByExerciseIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routines []*model.Routine `protobuf:"bytes,1,rep,name=routines,proto3" json:"routines,omitempty"`
}

func (x *GetRoutinesByExerciseIdResponse) Reset() {
	*x = GetRoutinesByExerciseIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exercise_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoutinesByExerciseIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoutinesByExerciseIdResponse) ProtoMessage() {}

func (x *GetRoutinesByExerciseIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exercise_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoutinesByExerciseIdResponse.ProtoReflect.Descriptor instead.
func (*GetRoutinesByExerciseIdResponse) Descriptor() ([]byte, []int) {
	return file_exercise_proto_rawDescGZIP(), []int{8}
}

func (x *GetRoutinesByExerciseIdResponse) GetRoutines() []*model.Routine {
	if x != nil {
		return x.Routines
	}
	return nil
}

var File_exercise_proto protoreflect.FileDescriptor

var file_exercise_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x52, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x44, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x08, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x09, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x08, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x22, 0x41, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x73, 0x42, 0x79, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69,
	0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65,
	0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x52, 0x08, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x32, 0x8b, 0x04, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65,
	0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73,
	0x65, 0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x2e, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x56, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x12,
	0x20, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x65, 0x73, 0x42, 0x79, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12,
	0x28, 0x2e, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x42, 0x79, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x65, 0x78, 0x65, 0x72,
	0x63, 0x69, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73,
	0x42, 0x79, 0x45, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x76, 0x69, 0x6f, 0x65, 0x73, 0x74, 0x65, 0x76, 0x65, 0x73,
	0x2f, 0x77, 0x69, 0x7a, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x65, 0x72, 0x63, 0x69, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exercise_proto_rawDescOnce sync.Once
	file_exercise_proto_rawDescData = file_exercise_proto_rawDesc
)

func file_exercise_proto_rawDescGZIP() []byte {
	file_exercise_proto_rawDescOnce.Do(func() {
		file_exercise_proto_rawDescData = protoimpl.X.CompressGZIP(file_exercise_proto_rawDescData)
	})
	return file_exercise_proto_rawDescData
}

var file_exercise_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_exercise_proto_goTypes = []any{
	(*CreateExerciseRequest)(nil),           // 0: exercise.CreateExerciseRequest
	(*GetExerciseRequest)(nil),              // 1: exercise.GetExerciseRequest
	(*UpdateExerciseRequest)(nil),           // 2: exercise.UpdateExerciseRequest
	(*DeleteExerciseRequest)(nil),           // 3: exercise.DeleteExerciseRequest
	(*GetAllExercisesRequest)(nil),          // 4: exercise.GetAllExercisesRequest
	(*GetAllExercisesResponse)(nil),         // 5: exercise.GetAllExercisesResponse
	(*ExerciseResponse)(nil),                // 6: exercise.ExerciseResponse
	(*GetRoutinesByExerciseIdRequest)(nil),  // 7: exercise.GetRoutinesByExerciseIdRequest
	(*GetRoutinesByExerciseIdResponse)(nil), // 8: exercise.GetRoutinesByExerciseIdResponse
	(*model.Exercise)(nil),                  // 9: model.Exercise
	(*model.Routine)(nil),                   // 10: model.Routine
	(*empty.Empty)(nil),                     // 11: google.protobuf.Empty
}
var file_exercise_proto_depIdxs = []int32{
	9,  // 0: exercise.CreateExerciseRequest.exercise:type_name -> model.Exercise
	9,  // 1: exercise.UpdateExerciseRequest.exercise:type_name -> model.Exercise
	9,  // 2: exercise.GetAllExercisesResponse.exercises:type_name -> model.Exercise
	9,  // 3: exercise.ExerciseResponse.exercise:type_name -> model.Exercise
	10, // 4: exercise.GetRoutinesByExerciseIdResponse.routines:type_name -> model.Routine
	0,  // 5: exercise.ExerciseService.CreateExercise:input_type -> exercise.CreateExerciseRequest
	1,  // 6: exercise.ExerciseService.GetExercise:input_type -> exercise.GetExerciseRequest
	2,  // 7: exercise.ExerciseService.UpdateExercise:input_type -> exercise.UpdateExerciseRequest
	3,  // 8: exercise.ExerciseService.DeleteExercise:input_type -> exercise.DeleteExerciseRequest
	4,  // 9: exercise.ExerciseService.GetAllExercises:input_type -> exercise.GetAllExercisesRequest
	7,  // 10: exercise.ExerciseService.GetRoutinesByExerciseId:input_type -> exercise.GetRoutinesByExerciseIdRequest
	6,  // 11: exercise.ExerciseService.CreateExercise:output_type -> exercise.ExerciseResponse
	6,  // 12: exercise.ExerciseService.GetExercise:output_type -> exercise.ExerciseResponse
	6,  // 13: exercise.ExerciseService.UpdateExercise:output_type -> exercise.ExerciseResponse
	11, // 14: exercise.ExerciseService.DeleteExercise:output_type -> google.protobuf.Empty
	5,  // 15: exercise.ExerciseService.GetAllExercises:output_type -> exercise.GetAllExercisesResponse
	8,  // 16: exercise.ExerciseService.GetRoutinesByExerciseId:output_type -> exercise.GetRoutinesByExerciseIdResponse
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_exercise_proto_init() }
func file_exercise_proto_init() {
	if File_exercise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exercise_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateExerciseRequest); i {
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
		file_exercise_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetExerciseRequest); i {
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
		file_exercise_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateExerciseRequest); i {
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
		file_exercise_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteExerciseRequest); i {
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
		file_exercise_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllExercisesRequest); i {
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
		file_exercise_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllExercisesResponse); i {
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
		file_exercise_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ExerciseResponse); i {
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
		file_exercise_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoutinesByExerciseIdRequest); i {
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
		file_exercise_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetRoutinesByExerciseIdResponse); i {
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
			RawDescriptor: file_exercise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exercise_proto_goTypes,
		DependencyIndexes: file_exercise_proto_depIdxs,
		MessageInfos:      file_exercise_proto_msgTypes,
	}.Build()
	File_exercise_proto = out.File
	file_exercise_proto_rawDesc = nil
	file_exercise_proto_goTypes = nil
	file_exercise_proto_depIdxs = nil
}
