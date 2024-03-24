package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name         string
	Addr         string
	User_Channel chan string
	conn         net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:         userAddr,
		Addr:         userAddr,
		User_Channel: make(chan string),
		conn:         conn,
		server:       server,
	}

	go user.ListenMessage()

	return user
}

// 用户上线业务
func (u *User) Online() {
	u.server.maplock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.maplock.Unlock()

	u.server.BroadCast(u, "已上线")
}

// 用户下线业务
func (u *User) Offline() {
	u.server.maplock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.maplock.Unlock()

	u.server.BroadCast(u, "已下线")
}

// 给当前User对应的客户端发送消息
func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

// 用户消息处理业务
func (u *User) DoMessage(msg string) {
	// 查询在线
	if msg == "who" {
		u.server.maplock.Lock()
		for _, cli := range u.server.OnlineMap {
			onlineMsg := "[" + cli.Addr + "]" + cli.Name + ":" + "在线...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.maplock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 更改用户名
		newName := strings.Split(msg, "|")[1]
		_, ok := u.server.OnlineMap[newName]
		if ok {
			u.SendMsg("当前名称已被占用\n")
		} else {
			u.server.maplock.Lock()
			delete(u.server.OnlineMap, u.Name)
			u.server.OnlineMap[newName] = u
			u.server.maplock.Unlock()
			u.Name = newName
			u.SendMsg("您的用户名更新为：" + u.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 私聊，to|用户名|消息
		// 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMsg("消息格式不正确，请使用\"to|张三|你好\".\n")
			return
		}
		//根据用户名获取User对象
		remoteUser, ok := u.server.OnlineMap[remoteName]
		if !ok {
			u.SendMsg("用户不存在\n")
			return
		}
		//获取消息内容，通过对方User对象将消息发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(u.Name + "对您说：" + content + "\n")
	} else {
		u.server.BroadCast(u, msg)
	}
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
