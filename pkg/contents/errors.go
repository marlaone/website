package contents

type ContentNotFound struct {
	URLPath string
}

func NewContentNotFoundError(urlPath string) *ContentNotFound {
	return &ContentNotFound{URLPath: urlPath}
}

func (e *ContentNotFound) Error() string {
	return "content not found for " + e.URLPath
}
