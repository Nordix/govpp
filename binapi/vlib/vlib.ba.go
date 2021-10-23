// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.02-rc0~100-g054ad8509
// source: /usr/share/vpp/api/core/vlib.api.json

// Package vlib contains generated bindings for API file vlib.api.
//
// Contents:
//   1 struct
//  20 messages
//
package vlib

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vlib"
	APIVersion = "1.0.0"
	VersionCrc = 0xbdb28485
)

// ThreadData defines type 'thread_data'.
type ThreadData struct {
	ID        uint32 `binapi:"u32,name=id" json:"id,omitempty"`
	Name      string `binapi:"string[64],name=name" json:"name,omitempty"`
	Type      string `binapi:"string[64],name=type" json:"type,omitempty"`
	PID       uint32 `binapi:"u32,name=pid" json:"pid,omitempty"`
	CPUID     uint32 `binapi:"u32,name=cpu_id" json:"cpu_id,omitempty"`
	Core      uint32 `binapi:"u32,name=core" json:"core,omitempty"`
	CPUSocket uint32 `binapi:"u32,name=cpu_socket" json:"cpu_socket,omitempty"`
}

// AddNodeNext defines message 'add_node_next'.
type AddNodeNext struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
	NextName string `binapi:"string[64],name=next_name" json:"next_name,omitempty"`
}

func (m *AddNodeNext) Reset()               { *m = AddNodeNext{} }
func (*AddNodeNext) GetMessageName() string { return "add_node_next" }
func (*AddNodeNext) GetCrcString() string   { return "2457116d" }
func (*AddNodeNext) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AddNodeNext) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	size += 64 // m.NextName
	return size
}
func (m *AddNodeNext) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	buf.EncodeString(m.NextName, 64)
	return buf.Bytes(), nil
}
func (m *AddNodeNext) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	m.NextName = buf.DecodeString(64)
	return nil
}

// AddNodeNextReply defines message 'add_node_next_reply'.
type AddNodeNextReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NextIndex uint32 `binapi:"u32,name=next_index" json:"next_index,omitempty"`
}

func (m *AddNodeNextReply) Reset()               { *m = AddNodeNextReply{} }
func (*AddNodeNextReply) GetMessageName() string { return "add_node_next_reply" }
func (*AddNodeNextReply) GetCrcString() string   { return "2ed75f32" }
func (*AddNodeNextReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AddNodeNextReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NextIndex
	return size
}
func (m *AddNodeNextReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NextIndex)
	return buf.Bytes(), nil
}
func (m *AddNodeNextReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NextIndex = buf.DecodeUint32()
	return nil
}

// Cli defines message 'cli'.
type Cli struct {
	CmdInShmem uint64 `binapi:"u64,name=cmd_in_shmem" json:"cmd_in_shmem,omitempty"`
}

func (m *Cli) Reset()               { *m = Cli{} }
func (*Cli) GetMessageName() string { return "cli" }
func (*Cli) GetCrcString() string   { return "23bfbfff" }
func (*Cli) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Cli) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.CmdInShmem
	return size
}
func (m *Cli) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.CmdInShmem)
	return buf.Bytes(), nil
}
func (m *Cli) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CmdInShmem = buf.DecodeUint64()
	return nil
}

// CliInband defines message 'cli_inband'.
type CliInband struct {
	Cmd string `binapi:"string[],name=cmd" json:"cmd,omitempty"`
}

func (m *CliInband) Reset()               { *m = CliInband{} }
func (*CliInband) GetMessageName() string { return "cli_inband" }
func (*CliInband) GetCrcString() string   { return "f8377302" }
func (*CliInband) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CliInband) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 + len(m.Cmd) // m.Cmd
	return size
}
func (m *CliInband) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Cmd, 0)
	return buf.Bytes(), nil
}
func (m *CliInband) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cmd = buf.DecodeString(0)
	return nil
}

// CliInbandReply defines message 'cli_inband_reply'.
type CliInbandReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Reply  string `binapi:"string[],name=reply" json:"reply,omitempty"`
}

