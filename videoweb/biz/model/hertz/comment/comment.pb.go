// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.5
// source: comment.proto

package comment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "videoweb/biz/model/api"
	common "videoweb/biz/model/hertz/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommentPublishReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid     *string `protobuf:"bytes,1,opt,name=Vid,proto3,oneof" form:"vid" json:"vid,omitempty"`
	Cid     *string `protobuf:"bytes,2,opt,name=Cid,proto3,oneof" form:"cid" json:"cid,omitempty"`
	Content string  `protobuf:"bytes,3,opt,name=Content,proto3" form:"content" json:"content,omitempty"`
}

func (x *CommentPublishReq) Reset() {
	*x = CommentPublishReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentPublishReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentPublishReq) ProtoMessage() {}

func (x *CommentPublishReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentPublishReq.ProtoReflect.Descriptor instead.
func (*CommentPublishReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *CommentPublishReq) GetVid() string {
	if x != nil && x.Vid != nil {
		return *x.Vid
	}
	return ""
}

func (x *CommentPublishReq) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

func (x *CommentPublishReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CommentListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid      *string `protobuf:"bytes,1,opt,name=Vid,proto3,oneof" json:"Vid,omitempty" query:"vid"`
	Cid      *string `protobuf:"bytes,2,opt,name=Cid,proto3,oneof" json:"Cid,omitempty" query:"cid"`
	PageSize *int64  `protobuf:"varint,3,opt,name=PageSize,proto3,oneof" json:"PageSize,omitempty" query:"page_size"`
	PageNum  *int64  `protobuf:"varint,4,opt,name=PageNum,proto3,oneof" json:"PageNum,omitempty" query:"page_num"`
}

func (x *CommentListReq) Reset() {
	*x = CommentListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListReq) ProtoMessage() {}

func (x *CommentListReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListReq.ProtoReflect.Descriptor instead.
func (*CommentListReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentListReq) GetVid() string {
	if x != nil && x.Vid != nil {
		return *x.Vid
	}
	return ""
}

func (x *CommentListReq) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

func (x *CommentListReq) GetPageSize() int64 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

func (x *CommentListReq) GetPageNum() int64 {
	if x != nil && x.PageNum != nil {
		return *x.PageNum
	}
	return 0
}

type CommentDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid *string `protobuf:"bytes,1,opt,name=Vid,proto3,oneof" form:"vid" json:"vid,omitempty"`
	Cid *string `protobuf:"bytes,2,opt,name=Cid,proto3,oneof" form:"cid" json:"cid,omitempty"`
}

func (x *CommentDeleteReq) Reset() {
	*x = CommentDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDeleteReq) ProtoMessage() {}

func (x *CommentDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDeleteReq.ProtoReflect.Descriptor instead.
func (*CommentDeleteReq) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentDeleteReq) GetVid() string {
	if x != nil && x.Vid != nil {
		return *x.Vid
	}
	return ""
}

func (x *CommentDeleteReq) GetCid() string {
	if x != nil && x.Cid != nil {
		return *x.Cid
	}
	return ""
}

type CommentItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string `protobuf:"bytes,1,opt,name=Uid,proto3" form:"uid" json:"uid,omitempty"`
	Vid        string `protobuf:"bytes,2,opt,name=Vid,proto3" form:"vid" json:"vid,omitempty"`
	Cid        string `protobuf:"bytes,3,opt,name=Cid,proto3" form:"cid" json:"cid,omitempty"`
	ParentId   string `protobuf:"bytes,4,opt,name=ParentId,proto3" form:"parent_id" json:"parent_id,omitempty"`
	LikeCount  int64  `protobuf:"varint,5,opt,name=LikeCount,proto3" form:"like_count" json:"like_count,omitempty"`
	ChildCount int64  `protobuf:"varint,6,opt,name=ChildCount,proto3" form:"child_count" json:"child_count,omitempty"`
	Content    string `protobuf:"bytes,7,opt,name=Content,proto3" form:"content" json:"content,omitempty"`
	CreatedAt  string `protobuf:"bytes,10,opt,name=CreatedAt,proto3" form:"created_at" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" form:"updated_at" json:"updated_at,omitempty"`
	DeletedAt  string `protobuf:"bytes,12,opt,name=DeletedAt,proto3" form:"deleted_at" json:"deleted_at,omitempty"`
}

func (x *CommentItems) Reset() {
	*x = CommentItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentItems) ProtoMessage() {}

func (x *CommentItems) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentItems.ProtoReflect.Descriptor instead.
func (*CommentItems) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentItems) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CommentItems) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

