package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("return faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		want := fastServer.URL
		got, err := Racer(fastServer.URL, slowServer.URL, 10*time.Millisecond)

		if err != nil {
			t.Fatalf("didn't expect an error but got %v", err)
		}

		if got != want {
			t.Errorf("want %q, got %q", want, got)
		}
	})

	t.Run("returns an error if server doesn't responds within 10s", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Millisecond)
		fastServer := makeDelayedServer(12 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		_, err := Racer(slowServer.URL, fastServer.URL, 10*time.Millisecond)
		if err == nil {
			t.Error("want error but didn't get any")
		}
	})
}

func makeDelayedServer(latency time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(latency)
		rw.WriteHeader(http.StatusOK)
	}))
}
