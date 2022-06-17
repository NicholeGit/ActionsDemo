package main

import "log"

var (
	appVersion   string
	appRevision  string
	appBuildDate string
)

func main() {
	log.Println("start ...")
	log.Println("appVersion:", appVersion)
	log.Println("appRevision:", appRevision)
	log.Println("appBuildDate:", appBuildDate)
}
