//
//Copyright 2018-2019 Mailgun Technologies Inc
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: gubernator.proto

package gubernator

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Algorithm int32

const (
	// Token bucket algorithm https://en.wikipedia.org/wiki/Token_bucket
	Algorithm_TOKEN_BUCKET Algorithm = 0
	// Leaky bucket algorithm https://en.wikipedia.org/wiki/Leaky_bucket
	Algorithm_LEAKY_BUCKET Algorithm = 1
)

// Enum value maps for Algorithm.
var (
	Algorithm_name = map[int32]string{
		0: "TOKEN_BUCKET",
		1: "LEAKY_BUCKET",
	}
	Algorithm_value = map[string]int32{
		"TOKEN_BUCKET": 0,
		"LEAKY_BUCKET": 1,
	}
)

func (x Algorithm) Enum() *Algorithm {
	p := new(Algorithm)
	*p = x
	return p
}

func (x Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_gubernator_proto_enumTypes[0].Descriptor()
}

func (Algorithm) Type() protoreflect.EnumType {
	return &file_gubernator_proto_enumTypes[0]
}

func (x Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Algorithm.Descriptor instead.
func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{0}
}

// A set of int32 flags used to control the behavior of a rate limit in gubernator
type Behavior int32

const (
	// BATCHING is the default behavior. This enables batching requests which protects the
	// service from thundering herd. IE: When a service experiences spikes of unexpected high
	// volume requests.
	//
	// Using this option introduces a small amount of latency depending on
	// the `batchWait` setting. Defaults to around 500 Microseconds of additional
	// latency in low throughput situations. For high volume loads, batching can reduce
	// the overall load on the system substantially.
	Behavior_BATCHING Behavior = 0 // <-- this is here because proto requires it, but has no effect if used
	// Disables batching. Use this for super low latency rate limit requests when
	// thundering herd is not a concern but latency of requests is of paramount importance.
	Behavior_NO_BATCHING Behavior = 1
	// Enables Global caching of the rate limit. Use this if the rate limit applies globally to
	// all ingress requests. (IE: Throttle hundreds of thousands of requests to an entire
	// datacenter or cluster of http servers)
	//
	// Using this option gubernator will continue to use a single peer as the rate limit coordinator
	// to increment and manage the state of the rate limit, however the result of the rate limit is
	// distributed to each peer and cached locally. A rate limit request received from any peer in the
	// cluster will first check the local cache for a rate limit answer, if it exists the peer will
	// immediately return the answer to the client and asynchronously forward the aggregate hits to
	// the peer coordinator. Because of GLOBALS async nature we lose some accuracy in rate limit
	// reporting, which may result in allowing some requests beyond the chosen rate limit. However we
	// gain massive performance as every request coming into the system does not have to wait for a
	// single peer to decide if the rate limit has been reached.
	Behavior_GLOBAL Behavior = 2
	// Changes the behavior of the `Duration` field. When `Behavior` is set to `DURATION_IS_GREGORIAN`
	// the `Duration` of the rate limit is reset whenever the end of selected GREGORIAN calendar
	// interval is reached.
	//
	// Given the following `Duration` values
	//   0 = Minutes
	//   1 = Hours
	//   2 = Days
	//   3 = Weeks
	//   4 = Months
	//   5 = Years
	//
	// Examples when using `Behavior = DURATION_IS_GREGORIAN`
	//
	// If  `Duration = 2` (Days) then the rate limit will expire at the end of the current day the
	// rate limit was created.
	//
	// If `Duration = 0` (Minutes) then the rate limit will expire at the end of the current minute
	// the rate limit was created.
	//
	// If `Duration = 4` (Months) then the rate limit will expire at the end of the current month
	// the rate limit was created.
	Behavior_DURATION_IS_GREGORIAN Behavior = 4
	// If this flag is set causes the rate limit to reset any accrued hits stored in the cache, and will
	// ignore any `Hit` values provided in the current request. The effect this has is dependent on
	// algorithm chosen. For instance, if used with `TOKEN_BUCKET` it will immediately expire the
	// cache value. For `LEAKY_BUCKET` it sets the `Remaining` to `Limit`.
	Behavior_RESET_REMAINING Behavior = 8
	// Enables rate limits to be pushed to other regions. Currently this is only implemented when using
	// 'member-list' peer discovery. Also requires GUBER_DATA_CENTER to be set to different values on at
	// least 2 instances of Gubernator.
	Behavior_MULTI_REGION Behavior = 16
)

