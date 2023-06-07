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

var InfoUser = Users{
	Name:  "Mary",
	Email: "maary@mail.com",
	Phone: 3669977,
}

func Index(rw http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles("templates/index.html")

	user := Users{InfoUser.Name, InfoUser.Email, InfoUser.Phone}

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

func bug() {

	var a int = 10
	var b int = 0
	result := a / b
	fmt.Println(result)
}

func name() {
	greet("John")
	greet("Jane")
}

func greet(name string) {
	fmt.Println("Hello, " + name)
	fmt.Println("How are you, " + name)
	fmt.Println("Nice to meet you, " + name)
}
