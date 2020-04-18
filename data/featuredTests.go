package data

 import (
	"gpitest/model"
	"log"
// 	// "sso/util"
// 	// "strings"
	"fmt"
// 	// "time"
)

func GetSubjectMatterCands(identity string, testid string, subjectmatter string) []model.TRUser {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	SubjectMatterCands := []model.TRUser{}
	log.Println(identity)
	fmt.Println(identity)
	results, err := Conn.Query("SELECT u.fname, u.lname, u.username, u.id FROM GPI_Oauth_clients.accessmentinfo a, GPI_Oauth_clients.users u where a.uid = u.username and u.company = '"+identity+"' and a.testid = '"+testid+"' and a.testtype = '"+subjectmatter+"' group by u.username") 
	//, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		for results.Next() {

			SubjectMatterCand := model.TRUser{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&SubjectMatterCand.Firstname, &SubjectMatterCand.Lastname, &SubjectMatterCand.Username, &SubjectMatterCand.ID)
			if err != nil {
				//panic(err.Error())
			} else {
				 //var sbtotal int
				 //row := Conn.QueryRow("SELECT count(distinct(uid)) as sbtotal FROM GPI_Oauth_clients.accessmentinfo a, GPI_Oauth_clients.users u where a.uid = u.username and u.company = '"+identity+"' and a.testid = '"+testid+"' and a.testtype = '"+subjectmatter+"'")
				 //row.Scan(&sbtotal)
					
				 	//SubjectMatterTest.SBTotal=sbtotal

					SubjectMatterCands = append(SubjectMatterCands, SubjectMatterCand)
				//}

			}
		}


	return SubjectMatterCands

}

 func GetSubjectMatterTest(identity string, tag string, subjectmatter string) []model.SB {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	SubjectMatterTests := []model.SB{}
	log.Println(identity)
	fmt.Println(identity)
	results, err := Conn.Query("select id, testname as name FROM managedtests where companyid='" + identity + "' and companytag='"+tag+"' and subjectmatter='"+subjectmatter+"'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		for results.Next() {

			SubjectMatterTest := model.SB{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&SubjectMatterTest.ID, &SubjectMatterTest.Name)
			if err != nil {
				//panic(err.Error())
			} else {
				 var sbtotal int
				 row := Conn.QueryRow("SELECT count(distinct(uid)) as sbtotal FROM GPI_Oauth_clients.accessmentinfo a, GPI_Oauth_clients.users u where a.uid = u.username and u.company = '"+identity+"' and a.testtype = '"+subjectmatter+"'")
				 row.Scan(&sbtotal)
					
				 	SubjectMatterTest.SBTotal=sbtotal

					SubjectMatterTests = append(SubjectMatterTests, SubjectMatterTest)
				//}

			}
		}


	return SubjectMatterTests

}

func GetSubjectMatter(identity string) []model.SB {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	SubjectMatters := []model.SB{}
	log.Println(identity)
	fmt.Println(identity)
	results, err := Conn.Query("select count(id) as id, subjectmatter as name FROM managedtests where companyid='" + identity + "' group by subjectmatter") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		for results.Next() {

			SubjectMatter := model.SB{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&SubjectMatter.ID, &SubjectMatter.Name)
			if err != nil {
				//panic(err.Error())
			} else {
				var sbtotal int
				row := Conn.QueryRow("SELECT count(distinct(uid)) as sbtotal FROM GPI_Oauth_clients.accessmentinfo a, GPI_Oauth_clients.users u where a.uid = u.username and u.company = '"+identity+"' and a.testtype = '"+SubjectMatter.Name+"'")
				row.Scan(&sbtotal)
					
					SubjectMatter.SBTotal=sbtotal

					SubjectMatters = append(SubjectMatters, SubjectMatter)
				//}

			}
		}


	return SubjectMatters

}