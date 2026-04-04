package main

import (
	"fmt"

	"github.com/NjiruClinton/chatroom/internal/chatroom"
)

func main() {
	fmt.Println("Starting server from cmd/server...")
	chatroom.StartServer()
}
