// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package bier

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vlib "github.com/edwarnicke/govpp/binapi/vlib"
)

// RPCService defines RPC service bier.
type RPCService interface {
	BierDispEntryAddDel(ctx context.Context, in *BierDispEntryAddDel) (*BierDispEntryAddDelReply, error)
	BierDispEntryDump(ctx context.Context, in *BierDispEntryDump) (RPCService_BierDispEntryDumpClient, error)
	BierDispTableAddDel(ctx context.Context, in *BierDispTableAddDel) (*BierDispTableAddDelReply, error)
	BierDispTableDump(ctx context.Context, in *BierDispTableDump) (RPCService_BierDispTableDumpClient, error)
	BierImpAdd(ctx context.Context, in *BierImpAdd) (*BierImpAddReply, error)
	BierImpDel(ctx context.Context, in *BierImpDel) (*BierImpDelReply, error)
	BierImpDump(ctx context.Context, in *BierImpDump) (RPCService_BierImpDumpClient, error)
	BierRouteAddDel(ctx context.Context, in *BierRouteAddDel) (*BierRouteAddDelReply, error)
	BierRouteDump(ctx context.Context, in *BierRouteDump) (RPCService_BierRouteDumpClient, error)
	BierTableAddDel(ctx context.Context, in *BierTableAddDel) (*BierTableAddDelReply, error)
	BierTableDump(ctx context.Context, in *BierTableDump) (RPCService_BierTableDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) BierDispEntryAddDel(ctx context.Context, in *BierDispEntryAddDel) (*BierDispEntryAddDelReply, error) {
	out := new(BierDispEntryAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierDispEntryDump(ctx context.Context, in *BierDispEntryDump) (RPCService_BierDispEntryDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BierDispEntryDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BierDispEntryDumpClient interface {
	Recv() (*BierDispEntryDetails, error)
	api.Stream
}

type serviceClient_BierDispEntryDumpClient struct {
	api.Stream
}

func (c *serviceClient_BierDispEntryDumpClient) Recv() (*BierDispEntryDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BierDispEntryDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BierDispTableAddDel(ctx context.Context, in *BierDispTableAddDel) (*BierDispTableAddDelReply, error) {
	out := new(BierDispTableAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierDispTableDump(ctx context.Context, in *BierDispTableDump) (RPCService_BierDispTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BierDispTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BierDispTableDumpClient interface {
	Recv() (*BierDispTableDetails, error)
	api.Stream
}

type serviceClient_BierDispTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_BierDispTableDumpClient) Recv() (*BierDispTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BierDispTableDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BierImpAdd(ctx context.Context, in *BierImpAdd) (*BierImpAddReply, error) {
	out := new(BierImpAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierImpDel(ctx context.Context, in *BierImpDel) (*BierImpDelReply, error) {
	out := new(BierImpDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierImpDump(ctx context.Context, in *BierImpDump) (RPCService_BierImpDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BierImpDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BierImpDumpClient interface {
	Recv() (*BierImpDetails, error)
	api.Stream
}

type serviceClient_BierImpDumpClient struct {
	api.Stream
}

func (c *serviceClient_BierImpDumpClient) Recv() (*BierImpDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BierImpDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BierRouteAddDel(ctx context.Context, in *BierRouteAddDel) (*BierRouteAddDelReply, error) {
	out := new(BierRouteAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierRouteDump(ctx context.Context, in *BierRouteDump) (RPCService_BierRouteDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BierRouteDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BierRouteDumpClient interface {
	Recv() (*BierRouteDetails, error)
	api.Stream
}

type serviceClient_BierRouteDumpClient struct {
	api.Stream
}

func (c *serviceClient_BierRouteDumpClient) Recv() (*BierRouteDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BierRouteDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) BierTableAddDel(ctx context.Context, in *BierTableAddDel) (*BierTableAddDelReply, error) {
	out := new(BierTableAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BierTableDump(ctx context.Context, in *BierTableDump) (RPCService_BierTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BierTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vlib.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BierTableDumpClient interface {
	Recv() (*BierTableDetails, error)
	api.Stream
}

type serviceClient_BierTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_BierTableDumpClient) Recv() (*BierTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BierTableDetails:
		return m, nil
	case *vlib.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}
