package znsq

import (
	"fmt"
	"net"
	"strings"
)

func getLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("nsq生产者 获取ip地址失败")
		return "", err
	}
	var ip []string
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = append(ip, ipnet.IP.String())
			}
		}
	}
	return strings.Join(ip, ";"), nil
}
