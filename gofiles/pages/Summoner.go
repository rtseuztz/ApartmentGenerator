package pages

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var summonerTpl *template.Template
var key string

func init() {
	summonerTpl = GetTemplate("summoner")
	key = "RGAPI-2af5b0aa-b5c5-4f4b-a548-6de8cadf84d1"
}

// HomeHandler renders the homepage view template
func SummonerHandler(w http.ResponseWriter, r *http.Request) {
	setCSS(w)

	pathVariables := mux.Vars(r)

	summoner := GetSummoner(pathVariables["name"])
	games := GetGames(summoner.Puuid)
	fmt.Printf("games: %v\n", games)

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
type Game struct {
	MetaData MetaData `json:"metadata"`
	Info     Info     `json:"info"`
}
type MetaData struct {
	MatchID      string   `json:"matchId"`
	Participants []string `json:"participants"`
}
type Info struct {
	GameDuration int `json:"gameDuration"`
}
type GameList struct {
	games string
}

func GetSummoner(name string) Summoner {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", name, key)
	summoner := new(Summoner)
	getJson(url, summoner)
	return *summoner
}
func GetGames(puuid string) []Game {
	url := fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?start=0&count=19&api_key=%s", puuid, key)
	var gameIDArr []string
	err := getJsonArr(url, &gameIDArr)
	if err != nil {
		log.Printf("Error getting game list")
		return nil
	}
	var games []Game
	for _, gameID := range gameIDArr {
		games = append(games, GetGame(gameID))
	}
	return games
}
func GetGame(gameID string) Game {
	url := fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/%s?api_key=%s", gameID, key)
	game := new(Game)
	getJson(url, game)
	return *game
}
