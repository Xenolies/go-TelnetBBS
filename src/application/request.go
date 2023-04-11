package application

import (
	"TelnetBBS/src/inface"
)

type Request struct {
	Conn inface.IConnection
	Msg  inface.IMessage
// 	Data string
}

func (r *Request) GetConnection() inface.IConnection {
	return r.Conn
}

func (r *Request) GetData() string {
	//return r.Msg.GetData()
	return r.Msg.GetData()
}

func (r *Request) GetCommand()string{
	return r.Msg.GetCommand()
}

//func (r *Request) GetMsg() inface.IMessage {
//	return r.Msg
//}
