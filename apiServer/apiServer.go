package main

import (
	"log"
	"net/http"
	"os"

	"ckoss/apiServer/heartbeat"
	"ckoss/apiServer/locate"
	"ckoss/apiServer/objects"
	"ckoss/apiServer/versions"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
