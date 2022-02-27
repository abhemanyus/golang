package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
}

func (ss *SpyStore) Fetch(ctx context.Context) (string, error) {
	ch := make(chan string, 1)

	go func() {
		var result string
		for _, c := range ss.response {
			select {
			case <-ctx.Done():
				log.Println("spystore got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		ch <- result
	}()
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res, nil
	}
}

func (ss *SpyStore) Cancel() {
	ss.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world!"
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if !store.cancelled {
			t.Errorf("store was not told to cancel")
		}
	})

	t.Run("returns data from the store", func(t *testing.T) {
		data := "hello, world!"
		store := SpyStore{response: data}
		svr := Server(&store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)
		got := response.Body.String()
		if got != data {
			t.Errorf("want %q, got %q", data, got)
		}

		if store.cancelled {
			t.Error("it should not have cancelled the store")
		}
	})
}
