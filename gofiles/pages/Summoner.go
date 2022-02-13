package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var summonerTpl *template.Template

func init() {
	summonerTpl = GetTemplate("summoner")
}

// HomeHandler renders the homepage view template
func SummonerHandler(w http.ResponseWriter, r *http.Request) {

	Push(w, "/static/home.css")
	Push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pathVariables := mux.Vars(r)

	GetSummoner(pathVariables["name"])
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(GetNavigationBarHTML()),
		"Name":          pathVariables["name"],
	}
	// x := homepageTpl
	// template := GetTemplate("index")
	Render(w, r, summonerTpl, "summoner", fullData)
}

type Summoner struct {
	Name        string `json:"name"`
	Puuid       string `json:"puuid"`
	Level       int    `json:"level"`
	ProfileIcon int    `json:"profileIconId"`
}

func GetSummoner(name string) string {
	key := "RGAPI-f402153b-d401-4803-b1bb-11c0ec723270"
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", name, key)
	summoner := new(Summoner)
	getJson(url, summoner)
	//summoner has the data
}