func (m *CliInbandReply) Reset()               { *m = CliInbandReply{} }
func (*CliInbandReply) GetMessageName() string { return "cli_inband_reply" }
func (*CliInbandReply) GetCrcString() string   { return "05879051" }
func (*CliInbandReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CliInbandReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                // m.Retval
	size += 4 + len(m.Reply) // m.Reply
	return size
}
func (m *CliInbandReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.Reply, 0)
	return buf.Bytes(), nil
}
func (m *CliInbandReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Reply = buf.DecodeString(0)
	return nil
}

// CliReply defines message 'cli_reply'.
type CliReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ReplyInShmem uint64 `binapi:"u64,name=reply_in_shmem" json:"reply_in_shmem,omitempty"`
}

func (m *CliReply) Reset()               { *m = CliReply{} }
func (*CliReply) GetMessageName() string { return "cli_reply" }
func (*CliReply) GetCrcString() string   { return "06d68297" }
func (*CliReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CliReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.ReplyInShmem
	return size
}
func (m *CliReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint64(m.ReplyInShmem)
	return buf.Bytes(), nil
}
func (m *CliReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ReplyInShmem = buf.DecodeUint64()
	return nil
}

// ControlPing defines message 'control_ping'.
type ControlPing struct{}

func (m *ControlPing) Reset()               { *m = ControlPing{} }
func (*ControlPing) GetMessageName() string { return "control_ping" }
func (*ControlPing) GetCrcString() string   { return "51077d14" }
func (*ControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ControlPing) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ControlPing) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ControlPing) Unmarshal(b []byte) error {
	return nil
}

// ControlPingReply defines message 'control_ping_reply'.
type ControlPingReply struct {
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ClientIndex uint32 `binapi:"u32,name=client_index" json:"client_index,omitempty"`
	VpePID      uint32 `binapi:"u32,name=vpe_pid" json:"vpe_pid,omitempty"`
}

func (m *ControlPingReply) Reset()               { *m = ControlPingReply{} }
func (*ControlPingReply) GetMessageName() string { return "control_ping_reply" }
func (*ControlPingReply) GetCrcString() string   { return "f6b0b8ca" }
func (*ControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ControlPingReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ClientIndex
	size += 4 // m.VpePID
	return size
}
func (m *ControlPingReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ClientIndex)
	buf.EncodeUint32(m.VpePID)
	return buf.Bytes(), nil
}
func (m *ControlPingReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ClientIndex = buf.DecodeUint32()
	m.VpePID = buf.DecodeUint32()
	return nil
}

// GetF64EndianValue defines message 'get_f64_endian_value'.
type GetF64EndianValue struct {
	F64One float64 `binapi:"f64,name=f64_one,default=1" json:"f64_one,omitempty"`
}

func (m *GetF64EndianValue) Reset()               { *m = GetF64EndianValue{} }
func (*GetF64EndianValue) GetMessageName() string { return "get_f64_endian_value" }
func (*GetF64EndianValue) GetCrcString() string   { return "809fcd44" }
func (*GetF64EndianValue) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetF64EndianValue) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.F64One
	return size
}
func (m *GetF64EndianValue) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(m.F64One)
	return buf.Bytes(), nil
}
func (m *GetF64EndianValue) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.F64One = buf.DecodeFloat64()
	return nil
}

// GetF64EndianValueReply defines message 'get_f64_endian_value_reply'.
type GetF64EndianValueReply struct {
	Retval       uint32  `binapi:"u32,name=retval" json:"retval,omitempty"`
	F64OneResult float64 `binapi:"f64,name=f64_one_result" json:"f64_one_result,omitempty"`
}

func (m *GetF64EndianValueReply) Reset()               { *m = GetF64EndianValueReply{} }
func (*GetF64EndianValueReply) GetMessageName() string { return "get_f64_endian_value_reply" }
func (*GetF64EndianValueReply) GetCrcString() string   { return "7e02e404" }
func (*GetF64EndianValueReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetF64EndianValueReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.F64OneResult
	return size
}
func (m *GetF64EndianValueReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Retval)
	buf.EncodeFloat64(m.F64OneResult)
	return buf.Bytes(), nil
}
func (m *GetF64EndianValueReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeUint32()
	m.F64OneResult = buf.DecodeFloat64()
	return nil
}

