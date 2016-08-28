package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"./swapico"
)

var chars = []swapico.Character{}

func main() {
	http.HandleFunc("/members", membersHandler)
	http.ListenAndServe(":8080", nil)
}

func membersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var char = swapico.People()
		j, _ := json.Marshal(char)
		w.Write(j)
	}

	if r.Method == "POST" {
		var c swapico.Character
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &c)
		chars = append(chars, c)
		j, _ := json.Marshal(c)
		w.Write(j)
	}
}