package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	//"regexp"
	"sso/model"
	"sso/data"
	"sso/util"
	"log"
	//"strings"
	"fmt"
)

func DoRegister(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in register")
	fmt.Printf("in register")
	founderrors := false
	//parse request parameters
	req.ParseForm()
	var ThisNewUserDetails model.NewGPIUser

	ThisNewUserDetails.Lastname = req.FormValue("lastname") //     string  `json:"lastname"`
	ThisNewUserDetails.Firstname = req.FormValue("othernames") //string  `json:"firstname"`
	ThisNewUserDetails.Dob = "" //      string `json:"dob"`
	ThisNewUserDetails.Sex = "" //      string `json:"sex"`
	ThisNewUserDetails.Email = req.FormValue("email") //        string  `json:"email"`
	ThisNewUserDetails.Address = req.FormValue("address") // string  `json:"address"`
	ThisNewUserDetails.City = req.FormValue("city") //      string `json:"city"`
	ThisNewUserDetails.Username = req.FormValue("username") //      string `json:"username"`
	ThisNewUserDetails.Password = req.FormValue("password") //	string `json:"username"`
	ThisNewUserDetails.RPassword = req.FormValue("Rpassword") //	string `json:"username"`

	if ThisNewUserDetails.Password != ThisNewUserDetails.RPassword {
		founderrors = true
	}


	thisUser := data.NewUserDetails(ThisNewUserDetails)

	//nameExists := data.UserNameExist()
	if data.UserNameExist(thisUser.Username) == true {
		founderrors = true		
	}

	hash, _ := util.HashPassword(thisUser.Password) // ignore error for the sake of simplicity

	_, _ = data.Conn.Exec("insert into users ( myemail,username,password,fname,lname,company)" +
		" values ( '" + ThisNewUserDetails.Email + "', '" + ThisNewUserDetails.Username + "', '" + hash + "', '" +
		ThisNewUserDetails.Firstname + "', '" + ThisNewUserDetails.Lastname + "', '" + "1" + "')")

	//myUserInsertId, err := lastIdVar.LastInsertId()
	err := false
	if err != false {
		founderrors = true
	} else {
		founderrors = false
		mypermission := "1"
		_, _ = data.Conn.Exec("insert into user_permission_matches ( user_id,permission_id )" + 
			" values ( '" + "1" + "', '" + mypermission + "')")

	}

	redirectUrl := "registration_confirmation"

	if founderrors {
		redirectUrl = "register"
		log.Println("redirect to url:" + redirectUrl)
		r.Redirect(redirectUrl)
		return
	}

	//redirectUrl = "subscriptions"

	log.Println("redirect to url:" + redirectUrl)
	r.Redirect(redirectUrl)
	return
}