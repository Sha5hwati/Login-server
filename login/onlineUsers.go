/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
Handles online users
*/
package login

import (
	"encoding/json"
	"loginServer/users"
	"net/http"
)

//List fo users currently online
var online []users.User

/***

 */
func  AddOnlineUser(user users.User){
	online = append(online, user)
}

/****
func (onlineUser OnLineUser) RemoveUser(username string)
The function removes users that are logged out from the list
of online people.
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func RemoveOnlineUser(username string){
	for i := len(online) - 1; i > 0; i-- {
		if online[i].GetName() == username{
			online = append(online[:i], online[i+1:]...)
			return
		}
	}
}

/***
func (onlineUser OnLineUser) GetUserHandler(w http.ResponseWriter, r *http.Request)
This function sends the list of all online users to the front end
in json format.
Parameters:
	response: http.ResponseWriter
	request: http.Request
 */
func GetUserHandler(w http.ResponseWriter, r *http.Request){
	jsn, _ := json.Marshal(online)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsn)
}


