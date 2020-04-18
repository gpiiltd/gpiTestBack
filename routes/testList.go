package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"gpitest/data"
	//"gpitest/model"
	//"strings"
	"fmt"
	//"encoding/json"
	"time"
	"strconv"
)

func ListTests(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ttype := req.FormValue("ttype")
	AT := data.ListTests(ttype)
	
	r.JSON(200, AT)
	return
}

func ListMyTests(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	ttype := req.FormValue("ttype")
	AT := data.ListMyTests(ttype)
	
	r.JSON(200, AT)
	return
}

func GetEnterpriseMembers(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	identity := req.FormValue("identity")
	EM := data.GetEnterpriseMembers(identity)
	
	r.JSON(200, EM)
	return
}

func GetEnterpriseTests(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	identity := req.FormValue("identity")
	ET := data.GetEnterpriseTests(identity)
	
	r.JSON(200, ET)
	return
}

func ListTestResults(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//username := req.FormValue("username")
	userid := req.FormValue("userid")
	//fmt.Printf(userid)
	//fmt.Printf("starr")
	TRL := data.ListTestResults(userid)
	//data.GetJSON(userid) //
	//m, _ := json.Marshal(TRL)
	
	r.JSON(200, TRL)
	return
}

func SaveManagedTests(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	result := make(map[string]interface{})
	fmt.Printf("%+v\n", req.Form)
	user_id := req.FormValue("user_id")
	testslug := req.FormValue("testslug")

	enterPrise := data.GetEnterprise(user_id)
	if enterPrise == nil {
		result["status"] = false
		result["description"] = "You are not authorized to perform this operation. Kindly Register your organization."

	} else {

		companytag := enterPrise.Company //req.FormValue("companytag")
		companyid := enterPrise.ID

		testname := req.FormValue("testname")
		duration := req.FormValue("duration")
		totalquestions := req.FormValue("totalquestions")
		passscore := req.FormValue("passscore")
		subjectmatter := req.FormValue("subjectmatter")
		instructions := req.FormValue("instructions")
		testquestionfile := string(time.Now().Format(time.RFC850))+"_"+user_id+".csv"
		//time.Now().Format(time.RFC850) //req.FormValue("testquestionfile")
		//fmt.Printf(userid)
		//fmt.Printf("starr")
		//user_id, companytag, testname, duration, totalquestions, passscore, subjectmatter, instructions, testquestionfile, created
		
		result["status"] = data.SaveManagedTests(user_id, companytag, testname, duration, totalquestions, passscore, subjectmatter, 
			instructions, testquestionfile,time.Now().Unix(),testslug, strconv.Itoa(companyid))
		result["description"] = testname + ": Saved Successfully"
		//data.GetJSON(userid) //
		//m, _ := json.Marshal(TRL)
	}
	SMT := result

	


	
	r.JSON(200, SMT)
	return
}