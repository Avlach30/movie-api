package taskscheduler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"movie-api/helper"
	"net/http"
)

type Location struct {
	Name      string `json:"name"`
	Localtime string `json:"localtime"`
}

type Current struct {
	Temperature int `json:"temperature"`
}

type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func FetchExternalAPI() {
	get := helper.GetEnvValue
	WEATHERSTACK_KEY := get("WEATHERSTACK_KEY")

	cities := []string{"sidoarjo", "malang", "surabaya", "jakarta"}

	min := 0
	max := len(cities)-1
	randIndex := rand.Intn(max - min + 1) + min

	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", WEATHERSTACK_KEY, cities[randIndex])

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	log.Printf("Current temperature at %s: %d celcius", responseObject.Location.Name, responseObject.Current.Temperature)
}

func NewSchedule() {
	newSchedule := ConfigTaskScheduler()

	newSchedule.Every(1).Minutes().Do(FetchExternalAPI)

	newSchedule.StartAsync()
}
