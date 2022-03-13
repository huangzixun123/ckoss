package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8088", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	request, _ := http.NewRequest("PUT", "http://localhost:8089/server2/test", r.Body)
	client := http.Client{}
	rr, e := client.Do(request)
	if e == nil && rr.StatusCode != http.StatusOK {
		e = fmt.Errorf("dataServer return http code %d", rr.StatusCode)
	}
}
