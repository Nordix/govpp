// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package interfaces

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/edwarnicke/govpp/binapi/memclnt"
)

// RPCService defines RPC service interface.
type RPCService interface {
	CollectDetailedInterfaceStats(ctx context.Context, in *CollectDetailedInterfaceStats) (*CollectDetailedInterfaceStatsReply, error)
	CreateLoopback(ctx context.Context, in *CreateLoopback) (*CreateLoopbackReply, error)
	CreateLoopbackInstance(ctx context.Context, in *CreateLoopbackInstance) (*CreateLoopbackInstanceReply, error)
	CreateSubif(ctx context.Context, in *CreateSubif) (*CreateSubifReply, error)
	CreateVlanSubif(ctx context.Context, in *CreateVlanSubif) (*CreateVlanSubifReply, error)
	DeleteLoopback(ctx context.Context, in *DeleteLoopback) (*DeleteLoopbackReply, error)
	DeleteSubif(ctx context.Context, in *DeleteSubif) (*DeleteSubifReply, error)
	GetBuffersStats(ctx context.Context, in *GetBuffersStats) (*GetBuffersStatsReply, error)
	HwInterfaceSetMtu(ctx context.Context, in *HwInterfaceSetMtu) (*HwInterfaceSetMtuReply, error)
	InterfaceNameRenumber(ctx context.Context, in *InterfaceNameRenumber) (*InterfaceNameRenumberReply, error)
	PcapTraceOff(ctx context.Context, in *PcapTraceOff) (*PcapTraceOffReply, error)
	PcapTraceOn(ctx context.Context, in *PcapTraceOn) (*PcapTraceOnReply, error)
	SwInterfaceAddDelAddress(ctx context.Context, in *SwInterfaceAddDelAddress) (*SwInterfaceAddDelAddressReply, error)
	SwInterfaceAddDelMacAddress(ctx context.Context, in *SwInterfaceAddDelMacAddress) (*SwInterfaceAddDelMacAddressReply, error)
	SwInterfaceAddressReplaceBegin(ctx context.Context, in *SwInterfaceAddressReplaceBegin) (*SwInterfaceAddressReplaceBeginReply, error)
	SwInterfaceAddressReplaceEnd(ctx context.Context, in *SwInterfaceAddressReplaceEnd) (*SwInterfaceAddressReplaceEndReply, error)
	SwInterfaceClearStats(ctx context.Context, in *SwInterfaceClearStats) (*SwInterfaceClearStatsReply, error)
	SwInterfaceDump(ctx context.Context, in *SwInterfaceDump) (RPCService_SwInterfaceDumpClient, error)
	SwInterfaceGetMacAddress(ctx context.Context, in *SwInterfaceGetMacAddress) (*SwInterfaceGetMacAddressReply, error)
	SwInterfaceGetTable(ctx context.Context, in *SwInterfaceGetTable) (*SwInterfaceGetTableReply, error)
	SwInterfaceRxPlacementDump(ctx context.Context, in *SwInterfaceRxPlacementDump) (RPCService_SwInterfaceRxPlacementDumpClient, error)
	SwInterfaceSetFlags(ctx context.Context, in *SwInterfaceSetFlags) (*SwInterfaceSetFlagsReply, error)
	SwInterfaceSetInterfaceName(ctx context.Context, in *SwInterfaceSetInterfaceName) (*SwInterfaceSetInterfaceNameReply, error)
	SwInterfaceSetIPDirectedBroadcast(ctx context.Context, in *SwInterfaceSetIPDirectedBroadcast) (*SwInterfaceSetIPDirectedBroadcastReply, error)
	SwInterfaceSetMacAddress(ctx context.Context, in *SwInterfaceSetMacAddress) (*SwInterfaceSetMacAddressReply, error)
	SwInterfaceSetMtu(ctx context.Context, in *SwInterfaceSetMtu) (*SwInterfaceSetMtuReply, error)
	SwInterfaceSetPromisc(ctx context.Context, in *SwInterfaceSetPromisc) (*SwInterfaceSetPromiscReply, error)
	SwInterfaceSetRxMode(ctx context.Context, in *SwInterfaceSetRxMode) (*SwInterfaceSetRxModeReply, error)
	SwInterfaceSetRxPlacement(ctx context.Context, in *SwInterfaceSetRxPlacement) (*SwInterfaceSetRxPlacementReply, error)
	SwInterfaceSetTable(ctx context.Context, in *SwInterfaceSetTable) (*SwInterfaceSetTableReply, error)
	SwInterfaceSetTxPlacement(ctx context.Context, in *SwInterfaceSetTxPlacement) (*SwInterfaceSetTxPlacementReply, error)
	SwInterfaceSetUnnumbered(ctx context.Context, in *SwInterfaceSetUnnumbered) (*SwInterfaceSetUnnumberedReply, error)
	SwInterfaceTagAddDel(ctx context.Context, in *SwInterfaceTagAddDel) (*SwInterfaceTagAddDelReply, error)
	SwInterfaceTxPlacementGet(ctx context.Context, in *SwInterfaceTxPlacementGet) (RPCService_SwInterfaceTxPlacementGetClient, error)
	WantInterfaceEvents(ctx context.Context, in *WantInterfaceEvents) (*WantInterfaceEventsReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) CollectDetailedInterfaceStats(ctx context.Context, in *CollectDetailedInterfaceStats) (*CollectDetailedInterfaceStatsReply, error) {
	out := new(CollectDetailedInterfaceStatsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CreateLoopback(ctx context.Context, in *CreateLoopback) (*CreateLoopbackReply, error) {
	out := new(CreateLoopbackReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CreateLoopbackInstance(ctx context.Context, in *CreateLoopbackInstance) (*CreateLoopbackInstanceReply, error) {
	out := new(CreateLoopbackInstanceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CreateSubif(ctx context.Context, in *CreateSubif) (*CreateSubifReply, error) {
	out := new(CreateSubifReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) CreateVlanSubif(ctx context.Context, in *CreateVlanSubif) (*CreateVlanSubifReply, error) {
	out := new(CreateVlanSubifReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DeleteLoopback(ctx context.Context, in *DeleteLoopback) (*DeleteLoopbackReply, error) {
	out := new(DeleteLoopbackReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) DeleteSubif(ctx context.Context, in *DeleteSubif) (*DeleteSubifReply, error) {
	out := new(DeleteSubifReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GetBuffersStats(ctx context.Context, in *GetBuffersStats) (*GetBuffersStatsReply, error) {
	out := new(GetBuffersStatsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) HwInterfaceSetMtu(ctx context.Context, in *HwInterfaceSetMtu) (*HwInterfaceSetMtuReply, error) {
	out := new(HwInterfaceSetMtuReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) InterfaceNameRenumber(ctx context.Context, in *InterfaceNameRenumber) (*InterfaceNameRenumberReply, error) {
	out := new(InterfaceNameRenumberReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PcapTraceOff(ctx context.Context, in *PcapTraceOff) (*PcapTraceOffReply, error) {
	out := new(PcapTraceOffReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) PcapTraceOn(ctx context.Context, in *PcapTraceOn) (*PcapTraceOnReply, error) {
	out := new(PcapTraceOnReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceAddDelAddress(ctx context.Context, in *SwInterfaceAddDelAddress) (*SwInterfaceAddDelAddressReply, error) {
	out := new(SwInterfaceAddDelAddressReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceAddDelMacAddress(ctx context.Context, in *SwInterfaceAddDelMacAddress) (*SwInterfaceAddDelMacAddressReply, error) {
	out := new(SwInterfaceAddDelMacAddressReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceAddressReplaceBegin(ctx context.Context, in *SwInterfaceAddressReplaceBegin) (*SwInterfaceAddressReplaceBeginReply, error) {
	out := new(SwInterfaceAddressReplaceBeginReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceAddressReplaceEnd(ctx context.Context, in *SwInterfaceAddressReplaceEnd) (*SwInterfaceAddressReplaceEndReply, error) {
	out := new(SwInterfaceAddressReplaceEndReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceClearStats(ctx context.Context, in *SwInterfaceClearStats) (*SwInterfaceClearStatsReply, error) {
	out := new(SwInterfaceClearStatsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceDump(ctx context.Context, in *SwInterfaceDump) (RPCService_SwInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceDumpClient interface {
	Recv() (*SwInterfaceDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceDumpClient) Recv() (*SwInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceGetMacAddress(ctx context.Context, in *SwInterfaceGetMacAddress) (*SwInterfaceGetMacAddressReply, error) {
	out := new(SwInterfaceGetMacAddressReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceGetTable(ctx context.Context, in *SwInterfaceGetTable) (*SwInterfaceGetTableReply, error) {
	out := new(SwInterfaceGetTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceRxPlacementDump(ctx context.Context, in *SwInterfaceRxPlacementDump) (RPCService_SwInterfaceRxPlacementDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceRxPlacementDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceRxPlacementDumpClient interface {
	Recv() (*SwInterfaceRxPlacementDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceRxPlacementDumpClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceRxPlacementDumpClient) Recv() (*SwInterfaceRxPlacementDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceRxPlacementDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SwInterfaceSetFlags(ctx context.Context, in *SwInterfaceSetFlags) (*SwInterfaceSetFlagsReply, error) {
	out := new(SwInterfaceSetFlagsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetInterfaceName(ctx context.Context, in *SwInterfaceSetInterfaceName) (*SwInterfaceSetInterfaceNameReply, error) {
	out := new(SwInterfaceSetInterfaceNameReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetIPDirectedBroadcast(ctx context.Context, in *SwInterfaceSetIPDirectedBroadcast) (*SwInterfaceSetIPDirectedBroadcastReply, error) {
	out := new(SwInterfaceSetIPDirectedBroadcastReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetMacAddress(ctx context.Context, in *SwInterfaceSetMacAddress) (*SwInterfaceSetMacAddressReply, error) {
	out := new(SwInterfaceSetMacAddressReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetMtu(ctx context.Context, in *SwInterfaceSetMtu) (*SwInterfaceSetMtuReply, error) {
	out := new(SwInterfaceSetMtuReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetPromisc(ctx context.Context, in *SwInterfaceSetPromisc) (*SwInterfaceSetPromiscReply, error) {
	out := new(SwInterfaceSetPromiscReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetRxMode(ctx context.Context, in *SwInterfaceSetRxMode) (*SwInterfaceSetRxModeReply, error) {
	out := new(SwInterfaceSetRxModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetRxPlacement(ctx context.Context, in *SwInterfaceSetRxPlacement) (*SwInterfaceSetRxPlacementReply, error) {
	out := new(SwInterfaceSetRxPlacementReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetTable(ctx context.Context, in *SwInterfaceSetTable) (*SwInterfaceSetTableReply, error) {
	out := new(SwInterfaceSetTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetTxPlacement(ctx context.Context, in *SwInterfaceSetTxPlacement) (*SwInterfaceSetTxPlacementReply, error) {
	out := new(SwInterfaceSetTxPlacementReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetUnnumbered(ctx context.Context, in *SwInterfaceSetUnnumbered) (*SwInterfaceSetUnnumberedReply, error) {
	out := new(SwInterfaceSetUnnumberedReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceTagAddDel(ctx context.Context, in *SwInterfaceTagAddDel) (*SwInterfaceTagAddDelReply, error) {
	out := new(SwInterfaceTagAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceTxPlacementGet(ctx context.Context, in *SwInterfaceTxPlacementGet) (RPCService_SwInterfaceTxPlacementGetClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SwInterfaceTxPlacementGetClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SwInterfaceTxPlacementGetClient interface {
	Recv() (*SwInterfaceTxPlacementDetails, error)
	api.Stream
}

type serviceClient_SwInterfaceTxPlacementGetClient struct {
	api.Stream
}

func (c *serviceClient_SwInterfaceTxPlacementGetClient) Recv() (*SwInterfaceTxPlacementDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SwInterfaceTxPlacementDetails:
		return m, nil
	case *SwInterfaceTxPlacementGetReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) WantInterfaceEvents(ctx context.Context, in *WantInterfaceEvents) (*WantInterfaceEventsReply, error) {
	out := new(WantInterfaceEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
