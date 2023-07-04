// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              23.10-rc0~68-gee48020ac
// source: /usr/share/vpp/api/tracedump.api.json

// Package tracedump contains generated bindings for API file tracedump.api.
//
// Contents:
//   1 enum
//  17 messages
//
package tracedump

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "tracedump"
	APIVersion = "0.2.0"
	VersionCrc = 0x56abf80a
)

// TraceFilterFlag defines enum 'trace_filter_flag'.
type TraceFilterFlag uint32

const (
	TRACE_FF_NONE               TraceFilterFlag = 0
	TRACE_FF_INCLUDE_NODE       TraceFilterFlag = 1
	TRACE_FF_EXCLUDE_NODE       TraceFilterFlag = 2
	TRACE_FF_INCLUDE_CLASSIFIER TraceFilterFlag = 3
	TRACE_FF_EXCLUDE_CLASSIFIER TraceFilterFlag = 4
)

var (
	TraceFilterFlag_name = map[uint32]string{
		0: "TRACE_FF_NONE",
		1: "TRACE_FF_INCLUDE_NODE",
		2: "TRACE_FF_EXCLUDE_NODE",
		3: "TRACE_FF_INCLUDE_CLASSIFIER",
		4: "TRACE_FF_EXCLUDE_CLASSIFIER",
	}
	TraceFilterFlag_value = map[string]uint32{
		"TRACE_FF_NONE":               0,
		"TRACE_FF_INCLUDE_NODE":       1,
		"TRACE_FF_EXCLUDE_NODE":       2,
		"TRACE_FF_INCLUDE_CLASSIFIER": 3,
		"TRACE_FF_EXCLUDE_CLASSIFIER": 4,
	}
)

func (x TraceFilterFlag) String() string {
	s, ok := TraceFilterFlag_name[uint32(x)]
	if ok {
		return s
	}
	return "TraceFilterFlag(" + strconv.Itoa(int(x)) + ")"
}

// TraceCapturePackets defines message 'trace_capture_packets'.
// InProgress: the message form may change in the future versions
type TraceCapturePackets struct {
	NodeIndex       uint32 `binapi:"u32,name=node_index" json:"node_index,omitempty"`
	MaxPackets      uint32 `binapi:"u32,name=max_packets" json:"max_packets,omitempty"`
	UseFilter       bool   `binapi:"bool,name=use_filter" json:"use_filter,omitempty"`
	Verbose         bool   `binapi:"bool,name=verbose" json:"verbose,omitempty"`
	PreCaptureClear bool   `binapi:"bool,name=pre_capture_clear" json:"pre_capture_clear,omitempty"`
}

func (m *TraceCapturePackets) Reset()               { *m = TraceCapturePackets{} }
func (*TraceCapturePackets) GetMessageName() string { return "trace_capture_packets" }
func (*TraceCapturePackets) GetCrcString() string   { return "9e791a9b" }
func (*TraceCapturePackets) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceCapturePackets) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.NodeIndex
	size += 4 // m.MaxPackets
	size += 1 // m.UseFilter
	size += 1 // m.Verbose
	size += 1 // m.PreCaptureClear
	return size
}
func (m *TraceCapturePackets) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.NodeIndex)
	buf.EncodeUint32(m.MaxPackets)
	buf.EncodeBool(m.UseFilter)
	buf.EncodeBool(m.Verbose)
	buf.EncodeBool(m.PreCaptureClear)
	return buf.Bytes(), nil
}
func (m *TraceCapturePackets) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeIndex = buf.DecodeUint32()
	m.MaxPackets = buf.DecodeUint32()
	m.UseFilter = buf.DecodeBool()
	m.Verbose = buf.DecodeBool()
	m.PreCaptureClear = buf.DecodeBool()
	return nil
}

// TraceCapturePacketsReply defines message 'trace_capture_packets_reply'.
// InProgress: the message form may change in the future versions
type TraceCapturePacketsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceCapturePacketsReply) Reset()               { *m = TraceCapturePacketsReply{} }
func (*TraceCapturePacketsReply) GetMessageName() string { return "trace_capture_packets_reply" }
func (*TraceCapturePacketsReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceCapturePacketsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceCapturePacketsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceCapturePacketsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceCapturePacketsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceClearCache defines message 'trace_clear_cache'.
// InProgress: the message form may change in the future versions
type TraceClearCache struct{}

func (m *TraceClearCache) Reset()               { *m = TraceClearCache{} }
func (*TraceClearCache) GetMessageName() string { return "trace_clear_cache" }
func (*TraceClearCache) GetCrcString() string   { return "51077d14" }
func (*TraceClearCache) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceClearCache) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceClearCache) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceClearCache) Unmarshal(b []byte) error {
	return nil
}

