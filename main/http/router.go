package main

import (
  "net/http"

  "bitbucket.org/server-monitor/checkrlight/app/handlers/checker"
)

type routeType struct {
  Path string
  Name string
  Handler func(writer http.ResponseWriter, request *http.Request)
}

var routes = []routeType{
  routeType{Path: "/api",  Name: checker.Name, Handler: checker.Handler},
}
