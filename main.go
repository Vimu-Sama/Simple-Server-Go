package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Following Error ocurred while parsing-> %v", err)
	}
	name := r.FormValue("UserName")
	address := r.FormValue("Email")
	fmt.Fprintf(w, "The values were posted successfully")
	fmt.Printf("Name-> %v\n", name)
	fmt.Printf("Email address-> %v", address)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Bad Request Made", http.StatusBadRequest)
		return
	}
	fmt.Printf("Hello Page Accessed\n")
	fmt.Fprintf(w, "This is the Home Page of the local browser created")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file!")
	}
	portNum := os.Getenv("PORT")
	http.Handle("/", fs)
	http.HandleFunc("/hello", homeHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Printf("The server is running...")
	if err := http.ListenAndServe(":"+portNum, nil); err != nil {
		fmt.Printf("There was error in showing the page")
		log.Fatal(err)
	}
}
