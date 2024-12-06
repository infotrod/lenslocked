package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Course struct {
	LanguageName string
	Price        float64
}
type User struct {
	Name             string
	Email            string
	SubscriptionEnds string
	Courses          []Course
	IsPrimaryMember  bool
}

func executeTemplate(w http.ResponseWriter, filepath string, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath, nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath, nil)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath, nil)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "productID")
	w.Write([]byte(fmt.Sprintf("<p>Product id = %s</p>", productId)))
	// fmt.Fprint(w, fmt.Sprintf("<h1>Product id = %s</h1>", productId))
}

func userHadler(w http.ResponseWriter, r *http.Request) {
	courses := []Course{
		{LanguageName: "Go", Price: 29.99},
		{LanguageName: "Python", Price: 19.99},
		{LanguageName: "JavaScript", Price: 24.99},
	}

	user := User{
		Name:             "Rustem Sharipov",
		Email:            "sharipov.rustem@gmail.com",
		SubscriptionEnds: "03/01/2025",
		Courses:          courses,
		IsPrimaryMember:  true,
	}

	tplPath := filepath.Join("templates", "user.gohtml")
	executeTemplate(w, tplPath, user)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	r.Get("/product/{productID}", productHandler)
	r.Get("/user", userHadler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
