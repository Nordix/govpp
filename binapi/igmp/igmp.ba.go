// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              23.10-rc0~68-gee48020ac
// source: /usr/share/vpp/api/igmp.api.json

// Package igmp contains generated bindings for API file igmp.api.
//
// Contents:
//   2 enums
//   2 structs
//  19 messages
//
package igmp

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	interface_types "github.com/edwarnicke/govpp/binapi/interface_types"
	ip_types "github.com/edwarnicke/govpp/binapi/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "igmp"
	APIVersion = "1.0.0"
	VersionCrc = 0x95a4ff29
)

// FilterMode defines enum 'filter_mode'.
type FilterMode uint32

const (
	EXCLUDE FilterMode = 0
	INCLUDE FilterMode = 1
)

var (
	FilterMode_name = map[uint32]string{
		0: "EXCLUDE",
		1: "INCLUDE",
	}
	FilterMode_value = map[string]uint32{
		"EXCLUDE": 0,
		"INCLUDE": 1,
	}
)

func (x FilterMode) String() string {
	s, ok := FilterMode_name[uint32(x)]
	if ok {
		return s
	}
	return "FilterMode(" + strconv.Itoa(int(x)) + ")"
}

// GroupPrefixType defines enum 'group_prefix_type'.
type GroupPrefixType uint32

const (
	ASM GroupPrefixType = 0
	SSM GroupPrefixType = 1
)

var (
	GroupPrefixType_name = map[uint32]string{
		0: "ASM",
		1: "SSM",
	}
	GroupPrefixType_value = map[string]uint32{
		"ASM": 0,
		"SSM": 1,
	}
)

func (x GroupPrefixType) String() string {
	s, ok := GroupPrefixType_name[uint32(x)]
	if ok {
		return s
	}
	return "GroupPrefixType(" + strconv.Itoa(int(x)) + ")"
}

// GroupPrefix defines type 'group_prefix'.
type GroupPrefix struct {
	Type   GroupPrefixType `binapi:"group_prefix_type,name=type" json:"type,omitempty"`
	Prefix ip_types.Prefix `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
}

// IgmpGroup defines type 'igmp_group'.
type IgmpGroup struct {
	Filter    FilterMode                     `binapi:"filter_mode,name=filter" json:"filter,omitempty"`
	NSrcs     uint8                          `binapi:"u8,name=n_srcs" json:"-"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Gaddr     ip_types.IP4Address            `binapi:"ip4_address,name=gaddr" json:"gaddr,omitempty"`
	Saddrs    []ip_types.IP4Address          `binapi:"ip4_address[n_srcs],name=saddrs" json:"saddrs,omitempty"`
}

// IgmpClearInterface defines message 'igmp_clear_interface'.
type IgmpClearInterface struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IgmpClearInterface) Reset()               { *m = IgmpClearInterface{} }
func (*IgmpClearInterface) GetMessageName() string { return "igmp_clear_interface" }
func (*IgmpClearInterface) GetCrcString() string   { return "f9e6675e" }
func (*IgmpClearInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpClearInterface) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *IgmpClearInterface) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IgmpClearInterface) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IgmpClearInterfaceReply defines message 'igmp_clear_interface_reply'.
type IgmpClearInterfaceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpClearInterfaceReply) Reset()               { *m = IgmpClearInterfaceReply{} }
func (*IgmpClearInterfaceReply) GetMessageName() string { return "igmp_clear_interface_reply" }
func (*IgmpClearInterfaceReply) GetCrcString() string   { return "e8d4e804" }
func (*IgmpClearInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpClearInterfaceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpClearInterfaceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpClearInterfaceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IgmpDetails defines message 'igmp_details'.
type IgmpDetails struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Saddr     ip_types.IP4Address            `binapi:"ip4_address,name=saddr" json:"saddr,omitempty"`
	Gaddr     ip_types.IP4Address            `binapi:"ip4_address,name=gaddr" json:"gaddr,omitempty"`
}

func (m *IgmpDetails) Reset()               { *m = IgmpDetails{} }
func (*IgmpDetails) GetMessageName() string { return "igmp_details" }
func (*IgmpDetails) GetCrcString() string   { return "38f09929" }
func (*IgmpDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 1 * 4 // m.Saddr
	size += 1 * 4 // m.Gaddr
	return size
}
func (m *IgmpDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.Saddr[:], 4)
	buf.EncodeBytes(m.Gaddr[:], 4)
	return buf.Bytes(), nil
}
func (m *IgmpDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.Saddr[:], buf.DecodeBytes(4))
	copy(m.Gaddr[:], buf.DecodeBytes(4))
	return nil
}

