package contents

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type ParserCacheItem struct {
	Content []byte
	Meta    MDMeta
}
type ParserCache map[string]*ParserCacheItem

func (m ParserCache) Exists(key string) bool {
	_, ok := m[key]
	return ok
}

func (m ParserCache) Set(key string, content []byte, meta MDMeta) {
	m[key] = &ParserCacheItem{
		Content: content,
		Meta:    meta,
	}
}

func (m ParserCache) Get(key string) *ParserCacheItem {
	ci, ok := m[key]
	if ok {
		return ci
	}
	return nil
}

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
	cache     ParserCache
}

func NewParser() *Parser {
	p := &Parser{
		// @TODO implement real cache with check for modification time or expiry
		cache: ParserCache{},
	}
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

	if p.cache.Exists(markdownPath) {
		item := p.cache.Get(markdownPath)
		return item.Content, item.Meta, nil
	}

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

	p.cache.Set(markdownPath, sanitized, metaData)

	return sanitized, metaData, nil
}
