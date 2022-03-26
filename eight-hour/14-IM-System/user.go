//===----------------------------- 14-IM-System/user.go - [eight-hour] -----------------------------------*- Go -*-===//
// Brief :
//
//
// Author: YongDu
// Date  : 2022-03-26
//===--------------------------------------------------------------------------------------------------------------===//

package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

func (u *User) Online() {
	// 用户上线，将用户加入到 onlineMap 中
	u.server.mapLock.Lock()
	u.server.OnlineUsers[u.Name] = u
	u.server.mapLock.Unlock()

	u.server.BroadCast(u, "already online.")
}

func (u *User) Offline() {
	u.server.mapLock.Lock()
	delete(u.server.OnlineUsers, u.Name)
	u.server.mapLock.Unlock()

	u.server.BroadCast(u, "already offline.")
}

func (u *User) SendMsg(msg string) {
	u.conn.Write([]byte(msg))
}

func (u *User) DoMessage(msg string) {
	if msg == "who" {
		u.server.mapLock.Lock()
		for _, user := range u.server.OnlineUsers {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ": online...\n"
			u.SendMsg(onlineMsg)
		}
		u.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		// 判断 newName 是否已存在
		if _, ok := u.server.OnlineUsers[newName]; ok {
			u.SendMsg("UserName: " + newName + " has been used.\n")
		} else {
			u.server.mapLock.Lock()
			delete(u.server.OnlineUsers, u.Name)
			u.server.OnlineUsers[newName] = u
			u.server.mapLock.Unlock()

			u.Name = newName
			u.SendMsg("Update username successed.\n")
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		// msg: to|username|message
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			u.SendMsg("UserName Is Empty.\n")
			return
		}
		remoteUser, ok := u.server.OnlineUsers[remoteName]
		if !ok {
			u.SendMsg("UserName [" + remoteName + "] does not exist.\n")
			return
		}
		content := strings.Split(msg, "|")[2]
		if content == "" {
			u.SendMsg("Message is empty.\n")
			return
		} else {
			remoteUser.SendMsg("[" + u.Name + "] to you: " + content)
		}
	} else {
		u.server.BroadCast(u, msg)
	}
}

// 监听User chan，有消息便发送给对端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
