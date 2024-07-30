package lib

import (
	"os"

	polygon "github.com/polygon-io/client-go/rest"
)

var PolygonClient polygon.Client

func IntiPolygon() {
	polygonApiKey := os.Getenv("POLYGON_API_KEY")
	client := polygon.New(polygonApiKey)

	PolygonClient = *client
}
