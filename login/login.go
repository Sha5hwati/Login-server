/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
This file has all the page and request handlers
required to load html pages and pass information to the front end.
 */
package login

import (
	"encoding/json"
	"log"
	"loginServer/filemanager"
	"loginServer/users"
	"net/http"
)

//List of users registered
var userList = users.UserList{}

/***
func LoginHandler
Handles the client request to "/login"
It takes in the username and password from the request. It loads the data from
the file. It checks if the username and password are not empty. If they are not
and they match then it sets the cookie. It also adds the user to a list of online users.
It then redirects the user to the homepage.
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func LoginHandler(w http.ResponseWriter, r *http.Request){

	name := r.FormValue("username")
	password := r.FormValue("password")

	redirectTarget := "/"
	manager := filemanager.Manager{}
	if err := manager.Load("./users/userInfo.json", &userList); err != nil {
		log.Fatalln(err)
	}

	if len(name) !=0 && len(password)!=0{
		if userList.MatchPassword(name, password) == true {
			redirectTarget = "/home"
			SetCookie(name, w)
			email := userList.GetEmail(name)
			user := users.User{}.MakeUser(name, password, email)
			AddOnlineUser(user)
		}
	}
	http.Redirect(w, r, redirectTarget, 302)
}

/***
func LogoutHandler
Removes the user from the list of online people and clears the cookie
before redirecting the user to the default login page
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func LogoutHandler(w http.ResponseWriter, r *http.Request){
	username := GetUserName(r)
	RemoveOnlineUser(username)
	ClearCookie(w)
	http.Redirect(w, r, "/", 302)
}

/***
func NewuserHandler
Loads the list of users from the file and adds the new user to the list.
Then the list is saved back to the file and the user is redirected to the
default login page.
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func NewuserHandler(w http.ResponseWriter, r *http.Request){
	name := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	redirectTarget := "/"
	manager := filemanager.Manager{}
	if err := manager.Load("./users/userInfo.json", &userList); err != nil {
		log.Fatalln(err)
	}
	if len(name) !=0 && len(password)!=0 && len(email) !=0 {
		userList.AddUser(name, password, email)
		if err := manager.Save("./users/userInfo.json", userList); err != nil {
			log.Fatalln(err)
		}
		redirectTarget = "/"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

/***
func GetUserNameHandler
Returns the username of the logged in user for the request "/getUserName"
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func GetUserNameHandler(w http.ResponseWriter, r *http.Request){
	jsn, _ := json.Marshal(GetUserName(r))
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}



