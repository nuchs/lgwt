package blogposts_test

import (
	"lgwt/blogposts"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Hello, TDD world!")},
		"hello-world2.md": {Data: []byte("Title: Hello twitchy world")},
	}

	posts := blogposts.NewPostsFromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	expectedPost := blogposts.Post{Title: "Hello, TDD world!"}
	if posts[0] != expectedPost {
		t.Errorf("got %#v, want %#v", posts[0], expectedPost)
	}
}
