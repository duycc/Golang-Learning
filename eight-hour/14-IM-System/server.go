//===--------------------------- 14-IM-System/server.go - [eight-hour] -----------------------------------*- Go -*-===//
// Brief :
//
//
// Author: YongDu
// Date  : 2022-03-26
//===--------------------------------------------------------------------------------------------------------------===//

package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineUsers map[string]*User // 所有在线用户
	mapLock     sync.RWMutex     // 读写锁
	Message     chan string      // 消息广播
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:          ip,
		Port:        port,
		OnlineUsers: make(map[string]*User),
		Message:     make(chan string),
	}
	return server
}

func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		s.mapLock.Lock()
		for _, cli := range s.OnlineUsers {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg
	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	// 业务逻辑
	// fmt.Println("connect succ...")

	user := NewUser(conn, s)
	user.Online()

	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Printf("Conn Read Error: %+v", err)
				return
			}
			msg := string(buf[:n-1]) // 去除 '\n'
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
			// 重置定时器
		case <-time.After(time.Second * 1000):
			user.SendMsg("You are timeout\n")
			close(user.C)
			conn.Close()
			runtime.Goexit()
		} // 阻塞当前handler
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
	}
	defer listener.Close()
	go s.ListenMessage()
	fmt.Printf("server start, listening at %s:%d\n", s.Ip, s.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err ", err)
			continue
		}
		go s.Handler(conn)
	}
}
