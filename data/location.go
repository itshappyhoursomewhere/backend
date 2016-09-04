package data 

import (
    "gopkg.in/dancannon/gorethink.v2"
)

type lookupResult struct {
    Location Location `gorethink:"doc"`
}


type Active struct {
    Day string `gorethink:"day" json:"day" xml:"Day"`
    StartHour int64 `gorethink:"start" json:"start" xml:"Start"`
    EndHour int64 `gorethink:"end" json:"end" xml:"End"`
}

type Deal struct {
    Active []Active `gorethink:"active" json:"active" xml:"Active"`
    Description string `gorethink:"description" json:"description" xml:"Description"`
}

type Location struct {
    Name string `gorethink:"name" json:"name" xml:"Name"`
    Location Point `gorethink:"location" json:"geo" xml:"Geo`
    Icon string `gorethink:"icon" json:"icon" xml:"Icon`
    Phone string `gorethink:"phone" json:"phone" xml:"Phone`
    Filters []string `gorethink:"filter" json:"filter" xml:"Filter`
    Description string `gorethink:"description" json:"description" xml:"Description`
    Website string `gorethink:"website" json:"website" xml:"Website`
    Deals []Deal `gorethink:"deals" json:"deals" xml:"Deals`
    TTL int64 `gorethink:"ttl" json:"ttl" xml:"TTL"`
}

type Emails struct {
    Email string `gorethink:"email" json:"email" xml:"Email"`
}

func (ctx *Context) GetLocations(lat float64, long float64, rad float64) ([]Location, error) {
    cursor, err := gorethink.DB("primary").Table("locations").GetNearest(
        gorethink.Point(long, lat), 
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

