package data 

import (
    "gopkg.in/dancannon/gorethink.v2"
    "gopkg.in/dancannon/gorethink.v2/types"
)

type lookupResult struct {
    Location Location `gorethink:"doc"`
}

type Location struct {
    Name string `gorethink:"name" json:"name" xml:"Name"`
    Location types.Point `gorethink:"location" json:"location" xml:"Location`
    TTL int64 `gorethink:"ttl" json:"ttl" xml:"TTL"`
}

func (ctx *Context) GetLocations(lat float64, long float64, rad float64) ([]Location, error) {
    cursor, err := gorethink.DB("primary").Table("locations").GetNearest(
        gorethink.Point(lat, long), 
        gorethink.GetNearestOpts{
            Index: "location", 
            MaxResults: 1000, 
            MaxDist: 2000,  
        },
    ).Run(ctx.Rethink)

    if err != nil {
        return nil, err
    }

    defer cursor.Close()

    locations := new([]lookupResult)
    err = cursor.All(locations)
    if err != nil {
        return nil, err
    }

    docs := make([]Location, len(*locations))

    for i, v := range *locations {
        docs[i] = v.Location;
    }

    return docs, nil
}

func (ctx *Context) PushLocations(locs ...Location) error {
    _, err := gorethink.DB("primary").Table("locations").Insert(&locs, gorethink.InsertOpts{Conflict: "update"}).RunWrite(ctx.Rethink);
    return err;
}

