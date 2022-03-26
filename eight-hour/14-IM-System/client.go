//===--------------------------- 14-IM-System/client.go - [eight-hour] -----------------------------------*- Go -*-===//
// Brief :
//
//
// Author: YongDu
// Date  : 2022-03-26
//===--------------------------------------------------------------------------------------------------------------===//

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
	conn       net.Conn
	flag       int
}

func (c *Client) menu() bool {
	var flag int
	fmt.Println("[ 1 ] Public Chat")
	fmt.Println("[ 2 ] Private Chat")
	fmt.Println("[ 3 ] Update UserName")
	fmt.Println("[ 0 ] Exit Chat")

	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		c.flag = flag
		return true
	} else {
		fmt.Println("Please input fightful num.")
		return false
	}
}

func (c *Client) Run() {
	for c.flag != 0 {
		for c.menu() != true {
		}
		switch c.flag {
		case 1:
			c.PublicChat()
			break
		case 2:
			c.PrivateChat()
			break
		case 3:
			c.UpdateUserName()
			break
		}
	}
}

func (c *Client) PublicChat() {
	var chatMsg string
	fmt.Println("Please chat content, enter 'exit' to quit.")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
				fmt.Printf("conn Write Error: %+v", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("Please chat content, enter 'exit' to quit.")
		fmt.Scanln(&chatMsg)
	}
}

func (c *Client) SearchUsers() {
	sendMsg := "who\n"
	if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
		fmt.Printf("conn.Write Error: %+v", err)
	}
}

func (c *Client) PrivateChat() {
	var (
		remoteName string
		chatMsg    string
	)

	c.SearchUsers()
	fmt.Println("Please input object user, enter 'exit' to quit.")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println("Please chat content, enter 'exit' to quit.")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
					fmt.Printf("conn.Write Error: %+v", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("Please chat content, enter 'exit' to quit.")
			fmt.Scanln(&chatMsg)
		}
		c.SearchUsers()
		fmt.Println("Please input object user, enter 'exit' to quit.")
		fmt.Scanln(&remoteName)
	}
}

func (c *Client) UpdateUserName() bool {
	fmt.Println("Please input new username: ")
	fmt.Scanln(&c.Name)

	sendMsg := "rename|" + c.Name + "\n"
	if _, err := c.conn.Write([]byte(sendMsg)); err != nil {
		fmt.Printf("conn.Write Error: %+v", err)
		return false
	}
	return true
}

func (c *Client) DealResponse() {
	io.Copy(os.Stdout, c.conn)
}

func NewClient(srvIp string, srvPort int) *Client {
	client := &Client{
		ServerIp:   srvIp,
		ServerPort: srvPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", srvIp, srvPort))
	if err != nil {
		fmt.Printf("net.Dial Error: %+v", err)
		return nil
	}
	client.conn = conn
	return client
}

var (
	srvIp   string
	srvPort int
)

func init() {
	flag.StringVar(&srvIp, "ip", "127.0.0.1", "Set Server IP")
	flag.IntVar(&srvPort, "port", 8888, "Set Server Port")
}

func main() {
	flag.Parse()
	client := NewClient(srvIp, srvPort)
	if client == nil {
		fmt.Println("Connect to server failed.")
		return
	}

	go client.DealResponse()

	fmt.Println("Connect to server successed.")
	client.Run()
}
