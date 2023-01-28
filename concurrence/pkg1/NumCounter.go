package pkg1

import "fmt"

func receiver(ch chan int) {
	for {
		data := <-ch
		if data == 0 {
			break
		}
		fmt.Println(data)
	}
	ch <- 0
}

func dataSender() {
	ch := make(chan int)

	go receiver(ch)

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	ch <- 0
	<-ch
}

func Main() {
	dataSender()
}
