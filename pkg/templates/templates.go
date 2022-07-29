package templates

import (
	"html/template"
	"io"

	"github.com/kataras/blocks"
	"github.com/marlaone/website/pkg/config"
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
	views := blocks.New(viewsPath)
	if config.IsDebug() {
		views.Reload(true)
	}
	return &Templates{
		blocks: views,
	}
}

func (t *Templates) Load() error {
	return t.blocks.Load()
}

func (t *Templates) Parse(w io.Writer, layout string, template string, data TemplateVars) error {
	return t.blocks.ExecuteTemplate(w, template, layout, data)
}