// TraceClearCacheReply defines message 'trace_clear_cache_reply'.
// InProgress: the message form may change in the future versions
type TraceClearCacheReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceClearCacheReply) Reset()               { *m = TraceClearCacheReply{} }
func (*TraceClearCacheReply) GetMessageName() string { return "trace_clear_cache_reply" }
func (*TraceClearCacheReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceClearCacheReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceClearCacheReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceClearCacheReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceClearCacheReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceClearCapture defines message 'trace_clear_capture'.
// InProgress: the message form may change in the future versions
type TraceClearCapture struct{}

func (m *TraceClearCapture) Reset()               { *m = TraceClearCapture{} }
func (*TraceClearCapture) GetMessageName() string { return "trace_clear_capture" }
func (*TraceClearCapture) GetCrcString() string   { return "51077d14" }
func (*TraceClearCapture) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceClearCapture) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceClearCapture) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceClearCapture) Unmarshal(b []byte) error {
	return nil
}

// TraceClearCaptureReply defines message 'trace_clear_capture_reply'.
// InProgress: the message form may change in the future versions
type TraceClearCaptureReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceClearCaptureReply) Reset()               { *m = TraceClearCaptureReply{} }
func (*TraceClearCaptureReply) GetMessageName() string { return "trace_clear_capture_reply" }
func (*TraceClearCaptureReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceClearCaptureReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceClearCaptureReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceClearCaptureReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceClearCaptureReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceDetails defines message 'trace_details'.
// InProgress: the message form may change in the future versions
type TraceDetails struct {
	ThreadID       uint32 `binapi:"u32,name=thread_id" json:"thread_id,omitempty"`
	Position       uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	MoreThisThread uint8  `binapi:"u8,name=more_this_thread" json:"more_this_thread,omitempty"`
	MoreThreads    uint8  `binapi:"u8,name=more_threads" json:"more_threads,omitempty"`
	Done           uint8  `binapi:"u8,name=done" json:"done,omitempty"`
	PacketNumber   uint32 `binapi:"u32,name=packet_number" json:"packet_number,omitempty"`
	TraceData      string `binapi:"string[],name=trace_data" json:"trace_data,omitempty"`
}

func (m *TraceDetails) Reset()               { *m = TraceDetails{} }
func (*TraceDetails) GetMessageName() string { return "trace_details" }
func (*TraceDetails) GetCrcString() string   { return "1553e9eb" }
func (*TraceDetails) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                    // m.ThreadID
	size += 4                    // m.Position
	size += 1                    // m.MoreThisThread
	size += 1                    // m.MoreThreads
	size += 1                    // m.Done
	size += 4                    // m.PacketNumber
	size += 4 + len(m.TraceData) // m.TraceData
	return size
}
func (m *TraceDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeUint8(m.MoreThisThread)
	buf.EncodeUint8(m.MoreThreads)
	buf.EncodeUint8(m.Done)
	buf.EncodeUint32(m.PacketNumber)
	buf.EncodeString(m.TraceData, 0)
	return buf.Bytes(), nil
}
func (m *TraceDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.MoreThisThread = buf.DecodeUint8()
	m.MoreThreads = buf.DecodeUint8()
	m.Done = buf.DecodeUint8()
	m.PacketNumber = buf.DecodeUint32()
	m.TraceData = buf.DecodeString(0)
	return nil
}

// TraceDump defines message 'trace_dump'.
// InProgress: the message form may change in the future versions
type TraceDump struct {
	ClearCache uint8  `binapi:"u8,name=clear_cache" json:"clear_cache,omitempty"`
	ThreadID   uint32 `binapi:"u32,name=thread_id" json:"thread_id,omitempty"`
	Position   uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	MaxRecords uint32 `binapi:"u32,name=max_records" json:"max_records,omitempty"`
}

