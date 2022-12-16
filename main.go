package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}
}


func formHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/form"{
		http.Error(w,"404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "POST"{
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"ParseForm() err: %v", err)
	}
	fmt.Fprintf(w, "Post request successful\n")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/",fileServer)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/form",formHandler)

	fmt.Println("starting server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}