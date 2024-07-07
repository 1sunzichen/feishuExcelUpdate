package main

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	var body [100]byte
	for true {
		_, err := conn.Read(body[:])
		if err != nil {
			break
		}
		fmt.Printf("收到消息：%s \n", string(body[:]))
		_, err = conn.Write(body[:])
		if err != nil {
			break
		}
		fmt.Printf("写消息：%s \n", string(body[:]))

	}
}
func main() {
	l, _ := net.Listen("tcp", ":8080")
	for true {
		conn, _ := l.Accept()
		// fmt.Println(conn, "conn")
		go handleConn(conn)
	}
}
