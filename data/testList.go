package data

import (
	//"gpitest/model"
	"log"
	"strings"
	//"time"
	// "gpitest/util"
	"strconv"
	 "fmt"
	"encoding/json"
)

type EGPIUser struct {
	Lastname        string  `json:"lastname"`
	Firstname string  `json:"firstname"`
	Dob       string `json:"dob",omitempty`
	Gender       string `json:"gender",omitempty`
	Email     string  `json:"myemail",omitempty`
	ID string  `json:"id",omitempty`
	City       string `json:"city",omitempty`
	Username       string `json:"username"`
	
}

type TestListParams struct {
	ID string `json:"id,omitempty"`
	Owner_ID string `json:"owner_id,omitempty"`
	Name string  `json:"name,omitempty"`
	Slug string  `json:"slug,omitempty"`
	Group string  `json:"groups,omitempty"`
	Username string `json:"username,omitempty"`
	TestUrl string `json:"testurl,omitempty"`
	Duration       string `json:"duration,omitempty"`
	TotalQuestions       string `json:"totalquestions,omitempty"`
	PassScore       string `json:"passscore,omitempty"`
	Instructions       string `json:"instructions,omitempty"`

}


type GetDisplayResutlsParams struct {			
	ID string `json:"id,omitempty"`
	Course string `json:"course,omitempty"`
	TestScore       string `json:"testscore,omitempty"`
	TakenOn string `json:"takenon,omitempty"`
	CreatedBy string `json:"createdby,omitempty"`
	Recommendation string `json:"recommendation,omitempty"`
	TestID string `json:"testid,omitempty"`
	UID string `json:"uid,omitempty"`
	TestType string `json:"testtype,omitempty"`
	TotalSMQuestions string `json:"totalsmquestions,omitempty"`
}

type GetEnterpriseTestsParams struct {
	ID int `json:"id,omitempty"`
	User_ID int `json:"user_id,omitempty"`
	CompanyTag string  `json:"companytag,omitempty"`
	CompanyID int `json:"companyid,omitempty"`
	Testname string  `json:"testname,omitempty"`
	Duration int  `json:"duration,omitempty"`
	TotalQuestions int `json:"totalquestions,omitempty"`
	PassScore int `json:"passscore,omitempty"`
	SubjectMatter string `json:"subjectmatter,omitempty"`
	Instructions string `json:"instruction,omitempty"`
	TestQuestionFile string `json:"testquestionfile,omitempty"`
	Created string `json:"created,omitempty"`
	TestSlug string `json:"testslug,omitempty"`
}

type ConcertoResults struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Value string `json:"value"`
	Session_id string `json:"session_id"`
}

type ConcertoResultsSM struct {
	Item_ID string `json:"item_id"`
	Response string `json:"response"`
	Time_Taken string `json:"time_taken"`
	Session_id string `json:"session_id"`
	Correct int `json:"correct"`
}

type TestResultsParams struct {
	Usersnames string `json:"usersnames,omitempty"`
	UserID string `json:"userid,omitempty"`
	TestID string `json:"testid,omitempty"`
	TestToken string `json:"testtoken,omitempty"`
	TestType string `json:"testtype,omitempty"`
	TestName string `json:"testname,omitempty"`	
	I1q string `json:"i1q,omitempty"`
	I2q string `json:"i2q,omitempty"`
	Req string `json:"req,omitempty"`
	Gwq string `json:"gwq,omitempty"`
	Seq string `json:"seq,omitempty"`
	Lc string `json:"lc,omitempty"`
	Paq string `json:"paq,omitempty"`
	Aaq string `json:"aaq,omitempty"`
	Rq string `json:"rq,omitempty"`
	Sq string `json:"sq,omitempty"`
	Iq string `json:"iq,omitempty"`
	Fq string `json:"fq,omitempty"`
	Cq string `json:"cq,omitempty"`
	Hq string `json:"hq,omitempty"`
	Aq string `json:"aq,omitempty"`
	Bwq string `json:"bwq,omitempty"`
	Score int `json:"score,omitempty"`

}

type TestResultsParamsSM struct {
	Usersnames string `json:"usersnames,omitempty"`
	UserID string `json:"userid,omitempty"`
	TestID string `json:"testid,omitempty"`
	TestToken string `json:"testtoken,omitempty"`
	TestType string `json:"testtype,omitempty"`
	TestName string `json:"testname,omitempty"`	
	
	Score int `json:"score,omitempty"`
	Correct int `json:"correct,omitempty"`
	Wrong int `json:"wrong,omitempty"`

}

type TestResultsParamsR struct {

	I1q int 
	I1qT string 

	I2q int 
	I2qT string 

	Req int
	ReqT string 
	
	Gwq int
	QwqT string

	Seq int
	SeqT string

	Lc int 
	LcT string

	Paq int 
	PaqT string

	Aaq int 
	AaqT string

	Rq int
	RqT string

	Sq int 
	SqT string

	Iq int 
	IqT string

	Fq int 
	FqT string

	Cq int 
	CqT string

	Hq int 
	HqT string

	Aq int 
	AqT string

	Bwq int
	BwqT string

	Score int
	ScoreT string

}

type TestResultsParamsmm struct {
	i1q string `json:"i1q"`
	i2q string `json:"i2q"`
	req string `json:"req"`
	gwq string `json:"gwq"`
	seq string `json:"sed"`
	lc string `json:"lc"`
	paq string `json:"paq"`
	aaq string `json:"aaq"`
	rq string `json:"rq"`
	sq string `json:"sq"`
	iq string `json:"iq"`
	fq string `json:"fq"`
	cq string `json:"cq"`
	hq string `json:"hq"`
	aq string `json:"aq"`
	bwq string `json:"bwq"`
}

func GetSessionID(token string, username string) string {
	fmt.Println(username)
	fmt.Println(token)
	row := ConnC.QueryRow("select session_id FROM concerto.default_data_table where value = '"+token+"' order by id desc limit 1")
	var SessionID string
	
	row.Scan(&SessionID)

	if SessionID != "" {
		fmt.Printf(SessionID)
		fmt.Println(" king")
		return SessionID
	} else {
		fmt.Printf(SessionID)
		fmt.Println(" queen")
		return ""
	}
}

