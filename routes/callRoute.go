package routes

import (
	"github.com/martini-contrib/render"
  	//"github.com/go-martini/martini"
	"net/http"
	//"net/url"
	"gpitest/data"
	"log"
	"fmt"
	"strconv"
	//"time"
	//"io/ioutil"
	//"gpitest/util"
	//"encoding/json"
)

func SaveMembersForTest(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
    log.Printf("%+v\n", req.Form)
    fmt.Println("entered SaveMembersForTest ")
    fmt.Println("%+v\n", req.Form)
    membersSelected := req.Form["subcheck"]
	result := make(map[string]interface{})
	testid := req.FormValue("testid")
	testname := req.FormValue("testname")
	testtype := req.FormValue("testtype")
	//username := req.FormValue("username")
	//usernames := req.FormValue("usernames")
	//fmt.Println("fparam = "+ttype)
	_, AT := data.Containsnew(membersSelected,testid,testname,testtype) //data.ListTests(ttype)
	
	//fmt.Println("this is ")
	//fmt.Println(AT)

	result["testid"] = testid
	result["testname"] = testname
	result["testtype"] = testtype
	result["data"] = AT
	r.JSON(200, result)
	return
}

func SaveTestsForMember(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
    log.Printf("%+v\n", req.Form)
    fmt.Printf("%+v\n", req.Form)
    fmt.Println("entered SaveTestsForMembers -")
    productsSelected := req.Form["subcheck"]
	result := make(map[string]interface{})
	username := req.FormValue("username")
	usernames := req.FormValue("usernames")
	user_id := req.FormValue("user_id")
	fmt.Println("userid - "+user_id)
	test_token, AT := data.Contains(productsSelected,usernames,username,user_id) //data.ListTests(ttype)
	
	
	result["username"] = username
	result["usernames"] = usernames
	result["data"] = AT
	result["test_token"] = test_token
	r.JSON(200, result)
	return
}

func PrepareTestResults(r render.Render, res http.ResponseWriter, req *http.Request) {
	//parse request parameters
	req.ParseForm()
	//filter := make(map[string]interface{})
	token := req.FormValue("testtoken")
	username := req.FormValue("username")
	test_id := req.FormValue("testid")
	//user_id := req.FormValue("userid")
	ttype := req.FormValue("testtype")
	names := req.FormValue("names")
	testname := req.FormValue("testname")
	returnresult := make(map[string]interface{})
	session_id := data.GetSessionID(token, username)
	if (len(session_id) == 0) {
		returnresult["status"] = "error"
		returnresult["data"] = "[]"
		returnresult["type"] = ttype
		returnresult["description"] = "Invalid Test token detected!"
		r.JSON(200, returnresult)
	}


	fmt.Println(session_id+" session id for " + ttype)
	//log.Println(session_id+" session id for " + ttype)
	
	 
	if (ttype == "Psychometrics") {
		results, boolval := data.GetTestResults(session_id)
		
		if (boolval) {
			//reformResult := data.ReformResults(results)
			results.Usersnames = names + " - (" + username+")"
			results.UserID = username
			results.TestID = test_id
			results.TestToken = token
			results.TestType = ttype
			results.TestName = testname

			i1q := (data.ConverToInt(results.Seq) + data.ConverToInt(results.Bwq) + data.ConverToInt(results.Lc))/3
			results.I1q = strconv.Itoa(i1q)
			i2q := (data.ConverToInt(results.Aaq) + data.ConverToInt(results.Paq) + data.ConverToInt(results.Fq) + data.ConverToInt(results.Cq))/4
			results.I2q = strconv.Itoa(i2q)
			
			reqd := (data.ConverToInt(results.Rq) + data.ConverToInt(results.Iq))/2
			results.Req = strconv.Itoa(reqd)
			
			gwq := (data.ConverToInt(results.Sq) + data.ConverToInt(results.Hq) + data.ConverToInt(results.Aq))/3
			results.Gwq = strconv.Itoa(gwq)

			results.Score = (i1q+i2q+reqd+gwq)/4

			wres := data.WriteTestResultsToDB(results)
			if (!wres) {
				returnresult["status"] = "warning"
				//returnresult["data"] = "[]"
				returnresult["type"] = ttype
				returnresult["description"] = "This result could not be recorded. Either it is already available or the network is down!"
			} else {
				//returnresult["data"] = ret //data.ReformResults(ret)
	 			returnresult["status"] = "success"
				returnresult["type"] = ttype
				returnresult["description"] = "Your test result was retrieved successfully, and is now available!"
			}
			ret, retboolval := data.GetDisplayResutls(username)
			if (retboolval) {
				returnresult["status"] = "success"
				returnresult["data"] = ret
			} else {
				returnresult["status"] = "warning"
				returnresult["data"] = ret
				returnresult["description"] = "There was an error retrieving your results. Please try again later."
			}
			

		} else {
			returnresult["status"] = "error"
			returnresult["data"] = "[]"
			returnresult["type"] = ttype
			returnresult["description"] = "Error encountered retrieving test results!"

		}
	} else if (ttype != "Psychometrics") {
		results, boolval := data.GetTestResultsSM(session_id)
		if (boolval) {
			results.Usersnames = names + " - (" + username+")"
			results.UserID = username
			results.TestID = test_id
			results.TestToken = token
			results.TestType = ttype
			results.TestName = testname

			//*****************

			wres := data.WriteTestResultsToDBSM(results)
			if (!wres) {
				returnresult["status"] = "warning"
				//returnresult["data"] = "[]"
				returnresult["type"] = ttype
				returnresult["description"] = "This result could not be recorded. Either it is already available or the network is down!"
			} else {
				//returnresult["data"] = ret //data.ReformResults(ret)
				returnresult["status"] = "success"
				returnresult["type"] = ttype
				returnresult["description"] = "Your test result was retrieved successfully, and is now available!"
			}
			ret, retboolval := data.GetDisplayResutls(username)
			if (retboolval) {
				returnresult["status"] = "success"
				returnresult["data"] = ret
			} else {
				returnresult["status"] = "warning"
				returnresult["data"] = ret
				returnresult["description"] = "There was an error retrieving your results. Please try again later."
			}
		} else {
			returnresult["status"] = "error"
			returnresult["data"] = "[]"
			returnresult["type"] = ttype
			returnresult["description"] = "Error encountered retrieving test results!"

		}
	} else {
		returnresult["status"] = "failed"
		returnresult["data"] = "[]"
		returnresult["type"] = ttype
		returnresult["description"] = "Result not ready for harvest"
		//r.JSON(200, "[{'status':error, 'data':'not avaiable'}]") //"[{'status':true, 'data':'written'}]") //"[{1},{2},{"+session_id+"}]")		
	}
	r.JSON(200, returnresult)
	return
}