func (m *TraceDump) Reset()               { *m = TraceDump{} }
func (*TraceDump) GetMessageName() string { return "trace_dump" }
func (*TraceDump) GetCrcString() string   { return "c7d6681f" }
func (*TraceDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.ClearCache
	size += 4 // m.ThreadID
	size += 4 // m.Position
	size += 4 // m.MaxRecords
	return size
}
func (m *TraceDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.ClearCache)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeUint32(m.MaxRecords)
	return buf.Bytes(), nil
}
func (m *TraceDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ClearCache = buf.DecodeUint8()
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.MaxRecords = buf.DecodeUint32()
	return nil
}

// TraceDumpReply defines message 'trace_dump_reply'.
// InProgress: the message form may change in the future versions
type TraceDumpReply struct {
	Retval         int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	LastThreadID   uint32 `binapi:"u32,name=last_thread_id" json:"last_thread_id,omitempty"`
	LastPosition   uint32 `binapi:"u32,name=last_position" json:"last_position,omitempty"`
	MoreThisThread uint8  `binapi:"u8,name=more_this_thread" json:"more_this_thread,omitempty"`
	MoreThreads    uint8  `binapi:"u8,name=more_threads" json:"more_threads,omitempty"`
	FlushOnly      uint8  `binapi:"u8,name=flush_only" json:"flush_only,omitempty"`
	Done           uint8  `binapi:"u8,name=done" json:"done,omitempty"`
}

func (m *TraceDumpReply) Reset()               { *m = TraceDumpReply{} }
func (*TraceDumpReply) GetMessageName() string { return "trace_dump_reply" }
func (*TraceDumpReply) GetCrcString() string   { return "e0e87f9d" }
func (*TraceDumpReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceDumpReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.LastThreadID
	size += 4 // m.LastPosition
	size += 1 // m.MoreThisThread
	size += 1 // m.MoreThreads
	size += 1 // m.FlushOnly
	size += 1 // m.Done
	return size
}
func (m *TraceDumpReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.LastThreadID)
	buf.EncodeUint32(m.LastPosition)
	buf.EncodeUint8(m.MoreThisThread)
	buf.EncodeUint8(m.MoreThreads)
	buf.EncodeUint8(m.FlushOnly)
	buf.EncodeUint8(m.Done)
	return buf.Bytes(), nil
}
func (m *TraceDumpReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.LastThreadID = buf.DecodeUint32()
	m.LastPosition = buf.DecodeUint32()
	m.MoreThisThread = buf.DecodeUint8()
	m.MoreThreads = buf.DecodeUint8()
	m.FlushOnly = buf.DecodeUint8()
	m.Done = buf.DecodeUint8()
	return nil
}

// TraceFilterFunctionDetails defines message 'trace_filter_function_details'.
// InProgress: the message form may change in the future versions
type TraceFilterFunctionDetails struct {
	Selected bool   `binapi:"bool,name=selected" json:"selected,omitempty"`
	Name     string `binapi:"string[],name=name" json:"name,omitempty"`
}

func (m *TraceFilterFunctionDetails) Reset()               { *m = TraceFilterFunctionDetails{} }
func (*TraceFilterFunctionDetails) GetMessageName() string { return "trace_filter_function_details" }
func (*TraceFilterFunctionDetails) GetCrcString() string   { return "28821359" }
func (*TraceFilterFunctionDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceFilterFunctionDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1               // m.Selected
	size += 4 + len(m.Name) // m.Name
	return size
}
func (m *TraceFilterFunctionDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Selected)
	buf.EncodeString(m.Name, 0)
	return buf.Bytes(), nil
}
func (m *TraceFilterFunctionDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Selected = buf.DecodeBool()
	m.Name = buf.DecodeString(0)
	return nil
}

// TraceFilterFunctionDump defines message 'trace_filter_function_dump'.
// InProgress: the message form may change in the future versions
type TraceFilterFunctionDump struct{}

func (m *TraceFilterFunctionDump) Reset()               { *m = TraceFilterFunctionDump{} }
func (*TraceFilterFunctionDump) GetMessageName() string { return "trace_filter_function_dump" }
func (*TraceFilterFunctionDump) GetCrcString() string   { return "51077d14" }
func (*TraceFilterFunctionDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceFilterFunctionDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *TraceFilterFunctionDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *TraceFilterFunctionDump) Unmarshal(b []byte) error {
	return nil
}

// TraceSetFilterFunction defines message 'trace_set_filter_function'.
// InProgress: the message form may change in the future versions
type TraceSetFilterFunction struct {
	FilterFunctionName string `binapi:"string[],name=filter_function_name" json:"filter_function_name,omitempty"`
}

