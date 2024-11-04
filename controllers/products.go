package controllers

import (
	"go-web-application/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func NewProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "NewProduct", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error:", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error:", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error:", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error:", err)
		}

		convertedAmount, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error:", err)
		}

		models.UpdateProduct(convertedId, name, description, convertedPrice, convertedAmount)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
