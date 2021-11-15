// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.02-rc0~165-g0edb85b37
// source: /usr/share/vpp/api/plugins/gtpu.api.json

// Package gtpu contains generated bindings for API file gtpu.api.
//
// Contents:
//  10 messages
//
package gtpu

import (
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
	APIFile    = "gtpu"
	APIVersion = "2.0.1"
	VersionCrc = 0x1462473
)

// GtpuAddDelTunnel defines message 'gtpu_add_del_tunnel'.
type GtpuAddDelTunnel struct {
	IsAdd          bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SrcAddress     ip_types.Address               `binapi:"address,name=src_address" json:"src_address,omitempty"`
	DstAddress     ip_types.Address               `binapi:"address,name=dst_address" json:"dst_address,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapNextIndex uint32                         `binapi:"u32,name=decap_next_index" json:"decap_next_index,omitempty"`
	Teid           uint32                         `binapi:"u32,name=teid" json:"teid,omitempty"`
	Tteid          uint32                         `binapi:"u32,name=tteid" json:"tteid,omitempty"`
}

func (m *GtpuAddDelTunnel) Reset()               { *m = GtpuAddDelTunnel{} }
func (*GtpuAddDelTunnel) GetMessageName() string { return "gtpu_add_del_tunnel" }
func (*GtpuAddDelTunnel) GetCrcString() string   { return "ca983a2b" }
func (*GtpuAddDelTunnel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GtpuAddDelTunnel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 1      // m.DstAddress.Af
	size += 1 * 16 // m.DstAddress.Un
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapNextIndex
	size += 4      // m.Teid
	size += 4      // m.Tteid
	return size
}
func (m *GtpuAddDelTunnel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.DstAddress.Af))
	buf.EncodeBytes(m.DstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapNextIndex)
	buf.EncodeUint32(m.Teid)
	buf.EncodeUint32(m.Tteid)
	return buf.Bytes(), nil
}
func (m *GtpuAddDelTunnel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.DstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapNextIndex = buf.DecodeUint32()
	m.Teid = buf.DecodeUint32()
	m.Tteid = buf.DecodeUint32()
	return nil
}

// GtpuAddDelTunnelReply defines message 'gtpu_add_del_tunnel_reply'.
type GtpuAddDelTunnelReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *GtpuAddDelTunnelReply) Reset()               { *m = GtpuAddDelTunnelReply{} }
func (*GtpuAddDelTunnelReply) GetMessageName() string { return "gtpu_add_del_tunnel_reply" }
func (*GtpuAddDelTunnelReply) GetCrcString() string   { return "5383d31f" }
func (*GtpuAddDelTunnelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GtpuAddDelTunnelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *GtpuAddDelTunnelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *GtpuAddDelTunnelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// GtpuOffloadRx defines message 'gtpu_offload_rx'.
type GtpuOffloadRx struct {
	HwIfIndex uint32 `binapi:"u32,name=hw_if_index" json:"hw_if_index,omitempty"`
	SwIfIndex uint32 `binapi:"u32,name=sw_if_index" json:"sw_if_index,omitempty"`
	Enable    uint8  `binapi:"u8,name=enable" json:"enable,omitempty"`
}

func (m *GtpuOffloadRx) Reset()               { *m = GtpuOffloadRx{} }
func (*GtpuOffloadRx) GetMessageName() string { return "gtpu_offload_rx" }
func (*GtpuOffloadRx) GetCrcString() string   { return "f0b08786" }
func (*GtpuOffloadRx) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GtpuOffloadRx) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.HwIfIndex
	size += 4 // m.SwIfIndex
	size += 1 // m.Enable
	return size
}
func (m *GtpuOffloadRx) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.HwIfIndex)
	buf.EncodeUint32(m.SwIfIndex)
	buf.EncodeUint8(m.Enable)
	return buf.Bytes(), nil
}
func (m *GtpuOffloadRx) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.HwIfIndex = buf.DecodeUint32()
	m.SwIfIndex = buf.DecodeUint32()
	m.Enable = buf.DecodeUint8()
	return nil
}

// GtpuOffloadRxReply defines message 'gtpu_offload_rx_reply'.
type GtpuOffloadRxReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GtpuOffloadRxReply) Reset()               { *m = GtpuOffloadRxReply{} }
func (*GtpuOffloadRxReply) GetMessageName() string { return "gtpu_offload_rx_reply" }
func (*GtpuOffloadRxReply) GetCrcString() string   { return "e8d4e804" }
func (*GtpuOffloadRxReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GtpuOffloadRxReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GtpuOffloadRxReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GtpuOffloadRxReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// GtpuTunnelDetails defines message 'gtpu_tunnel_details'.
type GtpuTunnelDetails struct {
	SwIfIndex      interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	SrcAddress     ip_types.Address               `binapi:"address,name=src_address" json:"src_address,omitempty"`
	DstAddress     ip_types.Address               `binapi:"address,name=dst_address" json:"dst_address,omitempty"`
	McastSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=mcast_sw_if_index" json:"mcast_sw_if_index,omitempty"`
	EncapVrfID     uint32                         `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	DecapNextIndex uint32                         `binapi:"u32,name=decap_next_index" json:"decap_next_index,omitempty"`
	Teid           uint32                         `binapi:"u32,name=teid" json:"teid,omitempty"`
	Tteid          uint32                         `binapi:"u32,name=tteid" json:"tteid,omitempty"`
}

func (m *GtpuTunnelDetails) Reset()               { *m = GtpuTunnelDetails{} }
func (*GtpuTunnelDetails) GetMessageName() string { return "gtpu_tunnel_details" }
func (*GtpuTunnelDetails) GetCrcString() string   { return "27f434ae" }
func (*GtpuTunnelDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GtpuTunnelDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.SwIfIndex
	size += 1      // m.SrcAddress.Af
	size += 1 * 16 // m.SrcAddress.Un
	size += 1      // m.DstAddress.Af
	size += 1 * 16 // m.DstAddress.Un
	size += 4      // m.McastSwIfIndex
	size += 4      // m.EncapVrfID
	size += 4      // m.DecapNextIndex
	size += 4      // m.Teid
	size += 4      // m.Tteid
	return size
}
func (m *GtpuTunnelDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.SrcAddress.Af))
	buf.EncodeBytes(m.SrcAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(uint8(m.DstAddress.Af))
	buf.EncodeBytes(m.DstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.McastSwIfIndex))
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.DecapNextIndex)
	buf.EncodeUint32(m.Teid)
	buf.EncodeUint32(m.Tteid)
	return buf.Bytes(), nil
}
func (m *GtpuTunnelDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.SrcAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.SrcAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.DstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.McastSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EncapVrfID = buf.DecodeUint32()
	m.DecapNextIndex = buf.DecodeUint32()
	m.Teid = buf.DecodeUint32()
	m.Tteid = buf.DecodeUint32()
	return nil
}

// GtpuTunnelDump defines message 'gtpu_tunnel_dump'.
type GtpuTunnelDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *GtpuTunnelDump) Reset()               { *m = GtpuTunnelDump{} }
func (*GtpuTunnelDump) GetMessageName() string { return "gtpu_tunnel_dump" }
func (*GtpuTunnelDump) GetCrcString() string   { return "f9e6675e" }
func (*GtpuTunnelDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GtpuTunnelDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *GtpuTunnelDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *GtpuTunnelDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// GtpuTunnelUpdateTteid defines message 'gtpu_tunnel_update_tteid'.
type GtpuTunnelUpdateTteid struct {
	DstAddress ip_types.Address `binapi:"address,name=dst_address" json:"dst_address,omitempty"`
	EncapVrfID uint32           `binapi:"u32,name=encap_vrf_id" json:"encap_vrf_id,omitempty"`
	Teid       uint32           `binapi:"u32,name=teid" json:"teid,omitempty"`
	Tteid      uint32           `binapi:"u32,name=tteid" json:"tteid,omitempty"`
}

func (m *GtpuTunnelUpdateTteid) Reset()               { *m = GtpuTunnelUpdateTteid{} }
func (*GtpuTunnelUpdateTteid) GetMessageName() string { return "gtpu_tunnel_update_tteid" }
func (*GtpuTunnelUpdateTteid) GetCrcString() string   { return "79f33816" }
func (*GtpuTunnelUpdateTteid) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GtpuTunnelUpdateTteid) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.DstAddress.Af
	size += 1 * 16 // m.DstAddress.Un
	size += 4      // m.EncapVrfID
	size += 4      // m.Teid
	size += 4      // m.Tteid
	return size
}
func (m *GtpuTunnelUpdateTteid) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.DstAddress.Af))
	buf.EncodeBytes(m.DstAddress.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.EncapVrfID)
	buf.EncodeUint32(m.Teid)
	buf.EncodeUint32(m.Tteid)
	return buf.Bytes(), nil
}
func (m *GtpuTunnelUpdateTteid) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.DstAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.DstAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.EncapVrfID = buf.DecodeUint32()
	m.Teid = buf.DecodeUint32()
	m.Tteid = buf.DecodeUint32()
	return nil
}

// GtpuTunnelUpdateTteidReply defines message 'gtpu_tunnel_update_tteid_reply'.
type GtpuTunnelUpdateTteidReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GtpuTunnelUpdateTteidReply) Reset()               { *m = GtpuTunnelUpdateTteidReply{} }
func (*GtpuTunnelUpdateTteidReply) GetMessageName() string { return "gtpu_tunnel_update_tteid_reply" }
func (*GtpuTunnelUpdateTteidReply) GetCrcString() string   { return "e8d4e804" }
func (*GtpuTunnelUpdateTteidReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GtpuTunnelUpdateTteidReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GtpuTunnelUpdateTteidReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GtpuTunnelUpdateTteidReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// SwInterfaceSetGtpuBypass defines message 'sw_interface_set_gtpu_bypass'.
type SwInterfaceSetGtpuBypass struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIPv6    bool                           `binapi:"bool,name=is_ipv6" json:"is_ipv6,omitempty"`
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
}

func (m *SwInterfaceSetGtpuBypass) Reset()               { *m = SwInterfaceSetGtpuBypass{} }
func (*SwInterfaceSetGtpuBypass) GetMessageName() string { return "sw_interface_set_gtpu_bypass" }
func (*SwInterfaceSetGtpuBypass) GetCrcString() string   { return "65247409" }
func (*SwInterfaceSetGtpuBypass) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceSetGtpuBypass) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIPv6
	size += 1 // m.Enable
	return size
}
func (m *SwInterfaceSetGtpuBypass) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsIPv6)
	buf.EncodeBool(m.Enable)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetGtpuBypass) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIPv6 = buf.DecodeBool()
	m.Enable = buf.DecodeBool()
	return nil
}

// SwInterfaceSetGtpuBypassReply defines message 'sw_interface_set_gtpu_bypass_reply'.
type SwInterfaceSetGtpuBypassReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SwInterfaceSetGtpuBypassReply) Reset() { *m = SwInterfaceSetGtpuBypassReply{} }
func (*SwInterfaceSetGtpuBypassReply) GetMessageName() string {
	return "sw_interface_set_gtpu_bypass_reply"
}
func (*SwInterfaceSetGtpuBypassReply) GetCrcString() string { return "e8d4e804" }
func (*SwInterfaceSetGtpuBypassReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceSetGtpuBypassReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SwInterfaceSetGtpuBypassReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SwInterfaceSetGtpuBypassReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_gtpu_binapi_init() }
func file_gtpu_binapi_init() {
	api.RegisterMessage((*GtpuAddDelTunnel)(nil), "gtpu_add_del_tunnel_ca983a2b")
	api.RegisterMessage((*GtpuAddDelTunnelReply)(nil), "gtpu_add_del_tunnel_reply_5383d31f")
	api.RegisterMessage((*GtpuOffloadRx)(nil), "gtpu_offload_rx_f0b08786")
	api.RegisterMessage((*GtpuOffloadRxReply)(nil), "gtpu_offload_rx_reply_e8d4e804")
	api.RegisterMessage((*GtpuTunnelDetails)(nil), "gtpu_tunnel_details_27f434ae")
	api.RegisterMessage((*GtpuTunnelDump)(nil), "gtpu_tunnel_dump_f9e6675e")
	api.RegisterMessage((*GtpuTunnelUpdateTteid)(nil), "gtpu_tunnel_update_tteid_79f33816")
	api.RegisterMessage((*GtpuTunnelUpdateTteidReply)(nil), "gtpu_tunnel_update_tteid_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceSetGtpuBypass)(nil), "sw_interface_set_gtpu_bypass_65247409")
	api.RegisterMessage((*SwInterfaceSetGtpuBypassReply)(nil), "sw_interface_set_gtpu_bypass_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GtpuAddDelTunnel)(nil),
		(*GtpuAddDelTunnelReply)(nil),
		(*GtpuOffloadRx)(nil),
		(*GtpuOffloadRxReply)(nil),
		(*GtpuTunnelDetails)(nil),
		(*GtpuTunnelDump)(nil),
		(*GtpuTunnelUpdateTteid)(nil),
		(*GtpuTunnelUpdateTteidReply)(nil),
		(*SwInterfaceSetGtpuBypass)(nil),
		(*SwInterfaceSetGtpuBypassReply)(nil),
	}
}