func (m *TraceSetFilterFunction) Reset()               { *m = TraceSetFilterFunction{} }
func (*TraceSetFilterFunction) GetMessageName() string { return "trace_set_filter_function" }
func (*TraceSetFilterFunction) GetCrcString() string   { return "616abb92" }
func (*TraceSetFilterFunction) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceSetFilterFunction) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 + len(m.FilterFunctionName) // m.FilterFunctionName
	return size
}
func (m *TraceSetFilterFunction) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.FilterFunctionName, 0)
	return buf.Bytes(), nil
}
func (m *TraceSetFilterFunction) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FilterFunctionName = buf.DecodeString(0)
	return nil
}

// TraceSetFilterFunctionReply defines message 'trace_set_filter_function_reply'.
// InProgress: the message form may change in the future versions
type TraceSetFilterFunctionReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceSetFilterFunctionReply) Reset()               { *m = TraceSetFilterFunctionReply{} }
func (*TraceSetFilterFunctionReply) GetMessageName() string { return "trace_set_filter_function_reply" }
func (*TraceSetFilterFunctionReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceSetFilterFunctionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceSetFilterFunctionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceSetFilterFunctionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceSetFilterFunctionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceSetFilters defines message 'trace_set_filters'.
// InProgress: the message form may change in the future versions
type TraceSetFilters struct {
	Flag                 TraceFilterFlag `binapi:"trace_filter_flag,name=flag" json:"flag,omitempty"`
	Count                uint32          `binapi:"u32,name=count" json:"count,omitempty"`
	NodeIndex            uint32          `binapi:"u32,name=node_index,default=4294967295" json:"node_index,omitempty"`
	ClassifierTableIndex uint32          `binapi:"u32,name=classifier_table_index,default=4294967295" json:"classifier_table_index,omitempty"`
}

func (m *TraceSetFilters) Reset()               { *m = TraceSetFilters{} }
func (*TraceSetFilters) GetMessageName() string { return "trace_set_filters" }
func (*TraceSetFilters) GetCrcString() string   { return "f522b44a" }
func (*TraceSetFilters) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceSetFilters) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Flag
	size += 4 // m.Count
	size += 4 // m.NodeIndex
	size += 4 // m.ClassifierTableIndex
	return size
}
func (m *TraceSetFilters) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Flag))
	buf.EncodeUint32(m.Count)
	buf.EncodeUint32(m.NodeIndex)
	buf.EncodeUint32(m.ClassifierTableIndex)
	return buf.Bytes(), nil
}
func (m *TraceSetFilters) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Flag = TraceFilterFlag(buf.DecodeUint32())
	m.Count = buf.DecodeUint32()
	m.NodeIndex = buf.DecodeUint32()
	m.ClassifierTableIndex = buf.DecodeUint32()
	return nil
}

// TraceSetFiltersReply defines message 'trace_set_filters_reply'.
// InProgress: the message form may change in the future versions
type TraceSetFiltersReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *TraceSetFiltersReply) Reset()               { *m = TraceSetFiltersReply{} }
func (*TraceSetFiltersReply) GetMessageName() string { return "trace_set_filters_reply" }
func (*TraceSetFiltersReply) GetCrcString() string   { return "e8d4e804" }
func (*TraceSetFiltersReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceSetFiltersReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *TraceSetFiltersReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *TraceSetFiltersReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// TraceV2Details defines message 'trace_v2_details'.
// InProgress: the message form may change in the future versions
type TraceV2Details struct {
	ThreadID  uint32 `binapi:"u32,name=thread_id" json:"thread_id,omitempty"`
	Position  uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	More      bool   `binapi:"bool,name=more" json:"more,omitempty"`
	TraceData string `binapi:"string[],name=trace_data" json:"trace_data,omitempty"`
}

func (m *TraceV2Details) Reset()               { *m = TraceV2Details{} }
func (*TraceV2Details) GetMessageName() string { return "trace_v2_details" }
func (*TraceV2Details) GetCrcString() string   { return "91f87d52" }
func (*TraceV2Details) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *TraceV2Details) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                    // m.ThreadID
	size += 4                    // m.Position
	size += 1                    // m.More
	size += 4 + len(m.TraceData) // m.TraceData
	return size
}
func (m *TraceV2Details) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeBool(m.More)
	buf.EncodeString(m.TraceData, 0)
	return buf.Bytes(), nil
}
func (m *TraceV2Details) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.More = buf.DecodeBool()
	m.TraceData = buf.DecodeString(0)
	return nil
}

