package assetmanager

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

type JSLoader struct {
	rule *regexp.Regexp
}

var _ Loader = &JSLoader{}

func NewJSLoader(rule *regexp.Regexp) *JSLoader {
	return &JSLoader{
		rule: rule,
	}
}
func (j *JSLoader) Test(asset Asset) bool {
	log.Println("hello", j.rule.MatchString(asset.Path()))
	return j.rule.MatchString(asset.Path())
}

func (j *JSLoader) Process(asset Asset) error {
	contents, err := ioutil.ReadFile(asset.Path())
	if err != nil {
		return fmt.Errorf("failed to load [%s]: %v", asset.Path(), err)
	}
	_, err = asset.Write(contents)
	if err != nil {
		return fmt.Errorf("failed to write [%s]: %v", asset.Path(), err)
	}
	return nil
}
