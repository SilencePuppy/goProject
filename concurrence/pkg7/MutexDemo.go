package pkg7

import (
	"fmt"
	"sync/atomic"
	"time"
)

var seq int64

func genId() int64 {
	 atomic.AddInt64(&seq,1)
	 fmt.Println(seq)
	 return seq
}

func Mutex1(){
	for i := 0; i < 10; i++ {
		go genId()
	}
	time.Sleep(1*time.Second)
	fmt.Println(genId())
}