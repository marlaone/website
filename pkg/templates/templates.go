package templates

import (
	"github.com/kataras/blocks"
	"html/template"
	"io"
)

type HTMLMeta struct {
	Title         string
	Description   string
	Keywords      []string
	OGTitle       string
	OGDescription string
	OGImage       string
	OGAuthor      string
	OGPublisher   string
	OGSiteName    string
}

type TemplateVars struct {
	Content template.HTML
	Meta    HTMLMeta
}

type Templates struct {
	blocks *blocks.Blocks
}

func NewTemplates(viewsPath string) *Templates {
	return &Templates{
		blocks: blocks.New(viewsPath),
	}
}

func (t *Templates) Load() error {
	return t.blocks.Load()
}

func (t *Templates) Parse(w io.Writer, layout string, template string, data TemplateVars) error {
	return t.blocks.ExecuteTemplate(w, template, layout, data)
}
