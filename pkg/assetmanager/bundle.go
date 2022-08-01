package assetmanager

type Bundle interface {
	Name() string
	Config() *BundleConfig
	AddJavaScript(js *JavaScriptAsset, opts ...JavaScriptAssetOption)
	AddCSS(css *CSSAsset, opts ...CSSAssetOption)
	JavaScriptAssets() []*JavaScriptAsset
	CSSAssets() []*CSSAsset
	Process() error
}

type Bundles map[string]Bundle

type AssetBundle struct {
	name        string
	config      *BundleConfig
	javascripts []*JavaScriptAsset
	css         []*CSSAsset
}

var _ Bundle = &AssetBundle{}

func NewBundle(name string, config *BundleConfig) *AssetBundle {
	return &AssetBundle{
		name:   name,
		config: config,
	}
}

func (b *AssetBundle) Config() *BundleConfig {
	return b.config
}

func (b *AssetBundle) Name() string {
	return b.name
}

func (b *AssetBundle) AddJavaScript(js *JavaScriptAsset, opts ...JavaScriptAssetOption) {
	for _, opt := range opts {
		opt(js)
	}

	b.javascripts = append(b.javascripts, js)
}

func (b *AssetBundle) AddCSS(css *CSSAsset, opts ...CSSAssetOption) {
	for _, opt := range opts {
		opt(css)
	}
	b.css = append(b.css, css)
}

func (b *AssetBundle) JavaScriptAssets() []*JavaScriptAsset {
	return b.javascripts
}

func (b *AssetBundle) CSSAssets() []*CSSAsset {
	return b.css
}

func (b *AssetBundle) Process() error {

	for _, l := range b.Config().Loaders() {
		for _, js := range b.JavaScriptAssets() {
			if l.Test(js) {
				if err := l.Process(js); err != nil {
					return err
				}
			}
		}

		for _, css := range b.CSSAssets() {
			if l.Test(css) {
				if err := l.Process(css); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
