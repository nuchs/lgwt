package blogrender

import (
	"fmt"
	"io"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>", p.Title, p.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "\nTags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range p.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "</ul>")
	if err != nil {
		return err
	}

	return err
}