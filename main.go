package main

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Msg string
	IP  string
}

func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

func getSchools(c *gin.Context) {
	ip, _ := GetLocalIP()
	var res = Result{Msg: "hello~", IP: ip}
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
	router := gin.Default()
	// url - 方法 映射
	router.GET("/index", getSchools)

	router.Run("localhost:8081")
}
