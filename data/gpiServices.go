package data

import (
	"gpitest/model"
	// "sso/util"
	// "strings"
	 "fmt"
	// "time"
)

func GetEnterprise(user_id string) *model.GPICompanyDetails {
	fmt.Println(user_id)
	row := Conn.QueryRow("select id, account_owner, account_id, company, companyname FROM companies where account_id='" + user_id + "'")
	var id int
	var account_owner string
	var account_id string
	var company string
	var companyname string
	
	row.Scan(&id, &account_owner, &account_id, &company, &companyname)

	if company != "" {
		var u = new(model.GPICompanyDetails)
		u.ID = id
		u.AccountOwner = account_owner
		u.AccountID = account_id
		u.Company = company
		u.CompanyName = companyname
		
		return u
	}
	
	return nil
}

func GetGPIService(serviceid string) *model.GPIService {
	//defer Conn.Close()
	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices limit 1")

	//var ServiceItems []GPIServices
	var ServiceItem model.GPIService

	//for rows.Next() {
	rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

		//ServiceItems = append(ServiceItems, ServiceItem)
	//}

	var ServiceItem_ = new(model.GPIService)

	ServiceItem_.ID = ServiceItem.ID
	ServiceItem_.Name = ServiceItem.Name
	ServiceItem_.Description = ServiceItem.Description
	ServiceItem_.DClass = ServiceItem.DClass
	ServiceItem_.IClass = ServiceItem.IClass

	return ServiceItem_

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