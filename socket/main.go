package main

import (
	"GoLearn/socket/server"
	"flag"
)

func main() {

	var ip, port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "ip")
	flag.StringVar(&port, "port", "8888", "端口")
	newServer := server.NewServer(ip, port)
	newServer.Run()
}
