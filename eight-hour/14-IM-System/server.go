package main

import (
	"fmt"
	"net"
	"sync"
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

	user := NewUser(conn)

	s.mapLock.Lock()
	s.OnlineUsers[user.Name] = user
	s.mapLock.Unlock()

	s.BroadCast(user, "already online.")

	select {} // 阻塞当前handler
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
