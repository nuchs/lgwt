package blogrender

import (
	"embed"
	"fmt"
	"html/template"
	"io"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func Render(w io.Writer, p Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	data, _ := postTemplates.ReadFile("templates/blog.gohtml")
	fmt.Printf("%q", string(data))

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
