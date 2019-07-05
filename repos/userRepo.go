package repos

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func Login(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "../view/login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, pwd FROM user WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "api/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "api/login", 301)
		return
	}

	res.Write([]byte("Hello" + databaseUsername))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../view/index.html")
}

func GetDB() *sql.DB {
	return db
}
