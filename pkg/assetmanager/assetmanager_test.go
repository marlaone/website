package assetmanager_test

import (
	"github.com/marlaone/website/pkg/assetmanager"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAssetManager(t *testing.T) {

	am := assetmanager.NewAssetManager()

	mainBundle := assetmanager.NewBundle("main", &assetmanager.BundleConfig{})
	am.AddBundle(mainBundle)

	assert.Nil(t, am.Process())
}
