package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}

// func Render(w io.Writer, p Post) error {
// 	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
// 	if err != nil {
// 		return err
// 	}

// 	if err := templ.Execute(w, p); err != nil {
// 		return err
// 	}

// 	return nil
// }
