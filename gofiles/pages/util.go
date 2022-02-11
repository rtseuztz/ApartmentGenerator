package pages

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

func Push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

// Render a template, or server error.
func Render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}
func GetTemplate(filename string) *template.Template {
	viewHTML := GetFileAsHTML(filename)
	viewTpl := template.Must(template.New(filename).Parse(viewHTML))
	return viewTpl
}
func GetFileAsHTML(filename string) string {
	content, fileErr := ioutil.ReadFile(fmt.Sprintf("../templates/%s.html", filename))
	if fileErr != nil {
		fmt.Printf("\nRender Error: %v\n", fileErr)
		return ""
	}
	return string(content)

}
func GetNavigationBarHTML() string {
	navigationBarHTML := GetFileAsHTML("navigation_bar") // THIS WORKS !!!
	return navigationBarHTML
}
