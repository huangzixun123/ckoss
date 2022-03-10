package main

import (
	"log"
	"net/http"
	"os"

	"ckoss/apiServer/heartbeat"
	"ckoss/apiServer/locate"
	"ckoss/apiServer/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
