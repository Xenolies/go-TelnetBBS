package inface


/*
Telent服务器接口
*/

type IServer interface {
	Start()
	Serve()
	Close()
	// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(router IRouter)
}
