package pkg8

import (
	"fmt"
	"sync"
)

func Demo1()  {
	var wg sync.WaitGroup
	for _,s :=range []string{"one","two","three"}{
		//m:=s
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(s)
		}()
	}
	wg.Wait()

}
