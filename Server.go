package main

import (
	"TelnetBBS/src/application"
	"TelnetBBS/src/inface"
	"fmt"
)

func main() {
	s := application.NewServer()

	s.AddRouter("@Login", &PingRouter{})

	s.Serve()

}

// PingRouter 自定义路由
type PingRouter struct {
	application.BaseRouter
}

// PreHandle 测试路由
func (pr *PingRouter) PreHandle(request inface.IRequest) {

	// 这里可以用一个线程来代替
	conn := request.GetConnection()
	cmd := request.GetCommand()
	data := request.GetData()
	user := conn.GetUser()

	// 后端命令解析
	switch cmd {
	case "@Login": // 登录用户名
		user.SetLoginID(data)

	case "@Pwd": // 密码
		user.SetPwd(data)
	}

	// fmt.Println("Call Router PreHandle...")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("Before Ping....  |  "))
	// if err != nil {
	// 	fmt.Println("Router PreHandle Write Error: ", err)
	// }

}

// Handle 测试路由
func (pr *PingRouter) Handle(request inface.IRequest) {
	fmt.Println("Call Router Handle...")

	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("....Ping....Ping....Ping....  |  "))
	//if err != nil {
	//	fmt.Println("Router Handle Write Error: ", err)
	//}
}

//
//// PostHandle 测试路由
//func (pr *PingRouter) PostHandle(request inface.IRequest) {
//	fmt.Println("Call Router PostHandle...")
//	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After Ping...."))
//	if err != nil {
//		fmt.Println("Router PostHandle Write Error: ", err)
//	}
//}

type PageIndex struct {
	inface.IPage
}
