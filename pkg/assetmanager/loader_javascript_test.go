package assetmanager_test

import (
	"github.com/marlaone/website/pkg/assetmanager"
	"github.com/stretchr/testify/assert"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"
)

func TestNewJSLoader(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	jsLoader := assetmanager.NewJSLoader(regexp.MustCompile(`(.*)\.js`))
	asset := assetmanager.JavaScriptFromPath(path.Join(filepath.Dir(filename), "./__mock__/read_test.js"))
	err := jsLoader.Process(asset)
	if err != nil {
		t.Fatal(err)
	}
	assert.Greater(t, len(asset.Bytes()), 0)
	assert.Equal(t, string(asset.Bytes()), "console.log('hello world')")
}
