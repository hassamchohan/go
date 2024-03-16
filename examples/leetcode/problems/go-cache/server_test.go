package go_cache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestGetUserHandler(t *testing.T) {
	s := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(s.handleGetUser))
	numOfReq := 1000
	wg := &sync.WaitGroup{}

	for i := 0; i < numOfReq; i++ {
		wg.Add(1)
		go func(i int) {
			id := i%100 + 1
			url := fmt.Sprintf("%s/?id=%d", ts.URL, id)
			resp, err := http.Get(url)
			if err != nil {
				t.Error(err)
			}
			user := &User{}
			if err := json.NewDecoder(resp.Body).Decode(user); err != nil {
				t.Error(err)
			}
			fmt.Printf("%+v\n", user)
			wg.Done()
		}(i)

		time.Sleep(time.Millisecond)
	}

	wg.Wait()

	fmt.Println("Times we hit the database: ", s.dbHit)
}
