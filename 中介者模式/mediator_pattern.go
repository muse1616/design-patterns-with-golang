package main

import (
	"fmt"
	"time"
)

// chatroom作为中介对象 使用饿汉单例生成
type Chatroom struct {
}

var chatroom = &Chatroom{}

func GetChatroom() *Chatroom {
	return chatroom
}

func (c *Chatroom) showMessage(u User, message string) {
	n := time.Now()
	fmt.Println(n.Format("2006-01-02 15:04:05"), "/  ", u.name, ":", message)
}

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{name: name}
}
func (u User) sendMessage(message string) {
	chatroom.showMessage(u, message)
}
