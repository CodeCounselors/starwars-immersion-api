package swapico

import ()
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"time"
)

const (
	URL = "http://swapi.co/api"
)

type People struct {
	Count      int `json:"count"`
	Next       string `json:"next"`
	Previous   interface{} `json:"previous"`
	Characters []Character `json:"results"`
}

type Character struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
	BirthYear string `json:"birth_year"`
	Gender    string `json:"gender"`
	Homeworld string `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string `json:"url"`
}

func AllCharacters() *[]Character {

	url := URL + "/people/"
	var r, everyone = loadPeoplePage([]Character{}, url)

	//TODO: MUX these all at the same time because we know the total #
	for len(r.Next) > 0  {
		r, everyone = loadPeoplePage(everyone, r.Next)
	}
	return &everyone;
}


func loadPeoplePage(appendTo []Character, url string) (*People, []Character) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	log.Println("Making request for: ", url)
	peeps, err := client.Do(req)
	if ( err != nil ) {
		log.Println(err)
	}
	defer peeps.Body.Close()
	body, _ := ioutil.ReadAll(peeps.Body)
	r := People{}
	json.Unmarshal(body, &r)
	allPeople := append(appendTo, r.Characters...)
	return &r, allPeople
}

