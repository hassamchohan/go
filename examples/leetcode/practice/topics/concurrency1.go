package main

import (
	"fmt"
	"sync"
	"time"
)

func getUserName() string {
	time.Sleep(time.Millisecond * 100)
	return "hassam.chohan343"
}

func getPosts(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "posts"
	wg.Done()
}

func getComments(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- "comments"
	wg.Done()
}

func getAge(username string, respCh chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respCh <- 30
	wg.Done()
}

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	respCh := make(chan any, 3)

	user := getUserName()

	wg.Add(3)
	go getPosts(user, respCh, wg)
	go getComments(user, respCh, wg)
	go getAge(user, respCh, wg)

	wg.Wait()
	close(respCh)

	for resp := range respCh {
		fmt.Println("resp: ", resp)
	}

	// fmt.Println("user: ", user)
	// fmt.Println("posts: ", posts)
	// fmt.Println("comments: ", comments)
	// fmt.Println("age: ", age)
	fmt.Println(time.Since(start))
}
