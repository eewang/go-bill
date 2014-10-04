package main

import (
  "net/http"
  "io"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html;charset=utf-8")
  io.WriteString(w, "hello world")
}