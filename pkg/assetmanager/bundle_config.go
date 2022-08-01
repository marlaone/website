package assetmanager

type BundleConfig struct {
	loaders []Loader
}

func NewBundleConfig() *BundleConfig {
	return &BundleConfig{
		loaders: []Loader{},
	}
}

func (c *BundleConfig) AddLoader(loader Loader) {
	c.loaders = append(c.loaders, loader)
}

func (c *BundleConfig) Loaders() []Loader {
	return c.loaders
}