func GetEnterpriseMembers(identity string) []EGPIUser {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	RegisteredMembers := []EGPIUser{}
	log.Println(identity)
	fmt.Println(identity)
	results, err := Conn.Query("select id, fname, lname, username, myemail, gender, city from GPI_Oauth_clients.users where company = '"+identity+"' order by id desc") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		for results.Next() {

			RegisteredMember := EGPIUser{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&RegisteredMember.ID, &RegisteredMember.Firstname, &RegisteredMember.Lastname, &RegisteredMember.Username, &RegisteredMember.Email, &RegisteredMember.Gender, &RegisteredMember.City)
			if err != nil {
				//panic(err.Error())
			} else {
					//AvailableTest.TestUrl="http://206.189.4.3/test/"+AvailableTest.Slug+"?"

					RegisteredMembers = append(RegisteredMembers, RegisteredMember)
				//}

			}
		}


	return RegisteredMembers

}

func GetEnterpriseTests(identity string) []GetEnterpriseTestsParams {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	AvailableTests := []GetEnterpriseTestsParams{}
	
	results, err := Conn.Query("select id, user_id, companytag, testname, duration, totalquestions, passscore, subjectmatter, instructions, testquestionsfile, created, testslug, companyid from GPI_Oauth_clients.managedtests where companyid = '"+identity+"' order by id desc") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
			fmt.Printf("%+v", err)
		 }

		for results.Next() {

			AvailableTest := GetEnterpriseTestsParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTest.ID, &AvailableTest.User_ID, &AvailableTest.CompanyTag, &AvailableTest.Testname, &AvailableTest.Duration, &AvailableTest.TotalQuestions, &AvailableTest.PassScore, &AvailableTest.SubjectMatter, &AvailableTest.Instructions, &AvailableTest.TestQuestionFile, &AvailableTest.Created, &AvailableTest.TestSlug, &AvailableTest.CompanyID)
			if err != nil {
				//panic(err.Error())
				fmt.Printf("%+v", err)
			} else {
					//AvailableTest.TestUrl="http://206.189.4.3/test/"+AvailableTest.Slug+"?"
				fmt.Printf(AvailableTest.Testname)

					AvailableTests = append(AvailableTests, AvailableTest)
				//}

			}
		}

	return AvailableTests

}

func ListTests(ttype string) []TestListParams {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	AvailableTests := []TestListParams{}
	log.Println(ttype)
	fmt.Println(ttype)
	results, err := Conn.Query("select id, user_id, testname, testslug, subjectmatter, companytag, duration, totalquestions, passscore, instructions from managedtests where companyid = '"+ttype+"' and publish=1") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		//  if results.Next() {
		// 	//return true
		// } else {
		// 	return AvailableTests
		// }

		for results.Next() {

			AvailableTest := TestListParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTest.ID, &AvailableTest.Owner_ID, &AvailableTest.Name, &AvailableTest.Slug, &AvailableTest.Group, &AvailableTest.Username, &AvailableTest.Duration, &AvailableTest.TotalQuestions, &AvailableTest.PassScore, &AvailableTest.Instructions )
			if err != nil {
				//panic(err.Error())
			} else {
					AvailableTest.TestUrl="http://test.my-gpi.com:8992/test/"+AvailableTest.Slug+"?test_id="+AvailableTest.ID

					AvailableTests = append(AvailableTests, AvailableTest)
				//}

			}
		}


	return AvailableTests

}


func ListMyTests(username string) []TestListParams {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	AvailableTests := []TestListParams{}
	//log.Println(ttype)
	//fmt.Println(ttype)
	results, err := Conn.Query("select m.id, m.user_id, m.testname, m.testslug, m.subjectmatter, m.companytag, m.duration, m.totalquestions, m.passscore, m.instructions from managedtests m, accessmentinfo a where a.uid = '"+username+"' and a.testid=m.id and a.status ='1'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }

		//  if results.Next() {
		// 	//return true
		// } else {
		// 	return AvailableTests
		// }

		for results.Next() {

			AvailableTest := TestListParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTest.ID, &AvailableTest.Owner_ID, &AvailableTest.Name, &AvailableTest.Slug, &AvailableTest.Group, &AvailableTest.Username, &AvailableTest.Duration, &AvailableTest.TotalQuestions, &AvailableTest.PassScore, &AvailableTest.Instructions )
			if err != nil {
				//panic(err.Error())
			} else {
					AvailableTest.TestUrl="http://test.my-gpi.com:8992/test/"+AvailableTest.Slug+"?test_id="+AvailableTest.ID

					AvailableTests = append(AvailableTests, AvailableTest)
				//}

			}
		}


	return AvailableTests

}



func GetUserByName(username string) string {
	//defer Conn.Close()
	row := Conn.QueryRow("select id from users where username = '" + username + "'")
	var userid string
	
	row.Scan(&userid)
	//fmt.Printf(userid)
	if userid != "" {
		
		return userid
	}
	return ""
}

func ListTestResult(username string, testid string) []TestResultsParams {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)"
	AvailableTestResults := []TestResultsParams{}
