package pages

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var summonerTpl *template.Template
var key string

func init() {
	summonerTpl = GetTemplate("summoner")
	key = "RGAPI-670007bb-22b7-43bf-8281-1be84c67fe57"
}

// HomeHandler renders the homepage view template
func SummonerHandler(w http.ResponseWriter, r *http.Request) {
	setCSS(w)

	pathVariables := mux.Vars(r)

	summoner := GetSummoner(pathVariables["name"])
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(GetNavigationBarHTML()),
		"Name":          summoner.Name,
		"Level":         summoner.Level,
		"ProfileIcon":   summoner.ProfileIcon,
		"ProfileURL":    fmt.Sprintf("/static/images/img/profileicon/%d.png", summoner.ProfileIcon),
	}
	// x := homepageTpl
	// template := GetTemplate("index")
	Render(w, r, summonerTpl, "summoner", fullData)
}

type Summoner struct {
	Name        string `json:"name"`
	Puuid       string `json:"puuid"`
	Level       int    `json:"summonerLevel"`
	ProfileIcon int    `json:"profileIconId"`
}

func GetSummoner(name string) Summoner {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", name, key)
	summoner := new(Summoner)
	getJson(url, summoner)
	//summoner has the data
	return *summoner
}
