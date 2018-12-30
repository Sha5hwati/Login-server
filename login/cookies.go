/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
This contains the functions to set and clear cookies.
 */
package login

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

//For security, encodes the coockie into a
//hash take and creates a secure cookie object
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

/***
func ClearCookie
Resets the existing cookie on the client.
Parameters:
	response http.ResponseWriter
 */
func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

/***
func SetCookiee
Sets the cookie with the value of encoded user name
Parameters:
	userName: string - user name of the user on the client
	response: http.ResponseWriter
 */
func SetCookie(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

/***
func GetUserName
This function gets the cookie from the request and identifies the user
with the username
Parameters:
	response: http.ResponseWriter
	request: http.Request
Returns
	string- username
 */
func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}






