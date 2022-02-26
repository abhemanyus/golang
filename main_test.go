package concurrency

import (
	"reflect"
	"testing"
)

func fakeWebsiteChecker(url string) bool {
	return url != "www.nosuchlink.shit"
}

func TestWebsiteChecker(t *testing.T) {
	urls := []string{
		"www.google.com",
		"www.facebook.com",
		"www.nosuchlink.shit",
	}
	want := map[string]bool{
		"www.google.com":      true,
		"www.facebook.com":    true,
		"www.nosuchlink.shit": false,
	}
	got := CheckWebsites(fakeWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
