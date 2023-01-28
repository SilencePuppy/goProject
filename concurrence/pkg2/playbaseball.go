package pkg2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			break
		}
		n := rand.Intn(100)
		if n%113 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}

func OpenBaseball() {
	court := make(chan int)
	wg.Add(2)

	go player("lxb", court)
	go player("cyn", court)

	court <- 1
	wg.Wait()
}
