package google

import (
	"fmt"
	json "github.com/bitly/go-simplejson"
	"log"
	"net/http"
)

const GOOGLE_GEOCODE_URI string = "https://maps.googleapis.com/maps/api/geocode/json"

// ReverseGeocode is the main API function for converting locations into addresses.
// This accepts a latitude and longitude as a float32 and returns the most likely
// address from Google's Geocode service.
func ReverseGeocode(latitude float64, longitude float64, ch chan string) {
	data, err := getGoogleResults(latitude, longitude)

	if err != nil {
		return
	}

	result, ok := parseResult(data)
	if ok {
		ch <- fmt.Sprintf("(Google) %v", result)
	}
}

func getGoogleResults(latitude float64, longitude float64) (*json.Json, error) {
	// Get reults from google, with a request such as:
	//https://maps.googleapis.com/maps/api/geocode/json?latlng=40.714224,-73.961452&key=YOUR_API_KEY
	latlng := fmt.Sprintf("%f,%f", latitude, longitude)
	//apiKey := "SOME_API_KEY"

	// We can handle all the different HTTP codes later. (if time)
	uri := fmt.Sprintf("%v?latlng=%v", GOOGLE_GEOCODE_URI, latlng)
	resp, err := http.Get(uri)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Failed to get a response from Google.", err)
		return nil, err
	}

	payload, err := json.NewFromReader(resp.Body)

	if err != nil {
		log.Println("Failed to parse Google response to JSON", err)
		return nil, err
	}

	return payload, nil
}

func parseResult(data *json.Json) (string, bool) {
	for i, _ := range data.Get("results").MustArray() {
		item := data.Get("results").GetIndex(i)
		address, ok := item.CheckGet("formatted_address")
		if ok {
			return address.MustString(), true
		}
	}
	return "", false
}
