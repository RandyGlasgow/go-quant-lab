package lib

import (
	"os"

	polygon "github.com/polygon-io/client-go/rest"
)

type PolyIo struct {
	Client polygon.Client
}

var PolyIoClient PolyIo

var PolygonClient polygon.Client

func IntiPolygon() {
	polygonApiKey := os.Getenv("POLYGON_API_KEY")
	client := polygon.New(polygonApiKey)

	PolyIoClient = PolyIo{Client: *client}
	PolygonClient = *client
}
