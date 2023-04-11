package inface

/*
消息接口,客户端向服务器传输消息
*/

type IMessage interface {
	GetData() string
	GetCommand() string
	GetDataLen() uint32
}
