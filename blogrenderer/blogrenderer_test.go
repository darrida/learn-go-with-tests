package blogrenderer_test

import (
	"bytes"
	"html/template"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"

	blogrenderer "github.com/darrida/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title: "hello world",
			Body: `
This is a post.

This is **bold**.

This is *italicized*.

This is a [link](https://www.google.com)`,
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
			HTML:        template.HTML(""),
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {

		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post\nThis is **bold**.\nThis is *italicized*.\nThis is a [link](https://www.google.com)",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
