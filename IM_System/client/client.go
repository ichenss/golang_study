package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	Conn       net.Conn
	flag       int // 当前客户端模式
}

func NewClient(serverIP string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIP,
		ServerPort: serverPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}
	client.Conn = conn
	return client
}

// 公聊模式
func (c *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>>>请输入聊天内容, exit退出")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := c.Conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("client write err: ", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>>请输入聊天内容, exit退出")
		fmt.Scanln(&chatMsg)
	}
}

// 处理server回应的消息
func (c *Client) DealResponse() {
	// 一旦 c.Conn 有数据，就直接copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, c.Conn)
}

func (c *Client) UpdateName() bool {
	fmt.Println(">>>>>请输入用户名：")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	_, err := c.Conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("client conn write err: ", err)
		return false
	}
	return true
}

func (c *Client) Menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字<<<<<")
		return false
	}
}

func (c *Client) Run() {
	for c.flag != 0 {
		for !c.Menu() {
		}
	Switch:
		switch c.flag {
		case 1:
			// fmt.Println("公聊模式选择...")
			c.PublicChat()
			break Switch
		case 2:
			fmt.Println("私聊模式选择...")
			break Switch
		case 3:
			// fmt.Println("更新用户名选择...")
			c.UpdateName()
			break Switch
		}
	}
}

var serverIP string
var serverPort int

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址(默认为127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8080, "设置服务器端口(默认8080)")
}

func main() {
	flag.Parse()

	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>>> 连接服务器失败...")
		return
	}

	// 单独开启一个goroutine处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>> 连接服务器成功")

	client.Run()
}
