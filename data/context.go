package data

import (
    "gopkg.in/dancannon/gorethink.v2"
)

type Context struct {
    Rethink *gorethink.Session    
}

func NewContext(rethink *gorethink.Session) *Context {
    return &Context{
        Rethink: rethink,
    }
}

func (ctx *Context) Initialize() {
    gorethink.DBCreate("primary").RunWrite(ctx.Rethink)
    gorethink.DB("primary").TableCreate("locations", gorethink.TableCreateOpts{PrimaryKey: "name"}).RunWrite(ctx.Rethink)
    gorethink.DB("primary").Table("locations").IndexCreate("location", gorethink.IndexCreateOpts{Geo: true}).RunWrite(ctx.Rethink)
}