package data

import (
	//"gpitest/model"
	//"log"
	// "sso/util"
	// "strings"
	 //"fmt"
	// "time"
)

func DashboardBreakingFigures(username string) interface{} {
	//*model.DashboardBreakingFigure {
	//make(map[string]interface{})
	
	// CountUsers string  `json:"countusers"`
	// CountCompanies string  `json:"countcompanies"`
	// TestTypes string  `json:"testtypes"`
	// GPIServices string `json:"gpiservices"`

	var dbf = make(map[string]interface{}) //new(model.DashboardBreakingFigure)
	row := Conn.QueryRow("SELECT count(id) as id FROM companies")
	var id string
	row.Scan(&id)
	dbf["CountCompanies"] = id

	row = Conn.QueryRow("SELECT count(id) as countUsers FROM users")
	var countUsers string
	row.Scan(&countUsers)
	dbf["CountUsers"] = countUsers
	

	row = Conn.QueryRow("SELECT count(distinct(testid)) as testTypes FROM accessmentinfo")
	var testTypes string
	row.Scan(&testTypes)
	dbf["TestTypes"] = testTypes

	row = Conn.QueryRow("SELECT count(id) as gpiServices FROM allservices")
	var gpiServices string
	row.Scan(&gpiServices)
	dbf["GPIServices"] = gpiServices
	
	row = Conn.QueryRow("SELECT count(distinct(subjectmatter)) as sb FROM GPI_Oauth_clients.managedtests")
	var sbgroups string
	row.Scan(&sbgroups)
	dbf["TotalSubGroups"] = sbgroups


	row = Conn.QueryRow("SELECT count(id) as totalsubtests FROM GPI_Oauth_clients.managedtests where subjectmatter <>'Psychometrics' and publish=1")
	var totalsubtests string
	row.Scan(&totalsubtests)
	dbf["TotalSubTests"] = totalsubtests

	//psychometrics
	row = Conn.QueryRow("SELECT count(id) as totalpsytests FROM GPI_Oauth_clients.managedtests where subjectmatter = 'Psychometrics' and publish=1")
	var totalpsytests string
	row.Scan(&totalpsytests)
	dbf["TotalPsyTests"] = totalpsytests

	row = Conn.QueryRow("select count(id) as pinnedTests FROM accessmentinfo where uid='"+username+"' and status=1 and published=1 and testtype='Psychometrics'")
	var pinnedTests string
	row.Scan(&pinnedTests)
	dbf["PinnedTests"] = pinnedTests

	row = Conn.QueryRow("select count(id) as testTaken FROM accessmentinfo where uid='"+username+"' and status='3' and published='1' and testtype='Psychometrics'")
	var testTaken string
	row.Scan(&testTaken)
	dbf["TestTaken"] = testTaken

	row = Conn.QueryRow("select avg(testscore) as AvgTestScore FROM accessmentinfo where uid='"+username+"' and status='3' and published='1' and testtype='Psychometrics'")
	var AvgTestScore string
	row.Scan(&AvgTestScore)
	dbf["AvgTestScore"] = AvgTestScore

	//subject matter
	row = Conn.QueryRow("select count(id) as spinnedTests FROM accessmentinfo where uid='"+username+"' and status=1 and published=1 and testtype<>'Psychometrics'")
	var spinnedTests string
	row.Scan(&spinnedTests)
	dbf["sPinnedTests"] = spinnedTests

	row = Conn.QueryRow("select count(id) as stestTaken FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype<>'Psychometrics'")
	var stestTaken string
	row.Scan(&stestTaken)
	dbf["sTestTaken"] = stestTaken

	row = Conn.QueryRow("select avg(testscore) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype<>'Psychometrics'")
	var sAvgTestScore string
	row.Scan(&sAvgTestScore)
	dbf["sAvgTestScore"] = sAvgTestScore

	row = Conn.QueryRow("select avg(`i1q`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var i1q string
	row.Scan(&i1q)
	dbf["i1q"] = i1q

	row = Conn.QueryRow("select avg(`i2q`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var i2q string
	row.Scan(&i2q)
	dbf["i2q"] = i2q

	row = Conn.QueryRow("select avg(`req`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var req string
	row.Scan(&req)
	dbf["req"] = req

	row = Conn.QueryRow("select avg(`gwq`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var gwq string
	row.Scan(&gwq)
	dbf["gwq"] = gwq

	row = Conn.QueryRow("select avg(`paq`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var paq string
	row.Scan(&paq)
	dbf["paq"] = paq

	row = Conn.QueryRow("select avg(`aaq`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var aaq string
	row.Scan(&aaq)
	dbf["aaq"] = aaq

	row = Conn.QueryRow("select avg(`fq`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var fq string
	row.Scan(&fq)
	dbf["fq"] = fq

	row = Conn.QueryRow("select avg(`cq`) as sAvgTestScore FROM accessmentinfo where uid='"+username+"' and status=3 and published=1 and testtype='Psychometrics'")
	var cq string
	row.Scan(&cq)
	dbf["cq"] = cq

	return dbf

	//for psychometrics
		//count piined test
			//select * FROM GPI_Oauth_clients.accessmentinfo where uid='coreadmin' and status=1 and published=1 and testtype='Psychometrics'
		//count all test taken
			//select * FROM GPI_Oauth_clients.accessmentinfo where uid='coreadmin' and status=3 and published=1 and testtype='Psychometrics'
		//average test score
			//select avg(testscore) as testscore FROM GPI_Oauth_clients.accessmentinfo where uid='coreadmin' and status=3 and published=1 and testtype='Psychometrics';



	// //defer Conn.Close()
	// rows, _ := Conn.Query("select id, sname, description, dclass, iclass from allservices limit 1")

	// //var ServiceItems []GPIServices
	// var ServiceItem model.GPIService

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