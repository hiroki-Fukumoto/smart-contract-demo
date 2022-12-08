package main

import (
	"time"

	"github.com/hiroki-Fukumoto/smart-contract-demo/route"
)

const location = "Asia/Tokyo"

// @title Geth DAPP DEMO
// @version 1.0
// @description  Geth DAPP DEMO
// @host localhost:8080
func main() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	r := route.SetupRouter()

	r.Run(":8080")
}
