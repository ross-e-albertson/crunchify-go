package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	myUrl := "localhost:8000"
	fmt.Println("Opening", myUrl)

	http.HandleFunc("/welcome", func(rw http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("views/welcome.tmpl")
		message := "<br><div style='text-align:center;'>" +
			"<h3>********** Hello World, Spring MVC Tutorial</h3>This message is coming from CrunchifyHelloWorld.java **********</div><br><br>"
		tmpl.Execute(rw, template.HTML(message))
	})

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("views/index.tmpl")
		tmpl.Execute(rw, nil)
	})

	if err := http.ListenAndServe(myUrl, nil); err != nil {
		panic(err)
	}

}
