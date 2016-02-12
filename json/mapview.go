package json

import (
	"mule/hexagon"
	"mule/overpower"
)

func LoadMapView(item overpower.MapView) *MapView {
	return &MapView{
		Gid:    item.Gid(),
		Fid:    item.Fid(),
		Center: item.Center(),
	}
}

func LoadMapViews(list []overpower.MapView) []*MapView {
	jList := make([]*MapView, len(list))
	for i, item := range list {
		jList[i] = LoadMapView(item)
	}
	return jList
}

type MapView struct {
	Gid    int           `json:"gid"`
	Fid    int           `json:"fid"`
	Center hexagon.Coord `json:"center"`
}
