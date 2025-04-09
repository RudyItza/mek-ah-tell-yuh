package render

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateManager struct {
	templates map[string]*template.Template
}

func NewTemplateManager(pattern string) (*TemplateManager, error) {
	tmpl := make(map[string]*template.Template)

	// Load all .tmpl files from the template directory
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		name := filepath.Base(file)
		tmpl[name] = template.Must(template.ParseFiles(file))
	}

	return &TemplateManager{templates: tmpl}, nil
}

func (tm *TemplateManager) Render(w http.ResponseWriter, name string, data any) error {
	t, ok := tm.templates[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return nil
	}
	return t.Execute(w, data)
}
