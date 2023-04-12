package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os/exec"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
)

// 启动服务
func start_webui() {
	// python -u webui.py --listen --port 7860
	cmd := exec.Command("python3", "-u", "webui.py", "--listen", "--port", "7860")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("程序启动成功\n")
}

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
		fmt.Printf("监听端口\n")
		// 监听端口启动
		wait_for_port(port)
	})
}

func initialize(c *gin.Context) {
	// 启动程序
	fmt.Printf("启动程序\n")
	start_webui()
}

func healthcheck(c *gin.Context) {
	// 健康检查
	c.String(http.StatusOK, "health")
}

func proxy(c *gin.Context) {

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

func main() {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	// Register the initializer handler.
	router.Any("/initialize", initialize)

	// healthcheck
	router.Any("/healthcheck", healthcheck)

	//Create a proxy route
	router.Any("/", proxy)

	router.Run(":9000")
}
