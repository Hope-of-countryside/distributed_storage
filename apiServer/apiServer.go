package main

import (
	"distributed_storage/apiServer/heartbeat"
	"distributed_storage/apiServer/locate"
	"distributed_storage/apiServer/objects"
	"distributed_storage/apiServer/temp"
	"distributed_storage/apiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
