package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title       string
	Body        string `accessor:"setter"`
	Description string
	Tags        []string
	HTML        template.HTML
}

func (p *Post) SetHTML(val string) {
	p.HTML = template.HTML(val)
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
	html := markdown.ToHTML([]byte(p.Body), nil, nil)
	p.SetHTML(string(html))

	if err := r.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