//	fmt.Println(userid)

	//fmt.Println(ttype)
		//results, err := Conn.Query("SELECT id,i1q2a,i2q2a,req2a,gwq2a,seq2a,lc2a,paq2a,aaq2a,rq2a,sq2a,iq2a,fq2a,cq2a,hq2a,aq2a,bwq2a from accessmentinfo where ") //, testid, testname, numques, duration, owner FROM featuredtests" )
	//var id, uid string
	results, err := Conn.Query("select i1q, i2q, req, gwq, seq, lc, paq, aaq, rq, sq, iq, fq, cq, hq, aq, bwq, testscore from GPI_Oauth_clients.accessmentinfo where uid = '"+username+"' and status='3' and published='1' and id='"+testid+"'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }
		for results.Next() {
			AvailableTestResult := TestResultsParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTestResult.I1q, &AvailableTestResult.I2q, &AvailableTestResult.Req, &AvailableTestResult.Gwq, &AvailableTestResult.Seq, &AvailableTestResult.Lc, &AvailableTestResult.Paq, &AvailableTestResult.Aaq, &AvailableTestResult.Rq, &AvailableTestResult.Sq, &AvailableTestResult.Iq, &AvailableTestResult.Fq, &AvailableTestResult.Cq, &AvailableTestResult.Hq, &AvailableTestResult.Aq, &AvailableTestResult.Bwq, &AvailableTestResult.Score)
			if err != nil {
				//panic(err.Error())
				fmt.Printf("%+v",err )
			} else {
					AvailableTestResults = append(AvailableTestResults, AvailableTestResult)
					//fmt.Printf("%+v", AvailableTestResults )
					//fmt.Printf("-------------" )
			}
		}
//		fmt.Printf("%+v", AvailableTestResults )
//		fmt.Printf("end")
	return AvailableTestResults
}


func ListTestResults(username string) []TestResultsParams {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)"
	AvailableTestResults := []TestResultsParams{}
//	fmt.Println(userid)

	//fmt.Println(ttype)
		//results, err := Conn.Query("SELECT id,i1q2a,i2q2a,req2a,gwq2a,seq2a,lc2a,paq2a,aaq2a,rq2a,sq2a,iq2a,fq2a,cq2a,hq2a,aq2a,bwq2a from accessmentinfo where ") //, testid, testname, numques, duration, owner FROM featuredtests" )
	//var id, uid string
	results, err := Conn.Query("select i1q, i2q, req, gwq, seq, lc, paq, aaq, rq, sq, iq, fq, cq, hq, aq, bwq, testscore from GPI_Oauth_clients.accessmentinfo where uid = '"+username+"' and status='3' and published='1'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			// 	panic(err.Error())
		 }
		for results.Next() {
			AvailableTestResult := TestResultsParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTestResult.I1q, &AvailableTestResult.I2q, &AvailableTestResult.Req, &AvailableTestResult.Gwq, &AvailableTestResult.Seq, &AvailableTestResult.Lc, &AvailableTestResult.Paq, &AvailableTestResult.Aaq, &AvailableTestResult.Rq, &AvailableTestResult.Sq, &AvailableTestResult.Iq, &AvailableTestResult.Fq, &AvailableTestResult.Cq, &AvailableTestResult.Hq, &AvailableTestResult.Aq, &AvailableTestResult.Bwq, &AvailableTestResult.Score)
			if err != nil {
				//panic(err.Error())
				fmt.Printf("%+v",err )
			} else {
					AvailableTestResults = append(AvailableTestResults, AvailableTestResult)
					//fmt.Printf("%+v", AvailableTestResults )
					//fmt.Printf("%+v", AvailableTestResult )
			}
		}
//		fmt.Printf("%+v", AvailableTestResults )
//		fmt.Printf("end")
	return AvailableTestResults
}


func GetDisplayResutls(username string) ([]GetDisplayResutlsParams, bool) {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)"
	AvailableTestResults := []GetDisplayResutlsParams{}

//	fmt.Println(userid)

	//fmt.Println(ttype)

		//results, err := Conn.Query("SELECT id,i1q2a,i2q2a,req2a,gwq2a,seq2a,lc2a,paq2a,aaq2a,rq2a,sq2a,iq2a,fq2a,cq2a,hq2a,aq2a,bwq2a from accessmentinfo where ") //, testid, testname, numques, duration, owner FROM featuredtests" )
	//var id, uid string
	results, err := Conn.Query("select ai.id, ai.uid,ai.testscore, ai.testid, ai.testname, ai.c_date, ai.testtype, mt.companytag, mt.passscore from GPI_Oauth_clients.accessmentinfo ai, GPI_Oauth_clients.managedtests mt where ai.testid=mt.id and ai.status = '3' and ai.uid = '"+username+"' order by ai.id desc") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			return AvailableTestResults, false
		 }
		for results.Next() {
			AvailableTestResult := GetDisplayResutlsParams{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTestResult.ID, &AvailableTestResult.UID, &AvailableTestResult.TestScore, &AvailableTestResult.TestID, &AvailableTestResult.Course, &AvailableTestResult.TakenOn,&AvailableTestResult.TestType, &AvailableTestResult.CreatedBy,&AvailableTestResult.Recommendation)
			if err != nil {
				//panic(err.Error())
				fmt.Printf("%+v",err )
			} else {
					AvailableTestResults = append(AvailableTestResults, AvailableTestResult)
					//fmt.Printf("%+v", AvailableTestResults )
					//fmt.Printf("%+v", AvailableTestResult )
			}
		}
//		fmt.Printf("%+v", AvailableTestResults )
//		fmt.Printf("end")
	return AvailableTestResults, true
}

func GetConcertoResults(session_id string) ([]ConcertoResults, bool) {
	 
	AvailableTestResults := []ConcertoResults{}

	
	results, err := ConnC.Query("SELECT id,name, value, session_id FROM concerto.default_data_table where session_id='"+session_id+"'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			return AvailableTestResults, false
		 }

		 fmt.Println("no errors")
//		 fmt.Printf("%+v", results)
		for results.Next() {
			//fmt.Printf("kingsley")

			AvailableTestResult := ConcertoResults{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTestResult.ID, &AvailableTestResult.Name, &AvailableTestResult.Value, &AvailableTestResult.Session_id)
			if err != nil {
				//panic(err.Error())
			} else {
				AvailableTestResults = append(AvailableTestResults, AvailableTestResult)
			}
		}
	return AvailableTestResults, true
}


