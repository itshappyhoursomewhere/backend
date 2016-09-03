package main

import (
    "net/http"
    "github.com/Sirupsen/logrus"
    "github.com/ThatsMrTalbot/scaffold"
    "github.com/ThatsMrTalbot/scaffold/encoding"
    "gopkg.in/dancannon/gorethink.v2"
    "github.com/itshappyhoursomewhere/backend/data"
)

type App struct {
    DataContext *data.Context
}

func NewApp() (*App, error) {
    rethink, err := gorethink.Connect(gorethink.ConnectOpts{
        Address: "database",
    })

    if err != nil {
        return nil, err
    }

    dataCtx := data.NewContext(rethink);
    dataCtx.Initialize()

    return &App {
        DataContext: dataCtx,
    }, nil
}

func (app *App) StartInternal() {
    dispatcher := scaffold.DefaultDispatcher()
    router := scaffold.New(dispatcher)
    router.AddHandlerBuilder(encoding.DefaultHandlerBuilder)

    router.Handle("/data.json", app.pushData)
    http.ListenAndServe(":6080", dispatcher)
}

func (app *App) StartExternal() {
    dispatcher := scaffold.DefaultDispatcher()
    router := scaffold.New(dispatcher)
    router.AddHandlerBuilder(encoding.DefaultHandlerBuilder)

    router.Handle("/data.json", app.getData)
    http.ListenAndServe(":80", dispatcher)
}

func (app *App) getData(req GetRequest) (GetResponse, error) {
    locations, err := app.DataContext.GetLocations(req.Lat, req.Long, 2000);
    return GetResponse{Locations: locations}, err
}

func (app *App) pushData(req PutRequest) (PutResponse, error) {
    err := app.DataContext.PushLocations(req.Locations...);
    return PutResponse{}, err
}

func (app *App) putEmail(req PutRequest) (PutResponse, error) {
    err := app.DataContext.PutEmail(req.Emails...);
    return PutResponse{}, err
}

func main() {
    app, err := NewApp()
    if err != nil {
        logrus.WithError(err).Fatal("Error initializing application");
    }

    go app.StartExternal()
    app.StartInternal()
}