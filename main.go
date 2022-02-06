// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
  Port = ":8080"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {  
  t, err := template.ParseFiles("index.html")
    if err != nil {
        fmt.Println(err)
    }
    items := struct {
        Country string
        City string
    }{
        Country: "France",
        City: "Paris",
    }
    t.Execute(w, items)
}

func main() {
        log.Print("starting server...")
        http.HandleFunc("/",serveStatic)

        // Determine port for HTTP service.
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("defaulting to port %s", port)
        }

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }
}

