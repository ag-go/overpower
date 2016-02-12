package json

import (
	"mule/hexagon"
	"mule/overpower"
)

func LoadShipView(item overpower.ShipView) *ShipView {
	return &ShipView{
		Gid:        item.Gid(),
		Fid:        item.Fid(),
		Sid:        item.Sid(),
		Turn:       item.Turn(),
		Controller: item.Controller(),
		Size:       item.Size(),
		Loc:        item.Loc(),
		Dest:       item.Dest(),
		Trail:      item.Trail(),
	}
}

func LoadShipViews(list []overpower.ShipView) []*ShipView {
	jList := make([]*ShipView, len(list))
	for i, item := range list {
		jList[i] = LoadShipView(item)
	}
	return jList
}

type ShipView struct {
	Gid        int               `json:"gid"`
	Fid        int               `json:"fid"`
	Turn       int               `json:"turn"`
	Sid        int               `json:"sid"`
	Controller int               `json:"controller"`
	Size       int               `json:"size"`
	Loc        hexagon.NullCoord `json:"loc"`
	Dest       hexagon.NullCoord `json:"dest"`
	Trail      hexagon.CoordList `json:"trail"`
}
