package data

import (
	"sso/model"
	//"sso/util"
	//"strings"
	//"fmt"
	//"time"
)

func NewUserDetails(details model.NewGPIUser) *model.NewGPIUser {
	//defer Conn.Close()
	detail := new(model.NewGPIUser)

	detail.Lastname = details.Lastname
	detail.Firstname = details.Firstname
	detail.Dob = details.Dob
	detail.Sex = details.Sex
	detail.Email = details.Email
	detail.Address = details.Address
	detail.City = details.City
	detail.Username = details.Username
	detail.Password = details.Password
	detail.RPassword = details.RPassword

	return detail

}

func UserNameExist(username string) bool {
	//defer Conn.Close()
	row := Conn.QueryRow("SELECT id FROM users WHERE username = '" + username + "'")
	var thisNewUser__id int
	
	row.Scan(&thisNewUser__id)

	if thisNewUser__id > 0 {
		return true
	}
	return false
}
