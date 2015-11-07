package openstreetmap

import (
  "net/http"
  "fmt"
  "log"
  json "github.com/bitly/go-simplejson"
)

const OSM_GEOCODE_URI string = "http://nominatim.openstreetmap.org/reverse"

func ReverseGeocode(latitude float64, longitude float64, ch chan string) {
  data, err := getOSMResults(latitude, longitude)

  if err != nil {
    return
  }
  
  result, ok := parseResult(data)
  if ok {
    ch <- fmt.Sprintf("(OSM) %v", result)
  }
}


func getOSMResults(latitude float64, longitude float64) (*json.Json, error) {
  // http://nominatim.openstreetmap.org/reverse?format=json&lat=52.5487429714954&lon=-1.81602098644987&zoom=18&addressdetails=1

  // We can handle all the different HTTP codes later. (if time)
  uri := fmt.Sprintf("%v?lat=%f&lon=%f&format=json", OSM_GEOCODE_URI, latitude, longitude)
  resp, err := http.Get(uri)
  defer resp.Body.Close()

  if err  != nil {
    log.Println("Failed to get a response from OSM.", err)
    return nil,  err
  }

  payload, err := json.NewFromReader(resp.Body)
 
  if err  != nil {
    log.Println("Failed to parse OSM response to JSON", err)
    return nil, err
  }

  return payload, nil
}

func parseResult(data *json.Json) (string, bool) {
  address, ok := data.CheckGet("display_name")
  if ok {
    return address.MustString(), true
  }
  return "", false
}
