package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "strings"
  "strconv"
)

func main() {
  for _, route := range routes {
    http.HandleFunc(route.Path, route.Handler)
  }

  public := http.FileServer(http.Dir("public"))
  http.Handle("/", public)

  port := 8080
  envPortStr := strings.TrimSpace(os.Getenv("PORT"))
  if envPortStr != "" {
    envPort, err := strconv.Atoi(envPortStr)
    if err != nil {
      log.Fatalf(
        "... str to int conversion failed, str: %#v, result: %#v, err: #%v",
        envPortStr, envPort, err,
      )
    }
    port = envPort
  }

  log.Printf("... listening on port %d", port)

  http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