func GetConcertoResultsSM(session_id string) ([]ConcertoResultsSM, bool) {
	 
		AvailableTestResults := []ConcertoResultsSM{}

	
		results, err := ConnC.Query("SELECT item_id,response, time_taken, session_id, correct FROM concerto.default_cat_response_table where session_id='"+session_id+"'") //, testid, testname, numques, duration, owner FROM featuredtests" )
		 if err != nil {
			return AvailableTestResults, false
		 }

		 fmt.Println("no errors")
//		 fmt.Printf("%+v", results)
		for results.Next() {
			//fmt.Printf("kingsley")

			AvailableTestResult := ConcertoResultsSM{}
			// for each row, scan the result into our tag composite object
			err = results.Scan(&AvailableTestResult.Item_ID, &AvailableTestResult.Response, &AvailableTestResult.Time_Taken, &AvailableTestResult.Session_id, &AvailableTestResult.Correct)
			if err != nil {
				//panic(err.Error())
			} else {
				//if (AvailableTestResult.name = ) {

				//}
				//fmt.Printf("%+v", AvailableTestResult)
				AvailableTestResults = append(AvailableTestResults, AvailableTestResult)
				
			}
		}


	return AvailableTestResults, true

}

func GetTestResultsSM(session_id string) (TestResultsParamsSM, bool) {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	//AvailableTestResults := []TestResultsParams{}

	smScores := TestResultsParamsSM{} //make(map[string]interface{})
	
		allResults, boolval := GetConcertoResultsSM(session_id)
		counter := 0
		//total := 0
		for _, d := range allResults {

        	switch d.Correct {
        		case 1:
					//fmt.Printf(d.Name)
					smScores.Correct = smScores.Correct + 1
				case 0:
					smScores.Wrong = smScores.Wrong + 1				
        		default:
        	}
        	counter++
        	//total=(total+ConverToInt(d.Value))/16
    	}
    	//psychometricScores.Score = total
    	smScores.Score = smScores.Correct
    	fmt.Println(counter)
    	fmt.Println("end of result pulling")

	return smScores, boolval

}

func GetTestResults(session_id string) (TestResultsParams, bool) {
	 //make(map[string]interface{})
	//var stag = new([]TestParams)
	//AvailableTestResults := []TestResultsParams{}

	psychometricScores := TestResultsParams{} //make(map[string]interface{})
	
		allResults, boolval := GetConcertoResults(session_id)
		counter := 0
		//total := 0
		for _, d := range allResults {

        	switch d.Name {
        		case "aggressive":
					//fmt.Printf(d.Name)
					psychometricScores.Aaq = d.Value
				case "anxiety":
					psychometricScores.Aq = d.Value
					//fmt.Printf(d.Name)
				case "control":
					psychometricScores.Cq = d.Value
					//fmt.Printf(d.Name)
				case "friendliness":
					psychometricScores.Fq = d.Value
					//fmt.Printf(d.Name)
				case "happiness":
					psychometricScores.Hq = d.Value
					//fmt.Printf(d.Name)
				case "impulsiveness":
					psychometricScores.Iq = d.Value
					//fmt.Printf(d.Name)
				case "locus_of_control":
					psychometricScores.Lc = d.Value
					//fmt.Printf(d.Name)
				case "passive":
					psychometricScores.Paq = d.Value
					//fmt.Printf(d.Name)
				case "rational_emotive":
					psychometricScores.Rq = d.Value
					//fmt.Printf(d.Name)
				case "self_esteem":
					psychometricScores.Seq = d.Value
					//fmt.Printf(d.Name)
				case "stress":
					psychometricScores.Sq = d.Value
					//fmt.Printf(d.Name)
				case "well_being":
					psychometricScores.Bwq = d.Value
					//fmt.Printf(d.Name)
        		default:
        	}
        	counter++
        	//total=(total+ConverToInt(d.Value))/16
    	}
    	//psychometricScores.Score = total
    	fmt.Println(counter)
    	fmt.Println("end of result pulling")

		
		//psychometricScores["i1q"] = allResults.



	return psychometricScores, boolval

}

func GetJSON(userid string) (string, error) {
	db := Conn
	sqlString := "select i1q,i2q,req,gwq,seq,lc,paq,aaq,rq,sq,iq,fq,cq,hq,aq,bwq from accessmentinfo where uid = '"+userid+"' and status = '3'"
  rows, err := db.Query(sqlString)
  if err != nil {
      return "", err
  }
  defer rows.Close()
  columns, err := rows.Columns()
  if err != nil {
      return "", err
  }
  count := len(columns)
  tableData := make([]map[string]interface{}, 0)
  values := make([]interface{}, count)
  valuePtrs := make([]interface{}, count)
  for rows.Next() {
      for i := 0; i < count; i++ {
          valuePtrs[i] = &values[i]
      }
      rows.Scan(valuePtrs...)
      entry := make(map[string]interface{})
      for i, col := range columns {
          var v interface{}
          val := values[i]
          b, ok := val.([]byte)
          if ok {
              v = string(b)
          } else {
              v = val
          }
          entry[col] = v
      }
      tableData = append(tableData, entry)
  }
  jsonData, err := json.Marshal(tableData)
  if err != nil {
      return "", err
  }
  fmt.Println(string(jsonData))
  return string(jsonData), nil 
}

func WriteTestResultsToDBSM(results TestResultsParamsSM) bool {

	fmt.Printf("writing results to db")
	//fmt.Printf("%+v",results)
	// if (len(results) == 0) {
	// 	return false
	// }
	returnresults := saveSMResults(results)
	if (!returnresults) {
		return false

	}

	return true
}

func WriteTestResultsToDB(results TestResultsParams) bool {

	fmt.Printf("writing results to db")
	//fmt.Printf("%+v",results)
	// if (len(results) == 0) {
	// 	return false
	// }
	returnresults := savePsyResults(results)
	if (!returnresults) {
		return false

	}

	return true
}

