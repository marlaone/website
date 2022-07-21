package contents

import (
	"bytes"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type SEOMeta struct {
	Title       string
	Description string
	Keywords    []string
}

type MDMeta struct {
	Layout string
	SEO    SEOMeta
}

func MDMetaFromMap(meta map[string]interface{}) MDMeta {
	mdMeta := MDMeta{}

	if title, ok := meta["title"]; ok {
		titleStr, ok := title.(string)
		if ok {
			mdMeta.SEO.Title = titleStr
		}
	}

	if desc, ok := meta["description"]; ok {
		descStr, ok := desc.(string)
		if ok {
			mdMeta.SEO.Description = descStr
		}
	}

	if keyw, ok := meta["keywords"]; ok {
		keywStr, ok := keyw.(string)
		if ok {
			mdMeta.SEO.Keywords = strings.Split(keywStr, ",")
		}
	}

	if layout, ok := meta["layout"]; ok {
		layoutStr, ok := layout.(string)
		if ok {
			mdMeta.Layout = layoutStr
		} else {
			mdMeta.Layout = "default"
		}
	} else {
		mdMeta.Layout = "default"
	}

	return mdMeta
}

type Parser struct {
	md        goldmark.Markdown
	sanitizer *bluemonday.Policy
}

func NewParser() *Parser {
	p := &Parser{}
	p.md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	p.sanitizer = bluemonday.UGCPolicy()
	p.sanitizer.AllowNoAttrs().Matching(regexp.MustCompile(`^marla-`)).OnElementsMatching(regexp.MustCompile(`^marla-`))
	return p
}

func (p *Parser) Parse(markdownPath string) ([]byte, MDMeta, error) {
	var buf bytes.Buffer
	var metaData MDMeta

	f, err := os.Open(markdownPath)
	if err != nil {
		return nil, metaData, fmt.Errorf("[contents/Parser] failed to open file: %v", err)
	}
	defer func() {
		_ = f.Close()
	}()

	markdownContents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, metaData, fmt.Errorf("[contents/Parser] failed to read file: %v", err)
	}

	context := parser.NewContext()
	if err := p.md.Convert(markdownContents, &buf, parser.WithContext(context)); err != nil {
		return nil, metaData, err
	}
	metaData = MDMetaFromMap(meta.Get(context))

	sanitized := p.sanitizer.SanitizeBytes(buf.Bytes())

	return sanitized, metaData, nil
}
