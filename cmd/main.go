package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/pfuz/goevents/utils"
)

type TemperatureEvent struct {
	ID   int
	Type string
	Data float64
}

type HumidityEvent struct {
	ID   int
	Type string
	Data float64
}

func generateTemperatureEvent(id int) TemperatureEvent {
	eventTypes := []string{"Information", "Warning", "Alert"}
	mintemp, maxtemp := 15.0, 35.0

	return TemperatureEvent{
		ID:   id,
		Type: eventTypes[rand.Intn(len(eventTypes))],
		Data: mintemp + rand.Float64()*(maxtemp-mintemp),
	}
}

func generateHumidityEvent(id int) HumidityEvent {
	eventTypes := []string{"Information", "Warning", "Alert"}
	minhum, maxhum := 30.0, 90.0

	return HumidityEvent{
		ID:   id,
		Type: eventTypes[rand.Intn(len(eventTypes))],
		Data: minhum + rand.Float64()*(maxhum-minhum),
	}
}

func tempStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	eventID := 1

	for {
		event := generateTemperatureEvent(eventID)
		fmt.Fprintf(w, "id: %d\ndata: {\"type\":\"%s\",\"data\":%.2f}\n\n", event.ID, event.Type, event.Data)
		flusher.Flush()
		eventID++
		time.Sleep(time.Second * 5)
	}
}

func humidityStreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	eventID := 1

	for {
		event := generateHumidityEvent(eventID)
		fmt.Fprintf(w, "id: %d\ndata: {\"type\":\"%s\",\"data\":%.2f}\n\n", event.ID, event.Type, event.Data)
		flusher.Flush()
		eventID++
		time.Sleep(time.Second * 5)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	parsedURL := utils.ParseURL(r.URL.Path)
	fmt.Println(parsedURL)
	cwd, _ := os.Getwd()
	staticPath := cwd + "/static/"

	http.ServeFile(w, r, staticPath+"index.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/temp", tempStreamHandler)
	http.HandleFunc("/humidity", humidityStreamHandler)
	fmt.Println("Server is listening on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
