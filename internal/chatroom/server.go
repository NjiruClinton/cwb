package chatroom

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func runServer() {
	chatRoom, err := NewChatRoom("./chatdata")
	if err != nil {
		fmt.Printf("Failed to initialize: %v\n ", err)
		return
	}
	defer chatRoom.shutdown()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nReceived shutdown signal")
		chatRoom.shutdown()
		os.Exit(0)
	}()

	go chatRoom.Run()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
	defer listener.Close()

	fmt.Println("Server started on :9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			continue
		}
		fmt.Println("New connection from: ", conn.RemoteAddr())
		go handleClient(conn, chatRoom)
	}

}

func (cr *Chatroom) shutdown() {
	fmt.Println("\nShutting down...")
	if err := cr.createSnapshot(); err != nil {
		fmt.Printf("Final snapshot failed: %v\n", err)
	}
	if cr.walFile != nil {
		cr.walFile.Close()
	}
	fmt.Println("Shutdown complete")
}

func StartServer() {
	runServer()
}
