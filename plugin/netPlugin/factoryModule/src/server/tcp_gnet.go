package server

import (
	"context"
	"fmt"

	"github.com/panjf2000/gnet"
)

type GNetServer struct {
	*gnet.EventServer
	sessionManager *SessionManager
}

func (gs *GNetServer) Start(
	ip string,
	port uint16,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool) error {
	return gnet.Serve(gs, fmt.Sprintf("tcp://%s:%d", ip, port), gnet.WithNumEventLoop(int(threadNum)))
}

func (gs *GNetServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	s := NewSession(c)
	gs.sessionManager.AddSession(s)
	return
}

func (gs *GNetServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	id := c.Context().(context.Context).Value("id").(int64)
	gs.sessionManager.RemoveSession(id)
	action = gnet.Close
	return
}

func (gs *GNetServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println(string(frame))
	id := c.Context().(context.Context).Value("id").(int64)
	gs.sessionManager.AddBuffer(id, frame)
	return
}
