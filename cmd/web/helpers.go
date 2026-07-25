package main

import (
	"net/http"
	"runtime/debug"
)

// renderTemplate looks up the template set for the given page, writes the
// HTTP status code, then executes it to the response.
func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, status int, page string, data templateData){

	// TODO: look up "page" in app.templateCache. If it isn't found, build
	// an error with fmt.Errorf() and pass it to app.serverError(), then return.

	// TODO: write the status code to the response with w.WriteHeader(status).

	// TODO: execute the "base" template from the template set into w. If
	// ExecuteTemplate() returns an error, pass it to app.serverError().
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {

	method := r.Method
	uri := r.URL
	trace := string(debug.Stack())

	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int) {

	http.Error(w, http.StatusText(status), status)

}
