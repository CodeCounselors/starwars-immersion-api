package swapico

import (
	//"fmt"
	"log"
	//"os"
	//"net/http"
	//"io/ioutil"
	//"encoding/json"
	"time"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

const (
	SquareV1URL = "https://connect.squareup.com/v1/"
	SquareV2URL = "https://connect.squareup.com/v2/"
)


type Character struct {
	FirstName string
	LastName  string
}

func People() *Character {
	return &Character{"Kylo", "Ren"}
}

func trace(s string) (string, time.Time) {
	//log.Println("START:", s)
	return s, time.Now()
}

func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}