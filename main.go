package main

import (
	"asynclawproject/websocket"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	fmt.Println(id)
	fmt.Printf("Type of id is %T\n", id)

	var temp = template.Must(template.ParseFiles("index.html"))
	temp.Execute(w, nil)

	statsHandler := func(w http.ResponseWriter, r *http.Request) {
		ws, err := websocket.Upgrade(w, r)
		if err != nil {
			fmt.Println(w, "%+v\n", err)
		}
		websocket.Writer(ws, id)
	}

	http.HandleFunc("/stats", statsHandler)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	setupRoutes()
}
