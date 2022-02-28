package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/abhemanyus/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody  = "Title: Post 1\nDescription: first post\nTags: tdd, go\n---\nOne Two \nThree"
		secondBody = "Title: Post 2\nDescription: second post\nTags: ooga, booga\n---\nFour Five Six"
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "first post",
		Tags:        []string{"tdd", "go"},
		Body:        "One Two \nThree",
	})
	assertPost(t, posts[1], blogposts.Post{
		Title:       "Post 2",
		Description: "second post",
		Tags:        []string{"ooga", "booga"},
		Body:        "Four Five Six",
	})
}

func assertPost(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, but want %+v", got, want)
	}
}
