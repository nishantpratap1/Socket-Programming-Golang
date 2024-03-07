package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func ReverseStr(str string) (result string) {
	for _, j := range str {
		result = string(j) + result
	}
	return result
}

type Response struct {
	Message string
}

func main() {
	fmt.Println("Server that listens to client and reverses the string and sends it back to client")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go HandleClientRequest(conn)
	}
}

func HandleClientRequest(conn net.Conn) {
	defer conn.Close()
	for {
		var request struct {
			Message string
		}
		decoder := gob.NewDecoder(conn)
		if err := decoder.Decode(&request); err != nil {
			fmt.Println("Error decoding request:", err)
			return
		}
		reversed := ReverseStr(request.Message)
		response := Response{Message: reversed}

		encoder := gob.NewEncoder(conn)
		if err := encoder.Encode(response); err != nil {
			fmt.Println("Error encoding response:", err)
			return
		}
	}
}
