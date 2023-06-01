package xhttp

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

// GetRemoteClientIp 获取远程客户端 IP
func GetRemoteClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr

	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}

	return remoteIp
}

// GetOutBoundIP 获取本地对外 IP，也就是公网 IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// GetLocalIP 获取内网ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ip, err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}
