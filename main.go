package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Widget struct {
	Color string
}

var ServingWidget = &Widget{Color: "blue"}

func loadWidget() error {
	js, err := ioutil.ReadFile("widget.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	var w Widget
	if err := json.Unmarshal(js, &w); err != nil {
		return err
	}

	ServingWidget = &w
	return nil
}

func serveWidget(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(ServingWidget)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling JSON: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	if err := loadWidget(); err != nil {
		log.Printf("Error loading widget: %s", err)
	}
	s := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(serveWidget),
	}
	log.Fatal(s.ListenAndServe())
}
