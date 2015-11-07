package main

import (
  "github.com/xmjw/locator/google"
  "github.com/xmjw/locator/openstreetmap"
  "log"
)

func main() {
  log.Println("Starting. Now what?")
  google := google.GoogleGeocode(0.1, 0.1)
  oms := openstreetmap.OpenStreetMapgGeocode(0.1, 0.1)
  log.Println(google)
  log.Println(oms)
}