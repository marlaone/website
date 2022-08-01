package assetmanager

type Loader interface {
	Test(asset Asset) bool
	Process(asset Asset) error
}
