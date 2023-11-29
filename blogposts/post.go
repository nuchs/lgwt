package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
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
	Body        string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	title := readline(scanner, PostTitle)
	description := readline(scanner, PostDesc)
	tags := strings.Split(readline(scanner, PostTags), ", ")
	body := readBody(scanner)

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // skip hyphons
	var buf bytes.Buffer
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func readline(scanner *bufio.Scanner, prefix string) string {
	scanner.Scan()
	return strings.TrimPrefix(scanner.Text(), prefix)
}
