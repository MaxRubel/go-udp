package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve the address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":41234")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on address:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 41234")

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			continue
		}

		fmt.Printf("Received %s from %s\n", string(buffer[:n]), clientAddr)

		response := "Hello from Go Server!"
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			continue
		}
	}
}
