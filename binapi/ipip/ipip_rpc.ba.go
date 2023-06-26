// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ipip

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/networkservicemesh/govpp/binapi/memclnt"
)

// RPCService defines RPC service ipip.
type RPCService interface {
	Ipip6rdAddTunnel(ctx context.Context, in *Ipip6rdAddTunnel) (*Ipip6rdAddTunnelReply, error)
	Ipip6rdDelTunnel(ctx context.Context, in *Ipip6rdDelTunnel) (*Ipip6rdDelTunnelReply, error)
	IpipAddTunnel(ctx context.Context, in *IpipAddTunnel) (*IpipAddTunnelReply, error)
	IpipDelTunnel(ctx context.Context, in *IpipDelTunnel) (*IpipDelTunnelReply, error)
	IpipTunnelDump(ctx context.Context, in *IpipTunnelDump) (RPCService_IpipTunnelDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) Ipip6rdAddTunnel(ctx context.Context, in *Ipip6rdAddTunnel) (*Ipip6rdAddTunnelReply, error) {
	out := new(Ipip6rdAddTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) Ipip6rdDelTunnel(ctx context.Context, in *Ipip6rdDelTunnel) (*Ipip6rdDelTunnelReply, error) {
	out := new(Ipip6rdDelTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpipAddTunnel(ctx context.Context, in *IpipAddTunnel) (*IpipAddTunnelReply, error) {
	out := new(IpipAddTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpipDelTunnel(ctx context.Context, in *IpipDelTunnel) (*IpipDelTunnelReply, error) {
	out := new(IpipDelTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) IpipTunnelDump(ctx context.Context, in *IpipTunnelDump) (RPCService_IpipTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IpipTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IpipTunnelDumpClient interface {
	Recv() (*IpipTunnelDetails, error)
	api.Stream
}

type serviceClient_IpipTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_IpipTunnelDumpClient) Recv() (*IpipTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IpipTunnelDetails:
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
