package internal

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateManager struct {
	TemplateCache map[string]*template.Template
}

func (t *TemplateManager) Render(w http.ResponseWriter, tmplName string, data interface{}) error {
	tmpl, exists := t.TemplateCache[tmplName]
	if !exists {
		return fmt.Errorf("Template %s not found", tmplName)
	}
	return tmpl.Execute(w, data)
}

func NewTemplateManager(pattern string) (*TemplateManager, error) {
	templates := make(map[string]*template.Template)

	tmplFiles, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, tmplFile := range tmplFiles {
		name := filepath.Base(tmplFile)
		tmpl, err := template.ParseFiles(tmplFile)
		if err != nil {
			return nil, err
		}
		templates[name] = tmpl
	}

	return &TemplateManager{TemplateCache: templates}, nil
}
