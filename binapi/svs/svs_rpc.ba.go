// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package svs

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	memclnt "github.com/networkservicemesh/govpp/binapi/memclnt"
)

// RPCService defines RPC service svs.
type RPCService interface {
	SvsDump(ctx context.Context, in *SvsDump) (RPCService_SvsDumpClient, error)
	SvsEnableDisable(ctx context.Context, in *SvsEnableDisable) (*SvsEnableDisableReply, error)
	SvsPluginGetVersion(ctx context.Context, in *SvsPluginGetVersion) (*SvsPluginGetVersionReply, error)
	SvsRouteAddDel(ctx context.Context, in *SvsRouteAddDel) (*SvsRouteAddDelReply, error)
	SvsTableAddDel(ctx context.Context, in *SvsTableAddDel) (*SvsTableAddDelReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SvsDump(ctx context.Context, in *SvsDump) (RPCService_SvsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SvsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SvsDumpClient interface {
	Recv() (*SvsDetails, error)
	api.Stream
}

type serviceClient_SvsDumpClient struct {
	api.Stream
}

func (c *serviceClient_SvsDumpClient) Recv() (*SvsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SvsDetails:
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

func (c *serviceClient) SvsEnableDisable(ctx context.Context, in *SvsEnableDisable) (*SvsEnableDisableReply, error) {
	out := new(SvsEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SvsPluginGetVersion(ctx context.Context, in *SvsPluginGetVersion) (*SvsPluginGetVersionReply, error) {
	out := new(SvsPluginGetVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SvsRouteAddDel(ctx context.Context, in *SvsRouteAddDel) (*SvsRouteAddDelReply, error) {
	out := new(SvsRouteAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SvsTableAddDel(ctx context.Context, in *SvsTableAddDel) (*SvsTableAddDelReply, error) {
	out := new(SvsTableAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
