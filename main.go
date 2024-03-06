package main

import (
  "fmt"
  "html/template"
 // "time"
  "log"
  "net/http"
 // "github.com/gofiber/fiber/v2"
)

func main(){
  fmt.Println("Hello web.")

  h1 := func(w http.ResponseWriter, r *http.Request){
    templ := template.Must(template.ParseFiles("index.html"))
    templ.Execute(w, nil)
  }

 
  http.HandleFunc("/", h1)
  log.Fatal(http.ListenAndServe(":8000", nil))
}

//bloating go text.
