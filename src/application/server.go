package application

import (
	"TelnetBBS/src/inface"
	"fmt"
	"net"
	"os"
)

type Server struct {
	IpVersion string
	IP        string
	Port      string
	//当前Server由用户绑定的回调router,也就是Server注册的链接对应的处理业务
	Router inface.IRouter
}

func (s *Server) AddRouter(router inface.IRouter) {
	s.Router = router
	fmt.Println("Router Add Success!!")
}

func NewServer() inface.IServer {
	return &Server{
		IpVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      "8899",
		Router:    nil,
	}
}

func (s *Server) Start() {
	//创建监听器
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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		dealConn := NewConnection(conn, ConnID, s.Router)
		ConnID++

		// 开启一个线程处理任务
		go dealConn.Start()

	}

}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) Close() {

}
