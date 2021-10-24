package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	FullName    string    `json:"full_name"`
	HtmlURL     string    `json:"html_url"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetRepositories(user string) []Repository {
	URL := apiURL + "/users/" + user + "/repos"
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var repos []Repository
	json.Unmarshal(body, &repos)

	return repos
}
