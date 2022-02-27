package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data := make(chan string, 1)
		go func() {
			d, err := store.Fetch(ctx)
			if err != nil {
				fmt.Printf("got error from Fetch %q", err)
			}
			data <- d
		}()
		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
			return
		}
	}
}
