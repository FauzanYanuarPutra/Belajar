package handler

import (
	"fmt"
	"golang-website-dev/entity"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title" : "learning golang",
	// 	"content" : "learning golang with Muhammad Fauzan",
	// }

	// data := entity.Product{ID: 1, Name: "Muhammad Fauzan", Price: 2000000, Stock: 3}

	data := []entity.Product{
		{ID: 1, Name: "Muhammad Fauzan", Price: 2000000, Stock: 20},
		{ID: 1, Name: "Yanuar Putra", Price: 4343242, Stock: 2},
		{ID: 1, Name: "Putra Fauzan", Price: 2323242, Stock: 8},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}



	// w.Write([]byte("Hello World"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {

	// }
	w.Write([]byte("About Me"))
}

func ProjectHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {

	// }
	w.Write([]byte("Project"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {

	// }
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Perodyct page"))

	// fmt.Fprintf(w, "Product Page : %d", idNumb)

	data := map[string]interface{} {
			"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}


func PostGet(w http.ResponseWriter, r *http.Request) {
	request := r.Method

	switch request{
	case "GET":
		w.Write([]byte("Ini Get"))
	case "POST":
		w.Write([]byte("Ini Post"))
	default:
		http.Error(w, "Error", http.StatusBadRequest)
	}
}

func FormData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		return
	}
	http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)

}

func Proccess(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
				return
			}

			

			tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
				return
			}

			nama := r.Form.Get("name")
			message := r.Form.Get("message")

			data := map[string]interface{}{
				"nama": nama,
				"message": message,
			}
			tmpl.Execute(w, data)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
				return
			}

			// w.Write([]byte(nama))
			
		}
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)

}