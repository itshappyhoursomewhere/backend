package data 

import (
    "gopkg.in/dancannon/gorethink.v2"
)

type Location struct {
    Name string `gorethink:"name"`
    Location gorethink.Term `gorethink:"location"`
    TTL int64 `gorethink:"ttl"`
}

func GetLocations(lat float64, long float64, rad float64) ([]Location, error) {
    return nil, nil
}

func PushLocations([]Location) error {
    return nil    
}

