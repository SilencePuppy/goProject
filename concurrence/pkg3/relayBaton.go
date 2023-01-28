package pkg3

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func runner2(baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d Running With Baton\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d please be ready\n", newRunner)
		go runner2(baton)
	}

	time.Sleep(2 * time.Second)

	if runner == 4 {
		fmt.Printf("Runner %d Finished,Race Over\n", runner)
		wg.Done()
		return
	}
	fmt.Printf("Runner %d over,Exchange With Runner %d\n", runner, newRunner)

	baton <- newRunner

}

func Starter() {
	baton := make(chan int)
	wg.Add(1)
	go runner2(baton)

	baton <- 1

	wg.Wait()
}