// Enum value maps for Behavior.
var (
	Behavior_name = map[int32]string{
		0:  "BATCHING",
		1:  "NO_BATCHING",
		2:  "GLOBAL",
		4:  "DURATION_IS_GREGORIAN",
		8:  "RESET_REMAINING",
		16: "MULTI_REGION",
	}
	Behavior_value = map[string]int32{
		"BATCHING":              0,
		"NO_BATCHING":           1,
		"GLOBAL":                2,
		"DURATION_IS_GREGORIAN": 4,
		"RESET_REMAINING":       8,
		"MULTI_REGION":          16,
	}
)

func (x Behavior) Enum() *Behavior {
	p := new(Behavior)
	*p = x
	return p
}

func (x Behavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Behavior) Descriptor() protoreflect.EnumDescriptor {
	return file_gubernator_proto_enumTypes[1].Descriptor()
}

func (Behavior) Type() protoreflect.EnumType {
	return &file_gubernator_proto_enumTypes[1]
}

func (x Behavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Behavior.Descriptor instead.
func (Behavior) EnumDescriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{1}
}

type Status int32

const (
	Status_UNDER_LIMIT Status = 0
	Status_OVER_LIMIT  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "UNDER_LIMIT",
		1: "OVER_LIMIT",
	}
	Status_value = map[string]int32{
		"UNDER_LIMIT": 0,
		"OVER_LIMIT":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_gubernator_proto_enumTypes[2].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_gubernator_proto_enumTypes[2]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{2}
}

// Must specify at least one Request
type GetRateLimitsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*RateLimitReq `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *GetRateLimitsReq) Reset() {
	*x = GetRateLimitsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateLimitsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateLimitsReq) ProtoMessage() {}

func (x *GetRateLimitsReq) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateLimitsReq.ProtoReflect.Descriptor instead.
func (*GetRateLimitsReq) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{0}
}

func (x *GetRateLimitsReq) GetRequests() []*RateLimitReq {
	if x != nil {
		return x.Requests
	}
	return nil
}

// RateLimits returned are in the same order as the Requests
type GetRateLimitsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*RateLimitResp `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *GetRateLimitsResp) Reset() {
	*x = GetRateLimitsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRateLimitsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRateLimitsResp) ProtoMessage() {}

func (x *GetRateLimitsResp) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRateLimitsResp.ProtoReflect.Descriptor instead.
func (*GetRateLimitsResp) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{1}
}

func (x *GetRateLimitsResp) GetResponses() []*RateLimitResp {
	if x != nil {
		return x.Responses
	}
	return nil
}

type RateLimitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the rate limit IE: 'requests_per_second', 'gets_per_minute`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Uniquely identifies this rate limit IE: 'ip:10.2.10.7' or 'account:123445'
	UniqueKey string `protobuf:"bytes,2,opt,name=unique_key,json=uniqueKey,proto3" json:"unique_key,omitempty"`
	// Rate limit requests optionally specify the number of hits a request adds to the matched limit. If Hit
	// is zero, the request returns the current limit, but does not increment the hit count.
	Hits int64 `protobuf:"varint,3,opt,name=hits,proto3" json:"hits,omitempty"`
	// The number of requests that can occur for the duration of the rate limit
	Limit int64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// The duration of the rate limit in milliseconds
	// Second = 1000 Milliseconds
	// Minute = 60000 Milliseconds
	// Hour = 3600000 Milliseconds
	Duration int64 `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// The algorithm used to calculate the rate limit. The algorithm may change on
	// subsequent requests, when this occurs any previous rate limit hit counts are reset.
	Algorithm Algorithm `protobuf:"varint,6,opt,name=algorithm,proto3,enum=pb.gubernator.Algorithm" json:"algorithm,omitempty"`
	// Behavior is a set of int32 flags that control the behavior of the rate limit in gubernator
	Behavior Behavior `protobuf:"varint,7,opt,name=behavior,proto3,enum=pb.gubernator.Behavior" json:"behavior,omitempty"`
}

func (x *RateLimitReq) Reset() {
	*x = RateLimitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitReq) ProtoMessage() {}

func (x *RateLimitReq) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitReq.ProtoReflect.Descriptor instead.
func (*RateLimitReq) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{2}
}

func (x *RateLimitReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitReq) GetUniqueKey() string {
	if x != nil {
		return x.UniqueKey
	}
	return ""
}

func (x *RateLimitReq) GetHits() int64 {
	if x != nil {
		return x.Hits
	}
	return 0
}

func (x *RateLimitReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RateLimitReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *RateLimitReq) GetAlgorithm() Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return Algorithm_TOKEN_BUCKET
}

func (x *RateLimitReq) GetBehavior() Behavior {
	if x != nil {
		return x.Behavior
	}
	return Behavior_BATCHING
}

type RateLimitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the rate limit.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=pb.gubernator.Status" json:"status,omitempty"`
	// The currently configured request limit (Identical to RateLimitRequest.rate_limit_config.limit).
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// This is the number of requests remaining before the limit is hit.
	Remaining int64 `protobuf:"varint,3,opt,name=remaining,proto3" json:"remaining,omitempty"`
	// This is the time when the rate limit span will be reset, provided as a unix timestamp in milliseconds.
	ResetTime int64 `protobuf:"varint,4,opt,name=reset_time,json=resetTime,proto3" json:"reset_time,omitempty"`
	// Contains the error; If set all other values should be ignored
	Error string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	// This is additional metadata that a client might find useful. (IE: Additional headers, corrdinator ownership, etc..)
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RateLimitResp) Reset() {
	*x = RateLimitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResp) ProtoMessage() {}

func (x *RateLimitResp) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResp.ProtoReflect.Descriptor instead.
func (*RateLimitResp) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{3}
}

func (x *RateLimitResp) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_UNDER_LIMIT
}

func (x *RateLimitResp) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *RateLimitResp) GetRemaining() int64 {
	if x != nil {
		return x.Remaining
	}
	return 0
}

func (x *RateLimitResp) GetResetTime() int64 {
	if x != nil {
		return x.ResetTime
	}
	return 0
}

func (x *RateLimitResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *RateLimitResp) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type HealthCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckReq) Reset() {
	*x = HealthCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckReq) ProtoMessage() {}

func (x *HealthCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckReq.ProtoReflect.Descriptor instead.
func (*HealthCheckReq) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{4}
}

type HealthCheckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid entries are 'healthy' or 'unhealthy'
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// If 'unhealthy', message indicates the problem
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The number of peers we know about
	PeerCount int32 `protobuf:"varint,3,opt,name=peer_count,json=peerCount,proto3" json:"peer_count,omitempty"`
}

func (x *HealthCheckResp) Reset() {
	*x = HealthCheckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gubernator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResp) ProtoMessage() {}

func (x *HealthCheckResp) ProtoReflect() protoreflect.Message {
	mi := &file_gubernator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResp.ProtoReflect.Descriptor instead.
func (*HealthCheckResp) Descriptor() ([]byte, []int) {
	return file_gubernator_proto_rawDescGZIP(), []int{5}
}

func (x *HealthCheckResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthCheckResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HealthCheckResp) GetPeerCount() int32 {
	if x != nil {
		return x.PeerCount
	}
	return 0
}

var File_gubernator_proto protoreflect.FileDescriptor

var file_gubernator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x4f, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0xf4, 0x01,
	0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x68, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e,
	0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12,
	0x33, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x22, 0xac, 0x02, 0x0a, 0x0d, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x46,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x10, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x22, 0x62, 0x0a, 0x0f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x2f, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f,
	0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x45, 0x41, 0x4b,
	0x59, 0x5f, 0x42, 0x55, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x01, 0x2a, 0x77, 0x0a, 0x08, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x08, 0x42, 0x41, 0x54, 0x43, 0x48, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x53,
	0x5f, 0x47, 0x52, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x41, 0x4e, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f,
	0x52, 0x45, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x08, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x4f,
	0x4e, 0x10, 0x10, 0x2a, 0x29, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0e,
	0x0a, 0x0a, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x32, 0xdd,
	0x01, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x70, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x65, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x67, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x76, 0x31, 0x2f, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x22,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x67, 0x75, 0x6e, 0x2f, 0x67, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x80,
	0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gubernator_proto_rawDescOnce sync.Once
	file_gubernator_proto_rawDescData = file_gubernator_proto_rawDesc
)

