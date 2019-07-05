package controller

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/sohaniwewelwala/go_auth_project/model"
	u "github.com/sohaniwewelwala/go_auth_project/utils"
)

// for GET
var LoginPageHandler = func(r http.ResponseWriter, w *http.Request) {
	var body, _ = u.LoadFile("view/login.html")
	fmt.Fprintf(r, body)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	// user := &model.User{}
	// err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	// if err != nil {
	// 	u.Respond(w, u.Message(false, "Invalid request"))
	// 	return
	// }
	username := r.FormValue("username")
	pass := r.FormValue("password")
	redirectTarget := "/"

	if !u.IsEmpty(username) && !u.IsEmpty(pass) {
		resp := model.Login(username, pass)

		if resp["status"] == true {

			// data := model.LoadData(resp)
			usrJSon, err := json.Marshal(resp["user"])
			if err != nil {
				panic(err)
			}
			redirectTarget = "/api/home"
			w.Header().Set("Content-Type", "application/json")
			w.Write(usrJSon)
			HomepageHandler(w, r)
		} else {
			redirectTarget = "/"
		}
	}
	// u.Respond(w, resp)
	http.Redirect(w, r, redirectTarget, 302)
}

// for GET
var HomepageHandler = func(response http.ResponseWriter, request *http.Request) {

	// user := &model.User{}
	// err := json.NewDecoder(request.Body).Decode(user) //decode the request body into struct and failed if any error occur
	// if err != nil {
	// 	u.Respond(response, u.Message(false, "Invalid request"))
	// 	return
	// }

	// fmt.Println(data["user"])

	var body, _ = u.LoadFile("view/home.html")
	fmt.Fprintf(response, body)
	// decoder := json.NewDecoder(r)
	// // userName := GetUserName(request)
	// fmt.Fprintf(response, indexBody, userName)
	// http.Redirect(response, request, "/", 302)
	// var indexBody, _ = u.LoadFile("view/home.html")
	// fmt.Fprintf(response, indexBody, user)
	// fmt.Println(response, indexBody)

	// decoder := json.NewDecoder(request.Body)
	// fmt.Println(decoder)
	// fmt.Println(decoder)
	// } else {
	//     http.Redirect(response, request, "/", 302)
	// }
}
