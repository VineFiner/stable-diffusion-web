package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
)

// 等待端口启动
func wait_for_port(port int) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", port), time.Second)
		if err == nil {
			fmt.Printf("端口 %d 已启动\n", port)
			conn.Close()
			return
		}
		time.Sleep(time.Second)
	}
}

func wait_for_port_once(port int) {
	once.Do(func() {
		wait_for_port(port)
	})
}

func proxy(c *gin.Context) {
	// 健康检查
	if c.Param("proxyPath") == "/healthcheck" {
		c.String(http.StatusOK, "health")
	} else {
		// 等待启动
		wait_for_port_once(7860)

		// 创建反向代理
		remote, err := url.Parse("http://0.0.0.0:7860")
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		//Define the director func
		//This is a good place to log, for example
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			// req.URL.Path = c.Param("proxyPath")
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	//Create a catchall route
	router.Any("/*proxyPath", proxy)

	router.Run(":9000")
}