func file_gubernator_proto_rawDescGZIP() []byte {
	file_gubernator_proto_rawDescOnce.Do(func() {
		file_gubernator_proto_rawDescData = protoimpl.X.CompressGZIP(file_gubernator_proto_rawDescData)
	})
	return file_gubernator_proto_rawDescData
}

var file_gubernator_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_gubernator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gubernator_proto_goTypes = []interface{}{
	(Algorithm)(0),            // 0: pb.gubernator.Algorithm
	(Behavior)(0),             // 1: pb.gubernator.Behavior
	(Status)(0),               // 2: pb.gubernator.Status
	(*GetRateLimitsReq)(nil),  // 3: pb.gubernator.GetRateLimitsReq
	(*GetRateLimitsResp)(nil), // 4: pb.gubernator.GetRateLimitsResp
	(*RateLimitReq)(nil),      // 5: pb.gubernator.RateLimitReq
	(*RateLimitResp)(nil),     // 6: pb.gubernator.RateLimitResp
	(*HealthCheckReq)(nil),    // 7: pb.gubernator.HealthCheckReq
	(*HealthCheckResp)(nil),   // 8: pb.gubernator.HealthCheckResp
	nil,                       // 9: pb.gubernator.RateLimitResp.MetadataEntry
}
var file_gubernator_proto_depIdxs = []int32{
	5, // 0: pb.gubernator.GetRateLimitsReq.requests:type_name -> pb.gubernator.RateLimitReq
	6, // 1: pb.gubernator.GetRateLimitsResp.responses:type_name -> pb.gubernator.RateLimitResp
	0, // 2: pb.gubernator.RateLimitReq.algorithm:type_name -> pb.gubernator.Algorithm
	1, // 3: pb.gubernator.RateLimitReq.behavior:type_name -> pb.gubernator.Behavior
	2, // 4: pb.gubernator.RateLimitResp.status:type_name -> pb.gubernator.Status
	9, // 5: pb.gubernator.RateLimitResp.metadata:type_name -> pb.gubernator.RateLimitResp.MetadataEntry
	3, // 6: pb.gubernator.V1.GetRateLimits:input_type -> pb.gubernator.GetRateLimitsReq
	7, // 7: pb.gubernator.V1.HealthCheck:input_type -> pb.gubernator.HealthCheckReq
	4, // 8: pb.gubernator.V1.GetRateLimits:output_type -> pb.gubernator.GetRateLimitsResp
	8, // 9: pb.gubernator.V1.HealthCheck:output_type -> pb.gubernator.HealthCheckResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gubernator_proto_init() }
func file_gubernator_proto_init() {
	if File_gubernator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gubernator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateLimitsReq); i {
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
		file_gubernator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRateLimitsResp); i {
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
		file_gubernator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitReq); i {
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
		file_gubernator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResp); i {
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
		file_gubernator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckReq); i {
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
		file_gubernator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResp); i {
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
			RawDescriptor: file_gubernator_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gubernator_proto_goTypes,
		DependencyIndexes: file_gubernator_proto_depIdxs,
		EnumInfos:         file_gubernator_proto_enumTypes,
		MessageInfos:      file_gubernator_proto_msgTypes,
	}.Build()
	File_gubernator_proto = out.File
	file_gubernator_proto_rawDesc = nil
	file_gubernator_proto_goTypes = nil
	file_gubernator_proto_depIdxs = nil
}
