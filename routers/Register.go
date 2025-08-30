package routers

import (
	"TelnetBBS/src/application"
	"TelnetBBS/src/inface"
	"fmt"
)

// UserRouter 自定义路由
type RegisterRouter struct {
	application.BaseRouter
}

func (registerR *RegisterRouter) Handle(request inface.IRequest) {
	fmt.Println("PwdRouter Handle")
	// 这里可以用一个线程来代替
	cmd := request.GetCommand()
	fmt.Println(cmd)

	// 后端命令解析
	switch cmd {
	case "@register": // 登录用户名

		//user.SetLoginID(data)
	}
}
