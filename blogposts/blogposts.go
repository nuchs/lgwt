package blogposts

import (
	"io/fs"
	"testing/fstest"
)

func NewPostsFromFS(files fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(files, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}

type Post struct {
}
