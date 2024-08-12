// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: budget.proto

package genprotos

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

type CreateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Budget *Budget `protobuf:"bytes,1,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *CreateBudgetRequest) Reset() {
	*x = CreateBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetRequest) ProtoMessage() {}

func (x *CreateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetRequest.ProtoReflect.Descriptor instead.
func (*CreateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type CreateBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBudgetResponse) Reset() {
	*x = CreateBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBudgetResponse) ProtoMessage() {}

func (x *CreateBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBudgetResponse.ProtoReflect.Descriptor instead.
func (*CreateBudgetResponse) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{1}
}

type Budget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CategoryId string  `protobuf:"bytes,3,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Period     string  `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	Amount     float32 `protobuf:"fixed32,5,opt,name=amount,proto3" json:"amount,omitempty"`
	StartDate  string  `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate    string  `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Budget) Reset() {
	*x = Budget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Budget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Budget) ProtoMessage() {}

func (x *Budget) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Budget.ProtoReflect.Descriptor instead.
func (*Budget) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{2}
}

func (x *Budget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Budget) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Budget) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

func (x *Budget) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Budget) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Budget) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Budget) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type UpdateBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Budget *Budget `protobuf:"bytes,1,opt,name=budget,proto3" json:"budget,omitempty"`
}

func (x *UpdateBudgetRequest) Reset() {
	*x = UpdateBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBudgetRequest) ProtoMessage() {}

func (x *UpdateBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBudgetRequest.ProtoReflect.Descriptor instead.
func (*UpdateBudgetRequest) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBudgetRequest) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

type UpdateBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateBudgetResponse) Reset() {
	*x = UpdateBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBudgetResponse) ProtoMessage() {}

func (x *UpdateBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBudgetResponse.ProtoReflect.Descriptor instead.
func (*UpdateBudgetResponse) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{4}
}

type GetBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBudgetRequest) Reset() {
	*x = GetBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetRequest) ProtoMessage() {}

func (x *GetBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetRequest.ProtoReflect.Descriptor instead.
func (*GetBudgetRequest) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{5}
}

