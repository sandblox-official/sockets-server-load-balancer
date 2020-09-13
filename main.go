package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	nodes         = 10
	nextNode      = 1
	nodesLocation = ".example.com"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/node", func(w http.ResponseWriter, r *http.Request) {
		if nextNode == nodes+1 {
			nextNode = 1
		}
		nextNodeAsByte := []byte(strconv.Itoa(nextNode))
		w.Write([]byte("node" + string(nextNodeAsByte) + string(nodesLocation)))
		nextNode++

	}).Methods(http.MethodGet)
	http.ListenAndServe(":8088", router)
}
