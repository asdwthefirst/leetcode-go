/*
*

	@author: jiangxi
	@since: 2023/7/16
	@desc: //TODO

*
*/

// 服务器端代码
package main

import (
	"fmt"
	"net"
)

func RunServer() {
	// 启动服务器并监听端口
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost:8080")

	for {
		// 接受客户端连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handleRequest(conn) // 启动一个goroutine处理客户端请求
	}
}

// 处理客户端请求
func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received message:", message)

	conn.Write([]byte("Message received."))
	conn.Close()
}
