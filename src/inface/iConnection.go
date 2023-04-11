package inface

import "net"

/*
链接接口
*/

type IConnection interface {
	Writer()
	Reader()

	Start()
	Stop()

	RemoteAddr() net.Addr

	// GetTCPConnection 获取当前链接绑定的 Socket Conn
	GetTCPConnection() *net.TCPConn

	GetUser() IUser

	// 命令处理线程
	// CommandHandle()
}
