package shopcontroller

import (
	"net/http"
	"pbkk-fp/entities"
	"pbkk-fp/models/shopmodel"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	shops := shopmodel.GetAll()
	data := map[string]any{
		"shops": shops,
	}

	temp, err := template.ParseFiles("views/shop/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/shop/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var shop entities.Shop

		shop.Name = r.FormValue("name")
		shop.Address = r.FormValue("address")
		shop.CreatedAt = time.Now()
		shop.UpdatedAt = time.Now()

		if ok := shopmodel.Create(shop); !ok {
			temp, _ := template.ParseFiles("views/shop/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/shops", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/shop/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		shop := shopmodel.Detail(id)
		data := map[string]any{
			"shop": shop,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var shop entities.Shop

		ifString := r.FormValue("id")
		id, err := strconv.Atoi(ifString)

		if err != nil {
			panic(err)
		}

		shop.Name = r.FormValue("name")
		shop.Address = r.FormValue("address")
		shop.UpdatedAt = time.Now()

		if ok := shopmodel.Update(id, shop); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		}

		http.Redirect(w, r, "/shops", http.StatusSeeOther)
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	shop := shopmodel.Detail(id)
	data := map[string]any{
		"shop": shop, // Passing the shop with its products
	}

	temp, err := template.ParseFiles("views/shop/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		panic(err)
	}

	if err := shopmodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/shops", http.StatusSeeOther)
}
