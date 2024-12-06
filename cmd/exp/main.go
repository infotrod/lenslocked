package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Rustem Sharipov",
		Bio:  `<script>alert("Haha, you have been hacked!");</script>`,
		Age:  111,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
