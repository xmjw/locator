package main

import (
  "github.com/xmjw/locator/google"
  osm "github.com/xmjw/locator/openstreetmap"
  "log"
  "flag"
  "time"
)

var (
  latitude  *float64 = flag.Float64("lat", 40.714224, "A Latitude")
  longitude *float64 = flag.Float64("long", -73.961452, "A Longitude")
)

func main() {
  log.Println("Starting. Now what?")
  flag.Parse()

  addressChan := make(chan string)

  // Let the lookup tools do some work...
  go osm.ReverseGeocode(*latitude, *longitude, addressChan)
  go google.ReverseGeocode(*latitude, *longitude, addressChan)
  go func() {
    time.Sleep(30 * time.Second)
    addressChan <- "Timed out. Could not get results in time."
  }()
  
  // take the first response.
  // Ultimately, this could be halting if nothing comes off the channel.
  // Need to add an n second timer to prevent this.
  address := <-addressChan
  
  log.Println(address)
}