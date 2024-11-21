package main

import (
	"log"
	"net/http"
	"pbkk-fp/config"
	"pbkk-fp/controllers/categorycontroller"
	"pbkk-fp/controllers/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
