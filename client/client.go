package client


import (
	"net"
	"flag"
	"fmt"
)


func main() {
	ip := flag.String("ip", 4590, "ip the server is on")
	port := flag.Int("port", 4590, "port Honey Bee is open on")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", (*ip, *addr))

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		panic(err.Error())
	}

	//fmt.Fprintf()
}