package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sohaniwewelwala/go_auth_project/controller"
)

var db *sql.DB
var err error

const (
	dbName = "goauthproject"
	dbPass = "1234"
	dbHost = "localhost"
	dbPort = "root"
)

func main() {

	//init router
	r := mux.NewRouter()

	//router handler
	r.HandleFunc("/", controller.LoginPageHandler)
	r.HandleFunc("/api/login", controller.Authenticate).Methods("POST")
	r.HandleFunc("/api/home", controller.HomepageHandler).Methods("POST")
	r.Handle("/", http.FileServer(http.Dir("./view/")))

	// r.Use(app.JwtAuthentication)

	fmt.Println("server running")
	log.Fatal(http.ListenAndServe(":3000", r))
}