func GetAllTestResults(r render.Render, res http.ResponseWriter, req *http.Request) {
	//parse request parameters
	req.ParseForm()
	//filter := make(map[string]interface{})
	
	username := req.FormValue("username")
	
	//user_id := req.FormValue("userid")
	
	//names := req.FormValue("names")
	
	returnresult := make(map[string]interface{})
	ret, retboolval := data.GetDisplayResutls(username)

	fmt.Printf(strconv.Itoa(len(ret)))

	//if (len(ret)  > 0) {
		if (retboolval) {
		
		returnresult["data"] = ret //data.ReformResults(ret)
		returnresult["status"] = "success"
		//returnresult["type"] = "All"
		returnresult["description"] = "Result ready for harvest"

	} else {
		returnresult["status"] = "failed"
		returnresult["data"] = "[]"
		//returnresult["type"] = ttype
		returnresult["description"] = "Result not ready for harvest"
		//r.JSON(200, "[{'status':error, 'data':'not avaiable'}]") //"[{'status':true, 'data':'written'}]") //"[{1},{2},{"+session_id+"}]")		
	}
	r.JSON(200, returnresult)
	return
}


func GetThisTestResults(r render.Render, res http.ResponseWriter, req *http.Request) {
	//parse request parameters
	req.ParseForm()
	//filter := make(map[string]interface{})
	
	username := req.FormValue("username")
	
	testid := req.FormValue("testid")
	
	//names := req.FormValue("names")
	
	returnresult := make(map[string]interface{})
	ret := data.ListTestResult(username, testid)
	fmt.Printf("%+v\n", ret)
	fmt.Printf(strconv.Itoa(len(ret)))
	if (len(ret)  > 0) {

		for _, frvals := range ret {
						fmt.Printf("--------------------hhhhhjjjff---------------")
			fmt.Printf("%+v\n",frvals)
			fmt.Printf("--------------------hhhhhjjjff---------------")
			//returnresult["data"] = data.ReformResults(ret)
			//returnresult["data"] = data.ReformResults(frvals)
			returnresult["data"] = data.ReformResults(frvals)
		}
		
		
		returnresult["status"] = "success"
		//returnresult["type"] = "All"
		returnresult["description"] = "Result ready for harvest"

	} else {
		returnresult["status"] = "failed"
		returnresult["data"] = "[]"
		//returnresult["type"] = ttype
		returnresult["description"] = "Result not ready for harvest"
		//r.JSON(200, "[{'status':error, 'data':'not avaiable'}]") //"[{'status':true, 'data':'written'}]") //"[{1},{2},{"+session_id+"}]")		
	}
	r.JSON(200, returnresult)
	return
}