func Containsn(slice []string, testid string, testname string, testtype string) bool {
    set := make(map[string]struct{}, len(slice))
    //c_date := time.Now().Unix()
    //end_date := time.Now().Unix()
    status := "1"
    items := ""
    for _, s := range slice {
        set[s] = struct{}{}

        
		sv := strings.Split(s, ":")
    	_, username, candname := sv[0], sv[1], sv[2]

		//sb := strconv.FormatInt(c_date, 10)
		fmt.Printf("start bugging")
    	row := Conn.QueryRow("select count(id) as cot from GPI_Oauth_clients.accessmentinfo where status = '1' and testid='" + testid +
		"' and uid ='" + username + "' and  testtype ='" + testtype + "'")
		var cot int
		row.Scan(&cot)
		log.Printf("CODE: ")
		log.Printf("%+v", cot)

		if (cot == 0) {

        	Conn.Exec("insert into GPI_Oauth_clients.accessmentinfo (uid, status, testid, candname, testtype, testname)" +
				" values ( '" + username + "', '" +status + "', '" + testid + "', '" +
			candname + "','" + testtype + "','" + testname + "')")
       }
    }
    _, ok := set[items]
    return ok
}

func Containsnew(slice []string, testid string, testname string, testtype string) (string, bool) {
	var cot string
    var boolval = false
    set := make(map[string]struct{}, len(slice))
    //c_date := time.Now().Unix()
    //end_date := time.Now().Unix()
    status := "1"
    items := ""
    for _, s := range slice {
        set[s] = struct{}{}

        
		sv := strings.Split(s, ":")
    	thistesttoken, username, candname := sv[0], sv[1], sv[2]
    	cot=thistesttoken
		//sb := strconv.FormatInt(c_date, 10)
		fmt.Printf("start bugging")
    	row := Conn.QueryRow("select id as cot from GPI_Oauth_clients.accessmentinfo where status = '1' and testid='" + testid +
		"' and uid ='" + username + "' and  testtype ='" + testtype + "'")
		var cots string
		row.Scan(&cots)
		log.Printf("CODE: ")
		log.Printf("%+v", cots)

		if (len(cots) == 0) {

        	Conn.Exec("insert into GPI_Oauth_clients.accessmentinfo (uid, status, testid, candname, testtype, testtoken, testname)" +
				" values ( '" + username + "', '" +status + "', '" + testid + "', '" +
			candname + "','" + testtype + "','" + thistesttoken + "','" + testname + "')")
			cot=  thistesttoken
        	boolval= true
       } else {
       		cot=  cots
        	boolval= false
       }
    }
    _, _ = set[items]
    return cot, boolval
}

func Contains(slice []string, item string, username string, uid string) (string, bool) {
    set := make(map[string]struct{}, len(slice))
    var cot string
    var boolval = false
    //c_date := time.Now().Unix()
    //end_date := time.Now().Unix()
    status := "1"
    items := ""

//fmt.Printf("%+v",slice)

    for _, s := range slice {
        set[s] = struct{}{}
        //fmt.Printf("%+v\n", set[s])
        //fmt.Printf(s)
        
        service_id := s

        
        //data.Conn.Exec("insert into GPI_Oauth_clients.subscription_history ( user_id, start_date, end_date, service_id, status)" +
		//	" values ( '" + uid + "', '" + string(start_date) + "', '" + string(end_date)+ "', '" +
		//	service_id + "', '" + status + "')")
sv := strings.Split(s, ":")
fmt.Println("%+v", sv)
    service_id, testname, testtype, this_test_token := sv[0], sv[1], sv[2], sv[3]
    cot=this_test_token
  row := Conn.QueryRow("select testtoken as cots from GPI_Oauth_clients.accessmentinfo where status = '1' and testid='" + service_id +
		"' and uid ='" + username + "' and  testtype ='" + testtype + "'")
	var cots string
	row.Scan(&cots)
	fmt.Printf("the code = ")
	fmt.Println("%+v", cot)
	fmt.Println(len(cot))

		if (len(cots) == 0) {

        Conn.Exec("insert into GPI_Oauth_clients.accessmentinfo (uid, status, testid, candname, testtype, testtoken, testname)" +
			" values ( '" + username + "', '" +status + "', '" + service_id + "', '" +
			item + "','" + testtype + "','" + this_test_token + "','" + testname + "')")
			//return cot, true
			cot = this_test_token
			boolval = true
        } else {
        	cot=  cots
        	boolval= false
        }
    }
    _, _ = set[items]
    return cot, boolval
}

