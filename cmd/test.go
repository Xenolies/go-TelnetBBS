package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func Serve() {
	//创建监听器：
	listener, err := net.Listen("tcp", ":23")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Listening on port 23...")
	fmt.Println(listener.Addr().String())

	//创建监听器：
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}

}

// 处理连接请求：
func handleConnection(conn net.Conn) {
	// 确保函数结束时关闭连接
	defer fmt.Println(conn.RemoteAddr(), "is QUIT")
	defer conn.Close()

	// 在连接建立后发送欢迎消息
	conn.Write([]byte("Welcome to my Telnet server!\r\n"))

	// 创建一个读取器以读取客户端输入
	reader := bufio.NewReader(conn)

	for {
		// 从客户端读取输入
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		// 处理输入，并将响应发送回客户端
		response := handleInput(input)
		conn.Write([]byte(response))

		// 如果客户端退出，则断开连接
		if strings.TrimSpace(string(input)) == "QUIT" {
			fmt.Println("Closing connection")
			return
		}
	}
}

// 处理连接请求：
func handleInput(input string) string {
	input = strings.TrimSpace(input)
	if input == "TIME" {
		return time.Now().String() + "\r\n"
	} else if input == "HELLO" {
		return "Hello, world!\r\n"
	} else {
		return "Invalid command\r\n"
	}
}

func main() {
	Serve()

	select {}
}
