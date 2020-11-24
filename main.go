package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `jsom:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	article := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Test content"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(article)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to TX")
}

func handleRequest() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/articles", allArticles).Methods("POST")
	router.HandleFunc("/addUser/{name}", NewUser).Methods("GET")
	router.HandleFunc("/updateUser/{name}", UpdateUser).Methods("GET")
	router.HandleFunc("/deleteUser/{name}", deleteUser).Methods("GET")
	router.HandleFunc("/getlistUser/{name}", AllUsers).Methods("GET")
	// router.HandleFunc("/newuser", user.NewUser).Methods("POST")
	router.HandleFunc("/connectdb", connectdb).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {
	// ac := accounting.Accounting{Symbol: "$", Precision: 2}
	// fmt.Println(ac.FormatMoney(123456789.213123))

	// c4f.Println("this is red color")

	// var a int = 100000
	// fmt.Println(a)

	// var myResult string = "Hưng nổ"
	// fmt.Println(myResult)

	// var data bool = false
	// fmt.Println(data)

	// number := 19

	// if number == 10 {
	// 	fmt.Println("number = 10")
	// } else {
	// 	fmt.Println("number != 10")
	// }

	// -------------------------------------------------------------------------------------------
	fmt.Println("localhost:8000")
	handleRequest()

}
