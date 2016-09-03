package data

import (
    "gopkg.in/dancannon/gorethink.v2/types"
)

type Point struct {
	Lon float64 `json:"long" xml:"Long"`
	Lat float64 `json:"lat" xml:"Lat"`
}
type Line []Point
type Lines []Line

func (p Point) Coords() interface{} {
	return []interface{}{p.Lon, p.Lat}
}

func (p Point) MarshalRQL() (interface{}, error) {
	return map[string]interface{}{
		"$reql_type$": "GEOMETRY",
		"coordinates": p.Coords(),
		"type":        "Point",
	}, nil
}

func (p *Point) UnmarshalRQL(data interface{}) error {
	g := &types.Geometry{}
	err := g.UnmarshalRQL(data)
	if err != nil {
		return err
	}
	if g.Type != "Point" {
		return fmt.Errorf("pseudo-type GEOMETRY object has type %s, expected type %s", g.Type, "Point")
	}

	p.Lat = g.Point.Lat
	p.Lon = g.Point.Lon

	return nil
}