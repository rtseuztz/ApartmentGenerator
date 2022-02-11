package pages

import (
	"html/template"
	"net/http"
)

// HomeHandler renders the homepage view template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Push(w, "/static/style.css")
	Push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(GetNavigationBarHTML()),
	}
	template := GetTemplate("index")
	Render(w, r, template, "index", fullData)
}
