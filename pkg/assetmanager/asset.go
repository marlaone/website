package assetmanager

import (
	"bytes"
	"io"
)

type Asset interface {
	io.Writer
	Path() string
	Bytes() []byte
}

type JavaScriptAssetOption func(js *JavaScriptAsset)

type JavaScriptAsset struct {
	path     string
	contents bytes.Buffer
}

var _ Asset = &JavaScriptAsset{}

func JavaScriptFromPath(path string) *JavaScriptAsset {
	return &JavaScriptAsset{
		path: path,
	}
}

func (js *JavaScriptAsset) Path() string {
	return js.path
}

func (js *JavaScriptAsset) Write(p []byte) (n int, err error) {
	return js.contents.Write(p)
}

func (js *JavaScriptAsset) Bytes() []byte {
	return js.contents.Bytes()
}

type CSSAssetOption func(css *CSSAsset)

type CSSAsset struct {
	path     string
	contents bytes.Buffer
}

var _ Asset = &CSSAsset{}

func (css *CSSAsset) Path() string {
	return css.path
}

func (css *CSSAsset) Write(p []byte) (n int, err error) {
	return css.contents.Write(p)
}

func (css *CSSAsset) Bytes() []byte {
	return css.contents.Bytes()
}

func CSSFromPath(path string) *CSSAsset {
	return &CSSAsset{
		path: path,
	}
}
