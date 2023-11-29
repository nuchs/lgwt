package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	PostTitle = "Title: "
	PostDesc  = "Description: "
	PostTags  = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	title := readline(scanner, PostTitle)
	description := readline(scanner, PostDesc)
	tags := strings.Split(readline(scanner, PostTags), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}
}

func readline(scanner *bufio.Scanner, prefix string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), prefix)
}