// GetF64IncrementByOne defines message 'get_f64_increment_by_one'.
type GetF64IncrementByOne struct {
	F64Value float64 `binapi:"f64,name=f64_value,default=1" json:"f64_value,omitempty"`
}

func (m *GetF64IncrementByOne) Reset()               { *m = GetF64IncrementByOne{} }
func (*GetF64IncrementByOne) GetMessageName() string { return "get_f64_increment_by_one" }
func (*GetF64IncrementByOne) GetCrcString() string   { return "b64f027e" }
func (*GetF64IncrementByOne) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetF64IncrementByOne) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.F64Value
	return size
}
func (m *GetF64IncrementByOne) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(m.F64Value)
	return buf.Bytes(), nil
}
func (m *GetF64IncrementByOne) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.F64Value = buf.DecodeFloat64()
	return nil
}

// GetF64IncrementByOneReply defines message 'get_f64_increment_by_one_reply'.
type GetF64IncrementByOneReply struct {
	Retval   uint32  `binapi:"u32,name=retval" json:"retval,omitempty"`
	F64Value float64 `binapi:"f64,name=f64_value" json:"f64_value,omitempty"`
}

func (m *GetF64IncrementByOneReply) Reset()               { *m = GetF64IncrementByOneReply{} }
func (*GetF64IncrementByOneReply) GetMessageName() string { return "get_f64_increment_by_one_reply" }
func (*GetF64IncrementByOneReply) GetCrcString() string   { return "d25dbaa3" }
func (*GetF64IncrementByOneReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetF64IncrementByOneReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.F64Value
	return size
}
func (m *GetF64IncrementByOneReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Retval)
	buf.EncodeFloat64(m.F64Value)
	return buf.Bytes(), nil
}
func (m *GetF64IncrementByOneReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeUint32()
	m.F64Value = buf.DecodeFloat64()
	return nil
}

// GetNextIndex defines message 'get_next_index'.
type GetNextIndex struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
	NextName string `binapi:"string[64],name=next_name" json:"next_name,omitempty"`
}

func (m *GetNextIndex) Reset()               { *m = GetNextIndex{} }
func (*GetNextIndex) GetMessageName() string { return "get_next_index" }
func (*GetNextIndex) GetCrcString() string   { return "2457116d" }
func (*GetNextIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNextIndex) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	size += 64 // m.NextName
	return size
}
func (m *GetNextIndex) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	buf.EncodeString(m.NextName, 64)
	return buf.Bytes(), nil
}
func (m *GetNextIndex) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	m.NextName = buf.DecodeString(64)
	return nil
}

// GetNextIndexReply defines message 'get_next_index_reply'.
type GetNextIndexReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NextIndex uint32 `binapi:"u32,name=next_index" json:"next_index,omitempty"`
}

func (m *GetNextIndexReply) Reset()               { *m = GetNextIndexReply{} }
func (*GetNextIndexReply) GetMessageName() string { return "get_next_index_reply" }
func (*GetNextIndexReply) GetCrcString() string   { return "2ed75f32" }
func (*GetNextIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNextIndexReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NextIndex
	return size
}
func (m *GetNextIndexReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NextIndex)
	return buf.Bytes(), nil
}
func (m *GetNextIndexReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NextIndex = buf.DecodeUint32()
	return nil
}

// GetNodeGraph defines message 'get_node_graph'.
type GetNodeGraph struct{}

func (m *GetNodeGraph) Reset()               { *m = GetNodeGraph{} }
func (*GetNodeGraph) GetMessageName() string { return "get_node_graph" }
func (*GetNodeGraph) GetCrcString() string   { return "51077d14" }
func (*GetNodeGraph) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNodeGraph) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *GetNodeGraph) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *GetNodeGraph) Unmarshal(b []byte) error {
	return nil
}

