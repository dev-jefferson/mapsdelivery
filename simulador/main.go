package main

import (
	"fmt"

	routeMap "github.com/dev-jefferson/mapsdelivery/app/route"
)

func main() {

	route := routeMap.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0])
}
