package main

import (
	"golang-website-dev/handler"
	"log"
	"net/http"
)


func main() {
		mux := http.NewServeMux()

		// certificate := func(w http.ResponseWriter, r *http.Request) {
		// 	w.Write([]byte("Certificare"))
		// }

		mux.HandleFunc("/", handler.DashboardHandler) 
		mux.HandleFunc("/about", handler.AboutHandler) 
		mux.HandleFunc("/project", handler.ProjectHandler) 
		mux.HandleFunc("/product", handler.ProductHandler) 
		mux.HandleFunc("/post-get", handler.PostGet) 
		mux.HandleFunc("/form", handler.FormData) 
		mux.HandleFunc("/proccess", handler.Proccess) 



		// mux.HandleFunc("/certificate", certificate) 
		// mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
		// 	w.Write([]byte("Contact Me"))
		// })

		fileServer := http.FileServer(http.Dir("assets"))
		mux.Handle("/static/", http.StripPrefix("/static", fileServer))


		println("Port 3000 Sedang Berlajan ")
		err := http.ListenAndServe(":3000", mux)
		log.Fatal(err)
}


// %d: Memformat sebagai bilangan bulat desimal (integer).
// %f: Memformat sebagai angka floating-point (desimal).
// %s: Memformat sebagai string.
// %t: Memformat sebagai nilai boolean.
// %c: Memformat sebagai karakter (rune).
// %x: Memformat sebagai bilangan heksadesimal.
// %o: Memformat sebagai bilangan oktal.
// %v: Memformat dengan format default sesuai dengan tipe data.