// IgmpDump defines message 'igmp_dump'.
type IgmpDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IgmpDump) Reset()               { *m = IgmpDump{} }
func (*IgmpDump) GetMessageName() string { return "igmp_dump" }
func (*IgmpDump) GetCrcString() string   { return "f9e6675e" }
func (*IgmpDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *IgmpDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IgmpDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IgmpEnableDisable defines message 'igmp_enable_disable'.
type IgmpEnableDisable struct {
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
	Mode      uint8                          `binapi:"u8,name=mode" json:"mode,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IgmpEnableDisable) Reset()               { *m = IgmpEnableDisable{} }
func (*IgmpEnableDisable) GetMessageName() string { return "igmp_enable_disable" }
func (*IgmpEnableDisable) GetCrcString() string   { return "b1edfb96" }
func (*IgmpEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	size += 1 // m.Mode
	size += 4 // m.SwIfIndex
	return size
}
func (m *IgmpEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint8(m.Mode)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IgmpEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.Mode = buf.DecodeUint8()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IgmpEnableDisableReply defines message 'igmp_enable_disable_reply'.
type IgmpEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpEnableDisableReply) Reset()               { *m = IgmpEnableDisableReply{} }
func (*IgmpEnableDisableReply) GetMessageName() string { return "igmp_enable_disable_reply" }
func (*IgmpEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*IgmpEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IgmpEvent defines message 'igmp_event'.
type IgmpEvent struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Filter    FilterMode                     `binapi:"filter_mode,name=filter" json:"filter,omitempty"`
	Saddr     ip_types.IP4Address            `binapi:"ip4_address,name=saddr" json:"saddr,omitempty"`
	Gaddr     ip_types.IP4Address            `binapi:"ip4_address,name=gaddr" json:"gaddr,omitempty"`
}

func (m *IgmpEvent) Reset()               { *m = IgmpEvent{} }
func (*IgmpEvent) GetMessageName() string { return "igmp_event" }
func (*IgmpEvent) GetCrcString() string   { return "85fe93ec" }
func (*IgmpEvent) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func (m *IgmpEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 4     // m.Filter
	size += 1 * 4 // m.Saddr
	size += 1 * 4 // m.Gaddr
	return size
}
func (m *IgmpEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint32(uint32(m.Filter))
	buf.EncodeBytes(m.Saddr[:], 4)
	buf.EncodeBytes(m.Gaddr[:], 4)
	return buf.Bytes(), nil
}
func (m *IgmpEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Filter = FilterMode(buf.DecodeUint32())
	copy(m.Saddr[:], buf.DecodeBytes(4))
	copy(m.Gaddr[:], buf.DecodeBytes(4))
	return nil
}

// IgmpGroupPrefixDetails defines message 'igmp_group_prefix_details'.
type IgmpGroupPrefixDetails struct {
	Gp GroupPrefix `binapi:"group_prefix,name=gp" json:"gp,omitempty"`
}

func (m *IgmpGroupPrefixDetails) Reset()               { *m = IgmpGroupPrefixDetails{} }
func (*IgmpGroupPrefixDetails) GetMessageName() string { return "igmp_group_prefix_details" }
func (*IgmpGroupPrefixDetails) GetCrcString() string   { return "259ccd81" }
func (*IgmpGroupPrefixDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpGroupPrefixDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Gp.Type
	size += 1      // m.Gp.Prefix.Address.Af
	size += 1 * 16 // m.Gp.Prefix.Address.Un
	size += 1      // m.Gp.Prefix.Len
	return size
}
func (m *IgmpGroupPrefixDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Gp.Type))
	buf.EncodeUint8(uint8(m.Gp.Prefix.Address.Af))
	buf.EncodeBytes(m.Gp.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Gp.Prefix.Len)
	return buf.Bytes(), nil
}
func (m *IgmpGroupPrefixDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Gp.Type = GroupPrefixType(buf.DecodeUint32())
	m.Gp.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Gp.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Gp.Prefix.Len = buf.DecodeUint8()
	return nil
}

// IgmpGroupPrefixDump defines message 'igmp_group_prefix_dump'.
type IgmpGroupPrefixDump struct{}

func (m *IgmpGroupPrefixDump) Reset()               { *m = IgmpGroupPrefixDump{} }
func (*IgmpGroupPrefixDump) GetMessageName() string { return "igmp_group_prefix_dump" }
func (*IgmpGroupPrefixDump) GetCrcString() string   { return "51077d14" }
func (*IgmpGroupPrefixDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpGroupPrefixDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IgmpGroupPrefixDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IgmpGroupPrefixDump) Unmarshal(b []byte) error {
	return nil
}

// IgmpGroupPrefixSet defines message 'igmp_group_prefix_set'.
type IgmpGroupPrefixSet struct {
	Gp GroupPrefix `binapi:"group_prefix,name=gp" json:"gp,omitempty"`
}

func (m *IgmpGroupPrefixSet) Reset()               { *m = IgmpGroupPrefixSet{} }
func (*IgmpGroupPrefixSet) GetMessageName() string { return "igmp_group_prefix_set" }
func (*IgmpGroupPrefixSet) GetCrcString() string   { return "5b14a5ce" }
func (*IgmpGroupPrefixSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpGroupPrefixSet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Gp.Type
	size += 1      // m.Gp.Prefix.Address.Af
	size += 1 * 16 // m.Gp.Prefix.Address.Un
	size += 1      // m.Gp.Prefix.Len
	return size
}
func (m *IgmpGroupPrefixSet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Gp.Type))
	buf.EncodeUint8(uint8(m.Gp.Prefix.Address.Af))
	buf.EncodeBytes(m.Gp.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Gp.Prefix.Len)
	return buf.Bytes(), nil
}
func (m *IgmpGroupPrefixSet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Gp.Type = GroupPrefixType(buf.DecodeUint32())
	m.Gp.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Gp.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Gp.Prefix.Len = buf.DecodeUint8()
	return nil
}

// IgmpGroupPrefixSetReply defines message 'igmp_group_prefix_set_reply'.
type IgmpGroupPrefixSetReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpGroupPrefixSetReply) Reset()               { *m = IgmpGroupPrefixSetReply{} }
func (*IgmpGroupPrefixSetReply) GetMessageName() string { return "igmp_group_prefix_set_reply" }
func (*IgmpGroupPrefixSetReply) GetCrcString() string   { return "e8d4e804" }
func (*IgmpGroupPrefixSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpGroupPrefixSetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpGroupPrefixSetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpGroupPrefixSetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IgmpListen defines message 'igmp_listen'.
type IgmpListen struct {
	Group IgmpGroup `binapi:"igmp_group,name=group" json:"group,omitempty"`
}

func (m *IgmpListen) Reset()               { *m = IgmpListen{} }
func (*IgmpListen) GetMessageName() string { return "igmp_listen" }
func (*IgmpListen) GetCrcString() string   { return "19a49f1e" }
func (*IgmpListen) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpListen) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Group.Filter
	size += 1     // m.Group.NSrcs
	size += 4     // m.Group.SwIfIndex
	size += 1 * 4 // m.Group.Gaddr
	for j2 := 0; j2 < len(m.Group.Saddrs); j2++ {
		var s2 ip_types.IP4Address
		_ = s2
		if j2 < len(m.Group.Saddrs) {
			s2 = m.Group.Saddrs[j2]
		}
		size += 1 * 4 // s2
	}
	return size
}
func (m *IgmpListen) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Group.Filter))
	buf.EncodeUint8(uint8(len(m.Group.Saddrs)))
	buf.EncodeUint32(uint32(m.Group.SwIfIndex))
	buf.EncodeBytes(m.Group.Gaddr[:], 4)
	for j1 := 0; j1 < len(m.Group.Saddrs); j1++ {
		var v1 ip_types.IP4Address // Saddrs
		if j1 < len(m.Group.Saddrs) {
			v1 = m.Group.Saddrs[j1]
		}
		buf.EncodeBytes(v1[:], 4)
	}
	return buf.Bytes(), nil
}
func (m *IgmpListen) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Group.Filter = FilterMode(buf.DecodeUint32())
	m.Group.NSrcs = buf.DecodeUint8()
	m.Group.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.Group.Gaddr[:], buf.DecodeBytes(4))
	m.Group.Saddrs = make([]ip_types.IP4Address, m.Group.NSrcs)
	for j1 := 0; j1 < len(m.Group.Saddrs); j1++ {
		copy(m.Group.Saddrs[j1][:], buf.DecodeBytes(4))
	}
	return nil
}

// IgmpListenReply defines message 'igmp_listen_reply'.
type IgmpListenReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpListenReply) Reset()               { *m = IgmpListenReply{} }
func (*IgmpListenReply) GetMessageName() string { return "igmp_listen_reply" }
func (*IgmpListenReply) GetCrcString() string   { return "e8d4e804" }
func (*IgmpListenReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpListenReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpListenReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpListenReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IgmpProxyDeviceAddDel defines message 'igmp_proxy_device_add_del'.
type IgmpProxyDeviceAddDel struct {
	Add       uint8                          `binapi:"u8,name=add" json:"add,omitempty"`
	VrfID     uint32                         `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IgmpProxyDeviceAddDel) Reset()               { *m = IgmpProxyDeviceAddDel{} }
func (*IgmpProxyDeviceAddDel) GetMessageName() string { return "igmp_proxy_device_add_del" }
func (*IgmpProxyDeviceAddDel) GetCrcString() string   { return "0b9be9ce" }
func (*IgmpProxyDeviceAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpProxyDeviceAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Add
	size += 4 // m.VrfID
	size += 4 // m.SwIfIndex
	return size
}
func (m *IgmpProxyDeviceAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.Add)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IgmpProxyDeviceAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Add = buf.DecodeUint8()
	m.VrfID = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IgmpProxyDeviceAddDelInterface defines message 'igmp_proxy_device_add_del_interface'.
type IgmpProxyDeviceAddDelInterface struct {
	Add       bool                           `binapi:"bool,name=add" json:"add,omitempty"`
	VrfID     uint32                         `binapi:"u32,name=vrf_id" json:"vrf_id,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *IgmpProxyDeviceAddDelInterface) Reset() { *m = IgmpProxyDeviceAddDelInterface{} }
func (*IgmpProxyDeviceAddDelInterface) GetMessageName() string {
	return "igmp_proxy_device_add_del_interface"
}
func (*IgmpProxyDeviceAddDelInterface) GetCrcString() string { return "1a9ec24a" }
func (*IgmpProxyDeviceAddDelInterface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IgmpProxyDeviceAddDelInterface) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Add
	size += 4 // m.VrfID
	size += 4 // m.SwIfIndex
	return size
}
func (m *IgmpProxyDeviceAddDelInterface) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Add)
	buf.EncodeUint32(m.VrfID)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IgmpProxyDeviceAddDelInterface) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Add = buf.DecodeBool()
	m.VrfID = buf.DecodeUint32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IgmpProxyDeviceAddDelInterfaceReply defines message 'igmp_proxy_device_add_del_interface_reply'.
type IgmpProxyDeviceAddDelInterfaceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpProxyDeviceAddDelInterfaceReply) Reset() { *m = IgmpProxyDeviceAddDelInterfaceReply{} }
func (*IgmpProxyDeviceAddDelInterfaceReply) GetMessageName() string {
	return "igmp_proxy_device_add_del_interface_reply"
}
func (*IgmpProxyDeviceAddDelInterfaceReply) GetCrcString() string { return "e8d4e804" }
func (*IgmpProxyDeviceAddDelInterfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpProxyDeviceAddDelInterfaceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpProxyDeviceAddDelInterfaceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpProxyDeviceAddDelInterfaceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IgmpProxyDeviceAddDelReply defines message 'igmp_proxy_device_add_del_reply'.
type IgmpProxyDeviceAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IgmpProxyDeviceAddDelReply) Reset()               { *m = IgmpProxyDeviceAddDelReply{} }
func (*IgmpProxyDeviceAddDelReply) GetMessageName() string { return "igmp_proxy_device_add_del_reply" }
func (*IgmpProxyDeviceAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*IgmpProxyDeviceAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IgmpProxyDeviceAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IgmpProxyDeviceAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IgmpProxyDeviceAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// WantIgmpEvents defines message 'want_igmp_events'.
type WantIgmpEvents struct {
	Enable uint32 `binapi:"u32,name=enable" json:"enable,omitempty"`
	PID    uint32 `binapi:"u32,name=pid" json:"pid,omitempty"`
}

func (m *WantIgmpEvents) Reset()               { *m = WantIgmpEvents{} }
func (*WantIgmpEvents) GetMessageName() string { return "want_igmp_events" }
func (*WantIgmpEvents) GetCrcString() string   { return "cfaccc1f" }
func (*WantIgmpEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantIgmpEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Enable
	size += 4 // m.PID
	return size
}
func (m *WantIgmpEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Enable)
	buf.EncodeUint32(m.PID)
	return buf.Bytes(), nil
}
func (m *WantIgmpEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeUint32()
	m.PID = buf.DecodeUint32()
	return nil
}

// WantIgmpEventsReply defines message 'want_igmp_events_reply'.
type WantIgmpEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantIgmpEventsReply) Reset()               { *m = WantIgmpEventsReply{} }
func (*WantIgmpEventsReply) GetMessageName() string { return "want_igmp_events_reply" }
func (*WantIgmpEventsReply) GetCrcString() string   { return "e8d4e804" }
func (*WantIgmpEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantIgmpEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantIgmpEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantIgmpEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_igmp_binapi_init() }
func file_igmp_binapi_init() {
	api.RegisterMessage((*IgmpClearInterface)(nil), "igmp_clear_interface_f9e6675e")
	api.RegisterMessage((*IgmpClearInterfaceReply)(nil), "igmp_clear_interface_reply_e8d4e804")
	api.RegisterMessage((*IgmpDetails)(nil), "igmp_details_38f09929")
	api.RegisterMessage((*IgmpDump)(nil), "igmp_dump_f9e6675e")
	api.RegisterMessage((*IgmpEnableDisable)(nil), "igmp_enable_disable_b1edfb96")
	api.RegisterMessage((*IgmpEnableDisableReply)(nil), "igmp_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*IgmpEvent)(nil), "igmp_event_85fe93ec")
	api.RegisterMessage((*IgmpGroupPrefixDetails)(nil), "igmp_group_prefix_details_259ccd81")
	api.RegisterMessage((*IgmpGroupPrefixDump)(nil), "igmp_group_prefix_dump_51077d14")
	api.RegisterMessage((*IgmpGroupPrefixSet)(nil), "igmp_group_prefix_set_5b14a5ce")
	api.RegisterMessage((*IgmpGroupPrefixSetReply)(nil), "igmp_group_prefix_set_reply_e8d4e804")
	api.RegisterMessage((*IgmpListen)(nil), "igmp_listen_19a49f1e")
	api.RegisterMessage((*IgmpListenReply)(nil), "igmp_listen_reply_e8d4e804")
	api.RegisterMessage((*IgmpProxyDeviceAddDel)(nil), "igmp_proxy_device_add_del_0b9be9ce")
	api.RegisterMessage((*IgmpProxyDeviceAddDelInterface)(nil), "igmp_proxy_device_add_del_interface_1a9ec24a")
	api.RegisterMessage((*IgmpProxyDeviceAddDelInterfaceReply)(nil), "igmp_proxy_device_add_del_interface_reply_e8d4e804")
	api.RegisterMessage((*IgmpProxyDeviceAddDelReply)(nil), "igmp_proxy_device_add_del_reply_e8d4e804")
	api.RegisterMessage((*WantIgmpEvents)(nil), "want_igmp_events_cfaccc1f")
	api.RegisterMessage((*WantIgmpEventsReply)(nil), "want_igmp_events_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IgmpClearInterface)(nil),
		(*IgmpClearInterfaceReply)(nil),
		(*IgmpDetails)(nil),
		(*IgmpDump)(nil),
		(*IgmpEnableDisable)(nil),
		(*IgmpEnableDisableReply)(nil),
		(*IgmpEvent)(nil),
		(*IgmpGroupPrefixDetails)(nil),
		(*IgmpGroupPrefixDump)(nil),
		(*IgmpGroupPrefixSet)(nil),
		(*IgmpGroupPrefixSetReply)(nil),
		(*IgmpListen)(nil),
		(*IgmpListenReply)(nil),
		(*IgmpProxyDeviceAddDel)(nil),
		(*IgmpProxyDeviceAddDelInterface)(nil),
		(*IgmpProxyDeviceAddDelInterfaceReply)(nil),
		(*IgmpProxyDeviceAddDelReply)(nil),
		(*WantIgmpEvents)(nil),
		(*WantIgmpEventsReply)(nil),
	}
}
