package Connectors

import (
	"testing"
)

func TestDetail(t *testing.T) {
	placeid := "ChIJKzfU3sSrQjQRCzAF0PADjQE"

	t.Logf("start testing get place detail for id: %v", placeid)
	t.Log(string(GetPlaceDetail(placeid)[:]))

}
func TestNearbySearch(t *testing.T) {
	geoX := "25.043302"
	geoY := "121.5495064"
	radius := "100"
	placetype := "restaurant"

	t.Log("start testing nearby search")
	t.Log(string(GetNearby(geoX, geoY, radius, placetype)))
}
