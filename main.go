package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request)  {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Name - %s\n", name)
	fmt.Fprintf(w, "Age - %s\n", age)
}

func chrząszczHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/chrząszcz" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method not allowed", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Chrząszcz!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/chrząszcz", chrząszczHandler)
	package main

	import(
		"fmt"
		"log"
		"net/http"
	)
	
	func formHandler(w http.ResponseWriter, r *http.Request)  {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "POST request successful")
		name := r.FormValue("name")
		age := r.FormValue("age")
		fmt.Fprintf(w, "Name - %s\n", name)
		fmt.Fprintf(w, "Age - %s\n", age)
	}
	
	func chrząszczHandler(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/chrząszcz" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
		if r.Method != "GET"{
			http.Error(w, "Method not allowed", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "Chrząszcz!")
	}
	
	func main(){
		fileServer := http.FileServer(http.Dir("./static"))
		http.Handle("/", fileServer)
		http.HandleFunc("/form", formHandler)
		http.HandleFunc("/chrząszcz", chrząszczHandler)
	
		fmt.Print("Starting server at port 8000\n")
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}