package main

import (
	"html/template"

	"github.com/anxious-aurelius/snippetbox/internal/models"
)

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}

// newTemplateCache builds an in-memory cache of parsed templates, keyed by
// page name, so templates are parsed once at startup rather than on every request.
func newTemplateCache() (map[string]*template.Template, error){

	cache := map[string]*template.Template{}

	// TODO: use filepath.Glob() to find all the page templates, e.g.
	// "./ui/html/pages/*.tmpl"

	// TODO: loop over the page templates. For each one:
	//   1. get the file name (filepath.Base) to use as the cache key
	//   2. build the slice of files to parse: base.tmpl, partials/nav.tmpl,
	//      and the page itself
	//   3. use template.ParseFiles(...) to parse them into a *template.Template
	//   4. store the result in the cache map, keyed by the file name

	return cache, nil

}