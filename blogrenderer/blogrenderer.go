package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

type PostViewModel struct {
	Post
	HTMLBody template.HTML
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: tmpl, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func newPostVM(p Post, r *PostRenderer) PostViewModel {
	vm := PostViewModel{Post: p}
	vm.HTMLBody = template.HTML((markdown.ToHTML([]byte(p.Body), r.mdParser, nil)))
	return vm
}