func (x *CommentItems) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *CommentItems) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *CommentItems) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *CommentItems) GetChildCount() int64 {
	if x != nil {
		return x.ChildCount
	}
	return 0
}

func (x *CommentItems) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentItems) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommentItems) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CommentItems) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type CommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *common.BaseResponse `protobuf:"bytes,1,opt,name=Base,proto3" form:"base" json:"base,omitempty"`
	Items []*CommentItems      `protobuf:"bytes,2,rep,name=Items,proto3" form:"items" json:"items,omitempty"`
}

func (x *CommentResponse) Reset() {
	*x = CommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentResponse) ProtoMessage() {}

func (x *CommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentResponse.ProtoReflect.Descriptor instead.
func (*CommentResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentResponse) GetBase() *common.BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *CommentResponse) GetItems() []*CommentItems {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8a, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x56, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x76, 0x69, 0x64, 0x48, 0x00, 0x52, 0x03, 0x56,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x03, 0x43, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x63, 0x69, 0x64, 0x48, 0x01, 0x52, 0x03, 0x43,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x56, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x43, 0x69, 0x64, 0x22, 0xd6, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1e, 0x0a, 0x03, 0x56, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xb2, 0xbb,
	0x18, 0x03, 0x76, 0x69, 0x64, 0x48, 0x00, 0x52, 0x03, 0x56, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1e, 0x0a, 0x03, 0x43, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xb2, 0xbb,
	0x18, 0x03, 0x63, 0x69, 0x64, 0x48, 0x01, 0x52, 0x03, 0x43, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x42, 0x0d, 0xb2, 0xbb, 0x18, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x48, 0x02, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2b, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x0c, 0xb2, 0xbb, 0x18, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x48, 0x03,
	0x52, 0x07, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04,
	0x5f, 0x56, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x43, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x50, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x61,
	0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x62, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x03, 0x56, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x76, 0x69, 0x64, 0x48,
	0x00, 0x52, 0x03, 0x56, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x03, 0x43, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x63, 0x69, 0x64, 0x48,
	0x01, 0x52, 0x03, 0x43, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x56, 0x69,
	0x64, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x43, 0x69, 0x64, 0x22, 0x9a, 0x03, 0x0a, 0x0c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x55, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x75, 0x69, 0x64,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x56, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xca, 0xbb, 0x18, 0x03, 0x76, 0x69, 0x64, 0x52, 0x03, 0x56, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x03, 0x43, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xca,
	0xbb, 0x18, 0x03, 0x63, 0x69, 0x64, 0x52, 0x03, 0x43, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xca,
	0xbb, 0x18, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x52, 0x08, 0x50, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x6c,
	0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x4c, 0x69, 0x6b, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0f, 0xca, 0xbb, 0x18, 0x0b, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xca, 0xbb, 0x18, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0xca, 0xbb, 0x18, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca,
	0xbb, 0x18, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xbb, 0x18,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7a, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0xca, 0xbb,
	0x18, 0x04, 0x62, 0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x42, 0x09, 0xca, 0xbb, 0x18, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x32, 0xf0, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x4e,
	0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x48,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0xca, 0xc1, 0x18, 0x0d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0xe2, 0xc1, 0x18, 0x0f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x77, 0x65,
	0x62, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x68, 0x65, 0x72, 0x74,
	0x7a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_comment_proto_goTypes = []interface{}{
	(*CommentPublishReq)(nil),   // 0: proto.CommentPublishReq
	(*CommentListReq)(nil),      // 1: proto.CommentListReq
	(*CommentDeleteReq)(nil),    // 2: proto.CommentDeleteReq
	(*CommentItems)(nil),        // 3: proto.CommentItems
	(*CommentResponse)(nil),     // 4: proto.CommentResponse
	(*common.BaseResponse)(nil), // 5: proto.BaseResponse
}
var file_comment_proto_depIdxs = []int32{
	5, // 0: proto.CommentResponse.Base:type_name -> proto.BaseResponse
	3, // 1: proto.CommentResponse.Items:type_name -> proto.CommentItems
	0, // 2: proto.Comment.Publish:input_type -> proto.CommentPublishReq
	1, // 3: proto.Comment.List:input_type -> proto.CommentListReq
	2, // 4: proto.Comment.Delete:input_type -> proto.CommentDeleteReq
	5, // 5: proto.Comment.Publish:output_type -> proto.BaseResponse
	4, // 6: proto.Comment.List:output_type -> proto.CommentResponse
	5, // 7: proto.Comment.Delete:output_type -> proto.BaseResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentPublishReq); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListReq); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDeleteReq); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentItems); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentResponse); i {
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
	file_comment_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_comment_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_comment_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}
