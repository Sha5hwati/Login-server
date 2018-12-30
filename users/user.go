/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
Handles a single user
 */
package users

/*
type User struct
A structure that contains the structure of a user
Parameters - name - storing the name of the account
			 password - storing the password of the account
			email - storing the email of the account
 */
type User struct{
	Name string
	Password string
	Email string
}

/*
func (user User) MakeUser(name string, password string, email string) User
Constructs a new user
Parameters - name - name of the user
			password - password of the user
			email - email of the user
returns - the new user
 */
func (user User) MakeUser(name string, password string, email string) User {
	user.Name = name
	user.Password = password
	user.Email = email
	return user
}


/*
func (user User) GetName() string
Returns the name of the user
 */
func (user User) GetName() string {
	return user.Name
}

/*
func (user User) GetPassword() string
Returns the password of the user
 */
func (user User) GetPassword() string {
	return user.Password
}

/*
func (user User) GetEmail() (string)
Returns the email of the user
 */
func (user User) GetEmail() string {
	return user.Email
}

/*
func (user User) SetName(newName string)
Sets the name of the user to a new value
parameter - newName - new name of the user
 */
func (user *User) SetName(newName string) {
	user.Name = newName
}

/*
func (user User) SetPassword(newPassword string)
Sets the password of the user to a new value
parameter - newPassword - new password of the user
 */
func (user *User) SetPassword(newPassword string) {
	user.Name = newPassword
}

/*
func (user User) SetEmail(newEmail string)
Sets the email of the user to a new value
parameter - newEmail - new email of the user
 */
func (user *User) SetEmail(newEmail string) {
	user.Name = newEmail
}

