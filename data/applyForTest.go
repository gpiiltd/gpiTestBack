package data

import (
	//"gpitest/model"
	//"log"
	// "gpitest/util"
	// "strings"
	 //"fmt"
	// "time"
)

// type FeaturedTestm struct {
// 	mw model.FeaturedTest
// }

type TestList_Params struct {
	ID string `json:"id,omitempty"`
	Owner_ID string `json:"owner_id,omitempty"`
	Name string  `json:"name,omitempty"`
	Slug string  `json:"slug,omitempty"`
	Group string  `json:"groups,omitempty"`
	//Owner string `json:"owner,omitempty"`
}

func ApplyForTest() []TestList_Params {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	AvailableTests := []TestList_Params{}
	
	results, err := ConnC.Query("select id, owner_id, name, slug, group from concerto.Test where slug like '"+"gpi"+"%'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		for results.Next() {

			AvailableTest := TestList_Params{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTest.ID, &AvailableTest.Owner_ID, &AvailableTest.Name, &AvailableTest.Slug, &AvailableTest.Group) //, &AvailableTest.Owner)
			if err != nil {
				//panic(err.Error())
			} else {
				//if ((tag.UID != r.Form.Get("uid")) || (len(tag.MobileAccessCode) < 1)) {

				//} else {
					//ttag := make(map[string]interface{})
					//ttag := new(TestParams)
					//ttag.ID = tag.ID
					// ttag.TestID = tag.TestID
					// ttag.NumQues = tag.NumQues
					// ttag.Duration = tag.Duration
					// ttag.Owner = tag.Owner
					// ttag.TestName = tag.TestName

					AvailableTests = append(AvailableTests, AvailableTest)
				//}

			}
		}


	return AvailableTests
	// //defer Conn.Close()
	// rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices limit 1")

	// //var ServiceItems []GPIServices
	// var ServiceItem model.GPIServicess

	// //for rows.Next() {
	// rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

	// 	//ServiceItems = append(ServiceItems, ServiceItem)
	// //}

	// var ServiceItem_ = new(model.GPIService)

	// ServiceItem_.ID = ServiceItem.ID
	// ServiceItem_.Name = ServiceItem.Name
	// ServiceItem_.Description = ServiceItem.Description
	// ServiceItem_.DClass = ServiceItem.DClass
	// ServiceItem_.IClass = ServiceItem.IClass

	// return ServiceItem_

}

// func GetAllGPIService() *[]model.GPIService {
// 	//defer Conn.Close()
// 	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices")

// 	var ServiceItems *[]model.GPIService
// 	var ServiceItem model.GPIService


// 	for rows.Next() {
// 		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

// 		ServiceItems = append(ServiceItems, ServiceItem)
// 	}

// 	return ServiceItems
// }