package api

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"API",
		"GET",
		"/api/",
		apiHandler,
	},
	Route{
		"Lamp resource",
		"GET",
		"/api/lamp/:lamp/",
		lampHandler,
	},
	Route{
		"Create a Lamp",
		"POST",
		"/api/lamp/",
		createLampHandler,
	},
	Route{
		"Lamp resource action",
		"GET",
		"/api/lamp/:lamp/:action",
		lampActionHandler,
	},
	Route{
		"Global lamp action",
		"GET",
		"/api/lamps/:action",
		allLampActionHandler,
	},
	Route{
		"List all Lamps",
		"GET",
		"/api/lamps",
		lampsHandler,
	},
}
