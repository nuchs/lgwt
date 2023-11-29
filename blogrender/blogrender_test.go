package blogrender_test

import (
	"bytes"
	"lgwt/blogrender"
	"testing"
)

func TestRender(t *testing.T) {
	aPost := blogrender.Post{
		Title:       "Hello, World!",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("Single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrender.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello, World!</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})
}
