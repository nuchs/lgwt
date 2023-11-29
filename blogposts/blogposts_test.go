package blogposts_test

import (
	"errors"
	"io/fs"
	"lgwt/blogposts"
	"testing"
	"testing/fstest"
)

func TestPostFromFS(t *testing.T) {
	t.Run("Happy Path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol`)},
			// "hello-world2.md": {Data: []byte("Title: Hello twitchy world")},
		}

		posts, err := blogposts.PostsFromFS(fs)

		if err != nil {
			t.Fatal("Failed to read files:", err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		expectedPost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
		}
		if posts[0] != expectedPost {
			t.Errorf("got %#v, want %#v", posts[0], expectedPost)
		}
	})

	t.Run("Failing file system", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})

		if err == nil {
			t.Errorf("expcted an error, didn't get one")
		}
	})
}

type FailingFS struct {
}

// Open implements fs.FS.
func (FailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Oh no!")
}
