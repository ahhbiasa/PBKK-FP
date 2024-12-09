package main

import (
	"log"
	"net/http"
	"pbkk-fp/config"
	"pbkk-fp/controllers/categorycontroller"
	"pbkk-fp/controllers/homecontroller"
	"pbkk-fp/controllers/productcontroller"
	"pbkk-fp/controllers/shopcontroller"
)

func main() {
	config.ConnectDB()
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	// 1.Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// 3. Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	// 4. Shops
	http.HandleFunc("/shops", shopcontroller.Index)
	http.HandleFunc("/shops/add", shopcontroller.Add)
	http.HandleFunc("/shops/detail", shopcontroller.Detail)
	http.HandleFunc("/shops/edit", shopcontroller.Edit)
	http.HandleFunc("/shops/delete", shopcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