// TraceV2Dump defines message 'trace_v2_dump'.
// InProgress: the message form may change in the future versions
type TraceV2Dump struct {
	ThreadID   uint32 `binapi:"u32,name=thread_id,default=4294967295" json:"thread_id,omitempty"`
	Position   uint32 `binapi:"u32,name=position" json:"position,omitempty"`
	Max        uint32 `binapi:"u32,name=max,default=50" json:"max,omitempty"`
	ClearCache bool   `binapi:"bool,name=clear_cache" json:"clear_cache,omitempty"`
}

func (m *TraceV2Dump) Reset()               { *m = TraceV2Dump{} }
func (*TraceV2Dump) GetMessageName() string { return "trace_v2_dump" }
func (*TraceV2Dump) GetCrcString() string   { return "83f88d8e" }
func (*TraceV2Dump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *TraceV2Dump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.ThreadID
	size += 4 // m.Position
	size += 4 // m.Max
	size += 1 // m.ClearCache
	return size
}
func (m *TraceV2Dump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.ThreadID)
	buf.EncodeUint32(m.Position)
	buf.EncodeUint32(m.Max)
	buf.EncodeBool(m.ClearCache)
	return buf.Bytes(), nil
}
func (m *TraceV2Dump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ThreadID = buf.DecodeUint32()
	m.Position = buf.DecodeUint32()
	m.Max = buf.DecodeUint32()
	m.ClearCache = buf.DecodeBool()
	return nil
}

func init() { file_tracedump_binapi_init() }
func file_tracedump_binapi_init() {
	api.RegisterMessage((*TraceCapturePackets)(nil), "trace_capture_packets_9e791a9b")
	api.RegisterMessage((*TraceCapturePacketsReply)(nil), "trace_capture_packets_reply_e8d4e804")
	api.RegisterMessage((*TraceClearCache)(nil), "trace_clear_cache_51077d14")
	api.RegisterMessage((*TraceClearCacheReply)(nil), "trace_clear_cache_reply_e8d4e804")
	api.RegisterMessage((*TraceClearCapture)(nil), "trace_clear_capture_51077d14")
	api.RegisterMessage((*TraceClearCaptureReply)(nil), "trace_clear_capture_reply_e8d4e804")
	api.RegisterMessage((*TraceDetails)(nil), "trace_details_1553e9eb")
	api.RegisterMessage((*TraceDump)(nil), "trace_dump_c7d6681f")
	api.RegisterMessage((*TraceDumpReply)(nil), "trace_dump_reply_e0e87f9d")
	api.RegisterMessage((*TraceFilterFunctionDetails)(nil), "trace_filter_function_details_28821359")
	api.RegisterMessage((*TraceFilterFunctionDump)(nil), "trace_filter_function_dump_51077d14")
	api.RegisterMessage((*TraceSetFilterFunction)(nil), "trace_set_filter_function_616abb92")
	api.RegisterMessage((*TraceSetFilterFunctionReply)(nil), "trace_set_filter_function_reply_e8d4e804")
	api.RegisterMessage((*TraceSetFilters)(nil), "trace_set_filters_f522b44a")
	api.RegisterMessage((*TraceSetFiltersReply)(nil), "trace_set_filters_reply_e8d4e804")
	api.RegisterMessage((*TraceV2Details)(nil), "trace_v2_details_91f87d52")
	api.RegisterMessage((*TraceV2Dump)(nil), "trace_v2_dump_83f88d8e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*TraceCapturePackets)(nil),
		(*TraceCapturePacketsReply)(nil),
		(*TraceClearCache)(nil),
		(*TraceClearCacheReply)(nil),
		(*TraceClearCapture)(nil),
		(*TraceClearCaptureReply)(nil),
		(*TraceDetails)(nil),
		(*TraceDump)(nil),
		(*TraceDumpReply)(nil),
		(*TraceFilterFunctionDetails)(nil),
		(*TraceFilterFunctionDump)(nil),
		(*TraceSetFilterFunction)(nil),
		(*TraceSetFilterFunctionReply)(nil),
		(*TraceSetFilters)(nil),
		(*TraceSetFiltersReply)(nil),
		(*TraceV2Details)(nil),
		(*TraceV2Dump)(nil),
	}
}
