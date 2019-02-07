package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func response(w http.ResponseWriter, r *http.Request) {
	ip := GetOutboundIP()
	hostname, _ := os.Hostname()
	output := "Hello from " + ip.String() + "! Hostname: " + hostname
	fmt.Println(`## INCOMING REQUEST (this may be duplicated if coming from a browser)`)
	fmt.Fprintln(w, output) // send data to client side
}

func main() {
	portPtr := flag.Int("p", 8080, "port")
	flag.Parse()
	port := fmt.Sprintf("%v", *portPtr)
	outbound := GetOutboundIP()
	fmt.Println("LISTENING ON " + outbound.String() + ":" + port)
	http.HandleFunc("/", response)
	err := http.ListenAndServe(":"+fmt.Sprintf("%v", *portPtr), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
