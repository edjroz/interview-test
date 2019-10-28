package rpc

import (
	"github.com/julienschmidt/httprouter"
)

// The "Route" structure defines the generalization of an api route.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// "Routes" is a function that returns all of the routes of the API.
func Routes() []Route {
	routes := []Route{
		{Name: "BlockByNumber", Method: "POST", Path: "/block/byNumber", HandlerFunc: BlockbyNumber},
	}
	return routes
}
