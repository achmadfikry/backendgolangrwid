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
	r.HandleFunc("/", HomeHandler)

	r.HandleFunc("/about", AboutHandler)

	r.HandleFunc("/search", SearchHandler).Methods("GET")
	r.HandleFunc("/login", LoginPageHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/dashboard", DashboardHandler) //tidak pakai .methods karena sudah pasti GET saja untuk sekarang

	http.Handle("/", r)
	fmt.Println("Server Ready")
	http.ListenAndServe(":8989", nil)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

// Menampilkan halaman login menggunakan ServeFile
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/dashboard.html")
}

// Handle form login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	// Ambil username & password dari form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Dummy login check (nanti bisa diganti dengan database)
	// if username == "admin" && password == "password123" {
	// 	fmt.Fprintf(w, "Login berhasil! Selamat datang, %s", username)
	// } else {
	// 	http.Error(w, "Username atau password salah", http.StatusUnauthorized)
	// }
	if username == "admin" && password == "password123" {
		http.Redirect(w, r, "/dashboard", http.StatusFound) //redirect url web
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Hello, World!"))
	http.ServeFile(w, r, "static/index.html")
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
