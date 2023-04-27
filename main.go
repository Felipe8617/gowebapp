package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Users struct {
	Name  string
	Email string
	Phone int
}

func Index(rw http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("templates/index.html")

	user := Users{"Felipe Carvajal", "felipe@mail.com", 123456789}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, user)
	}
}

func main() {

	http.HandleFunc("/", Index)

	//creating the server
	fmt.Println("server is runing on port 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe(":3000", nil)

}
