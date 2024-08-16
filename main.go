package main

import (
	"fmt"
	"net"
	"os"

	"github.com/akamensky/argparse"
	"github.com/mingrammer/cfmt"
)

func ping(ip string, port int, pool chan int) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		pool <- 0
		return
	}
	defer conn.Close()
	pool <- port
}

func main() {
	parser := argparse.NewParser("端口扫描", "端口扫描工具")
	ip := parser.String("i", "ip", &argparse.Options{Required: true, Help: "要扫描的IP地址"})
	startPort := parser.Int("s", "startPort", &argparse.Options{Required: true, Help: "起始端口"})
	endPort := parser.Int("e", "endPort", &argparse.Options{Required: true, Help: "结束端口"})
	err := parser.Parse(os.Args)
	if err != nil {
		cfmt.Info(parser.Usage(err))
		return
	}

	if *startPort > *endPort {
		cfmt.Errorf("起始端口不能大于结束端口\n")
		return
	}

	if *startPort < 1 || *endPort > 65535 {
		cfmt.Errorf("端口范围必须在1到65535之间\n")
		return
	}

	if *startPort == *endPort {
		cfmt.Errorf("起始端口和结束端口不能相同\n")
		return
	}

	var pool chan int = make(chan int, *endPort-*startPort+1)
	cfmt.Infof("开始扫描 %s 从 %d 到 %d 端口的开放情况\n", *ip, *startPort, *endPort)
	for i := *startPort; i <= *endPort; i++ {
		go ping(*ip, i, pool)
	}
	for i := *startPort; i <= *endPort; i++ {
		if port := <-pool; port != 0 {
			cfmt.Successf("%d 端口开放\n", port)
		}
	}
	cfmt.Successf("端口扫描结束\n")
}
