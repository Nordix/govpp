// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              22.02-release
// source: /usr/share/vpp/api/core/ethernet_types.api.json

// Package ethernet_types contains generated bindings for API file ethernet_types.api.
//
// Contents:
//   1 alias
//
package ethernet_types

import (
	"net"

	api "git.fd.io/govpp.git/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// MacAddress defines alias 'mac_address'.
type MacAddress [6]uint8

func ParseMacAddress(s string) (MacAddress, error) {
	var macaddr MacAddress
	mac, err := net.ParseMAC(s)
	if err != nil {
		return macaddr, err
	}
	copy(macaddr[:], mac[:])
	return macaddr, nil
}
func (x MacAddress) ToMAC() net.HardwareAddr {
	return net.HardwareAddr(x[:])
}
func (x MacAddress) String() string {
	return x.ToMAC().String()
}
func (x *MacAddress) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *MacAddress) UnmarshalText(text []byte) error {
	mac, err := ParseMacAddress(string(text))
	if err != nil {
		return err
	}
	*x = mac
	return nil
}
