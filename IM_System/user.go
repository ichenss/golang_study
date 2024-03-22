package main

import (
	"fmt"
	"net"
)

type User struct {
	Name string
	Addr string
	User_Channel chan string
	conn net.Conn
}

func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		User_Channel:    make(chan string),
		conn: conn,
	}
	
	go user.ListenMessage()

	return user
}

// ListenMessage 监听当前User channel 的方法，一旦有消息，就直接发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.User_Channel
		_, err := u.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("conn write error: ", err)
			return
		}
	}
}
