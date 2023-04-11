package main

import (
	"TelnetBBS/src/application"
	"TelnetBBS/src/inface"
)

func main() {
	s := application.NewServer()

	s.AddRouter("@Login", &LoginRouter{})
	s.AddRouter("@Pwd", &PwdRouter{})
	s.AddRouter("@User", &UserRouter{})

	s.Serve()

}

// LoginRouter 自定义路由
type LoginRouter struct {
	application.BaseRouter
}

// PreHandle 测试路由
func (pr *LoginRouter) PreHandle(request inface.IRequest) {

	// 这里可以用一个线程来代替
	conn := request.GetConnection()
	cmd := request.GetCommand()
	data := request.GetData()
	user := conn.GetUser()

	// 后端命令解析
	switch cmd {
	case "@Login": // 登录用户名
		user.SetLoginID(data)
	}
}

// PwdRouter 自定义路由
type PwdRouter struct {
	application.BaseRouter
}

// PreHandle 测试路由
func (pr *PwdRouter) PreHandle(request inface.IRequest) {

	// 这里可以用一个线程来代替
	conn := request.GetConnection()
	cmd := request.GetCommand()
	data := request.GetData()
	user := conn.GetUser()

	// 后端命令解析
	switch cmd {
	case "@Pwd": // 登录用户名
		user.SetPwd(data)
	}
}

// UserRouter 自定义路由
type UserRouter struct {
	application.BaseRouter
}

// PreHandle 测试路由
func (ur *UserRouter) PreHandle(request inface.IRequest) {

	// 这里可以用一个线程来代替
	conn := request.GetConnection()
	cmd := request.GetCommand()
	user := conn.GetUser()

	// 后端命令解析
	switch cmd {
	case "@User": // 登录用户名
		loginId := "LoginID: " + user.GetLoginID() + "\n\r"
		conn.GetTCPConnection().Write([]byte(loginId))
		loginPwd := "loginPwd: " + user.GetPwd() + "\n\r"
		conn.GetTCPConnection().Write([]byte(loginPwd))

	}
}

type PageIndex struct {
	inface.IPage
}
