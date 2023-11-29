package blogposts_test

import (
	"errors"
	"io/fs"
	"lgwt/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestPostFromFS(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol
Tags: tdd, go
---
Hello
World`)},
		}

		posts, err := blogposts.PostsFromFS(fs)

		if err != nil {
			t.Fatal("Failed to read files:", err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		})
	})

	t.Run("Failing file system", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})

		if err == nil {
			t.Errorf("expcted an error, didn't get one")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no!")
}
