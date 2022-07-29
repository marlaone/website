package contents

import (
	"os"
	"path"
	"strings"

	"github.com/marlaone/website/pkg/config"
	"github.com/spf13/viper"
)

func GetContentsPath(urlPath string) (string, error) {
	contentsDirectory := viper.GetString(config.KeyAppContents)

	urlPath = strings.TrimSuffix(urlPath, "/")
	p := path.Dir(urlPath)
	bn := path.Base(urlPath)

	if p == bn {
		bn = ""
	}

	fileName := bn
	if bn == "" {
		fileName = "index"
	}

	fileName = fileName + ".md"

	contentsPath := path.Join(contentsDirectory, p, fileName)

	if _, err := os.Stat(contentsPath); os.IsNotExist(err) {
		contentsPath = path.Join(contentsDirectory, p, bn, "index.md")
	}

	if _, err := os.Stat(contentsPath); os.IsNotExist(err) {
		return "", NewContentNotFoundError(urlPath)
	}
	return contentsPath, nil
}
