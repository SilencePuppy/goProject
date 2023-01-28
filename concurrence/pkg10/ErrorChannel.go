package pkg10

import (
	"fmt"
	"net/http"
)

func checkStatus(done <-chan interface{}, urls ...string) <-chan *http.Response {
	responses := make(chan *http.Response)
	go func() {
		defer close(responses)
		for _, url := range urls {
			response, err := http.Get(url)

			if err != nil {
				fmt.Println(err)
				continue
			}

			select {
			case <-done:
			case responses <- response:
			}
		}
	}()

	return responses
}

func GetUrlResponseWithoutError() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://badlocahost"}
	responses := checkStatus(done, urls...)
	for response := range responses {
		fmt.Printf("Response: %v\n", response.Status)
	}
}

type ResponseResult struct {
	Error    error
	Response *http.Response
}

func CheckUrlStatus(done <-chan interface{}, urls ...string) <-chan ResponseResult {
	responses := make(chan ResponseResult)
	go func() {
		defer close(responses)
		for _, url := range urls {
			resp, err := http.Get(url)
			result := ResponseResult{err, resp}
			select {
			case <-done:
			case responses <- result:
			}
		}
	}()
	return responses
}

func GetUrlResponseWithError()  {
	done:=make(chan interface{})
	urls:=[]string{"https://www.baidu.com","https://www.badlocalhost.com"}
	responseStatus := CheckUrlStatus(done, urls...)
	defer close(done)
	for result := range responseStatus {
		if result.Error!=nil {
			fmt.Println("error:%v",result.Error)
			continue
		}
		fmt.Printf("Response:%v\n",result.Response.Status)
	}
}
