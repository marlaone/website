package contents

import (
	"github.com/marlaone/website/pkg/templates"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

func Handler(logger *zap.Logger) http.Handler {
	parser := NewParser()
	views := templates.NewTemplates(viper.GetString("app.views"))

	if err := views.Load(); err != nil {
		logger.Fatal("failed to load app views", zap.Error(err))
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		contentsPath, err := GetContentsPath(r.URL.Path)
		if _, ok := err.(*ContentNotFound); ok {
			http.Error(w, "page not found", http.StatusNotFound)
			return
		}

		content, mdMeta, err := parser.Parse(contentsPath)
		if err != nil {
			logger.Error("parsing content failed", zap.String("path", contentsPath), zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := views.Parse(w, mdMeta.Layout, "content", templates.TemplateVars{
			Content: template.HTML(content),
		}); err != nil {
			logger.Error("parsing template failed", zap.String("path", contentsPath), zap.Error(err))
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

	})
}
