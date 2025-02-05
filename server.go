package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//Handle root / default route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About"))
	})

	r.HandleFunc("/search", SearchHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("Server Ready")
	http.ListenAndServe(":8989", nil)

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// query := r.URL.Query().Get("q")
	// a := r.URL.Query().Get("a")
	// b := r.URL.Query().Get("b")
	vars := r.URL.Query()
	query := vars.Get("q")
	strA := vars.Get("a")
	strB := vars.Get("b")

	a, erra := strconv.Atoi(strA)
	if erra != nil {
		// panic(erra)
		http.Error(w, "Parameter A harus berupa bilagan", http.StatusBadRequest)
		return
	}
	b, errb := strconv.Atoi(strB)
	if errb != nil {
		// panic(errb)
		http.Error(w, "Parameter B harus berupa bilagan", http.StatusBadRequest)
		return
	}

	c := a + b
	responseMessage := fmt.Sprintf("Hasil Pencarian untuk : %s \nPenjumlahan : %d + %d = %s", query, a, b, c)
	w.Write([]byte(responseMessage))
}
