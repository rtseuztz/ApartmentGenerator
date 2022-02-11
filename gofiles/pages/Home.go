package pages

import (
	"html/template"
	"net/http"
)

var homepageTpl *template.Template

func init() {
	homepageTpl = GetTemplate("index")
}

// HomeHandler renders the homepage view template
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	Push(w, "/static/home.css")
	Push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(GetNavigationBarHTML()),
	}
	// x := homepageTpl
	// template := GetTemplate("index")
	Render(w, r, homepageTpl, "index", fullData)
}
