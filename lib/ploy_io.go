package lib

import (
	polygon "github.com/polygon-io/client-go/rest"
	"os"
)

type PolyIo struct {
	Client polygon.Client
}

var PolyIoClient PolyIo

func IntiPolygon() {
	polygonApiKey := os.Getenv("POLYGON_API_KEY")
	client := polygon.New(polygonApiKey)

	PolyIoClient = PolyIo{Client: *client}
}
