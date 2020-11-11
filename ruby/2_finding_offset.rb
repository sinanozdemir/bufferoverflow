require 'socket'

buff = #Based on at what bytes you get to EIP, generate your pattern and insert it here, example "AAAA"
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

begin
    s = TCPSocket.open(RHOST, RPORT)
    s.puts buff + '\r\n'
    s.close
rescue Interrupt
    puts "Debugger crashed at #{buffSize} bytes"
end
