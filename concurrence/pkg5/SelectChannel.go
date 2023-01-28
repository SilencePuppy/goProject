package pkg5

import (
	"errors"
	"fmt"
	"time"
)

func RPCServer(ch chan string) {
	for {
		req := <-ch

		fmt.Printf("Server receive message:%s\n", req)

		ch <- "roger"
	}
}

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case rep := <-ch:
		fmt.Printf("Server response:%s\n", rep)
		return rep, nil
	case <-time.After(3 * time.Second):
		return "", errors.New("time out")
	}
}

func DemoDate() {
	ch := make(chan string)

	go RPCServer(ch)

	rep, err := RPCClient(ch, "lxb")
	if err != nil {
		fmt.Println("超时无响应")
	} else {
		fmt.Printf("服务器正常响应，返回结果：%s\n",rep)
	}
}

func MyRandom() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 1:
		case ch <- 0:
		}
	}
}
