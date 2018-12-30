/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
The file that handles all the users
 */
package users

/*
type UserList struct
A structure containing the list of users
users[] - an array of users
 */
type UserList struct{
	Users []User
}

/*
func (userList *UserList) AddUser(user User) []User
Adds the user to the user list and returns the list.
If the user already has an account then he is not
added to the user list.
parameter - userName - the entered username
			password - the entered password
return - bool - true if a user matches
				false otherwise
 */
func (userList *UserList) AddUser(name string, password string, email string) []User  {
	user := User{}.MakeUser(name, password, email)

	for _, aUser := range userList.Users{
		if aUser == user {
			return userList.Users
		}
	}
	userList.Users = append(userList.Users, user)
	return userList.Users
}

/*
func (userList UserList) MatchPassword(userName string, password string) bool
The MatchPassword function loops through the lists of users who have created
and account and looks for the user with the same username and password.
parameter - userName - the entered username
			password - the entered password
return - bool - true if a user matches
				false otherwise
 */
func (userList UserList) MatchPassword(userName string, password string) bool{
	for _, aUser := range userList.Users{
		u := aUser.GetName()
		p := aUser.GetPassword()
		if u == userName && password == p{
			return true
		}
	}
	return false
}

/*
func (userlist UserList) GetEmail(username string)string
Gets the email of the user based on the username
parameter - username - the name of the user
return - email - if the the user has an account
		 "" - if the user does not have an account
 */
func (userlist UserList) GetEmail(username string)string{
	for _, aUser := range userlist.Users{
		if aUser.Name == username{
			return aUser.Email
		}
	}
	return ""
}

