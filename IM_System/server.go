package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	maplock   sync.Mutex

	Server_Message_Channel chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Server_Message_Channel:   make(chan string),
	}
	return server
}

func (s *Server) ListenMessager() {
	for {
		msg := <-s.Server_Message_Channel

		s.maplock.Lock()
		for _, cli := range s.OnlineMap {
			cli.User_Channel <- msg
		}
		s.maplock.Unlock()
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Server_Message_Channel <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn, s)

	user.Online()

	isLive := make(chan bool)
	// 用户消息广播
	go func(){
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF{
				fmt.Println("read error: ", err)
				return
			}
			msg := string(buffer[:n-1])

			user.DoMessage(msg)

			isLive <- true
		}
	}()

	// 超时踢出
	for {
		select {
		case <- isLive:

		case <- time.After(time.Second * 15):
			user.SendMsg("你已被踢出")
			close(user.User_Channel)
			user.conn.Close()
			return
		}
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	defer listener.Close()

	go s.ListenMessager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}

		go s.Handler(conn)
	}
}
