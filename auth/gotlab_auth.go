package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://gitlab.com"
)

var (
	username = "bobclarke"
	password = "808state"
)

type App struct {
	Client *http.Client
}

type AuthenticityToken struct {
	Token string
}

type Project struct {
	Name string
}

func (app *App) getToken() AuthenticityToken {
	loginURL := baseURL + "/users/sign_in"
	client := app.Client

	response, err := client.Get(loginURL)

	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	token, _ := document.Find("input[name='authenticity_token']").Attr("value")

	authenticityToken := AuthenticityToken{
		Token: token,
	}

	return authenticityToken
}

func main() {

}
