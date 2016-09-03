package main

import (
    "github.com/itshappyhoursomewhere/backend/data"
)

type GetRequest struct {
    Lat float64 `json:"lat" xml:"Lat"`
    Long float64 `json:"long" xml:"Long"`
}

type GetResponse struct {
    Locations []data.Location `json:"locations" xml:"Locations"`
}

type PutRequest struct {
    Locations []data.Location `json:"locations" xml:"Locations"`
}

type PutResponse struct {}