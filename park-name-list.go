package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//var endpoint = "https://developer.nps.gov/api/v1/parks?limit=600"
//var ep = "https://developer.nps.gov/api/v1/parks?limit=10&start=11"
var ep = "https://developer.nps.gov/api/v1/parks?limit=10"
var key = "API-KEY-HERE"

func main() {
	p, _ := ParkList()
	//if err != nil {
	//	panic(err)
	//}

	fmt.Printf("%s", p)
}

func ParkList() ([]ParkID, error) {
	payloadUrl := fmt.Sprintf(ep)
	req, err := http.NewRequest(http.MethodGet, payloadUrl, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Api-Key", key)

	mc := http.Client{Timeout: 100 * time.Minute}
	resp, err := mc.Do(req)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("response is nil")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pr := &ParkResponse{}
	err = json.Unmarshal(body, &pr)
	var pids []ParkID
	if pr != nil {
		for _, p := range pr.Data {
			pid := &ParkID{
				Name: p.Name,
				Code: p.ParkCode,
			}
			pids = append(pids, *pid)
		}
	}
	return pids, err
}

type ParkID struct {
	Name string
	Code string
}

type ParkResponse struct {
	Total string `json:"total"`
	Data  []struct {
		Contacts struct {
			PhoneNumbers []struct {
				PhoneNumber string `json:"phoneNumber"`
				Description string `json:"description"`
				Extension   string `json:"extension"`
				Type        string `json:"type"`
			} `json:"phoneNumbers"`
			EmailAddresses []struct {
				Description  string `json:"description"`
				EmailAddress string `json:"emailAddress"`
			} `json:"emailAddresses"`
		} `json:"contacts"`
		States       string `json:"states"`
		Longitude    string `json:"longitude"`
		EntranceFees []struct {
			Cost        string `json:"cost"`
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"entranceFees"`
		DirectionsInfo string `json:"directionsInfo"`
		EntrancePasses []struct {
			Cost        string `json:"cost"`
			Description string `json:"description"`
			Title       string `json:"title"`
		} `json:"entrancePasses"`
		DirectionsURL  string `json:"directionsUrl"`
		URL            string `json:"url"`
		WeatherInfo    string `json:"weatherInfo"`
		Name           string `json:"name"`
		OperatingHours []struct {
			Exceptions    []interface{} `json:"exceptions"`
			Description   string        `json:"description"`
			StandardHours struct {
				Wednesday string `json:"wednesday"`
				Monday    string `json:"monday"`
				Thursday  string `json:"thursday"`
				Sunday    string `json:"sunday"`
				Tuesday   string `json:"tuesday"`
				Friday    string `json:"friday"`
				Saturday  string `json:"saturday"`
			} `json:"standardHours"`
			Name string `json:"name"`
		} `json:"operatingHours"`
		LatLong     string `json:"latLong"`
		Description string `json:"description"`
		Images      []struct {
			Credit  string `json:"credit"`
			AltText string `json:"altText"`
			Title   string `json:"title"`
			ID      string `json:"id"`
			Caption string `json:"caption"`
			URL     string `json:"url"`
		} `json:"images"`
		Designation string `json:"designation"`
		ParkCode    string `json:"parkCode"`
		Addresses   []struct {
			PostalCode string `json:"postalCode"`
			City       string `json:"city"`
			StateCode  string `json:"stateCode"`
			Line1      string `json:"line1"`
			Type       string `json:"type"`
			Line3      string `json:"line3"`
			Line2      string `json:"line2"`
		} `json:"addresses"`
		ID       string `json:"id"`
		FullName string `json:"fullName"`
		Latitude string `json:"latitude"`
	} `json:"data"`
	Limit string `json:"limit"`
	Start string `json:"start"`
}
