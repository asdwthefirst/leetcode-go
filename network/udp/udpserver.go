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
	// 监听UDP端口
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err.Error())
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Server listening on localhost:8080")

	buffer := make([]byte, 1024)
	for {
		// 读取客户端发送的数据
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		message := string(buffer[:n])
		fmt.Println("Received message:", message)

		// 向客户端发送响应数据
		response := []byte("Message received.")
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
