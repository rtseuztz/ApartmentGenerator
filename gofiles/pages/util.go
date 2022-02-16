package pages

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 5 * time.Second}
var navigationBarHTML string

func init() {
	navigationBarHTML = GetFileAsHTML("navigation_bar") // THIS WORKS !!!
}
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
func setCSS(w http.ResponseWriter) {
	Push(w, "/static/css/home.css")
	Push(w, "/static/css/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	return navigationBarHTML
}
func GET(url string) *http.Response {
	if len(url) <= 0 {
		return nil
	}
	r, err := myClient.Get(url)
	if err != nil {
		log.Printf("GET Failed")
		return nil
	}
	if r.StatusCode == 429 {
		time.Sleep(time.Second)
		r, err = myClient.Get(url)
		if err != nil {
			return nil
		}
	}
	if r.StatusCode != 200 {
		log.Printf("GET Failed")
		return nil
	}
	//defer r.Body.Close()
	return r
}

/*
Morphs the json into the given object
*/
func getJson(url string, target interface{}) error {
	resp := GET(url)
	if resp == nil {
		return errors.New("failed riot get")
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
