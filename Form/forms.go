package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// FormData represents the data structure for the form
type FormData struct {
	Username string
	Password string
}

func handleFormSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusInternalServerError)
			return
		}

		formData := FormData{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		fmt.Printf("Received form data: %+v\n", formData)

		http.Redirect(w, r, "/success", http.StatusSeeOther)
	}
}

func handleFormPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("form.html")
	if err != nil {
		http.Error(w, "Error parsing form template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handleSuccessPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Form submitted successfully!")
}

func main() {
	http.HandleFunc("/submit", handleFormSubmit)
	http.HandleFunc("/", handleFormPage)
	http.HandleFunc("/success", handleSuccessPage)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
