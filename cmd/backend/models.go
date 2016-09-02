package main

type LocationResponse struct {
    Name string `json:"name" xml:"Name"`
    Lat float64 `json:"lat" xml:"Lat"`
    Long float64 `json:"long" xml:"Long"`
}

type GetRequest struct {
    Lat float64 `json:"lat" xml:"Lat"`
    Long float64 `json:"long" xml:"Long"`
}

type GetResponse struct {
    Locations []LocationResponse `json:"locations" xml:"Locations"`
}