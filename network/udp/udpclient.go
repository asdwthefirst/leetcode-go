/*
*

	@author: jiangxi
	@since: 2023/7/16
	@desc: //TODO

*
*/
// 客户端代码
package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	message := []byte("Hello, server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	response := string(buffer[:n])
	fmt.Println("Server response:", response)
}
