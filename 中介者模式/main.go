package main

func main() {
	GetChatroom()
	u1 := NewUser("muse")
	u2 := NewUser("James")
	u1.sendMessage("hello James")
	u2.sendMessage("hello Muse")
}
