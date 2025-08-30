package routers

import (
	"TelnetBBS/src/application"
	"TelnetBBS/src/inface"
	"fmt"
)

// PwdRouter 自定义路由
type PwdRouter struct {
	application.BaseRouter
}

func (pwdR *PwdRouter) Handle(request inface.IRequest) {
	fmt.Println("PwdRouter Handle")
	// 这里可以用一个线程来代替
	cmd := request.GetCommand()
	fmt.Println(cmd)

	// 后端命令解析
	switch cmd {
	case "@pwd": // 登录用户名
		loginId := "LoginID: " + request.GetData() + "\n\r"
		fmt.Println("func (pr *LoginRouter) Handle(request inface.IRequest @Login: " + loginId)
		//user.SetLoginID(data)
	}
}
