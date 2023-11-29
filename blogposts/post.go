package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	title := readline(scanner, "Title: ")
	description := readline(scanner, "Description: ")
	tags := strings.Split(readline(scanner, "Tags: "), ", ")

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
