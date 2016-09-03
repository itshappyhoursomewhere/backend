package main

import (
    "github.com/itshappyhoursomewhere/backend/data"
)

type GetLocationRequest struct {
    Lat float64 `json:"lat" xml:"Lat"`
    Long float64 `json:"long" xml:"Long"`
}

type GetLocationResponse struct {
    Locations []data.Location `json:"locations" xml:"Locations"`
}

type PutLocationRequest struct {
    Locations []data.Location `json:"locations" xml:"Locations"`
}

type PutEmailRequest struct {
    Email data.Email `json:"email" xml:"Email"`
}

type PutResponse struct {}