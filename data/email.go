package data

import (
    "gopkg.in/dancannon/gorethink.v2"
)

type Email string

func (ctx *Context) PutEmail(email Email) error {
    _, err := gorethink.DB("primary").Table("emails").Insert(map[string]Email{"email": email}, gorethink.InsertOpts{Conflict: "update"}).RunWrite(ctx.Rethink);
    return err;
}

