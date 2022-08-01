package assetmanager

type AssetManager struct {
	bundles Bundles
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		bundles: Bundles{},
	}
}

func (am *AssetManager) AddBundle(bundle Bundle) {
	am.bundles[bundle.Name()] = bundle
}

func (am *AssetManager) Bundle(name string) Bundle {
	return am.bundles[name]
}

func (am *AssetManager) Bundles() Bundles {
	return am.bundles
}

func (am *AssetManager) Process() error {
	for _, b := range am.Bundles() {
		if err := b.Process(); err != nil {
			return err
		}
	}
	return nil
}
