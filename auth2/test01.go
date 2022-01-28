package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	baseURL = "https://gitlab.com"
)

var (
	username = "bobmclarke@gmail.com"
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

	fmt.Printf("url: %+v\n", loginURL)
	fmt.Printf("resp: %+v\n", response)

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

func (app *App) login() {
	client := app.Client

	authenticityToken := app.getToken()
	fmt.Printf("authenticityToken: %s\n", authenticityToken)

	loginURL := baseURL + "/users/sign_in"
	fmt.Printf("baseURL: %s\n", baseURL)

	data := url.Values{
		"authenticity_token": {authenticityToken.Token},
		"user[login]":        {username},
		"user[password]":     {password},
	}
	fmt.Printf("data: %s\n", data)

	response, err := client.PostForm(loginURL, data)
	fmt.Printf("responseBody: %s\n", response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
}

func (app *App) getProjects() []Project {
	projectsURL := baseURL + "/dashboard/projects"
	client := app.Client

	response, err := client.Get(projectsURL)

	if err != nil {
		log.Fatalln("Error fetching response. ", err)
	}

	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	var projects []Project

	document.Find(".project-name").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		project := Project{
			Name: name,
		}

		projects = append(projects, project)
	})

	return projects
}

func main() {
	jar, _ := cookiejar.New(nil)

	app := App{
		Client: &http.Client{Jar: jar},
	}

	app.getToken()

	//app.login()

	//projects := app.getProjects()

	//for index, project := range projects {
	//	fmt.Printf("%d: %s\n", index+1, project.Name)
	//}
}
