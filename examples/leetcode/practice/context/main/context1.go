package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userId := 10

	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", val)
	fmt.Println("took:  ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respCh := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffThatCanBeSlow()
		respCh <- Response{
			val,
			err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party takes too long")
		case resp := <-respCh:
			return resp.value, resp.err

		}
	}
}

func fetchThirdPartyStuffThatCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100)
	return 1, nil
}