func (x *GetBudgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Budget     *Budget   `protobuf:"bytes,1,opt,name=budget,proto3" json:"budget,omitempty"`
	Categories *Category `protobuf:"bytes,2,opt,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetBudgetResponse) Reset() {
	*x = GetBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBudgetResponse) ProtoMessage() {}

func (x *GetBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBudgetResponse.ProtoReflect.Descriptor instead.
func (*GetBudgetResponse) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{6}
}

func (x *GetBudgetResponse) GetBudget() *Budget {
	if x != nil {
		return x.Budget
	}
	return nil
}

func (x *GetBudgetResponse) GetCategories() *Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type DeleteBudgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBudgetRequest) Reset() {
	*x = DeleteBudgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBudgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBudgetRequest) ProtoMessage() {}

func (x *DeleteBudgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBudgetRequest.ProtoReflect.Descriptor instead.
func (*DeleteBudgetRequest) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBudgetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteBudgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBudgetResponse) Reset() {
	*x = DeleteBudgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBudgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBudgetResponse) ProtoMessage() {}

func (x *DeleteBudgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBudgetResponse.ProtoReflect.Descriptor instead.
func (*DeleteBudgetResponse) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{8}
}

type ListBudgetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListBudgetsRequest) Reset() {
	*x = ListBudgetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBudgetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsRequest) ProtoMessage() {}

func (x *ListBudgetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsRequest.ProtoReflect.Descriptor instead.
func (*ListBudgetsRequest) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{9}
}

func (x *ListBudgetsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListBudgetsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListBudgetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Budgets []*GetBudgetResponse `protobuf:"bytes,1,rep,name=budgets,proto3" json:"budgets,omitempty"`
}

func (x *ListBudgetsResponse) Reset() {
	*x = ListBudgetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_budget_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBudgetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBudgetsResponse) ProtoMessage() {}

func (x *ListBudgetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_budget_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBudgetsResponse.ProtoReflect.Descriptor instead.
func (*ListBudgetsResponse) Descriptor() ([]byte, []int) {
	return file_budget_proto_rawDescGZIP(), []int{10}
}

func (x *ListBudgetsResponse) GetBudgets() []*GetBudgetResponse {
	if x != nil {
		return x.Budgets
	}
	return nil
}

var File_budget_proto protoreflect.FileDescriptor

var file_budget_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x06, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xbc, 0x01,
	0x0a, 0x06, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3d, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06,
	0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x06, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x4a, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x07, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x73, 0x32, 0x84, 0x03, 0x0a, 0x0d, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12,
	0x1b, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73, 0x12,
	0x1a, 0x2e, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_budget_proto_rawDescOnce sync.Once
	file_budget_proto_rawDescData = file_budget_proto_rawDesc
)

func file_budget_proto_rawDescGZIP() []byte {
	file_budget_proto_rawDescOnce.Do(func() {
		file_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_budget_proto_rawDescData)
	})
	return file_budget_proto_rawDescData
}

var file_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_budget_proto_goTypes = []interface{}{
	(*CreateBudgetRequest)(nil),  // 0: budget.CreateBudgetRequest
	(*CreateBudgetResponse)(nil), // 1: budget.CreateBudgetResponse
	(*Budget)(nil),               // 2: budget.Budget
	(*UpdateBudgetRequest)(nil),  // 3: budget.UpdateBudgetRequest
	(*UpdateBudgetResponse)(nil), // 4: budget.UpdateBudgetResponse
	(*GetBudgetRequest)(nil),     // 5: budget.GetBudgetRequest
	(*GetBudgetResponse)(nil),    // 6: budget.GetBudgetResponse
	(*DeleteBudgetRequest)(nil),  // 7: budget.DeleteBudgetRequest
	(*DeleteBudgetResponse)(nil), // 8: budget.DeleteBudgetResponse
	(*ListBudgetsRequest)(nil),   // 9: budget.ListBudgetsRequest
	(*ListBudgetsResponse)(nil),  // 10: budget.ListBudgetsResponse
	(*Category)(nil),             // 11: category.Category
}
var file_budget_proto_depIdxs = []int32{
	2,  // 0: budget.CreateBudgetRequest.budget:type_name -> budget.Budget
	2,  // 1: budget.UpdateBudgetRequest.budget:type_name -> budget.Budget
	2,  // 2: budget.GetBudgetResponse.budget:type_name -> budget.Budget
	11, // 3: budget.GetBudgetResponse.categories:type_name -> category.Category
	6,  // 4: budget.ListBudgetsResponse.budgets:type_name -> budget.GetBudgetResponse
	0,  // 5: budget.BudgetService.CreateBudget:input_type -> budget.CreateBudgetRequest
	3,  // 6: budget.BudgetService.UpdateBudget:input_type -> budget.UpdateBudgetRequest
	7,  // 7: budget.BudgetService.DeleteBudget:input_type -> budget.DeleteBudgetRequest
	5,  // 8: budget.BudgetService.GetBudget:input_type -> budget.GetBudgetRequest
	9,  // 9: budget.BudgetService.ListBudgets:input_type -> budget.ListBudgetsRequest
	1,  // 10: budget.BudgetService.CreateBudget:output_type -> budget.CreateBudgetResponse
	4,  // 11: budget.BudgetService.UpdateBudget:output_type -> budget.UpdateBudgetResponse
	8,  // 12: budget.BudgetService.DeleteBudget:output_type -> budget.DeleteBudgetResponse
	6,  // 13: budget.BudgetService.GetBudget:output_type -> budget.GetBudgetResponse
	10, // 14: budget.BudgetService.ListBudgets:output_type -> budget.ListBudgetsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_budget_proto_init() }
func file_budget_proto_init() {
	if File_budget_proto != nil {
		return
	}
	file_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_budget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBudgetRequest); i {
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
		file_budget_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBudgetResponse); i {
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
		file_budget_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Budget); i {
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
		file_budget_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBudgetRequest); i {
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
		file_budget_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBudgetResponse); i {
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
		file_budget_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBudgetRequest); i {
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
		file_budget_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBudgetResponse); i {
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
		file_budget_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBudgetRequest); i {
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
		file_budget_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBudgetResponse); i {
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
		file_budget_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBudgetsRequest); i {
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
		file_budget_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBudgetsResponse); i {
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
			RawDescriptor: file_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_budget_proto_goTypes,
		DependencyIndexes: file_budget_proto_depIdxs,
		MessageInfos:      file_budget_proto_msgTypes,
	}.Build()
	File_budget_proto = out.File
	file_budget_proto_rawDesc = nil
	file_budget_proto_goTypes = nil
	file_budget_proto_depIdxs = nil
}