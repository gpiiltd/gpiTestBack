package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"gpitest/data"
	//"gpitest/model"
	//"strings"
	//"log"
	"strconv"
	"fmt"
)

func GetSubjectMatterCands(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	result := make(map[string]interface{})
	identity := req.FormValue("identity")
	subjectmatter := req.FormValue("subjectmatter")
	testid := req.FormValue("testid")
	fmt.Printf("%+v\n", req.Form)
	fmt.Printf("---")
	fmt.Printf(identity)
	fmt.Printf("---")
	fmt.Printf(testid)
	fmt.Printf("---")
	fmt.Printf(subjectmatter)
	enterPrise := data.GetEnterprise(identity)
	if enterPrise == nil {
		result["status"] = false
		result["description"] = "You are not authorized to perform this operation. Kindly Register your organization."

	} else {

		result["status"] = true
		result["description"] = data.GetSubjectMatterCands(strconv.Itoa(enterPrise.ID), testid, subjectmatter)

		
	}

	SM := result
	
	r.JSON(200, SM)
	return



}
	
func GetSubjectMatter(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	result := make(map[string]interface{})
	identity := req.FormValue("identity")
	enterPrise := data.GetEnterprise(identity)
	if enterPrise == nil {
		result["status"] = false
		result["description"] = "You are not authorized to perform this operation. Kindly Register your organization."

	} else {

		result["status"] = true
		result["description"] = data.GetSubjectMatter(strconv.Itoa(enterPrise.ID))

		
	}

	SM := result
	
	r.JSON(200, SM)
	return

}

func GetSubjectMatterTest(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	result := make(map[string]interface{})
	identity := req.FormValue("identity")
	subjectmatter := req.FormValue("subjectmatter")
	enterPrise := data.GetEnterprise(identity)
	if enterPrise == nil {
		result["status"] = false
		result["description"] = "You are not authorized to perform this operation. Kindly Register your organization."

	} else {

		result["status"] = true
		result["description"] = data.GetSubjectMatterTest(strconv.Itoa(enterPrise.ID), enterPrise.Company, subjectmatter)

		
	}

	SM := result
	
	r.JSON(200, SM)
	return



}

func GetTestCandidates(r render.Render, res http.ResponseWriter, req *http.Request) {


}

func FeaturedTests(r render.Render, res http.ResponseWriter, req *http.Request) {
	//ft := data.FeaturedTests()
	//items, err := item.GetAllServices()
	//res.Header().Set("Access-Control-Allow-Origin", "*")
	//res.Header().Add("Content-Type", "application/json")

	//r.JSON(200, ft)
	return


	//return ServiceItem_
	// log.Println("in ServiceValidate")
	// //parse request parameters
	// req.ParseForm()
	// service := req.FormValue("service")
	// ticket := req.FormValue("ticket")

	// log.Println("parsed params:service:" + service + ",ticket:" + ticket)
	// result := make(map[string]interface{})
	// st := data.FindST(ticket)
	// if st == nil {
	// 	result["result"] = false
	// 	result["code"] = 401
	// 	result["msg"] = "invalid st"
	// 	log.Println("invalid st:" + ticket)
	// 	r.JSON(401, result)
	// 	return
	// }
	// valid := strings.Contains(service, st.Service)
	// if valid {
	// 	tgt := data.FindTGT(st.Tgt)
	// 	if tgt == nil {
	// 		result["result"] = false
	// 		result["code"] = 401
	// 		result["msg"] = "can not found tgt with given st"
	// 		log.Println("can not found tgt" + st.Tgt + " with given st:" + ticket)
	// 		r.JSON(401, result)
	// 		return
	// 	} else {
	// 		result["result"] = true
	// 		result["code"] = 200
	// 		result["msg"] = "service validated"
	// 		result["username"] = tgt.Username
	// 		log.Println("service validated with username:" + tgt.Username)
	// 		r.JSON(200, result)
	// 		return
	// 	}

	// } else {
	// 	result["result"] = false
	// 	result["code"] = 401
	// 	result["msg"] = "service dose not match st"
	// 	log.Println("service dose not match st,service" + service + ", service in st:" + st.Service)
	// 	r.JSON(401, result)
	// 	return
	// }

}
