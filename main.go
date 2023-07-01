package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	GenerateHTML()
	IPAddressLists()
	fmt.Println("Port Number : ", SetPort())
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.ListenAndServe(":"+SetPort(), nil)
}

func SetPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4070"
	}
	return port
}

func GenerateHTML() {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	tmp := "<a href='%s'>%s</a><br>"
	html := ""
	for i, file := range files {
		namef := strconv.Itoa(i+1) + ". " + file.Name()
		html += fmt.Sprintf(tmp, file.Name(), namef)
	}
	d1 := []byte(html)
	err = os.WriteFile("index.html", d1, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func IPAddressLists() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			fmt.Println(ip.String())
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
