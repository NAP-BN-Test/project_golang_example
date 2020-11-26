package main

import (
	"encoding/json"
	"go-module/c4f"
	"go-module/table/account"

	_ "github.com/denisenkom/go-mssqldb"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})

}

func handleRequest() {
	// router := mux.NewRouter().StrictSlash(true)

	// // router.HandleFunc("/", homePage)
	// router.HandleFunc("/articles", allArticles).Methods("GET")
	// router.HandleFunc("/articles", allArticles).Methods("POST")
	// router.HandleFunc("/addUser/{name}", user.NewUser).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8000", router))

	router := gin.Default()
	// router.GET("/connect", connectdb)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/home", homePage)
	router.GET("/user", account.AllUsers)
	router.Run()
}

func main() {
	// ac := accounting.Accounting{Symbol: "$", Precision: 2}
	// fmt.Println(ac.FormatMoney(123456789.213123))

	c4f.Println("this is red color")
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
