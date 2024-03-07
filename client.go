package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
)

type Response struct {
	Message string
}

type Request struct {
	Message string
}

func main() {
	fmt.Println("Client that sends string to server and waits for the response from server for the reversed string")
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter Your String: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input) // Remove trailing newline

		request := Request{Message: input}
		encoder := gob.NewEncoder(conn)
		if err := encoder.Encode(request); err != nil {
			fmt.Println("Error encoding request:", err)
		}

		var response Response
		decoder := gob.NewDecoder(conn)
		if err := decoder.Decode(&response); err != nil {
			fmt.Println("Error decoding response:", err)
		}

		fmt.Println("Reply from Server : ", response.Message)
	}
}
