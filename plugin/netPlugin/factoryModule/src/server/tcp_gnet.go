package server

import (
	"context"
	"fmt"
	"github.com/ArkNX/ark-go/base"

	"github.com/panjf2000/gnet"
)

type GNetServer struct {
	*gnet.EventServer

	hl             base.HeadLength
	sessionManager *ServerService
}

func NewGNetServer(ss *ServerService) *GNetServer {
	return &GNetServer{sessionManager: ss}
}

func (gs *GNetServer) Start(
	hl uint32,
	ip string,
	port uint16,
	threadNum uint8,
	maxClient uint32,
	isIpv6 bool) error {
	gs.hl = base.HeadLength(hl)
	return gnet.Serve(gs, fmt.Sprintf("tcp://%s:%d", ip, port), gnet.WithNumEventLoop(int(threadNum)))
}

func (gs *GNetServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	s := base.NewSession(gs.hl, c)
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
