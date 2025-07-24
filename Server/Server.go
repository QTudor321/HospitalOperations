package main

import (
	"HospitalQOps/errorspacket"
	"HospitalQOps/model"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

var (
	connectedClients = make(map[string]net.Conn)
	mutex            sync.RWMutex
)

func main() {
	fmt.Println("--QHospital Central Server--")
	ipAddress := "127.0.0.1:50000"
	listener, err := net.Listen("tcp", ipAddress)
	if err != nil {
		fmt.Println(errorspacket.ListenerNetworkError)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening to: ", listener.Addr())
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(errorspacket.NetworkConnectionError)
			return
		}
		go handleConnection(connection)
	}
}
func handleConnection(con net.Conn) {
	defer con.Close()
	fmt.Println("New client connected:", con.RemoteAddr())

	var userEmail string

	for {
		buffer := make([]byte, 4096)
		n, err := con.Read(buffer)
		if err != nil {
			fmt.Println("Client disconnected:", con.RemoteAddr())
			if userEmail != "" {
				mutex.Lock()
				delete(connectedClients, userEmail)
				mutex.Unlock()
			}
			return
		}

		var msg model.Message
		err = json.Unmarshal(buffer[:n], &msg)
		if err != nil {
			fmt.Println("Invalid message format:", err)
			continue
		}
		if userEmail == "" && msg.Sender != "" {
			userEmail = msg.Sender
			mutex.Lock()
			connectedClients[userEmail] = con
			mutex.Unlock()
			fmt.Printf("[Registered connection] %s => %s\n", userEmail, con.RemoteAddr())
		}
		mutex.RLock()
		receiverConn, ok := connectedClients[msg.Receiver]
		mutex.RUnlock()

		if ok {
			encodedMsg, _ := json.Marshal(msg)
			receiverConn.Write(encodedMsg)
			fmt.Printf("[Message forwarded] From: %s To: %s\n", msg.Sender, msg.Receiver)
		} else {
			fmt.Printf("[Message NOT delivered] Receiver %s not online\n", msg.Receiver)
		}
	}
}
