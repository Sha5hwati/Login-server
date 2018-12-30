/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
Handles different pages
*/
package login

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/****
func LoadFile(fileName string) (string, error)
This function takes in filename as a string and reads it from the specified
folder.
Parameters:
	fileName: string - file to be loaded
Return
	string - data read in file
	error - any errors when loading the file
 */
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile("./public/" + fileName)
	if err != nil {
		return "", err

	}
	return string(bytes), nil
}

/***
func LoginPageHandler(response http.ResponseWriter, request *http.Request)
This function handles the request for the login page ("/")
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = LoadFile("login.html")
	fmt.Fprintf(response, body)
}

/***
func HomePageHandler(response http.ResponseWriter, request *http.Request)
It handles the "/home" request from the client.
Gets the username from the cookie to ensure that a verified user
is attempting to request teh homepage and then load the homepage
or else redirect to the login page.
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func HomePageHandler(response http.ResponseWriter, request *http.Request){
	username := GetUserName(request)
	if len(username) != 0 {
		var body, _= LoadFile("home.html")
		fmt.Fprintf(response, body)
	}else{
		http.Redirect(response, request, "/login", 302)
	}
}
