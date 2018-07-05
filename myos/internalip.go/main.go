package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
)

/*
* net.InterfaceAddrs()本地以太网适配器 [fe80::d93f:8a83:170c:8c27/64 192.168.6.233/24 ::1/128 127.0.0.1/8 fe80::5efe:c0a8:6e9/128]
* a.(*net.IPNet) Ipv6 fe80::d93f:8a83:170c:8c27/64
*ipnet.IP.IsLoopback() 回环地址 false
*ipnet.IP.String() 自己内网地址
*http://myexternalip.com/raw 获取外网ip的地址
*
 */
func main() {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	fmt.Println("addrs", addrs)
	if err != nil {
		os.Stderr.WriteString("InternalIPs Err:" + err.Error())
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			fmt.Println("a.(*net.IPNet)", a.(*net.IPNet))
			fmt.Println("ipnet.IP.IsLoopback()", ipnet.IP.IsLoopback())
			if ipnet.IP.To4() != nil {
				fmt.Println("ipnet.IP.To4", ipnet.IP.To4)
				ips = append(ips, ipnet.IP.String())
				fmt.Println("ipnet.IP.String()", ipnet.IP.String())
			}
		}
	}
	fmt.Println("ips", ips)

	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}

//获取外网Ip
func get_external() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}

//获取内网Ip
func get_internal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
