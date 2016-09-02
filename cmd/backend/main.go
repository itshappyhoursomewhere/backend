package main

import (
    "net/http"
    "github.com/ThatsMrTalbot/scaffold"
    "github.com/ThatsMrTalbot/scaffold/encoding"
)

func getData(req GetRequest) (GetResponse, error) {
    return GetResponse{}, nil
}

func external() {
    dispatcher := scaffold.DefaultDispatcher()
    router := scaffold.New(dispatcher)
    router.AddHandlerBuilder(encoding.DefaultHandlerBuilder)

    router.Get("/data.json", getData)
    http.ListenAndServe(":80", dispatcher)
}

func internal() {
    dispatcher := scaffold.DefaultDispatcher()
    router := scaffold.New(dispatcher)
    router.AddHandlerBuilder(encoding.DefaultHandlerBuilder)

    http.ListenAndServe(":6080", dispatcher)
}

func main() {
    go external();
    internal();
}