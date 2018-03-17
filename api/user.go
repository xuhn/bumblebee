package api

import (
	"optimusprime/controller"
)

type Stuff struct {
	Foo string ` json:"foo" `
	Bar int    ` json:"bar" `
}

type User struct {
	*controller.Controller
}

func (c User) Show(suite string) controller.Result {
	data := make(map[string]interface{})
	data["errormsg"] = nil
	stuff := Stuff{Foo: "xyz", Bar: 999}
	data["stuff"] = stuff
	data["param"] = suite
	return c.RenderJSON(data)
}
