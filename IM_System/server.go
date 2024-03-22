package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	maplock   sync.Mutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (s *Server) ListenMessager() {
	for {
		msg := <-s.Message

		s.maplock.Lock()
		for _, cli := range s.OnlineMap {
			cli.c <- msg
		}
		s.maplock.Unlock()
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn)

	s.maplock.Lock()
	s.OnlineMap[user.Name] = user
	s.maplock.Unlock()

	s.BroadCast(user, "已上线")

	// 当前 handler 阻塞
	select {}
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
