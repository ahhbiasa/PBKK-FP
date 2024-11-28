package productcontroller

import (
	"html/template"
	"net/http"
	"pbkk-fp/entities"
	"pbkk-fp/models/categorymodel"
	"pbkk-fp/models/productmodel"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.html")
		if err != nil {
			panic(err)
		}

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		product.Name = r.FormValue("name")

		// Validate and parse category_id
		categoryIDStr := r.FormValue("category_id")
		if categoryIDStr == "" {
			http.Error(w, "Category ID is required", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}
		product.Category.Id = categoryID

		// Validate and parse stock
		stockStr := r.FormValue("stock")
		// fmt.Println(stockStr)
		if stockStr == "" {
			http.Error(w, "Stock is required", http.StatusBadRequest)
			return
		}
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}
		product.Stock = stock

		product.Description = r.FormValue("description") // Fixed typo in "descripton"
		product.Created_At = time.Now()
		product.Updated_At = time.Now()

		// Call the model's Create function
		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}

}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
			"product":    product,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")

		// Validate and parse category_id
		categoryIDStr := r.FormValue("category_id")
		if categoryIDStr == "" {
			http.Error(w, "Category ID is required", http.StatusBadRequest)
			return
		}
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid Category ID", http.StatusBadRequest)
			return
		}
		product.Category.Id = categoryID

		// Validate and parse stock
		stockStr := r.FormValue("stock")
		// fmt.Println(stockStr)
		if stockStr == "" {
			http.Error(w, "Stock is required", http.StatusBadRequest)
			return
		}
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			http.Error(w, "Invalid stock value", http.StatusBadRequest)
			return
		}
		product.Stock = stock

		product.Description = r.FormValue("description") // Fixed typo in "descripton"
		product.Updated_At = time.Now()

		// Call the model's Create function
		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// fmt.Println(idString)
	if err := productmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/products", http.StatusSeeOther)
}
