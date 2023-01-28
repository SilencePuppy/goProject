package pkg6

import (
	"fmt"
	"time"
)

func TimerDemo() {
	exit := make(chan int)

	fmt.Println("start")

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("One second after")
		exit <- 0
	})

	<-exit
}

func TimerAndTicker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	timer := time.NewTimer(time.Second * 2)
	var i int
	for {
		select {
		case <-ticker.C:
			i++
			fmt.Printf("我是第%d次打点\n",i)
		case <-timer.C:
			fmt.Println("stop and goto stoplabel")
			goto stopLabel
		}
	}

	stopLabel:
		fmt.Println("all over")
}
