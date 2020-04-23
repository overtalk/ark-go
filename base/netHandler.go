package base

type NetMsgHandler func(msg *NetMsg, sessionID int64)
