//The main file handling the login server
package main

/*
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
Program - Login Server
Backend - Go Language (Go 1.11)
Frontend - HTML, CSS and JavaScript
Description - To demonstrate the power of Go language, we
decided to implement a Login Server. The login server allows
users to create accounts, login, view active users and logout.
In this program, we used methods that use goroutines for data
synchronization. Keeping that in mind, the users have been made
public.
Usage - Build and run the main program
		Go to http://localhost:8080/ to see the server
Assumption - A user can have multiple accounts with the same name
			 A user can have multiple login sessions
			 No password security
Note - To make this program work accurately, login users on different
computers on http:/<ip address>:8080/. The ip address should of the
server computer. This is because is we login different users on same
computer the server just overwrites the existing cookie and logout
functionality fails
 */

import (
	"loginServer/login"

	//implements a request router and dispatcher for matching
	//incoming requests to their respective handler
	"github.com/gorilla/mux"

	//logging package
	"log"

	//provides HTTP client and server implementations
	"net/http"
)

/*
func main()
It sets up a router which handles different functions
associated with each page. The index page for the server
is the login page. If the user logs in correctly, it takes
him to the home page, where the server identifies the user.
The user can get a list of active users by clicking on the
view all online users button or can logout.
If the user is new, he can create an account on the page and
then login to view the service
 */
func main(){
	var router = mux.NewRouter()

	//Loads the Login Page
	router.HandleFunc("/", login.LoginPageHandler)

	//Loads the Home Page
	router.HandleFunc("/home", login.HomePageHandler)

	//Handles the login button functions
	router.HandleFunc("/login", login.LoginHandler).Methods("POST")

	//Handles the logout button functions
	router.HandleFunc("/logout", login.LogoutHandler).Methods("POST")

	//Handles adding a new user functions
	router.HandleFunc("/newUser", login.NewuserHandler).Methods("POST")

	//Handles getting the list of active users functions
	router.HandleFunc("/getuser", login.GetUserHandler).Methods("POST")

	//Handles getting the current user's Username information functions
	router.HandleFunc("/getUserName", login.GetUserNameHandler).Methods("POST")

	//handles requests made on the router
	http.Handle("/", router)
	fileServer := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	//Listens to TCP network address and handles requests on incoming connections
	err:= http.ListenAndServe(":8080", nil)

	//checks for errors
	if err != nil {
		log.Fatal("Listen and serve at 8080 ", err)
	}
}