func saveSMResults(results TestResultsParamsSM) bool {
	if (len(results.Usersnames) == 0) {
		return false
	}
	row := Conn.QueryRow("select count(id) as cot from GPI_Oauth_clients.accessmentinfo where status = '1' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + "' and testtype ='" + results.TestType + "'")
	var cot int
	row.Scan(&cot)
	if (cot == 1) {
		//fmt.Printf("kingsmanley")
		Conn.Exec("update GPI_Oauth_clients.accessmentinfo set i1q='"+strconv.Itoa(results.Correct)+"', i2q='"+strconv.Itoa(results.Wrong)+
			"', uid='" + results.UserID + 
			"', status='3', testid='" + results.TestID + 
			"', candname ='" + results.Usersnames + "', testscore='" + strconv.Itoa(results.Score) + 
			"', testtoken ='" + results.TestToken + "', testtype ='" + results.TestType + 
			"', testname ='" + results.TestName + 
			"' where status = '"+"1"+"' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + 
		"' and testtype ='" + results.TestType + "'")
	} else {
		row := Conn.QueryRow("select count(id) as cot from GPI_Oauth_clients.accessmentinfo where status = '3' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + "' and testtype ='" + results.TestType + "'")
	var cot int
	row.Scan(&cot)

		if (cot == 0) {
			Conn.Exec("insert into GPI_Oauth_clients.accessmentinfo ( i1q,i2q,uid, status, testid, candname, testscore, testtoken, testtype, testname)" +
			" values ('"+strconv.Itoa(results.Correct)+"','"+strconv.Itoa(results.Wrong)+"', '" + results.UserID + "', '" +"3" + "', '" + results.TestID + "', '" +
			results.Usersnames + "', '" + strconv.Itoa(results.Score) + "', '" + results.TestToken + "', '" + results.TestType + "','" + results.TestName + "')")
		} else {
			Conn.Exec("update GPI_Oauth_clients.accessmentinfo set i1q='"+strconv.Itoa(results.Correct)+"', i2q='"+strconv.Itoa(results.Wrong) +
				"', uid='" + results.UserID + 
			"', status='3', testid='" + results.TestID + 
			"', candname ='" + results.Usersnames + "', testscore='" + strconv.Itoa(results.Score) + 
			"', testtoken ='" + results.TestToken + "', testtype ='" + results.TestType + 
			"', testname ='" + results.TestName + 
			"' where status = '"+"3"+"' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + 
		"' and testtype ='" + results.TestType + "'")
		
		}
	}
	return true
}


func savePsyResults(results TestResultsParams) bool {
	if (len(results.Usersnames) == 0) {
		return false
	}
	row := Conn.QueryRow("select count(id) as cot from GPI_Oauth_clients.accessmentinfo where status = '1' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + "' and testtype ='" + results.TestType + "'")
	var cot int
	row.Scan(&cot)
	if (cot == 1) {
		//fmt.Printf("kingsmanley")
		Conn.Exec("update GPI_Oauth_clients.accessmentinfo set i1q='"+results.I1q+"', i2q='"+results.I2q+
			"',req='"+results.Req+"',gwq='"+results.Gwq+"',seq='"+results.Seq+"',lc='"+results.Lc+"',paq='"+
			results.Paq+"',aaq='"+results.Aaq+"',rq='"+results.Rq+"',sq='"+results.Sq+"',iq='"+results.Iq+
			"',fq='"+results.Fq+"',cq='"+results.Cq+"',hq='"+results.Hq+"', aq='" + results.Aq+"',bwq='"+results.Bwq+"', uid='" + results.UserID + 
			"', status='3', testid='" + results.TestID + 
			"', candname ='" + results.Usersnames + "', testscore='" + strconv.Itoa(results.Score) + 
			"', testtoken ='" + results.TestToken + "', testtype ='" + results.TestType + 
			"', testname ='" + results.TestName + 
			"' where status = '"+"1"+"' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + 
		"' and testtype ='" + results.TestType + "'")
	} else {
		row := Conn.QueryRow("select count(id) as cot from GPI_Oauth_clients.accessmentinfo where status = '3' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + "' and testtype ='" + results.TestType + "'")
	var cot int
	row.Scan(&cot)

		if (cot == 0) {
			Conn.Exec("insert into GPI_Oauth_clients.accessmentinfo ( i1q,i2q,req,gwq,seq,lc,paq,aaq,rq,sq,iq,fq,cq,hq,aq,bwq,uid, status, testid, candname, testscore, testtoken, testtype, testname)" +
			" values ('"+results.I1q+"','"+results.I2q+"','"+results.Req+"','"+results.Gwq+"','"+results.Seq+"','"+results.Lc+"','"+results.Paq+"','"+results.Aaq+"','"+results.Rq+"','"+results.Sq+"','"+results.Iq+"','"+results.Fq+"','"+results.Cq+"','"+results.Hq+"','"+results.Aq+"','"+results.Bwq+"', '" + results.UserID + "', '" +"3" + "', '" + results.TestID + "', '" +
			results.Usersnames + "', '" + strconv.Itoa(results.Score) + "', '" + results.TestToken + "', '" + results.TestType + "','" + results.TestName + "')")
		} else {
			Conn.Exec("update GPI_Oauth_clients.accessmentinfo set i1q='"+results.I1q+"', i2q='"+results.I2q+
			"',req='"+results.Req+"',gwq='"+results.Gwq+"',seq='"+results.Seq+"',lc='"+results.Lc+"',paq='"+
			results.Paq+"',aaq='"+results.Aaq+"',rq='"+results.Rq+"',sq='"+results.Sq+"',iq='"+results.Iq+
			"',fq='"+results.Fq+"',cq='"+results.Cq+"',hq='"+results.Hq+"', aq='" + results.Aq+"',bwq='"+results.Bwq+"', uid='" + results.UserID + 
			"', status='3', testid='" + results.TestID + 
			"', candname ='" + results.Usersnames + "', testscore='" + strconv.Itoa(results.Score) + 
			"', testtoken ='" + results.TestToken + "', testtype ='" + results.TestType + 
			"', testname ='" + results.TestName + 
			"' where status = '"+"3"+"' and testid='" + results.TestID +
		"' and uid ='" + results.UserID + "' and testtoken ='" + results.TestToken + 
		"' and testtype ='" + results.TestType + "'")
		
		}
	}
	return true
}



func ConverToInt(value string) int {
	i, err := strconv.Atoi(value)
	if (err == nil) {
		return i
	} 
	return 0
}
func ReformResults(results TestResultsParams) TestResultsParamsR {
	returnResult := TestResultsParamsR{}
	returnResult.I1q = (ConverToInt(results.Seq) + ConverToInt(results.Bwq) + ConverToInt(results.Lc))/3
	returnResult.I1qT = getInterpretation("I1q", returnResult.I1q) 

	returnResult.I2q = (ConverToInt(results.Aaq) + ConverToInt(results.Paq) + ConverToInt(results.Fq) + ConverToInt(results.Cq))/4
	returnResult.I2qT = getInterpretation("I2q", returnResult.I2q) 

	returnResult.Req = (ConverToInt(results.Rq) + ConverToInt(results.Iq))/2
	returnResult.ReqT = getInterpretation("Req", returnResult.Req) 
	
	returnResult.Gwq = (ConverToInt(results.Sq) + ConverToInt(results.Hq) + ConverToInt(results.Aq))/3
	returnResult.QwqT = getInterpretation("Gwq", returnResult.Gwq)

	returnResult.Seq = ConverToInt(results.Seq)
	returnResult.SeqT = getInterpretation("Seq", returnResult.Seq)

	returnResult.Lc = ConverToInt(results.Lc) 
	returnResult.LcT = getInterpretation("Lc", returnResult.Lc)

	returnResult.Paq = ConverToInt(results.Paq) 
	returnResult.PaqT = getInterpretation("Paq", returnResult.Paq)

	returnResult.Aaq = ConverToInt(results.Aaq)
	returnResult.AaqT = getInterpretation("Aaq", returnResult.Aaq)

	returnResult.Rq = ConverToInt(results.Rq)
	returnResult.RqT = getInterpretation("Rq", returnResult.Rq)

	returnResult.Sq = ConverToInt(results.Sq)
	returnResult.SqT = getInterpretation("Sq", returnResult.Sq)

	returnResult.Iq = ConverToInt(results.Iq) 
	returnResult.IqT = getInterpretation("Iq", returnResult.Iq)

	returnResult.Fq = ConverToInt(results.Fq) 
	returnResult.FqT = getInterpretation("Fq", returnResult.Fq)

	returnResult.Cq = ConverToInt(results.Cq)
	returnResult.CqT = getInterpretation("Cq", returnResult.Cq)

	returnResult.Hq = ConverToInt(results.Hq)
	returnResult.HqT = getInterpretation("Hq", returnResult.Hq)

	returnResult.Aq = ConverToInt(results.Aq)
	returnResult.AqT = getInterpretation("Aq", returnResult.Aq)

	returnResult.Bwq = ConverToInt(results.Bwq)
	returnResult.BwqT = getInterpretation("Bwq", returnResult.Bwq)

	returnResult.Score = results.Score
	returnResult.ScoreT = getInterpretation("", returnResult.Score)

	return returnResult
}

func getInterpretation(id string, value int) string{
	if (id == "I1q") {
		log.Printf("i1q taken")
		if (value >= 0 && value <= 9 ) {
			return "Your scores show that there is a big gap between the degree of Intrapersonal Quotient that is present and that is desirable. Your response suggests that you do not, at all, percieve yourself as being confident or deserving. Therefore, a lot of effort is needed to develop a more confident self."
		} else if (value >= 10 && value <= 19) {
			return "Your reponse suggests that your confidence levels and sense of selfesteem lies in the median range. So while it is good that you have a fair amount of confidence and self-direction, it will be good if some more work is put in to develop a stronger sense of self."
		} else if (value >= 20 && value <= 40) {
			return "Your response suggests that the sense of self is adequate. Confidence and self-direction help you towards achieving most of your goals. It will be good if you maintain and build on the worldview and life process that have created this sense of self."
		}
	}

	if (id == "I2q") {
		log.Printf("i2q taken")
		if (value >= 0 && value <= 9 ) {
			return "A noticeable gap exists between the degree of Interpersonal Quotient that is present and that is desirable. Your responses suggest that you face significant difficulty in your work and interactions with others. A fair amount of effort needs to be put in by you alone to increase the degree of efficiency of your relationship with others."
		} else if (value >= 10 && value <= 19) {
			return "Your reponses suggest that your Interpersonal Quotient scores and hence your interpersonal skills are in the median range. While on an average, you work well with others; it is suggested you take some time to improve your interpersonal skills considering the enormous importance of these skills."
		} else if (value >= 20 && value <= 40) {
			return "Your reponses suggest that your interpersonal skill is adequate. And this capability will make you come across as amicable to others."
		}
	}

	if (id == "Req") {
		log.Printf("req taken")
		if (value >= 0 && value <=9 ) {
			return "The scores show that there is a big gap between the degree of Rational Emotive Quotient that is present and that is desirable. Your responses suggest that you are completely unable to manage your emotions in a way that leads to an effective life. Therefore, a tremendous amount of effort needs to put in to develop the ability to manage emotions and act rationally in given situations."
		} else if (value >= 10 && value <=16) {
			return "A significant gap exists between the degree of Rational Emotive Quotient that is present and that is desirable. Your responses suggest that in many situations you lack the ability to formulate emotional responses that are rational. Therefore a fair amount of attention needs to be paid to develop this ability."
		} else if (value >= 17 && value <=40) {
			if (value >= 17 && value <=23) {
				return "Your responses suggest that your ability to manage your emotions lies in the median range. So while it is good that you have a fair amount of rationality present in your emotional responses, it will be even better if you put in more efforts to develop a greater rational emotive capacity."
			}
			return "Your responses suggest that you have an adequate amount of rational emotiveness. This ability to manage your emotions well will help you towards achieving your goals. It will be good if you maintain and build on the worldview and life process that have created this ability."
		}
	}

	if (id == "Gwq") {
		log.Printf("req taken")
		if (value >= 0 && value <=9 ) {
			if (value >= 0 && value <=5) {
				return "Your scores show that there is a big gap between the degree of the feeling of General Well Being that is present and that is desirable. Your responses suggest that your feeling of general well-being is not satisfactory at all. Therefore, a lot of effort needs to be put in to improve the situation."
			}
			return "A significant gap exists between the degree of feeling of General Well Being that is present and that is desirable. Your responses suggest that you are under a fair bit of stress and desperately seek some improved quality of life. Therefore there is the need you learn to cope with the inevitable stresses of daily life, adapt and manage in times of change and uncertainty and also learn to build and maintain good relationships with others"
		} else if (value >= 10 && value <=16) {
			return "Your responses suggest that your general well-being lies in the median range. So while it is good that you feel fairly well in your life, it will be good if some more work is put in to develop a stronger sense of well-being."
		} else if (value >= 17 && value <=40) {
			if (value >= 17 && value <=23) {
				return "Your reponses suggests that your feeling of general well-being is fairly adequate. So while it is good that you feel fairly well in your life, some more work has to be put in in order to create a stronger sense of well-being."
			}
			return "Your responses suggest that your sense of well-being is adequate. A general feeling of happiness and good-humor will help you towards achieving your goals. It will be good if you do maintain and build on the worldview and life process that have created this sense of well-being."
		}
	}

	return "We are processing your result score for interpretation."
}

func getInterpretationbn(id string, value int) string{
	if (id == "I1q") {
		log.Printf("i1q taken")
		if (value >= 0 && value <= 9 ) {
			return "The scores show that there is a big gap between the degree of Intrapersonal Quotient that is present and that is desirable. The individual's response suggests that they don't at all percieve themselves as confident, deserving or achievers. Therefore, a lot of effort is needed to develop a more confident self."
		} else if (value >= 10 && value <= 19) {
			return "The individual's reponse suggests that their confidence levels and sense of selfesteem lies in the median range. So while it is good that they have a fair amount of confidence and self-direction, it will be good if some more work is put in to develop a stronger sense of self."
		} else if (value >= 20 && value <= 40) {
			return "The individual's response suggests that the sense of self is adequate. Confidence and self-direction help them towards achieving their goals. It will be good if the person maintains and builds on the worldview and life process that have created this sense of self."
		}
	}

	if (id == "I2q") {
		log.Printf("i2q taken")
		if (value >= 0 && value <= 9 ) {
			return "A noticeable gap exists between the degree of Interpersonal Quotient that is present and that is desirable. The individuals' responses suggest that they face significant difficulty in how they interact and work with others. A fair amount of effort needs to be put in by the individuals to increase the degree of efficiency of their interactions with others."
		} else if (value >= 10 && value <= 19) {
			return "The individuals' reponses suggests that their Interpersonal Quotient scores and hence their interpersonal skills are in the median range. While on an average, they work well with others, it is suggested they take some time to improve their interpersonal skills considering the enormous importance of these skills."
		} else if (value >= 20 && value <= 40) {
			return "The individuals' reponses suggests that their interpersonal skills is adequate. And this capability will make them come across as amicable to others."
		}
	}

	if (id == "Req") {
		log.Printf("req taken")
		if (value >= 0 && value <=9 ) {
			return "The scores show that there is a big gap between the degree of Rational Emotive Quotient that is present and that is desirable. The individual's response suggests that they they are completely unable to manage their emotions in a way that leads to an effective life. Therefore, a tremendous amount of effort needs to put in to develop the ability to manage emotions and act in situations rationally."
		} else if (value >= 10 && value <=16) {
			return "A significant gap exists between the degree of REQ that is present and that is desirable. The individual's response suggests that in many situations they lack the ability to formulate emotional responses that are rational. Therefore a fair amount of attention needs to be paid to develop this ability."
		} else if (value >= 17 && value <=40) {
			if (value >= 17 && value <=23) {
				return "The individual's response suggests that their ability to manage their emotions lies in the median range. So while it is good that they have a fair amount of rationality present in their emotional responses, it will be even better if some more work is put in to develop a greater rational emotive capacity."
			}
			return "The REQ suggests that they have an adequate amount of rational emotiveness. This ability to manage their emotions well will help them towards achieving their goals. It will be good if the person maintains and builds on the worldview and life process that have created this ability."
		}
	}

	if (id == "Gwq") {
		log.Printf("req taken")
		if (value >= 0 && value <=9 ) {
			if (value >= 0 && value <=5) {
				return "The scores show that there is a big gap between the degree of GWQ that is present and that is desirable. The individual's response suggests that their general well-being is not satisfactory at all. Therefore, a lot of effort needs to be put in to improve the situation."
			}
			return "A significant gap exists between the degree of GWQ that is present and that is desirable. The individual's response suggests they are under a fair bit of stress and are unhappy with their lives. Therefore a fair amount of attention needs to be paid to improve the situation."
		} else if (value >= 10 && value <=16) {
			return "The individual's reponse suggests that their general well-being lies in the median range. So while it is good that they feel fairly well in their lives, it will be good if some more work is put in to develop a stronger sense of well-being."
		} else if (value >= 17 && value <=40) {
			if (value >= 17 && value <=23) {
				return "The individual's reponse suggests that their general well-being is fairly adequate. So while it is good that they feel fairly well in their lives, some more work has to be put in in order to create a stronger sense of well-being."
			}
			return "The individual's response suggests that the sense of well-being is adequate. A general feeling of happiness and good-humor will help them towards achieving their goals. It will be good if the person maintains and builds on the worldview and life process that have created this sense of well-being."
		}
	}

	return "We are processing your result score for interpretation."
}

func getInterpretationI1q(id string, value int) string{
	return ""
//	return ""
}

func SaveManagedTests(user_id string, companytag string, testname string, duration string, totalquestions string, passscore string, subjectmatter string, instructions string, testquestionfile string, created int64, testslug string, companyid string) bool {	
//params_out := make(map[string]interface{})
	var params_out bool
	//results, err := Conn.pre("insert into managedtests ( user_id, companytag, testname, duration, totalquestions, passscore, subjectmatter, instructions, testquestionfile, created)" + " values ( '" + user_id + "', '" + companytag + "', '" + testname + "', '" + duration + "', '" + totalquestions + "', '" + passscore + "','" + subjectmatter + "','" + instructions + "','" + testquestionfile + "','" + string(time.Now().Unix()) + "')")
	

	stmt, _ := Conn.Prepare("insert into managedtests ( user_id, companytag, testname,duration,totalquestions,passscore, subjectmatter, instructions, testslug, companyid) values(?,?,?,?,?,?,?,?,?,?)")
	//, companytag, testname, duration, totalquestions, passscore, subjectmatter, instructions, testquestionfile, created) values(?,?,?,?,?,?,?,?,?,?)")
        

        res, _ := stmt.Exec(user_id,companytag,testname,duration,totalquestions,passscore, subjectmatter, instructions, testslug, companyid)
        //, "companytag", "testname", "duration", "totalquestions", "passscore", "subjectmatter", "instructions", "testquestionfile", "created")
        

        _, err := res.LastInsertId()
        



	if err != nil {
		params_out = false //[]byte(`{"Status":"Error","Data":`+"Error"+`}`)
	}
	params_out = true //[]byte(`{"Status":"Success","Data":`+string(id)+`}`)
	return params_out
}
