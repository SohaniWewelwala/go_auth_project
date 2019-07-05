package model

import (
	"database/sql"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	u "github.com/sohaniwewelwala/go_auth_project/utils"
)

var db *sql.DB
var err error

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type User struct {
	Username  string `json:"username"`
	Password  string `json:"pwd"`
	Longitude string `json:"long"`
	Latitude  string `json:"lat"`
}

func Login(username, password string) map[string]interface{} {

	db, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/goauthproject")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Ping to database successful, connection is still alive")

	user := &User{}

	fmt.Println(username + " " + password)

	// var databaseUsername, databasePassword, dblong, dblat string

	err := db.QueryRow("SELECT username, pwd, longit, lat FROM user WHERE username=?", username).Scan(&user.Username, &user.Password, &user.Longitude, &user.Latitude)

	fmt.Println(user.Username + " ============ " + user.Password)

	if err != nil {
		if err != nil {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	if user.Username != username {
		return u.Message(false, "Invalid Username. Please try again")
	} else if user.Password != password {
		return u.Message(false, "Invalid password. Please try again")
	}

	//Create JWT token
	// tk := &Token{UserId: 5}
	// token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	// user.Token = tokenString //Store the token in the response
	// fmt.Println("Token=" + user.Token)

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func LoadData(data map[string]interface{}) interface{} {
	// func LoadData(username, logitude, latitude string) map[string]interface{} {

	return data["user"]
}
