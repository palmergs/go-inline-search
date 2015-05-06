package tokensearch

import (
	"fmt"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In search.")
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In insert.")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In delete.")
}

func main() {
	http.HandleFunc("/search/", searchHandler)
	http.HandleFunc("/insert/", insertHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.ListenAndServe(":6060", nil)
}