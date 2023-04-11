package application

import (
	"TelnetBBS/src/inface"
	"TelnetBBS/src/utils"
	"fmt"
	"net"
	"os"
)

type Server struct {
	Name      string
	IpVersion string
	IP        string
	Port      string
	//当前Server由用户绑定的回调router,也就是Server注册的链接对应的处理业务
	Router inface.IRouter

	// 该Server的链接控制器
	ConnManager inface.IConnectionManager

	// 当前 Server 的消息处理模块, 用来绑定 Command 和对应的处理业务API
	MsgHandler inface.IMsgHandler
}

func (s *Server) AddRouter(cmd string, router inface.IRouter) {
	s.MsgHandler.AddRouter(cmd, router)
	fmt.Println("Router Add Success!!")
}

func NewServer() inface.IServer {
	return &Server{
		IpVersion:   "tcp4",
		IP:          utils.Gc.IP,
		Port:        utils.Gc.Port,
		Router:      nil,
		ConnManager: NewConnManager(),
		MsgHandler:  NewMsgHandle(),
	}
}

func (s *Server) Start() {
	s.MsgHandler.StartWorkPool()
	//创建监听器
	go func() {

		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%s", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveIPAddr Error : ", err)
			return
		}
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}

		//defer listener.Close()

		fmt.Println("Listening on port ", fmt.Sprintf("telnet %s %s", s.IP, s.Port))
		var ConnID uint32
		ConnID = 0

		// 阻塞等待客户端链接
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Error accepting connection: ", err.Error())
				continue
			}

			dealConn := NewConnection(s, conn, ConnID, s.Router, s.MsgHandler)
			ConnID++

			// 建立链接前判断是否超过最大链接个数
			// 超过就关闭
			if s.ConnManager.Len() > utils.Gc.MaxConn {
				fmt.Println("[MAX]Too Many Connection!!")
				fmt.Println("Now Conn: ", s.ConnManager.Len(), " MaxConn Setting: ", utils.Gc.MaxConn)
				conn.Write([]byte("Too Many Connection!!"))
				conn.Close()
				continue
			}
			// 开启一个线程处理任务
			go dealConn.Start()

		}
	}()

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) Close() {

}

// GetConnMgr 返回当前Server中的ConnManager
func (s *Server) GetConnMgr() inface.IConnectionManager {
	return s.ConnManager
}
