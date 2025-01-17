package application

import (
	"TelnetBBS/src/inface"
	"TelnetBBS/src/utils"
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Connection struct {
	// 当前连接所在Server
	TcpServer inface.IServer

	Conn *net.TCPConn

	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32

	isClosed bool

	ExitChan chan bool

	MsgChan chan string

	//该连接的处理方法router
	Router inface.IRouter

	// 用户属性对象
	User inface.IUser

	// 消息处理 Command 和对应的处理业务的关系
	MsgHandler inface.IMsgHandler
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func NewConnection(server inface.IServer, conn *net.TCPConn, connID uint32, router inface.IRouter, msgHandler inface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan string),
		User:       NewUser(conn),
		Router:     router,
		MsgHandler: msgHandler,
	}
	// 将Conn加入到ConnManager属性中
	c.TcpServer.GetConnMgr().Add(c)

	return c
}

func (c *Connection) GetUser() inface.IUser {
	return c.User
}

func (c *Connection) Writer() {
	fmt.Println("Writer Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn writer exit!")

	// 阻塞等待 Channel 消息
	for {
		select {
		case data := <-c.MsgChan:
			// 如果有数据将数据写入客户端
			if _, err := c.Conn.Write([]byte(data)); err != nil {
				fmt.Println("Write Data Error: ", err)
				return
			}

		case <-c.ExitChan:
			// ExitChan 告知退出 Writer也要退出
			return

		}
	}
}

func (c *Connection) Reader() {

	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	c.Conn.Write([]byte("Welcome to my Telnet server!\r\n"))

	reader := bufio.NewReader(c.Conn)
	for {
		// 从客户端读取输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		if strings.TrimSpace(input) == "QUIT" {
			fmt.Println("Closing connection")
			c.Stop()
			return
		}

		// 解析传输过来的消息
		cmd, cmdData, _ := utils.GetCommand(input)
		msg := NewMessage(cmd, cmdData)
		req := Request{
			Conn: c,
			Msg:  msg,
			// Data: cmdData,
		}
		fmt.Println(req.Conn.RemoteAddr().String())
		fmt.Println("Command:", msg.GetCommand())
		fmt.Println("CommandData:", msg.GetData())

		// // 从路由Routers 中找到注册绑定Conn的对应Handle
		//go func(request inface.IRequest) {
		//	//执行注册的路由方法
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
		// 判断工作池是否存在,存在则将消息交给工作池处理
		if utils.Gc.WorkerPoolSize > 0 {
			// 开启工作处将消息交给工作池处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由 Routers 中找到注册绑定Conn的对应Handle
			// 根据绑定好的MsgID, 传入消息处理模块 找到对应的处理业务
			go c.MsgHandler.DoMsgHandler(&req)

		}

	}

}

func (c *Connection) Start() {
	fmt.Println("Connection START...")

	// 启动当前连接读数据的业务
	go c.Reader()

	// 启动当前连接写数据的业务
	// go c.Writer()

}

func (c *Connection) Stop() {
	fmt.Printf("Connection STOP.... \n")

	//如果当前连接已经关闭
	if c.isClosed {
		return
	}

	c.isClosed = true

	// 关闭链接
	c.Conn.Close()

	// 告知Writer关闭
	c.ExitChan <- true

	// 将Channel关闭
	close(c.ExitChan)
	close(c.MsgChan)
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