// GetNodeGraphReply defines message 'get_node_graph_reply'.
type GetNodeGraphReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ReplyInShmem uint64 `binapi:"u64,name=reply_in_shmem" json:"reply_in_shmem,omitempty"`
}

func (m *GetNodeGraphReply) Reset()               { *m = GetNodeGraphReply{} }
func (*GetNodeGraphReply) GetMessageName() string { return "get_node_graph_reply" }
func (*GetNodeGraphReply) GetCrcString() string   { return "06d68297" }
func (*GetNodeGraphReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNodeGraphReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.ReplyInShmem
	return size
}
func (m *GetNodeGraphReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint64(m.ReplyInShmem)
	return buf.Bytes(), nil
}
func (m *GetNodeGraphReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ReplyInShmem = buf.DecodeUint64()
	return nil
}

// GetNodeIndex defines message 'get_node_index'.
type GetNodeIndex struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
}

func (m *GetNodeIndex) Reset()               { *m = GetNodeIndex{} }
func (*GetNodeIndex) GetMessageName() string { return "get_node_index" }
func (*GetNodeIndex) GetCrcString() string   { return "f1984c64" }
func (*GetNodeIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNodeIndex) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	return size
}
func (m *GetNodeIndex) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	return buf.Bytes(), nil
}
func (m *GetNodeIndex) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	return nil
}

// GetNodeIndexReply defines message 'get_node_index_reply'.
type GetNodeIndexReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NodeIndex uint32 `binapi:"u32,name=node_index" json:"node_index,omitempty"`
}

func (m *GetNodeIndexReply) Reset()               { *m = GetNodeIndexReply{} }
func (*GetNodeIndexReply) GetMessageName() string { return "get_node_index_reply" }
func (*GetNodeIndexReply) GetCrcString() string   { return "a8600b89" }
func (*GetNodeIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNodeIndexReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NodeIndex
	return size
}
func (m *GetNodeIndexReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NodeIndex)
	return buf.Bytes(), nil
}
func (m *GetNodeIndexReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NodeIndex = buf.DecodeUint32()
	return nil
}

// ShowThreads defines message 'show_threads'.
type ShowThreads struct{}

func (m *ShowThreads) Reset()               { *m = ShowThreads{} }
func (*ShowThreads) GetMessageName() string { return "show_threads" }
func (*ShowThreads) GetCrcString() string   { return "51077d14" }
func (*ShowThreads) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ShowThreads) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ShowThreads) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ShowThreads) Unmarshal(b []byte) error {
	return nil
}

