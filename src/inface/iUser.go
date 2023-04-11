package inface

import "net"


/*
用户接口,每一个链接创建一个User
*/

type IUser interface {
	GetLevel() uint
	SetLevel(LevenNum uint)

	GetAddr() string
	SetAddr(conn *net.TCPConn)

	GetLoginID() string
	SetLoginID(loginID string)

	GetName() string
	SetName(Name string)

	GetPwd() string
	SetPwd(pwd string)
}