package inface

/*
请求接口
*/

type IRequest interface {
	// GetConnection 得到当前链接
	GetConnection() IConnection
	// GetData 得到请求的消息
	GetData() string

	GetCommand() string

	GetMsg() IMessage
}
