package blogrender_test

import (
	"bytes"
	"io"
	"lgwt/blogrender"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrender.Post{
		Title:       "Hello, World!",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Single post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("Render index", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrender.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRedner(b *testing.B) {
	aPost := blogrender.Post{
		Title:       "Hello, World!",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrender.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