// ShowThreadsReply defines message 'show_threads_reply'.
type ShowThreadsReply struct {
	Retval     int32        `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count      uint32       `binapi:"u32,name=count" json:"-"`
	ThreadData []ThreadData `binapi:"thread_data[count],name=thread_data" json:"thread_data,omitempty"`
}

func (m *ShowThreadsReply) Reset()               { *m = ShowThreadsReply{} }
func (*ShowThreadsReply) GetMessageName() string { return "show_threads_reply" }
func (*ShowThreadsReply) GetCrcString() string   { return "efd78e83" }
func (*ShowThreadsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ShowThreadsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Count
	for j1 := 0; j1 < len(m.ThreadData); j1++ {
		var s1 ThreadData
		_ = s1
		if j1 < len(m.ThreadData) {
			s1 = m.ThreadData[j1]
		}
		size += 4  // s1.ID
		size += 64 // s1.Name
		size += 64 // s1.Type
		size += 4  // s1.PID
		size += 4  // s1.CPUID
		size += 4  // s1.Core
		size += 4  // s1.CPUSocket
	}
	return size
}
func (m *ShowThreadsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.ThreadData)))
	for j0 := 0; j0 < len(m.ThreadData); j0++ {
		var v0 ThreadData // ThreadData
		if j0 < len(m.ThreadData) {
			v0 = m.ThreadData[j0]
		}
		buf.EncodeUint32(v0.ID)
		buf.EncodeString(v0.Name, 64)
		buf.EncodeString(v0.Type, 64)
		buf.EncodeUint32(v0.PID)
		buf.EncodeUint32(v0.CPUID)
		buf.EncodeUint32(v0.Core)
		buf.EncodeUint32(v0.CPUSocket)
	}
	return buf.Bytes(), nil
}
func (m *ShowThreadsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.ThreadData = make([]ThreadData, m.Count)
	for j0 := 0; j0 < len(m.ThreadData); j0++ {
		m.ThreadData[j0].ID = buf.DecodeUint32()
		m.ThreadData[j0].Name = buf.DecodeString(64)
		m.ThreadData[j0].Type = buf.DecodeString(64)
		m.ThreadData[j0].PID = buf.DecodeUint32()
		m.ThreadData[j0].CPUID = buf.DecodeUint32()
		m.ThreadData[j0].Core = buf.DecodeUint32()
		m.ThreadData[j0].CPUSocket = buf.DecodeUint32()
	}
	return nil
}

func init() { file_vlib_binapi_init() }
func file_vlib_binapi_init() {
	api.RegisterMessage((*AddNodeNext)(nil), "add_node_next_2457116d")
	api.RegisterMessage((*AddNodeNextReply)(nil), "add_node_next_reply_2ed75f32")
	api.RegisterMessage((*Cli)(nil), "cli_23bfbfff")
	api.RegisterMessage((*CliInband)(nil), "cli_inband_f8377302")
	api.RegisterMessage((*CliInbandReply)(nil), "cli_inband_reply_05879051")
	api.RegisterMessage((*CliReply)(nil), "cli_reply_06d68297")
	api.RegisterMessage((*ControlPing)(nil), "control_ping_51077d14")
	api.RegisterMessage((*ControlPingReply)(nil), "control_ping_reply_f6b0b8ca")
	api.RegisterMessage((*GetF64EndianValue)(nil), "get_f64_endian_value_809fcd44")
	api.RegisterMessage((*GetF64EndianValueReply)(nil), "get_f64_endian_value_reply_7e02e404")
	api.RegisterMessage((*GetF64IncrementByOne)(nil), "get_f64_increment_by_one_b64f027e")
	api.RegisterMessage((*GetF64IncrementByOneReply)(nil), "get_f64_increment_by_one_reply_d25dbaa3")
	api.RegisterMessage((*GetNextIndex)(nil), "get_next_index_2457116d")
	api.RegisterMessage((*GetNextIndexReply)(nil), "get_next_index_reply_2ed75f32")
	api.RegisterMessage((*GetNodeGraph)(nil), "get_node_graph_51077d14")
	api.RegisterMessage((*GetNodeGraphReply)(nil), "get_node_graph_reply_06d68297")
	api.RegisterMessage((*GetNodeIndex)(nil), "get_node_index_f1984c64")
	api.RegisterMessage((*GetNodeIndexReply)(nil), "get_node_index_reply_a8600b89")
	api.RegisterMessage((*ShowThreads)(nil), "show_threads_51077d14")
	api.RegisterMessage((*ShowThreadsReply)(nil), "show_threads_reply_efd78e83")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AddNodeNext)(nil),
		(*AddNodeNextReply)(nil),
		(*Cli)(nil),
		(*CliInband)(nil),
		(*CliInbandReply)(nil),
		(*CliReply)(nil),
		(*ControlPing)(nil),
		(*ControlPingReply)(nil),
		(*GetF64EndianValue)(nil),
		(*GetF64EndianValueReply)(nil),
		(*GetF64IncrementByOne)(nil),
		(*GetF64IncrementByOneReply)(nil),
		(*GetNextIndex)(nil),
		(*GetNextIndexReply)(nil),
		(*GetNodeGraph)(nil),
		(*GetNodeGraphReply)(nil),
		(*GetNodeIndex)(nil),
		(*GetNodeIndexReply)(nil),
		(*ShowThreads)(nil),
		(*ShowThreadsReply)(nil),
	}
}
