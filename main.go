package main

import (
	//"net/http"
	"os"
	"log"
	//"fmt"
	"github.com/martini-contrib/cors"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gpitest/routes"
	"gpitest/util"
)

func init() {
	file, err := os.Create(util.LOG_FILE)
	if err != nil {
		log.Println("error create logFile")
		return
	} else {
		log.SetOutput(file)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("gpitest:")

}

func main() {
	/*http.HandleFunc("/test.png", serveimage)
	func serveimage(w http.ResponseWriter, r *http.Request) {
     http.ServeFile(w, r, "test.png")
	}*/
	//fs := http.FileServer(http.Dir("templates"))
	//http.Handle("/templates", http.StripPrefix("/templates/", fs))
  	//http.Handle("/", fs)
  	//m.Get("/templates",fs)
  	//m.Get("/templates", http.StripPrefix("/templates/", fs))
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		//IndentJSON: true,                       // Output human readable JSON
		//IndentXML:  true,                       // Output human readable XML
	}))
	m.Use(martini.Static("templates"))

	m.Use(cors.Allow(&cors.Options{
	    AllowOrigins:     []string{"*"},
	    AllowMethods:     []string{"GET"},
	    AllowHeaders:     []string{"Origin, Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With"},
	    ExposeHeaders:    []string{"Content-Length, Content-Type"},
	    AllowCredentials: false,
  	}))
	//m.Use(martini.Static("publicam"))

	//m.Get("/", func() string {
    //return "Hello world!"
  //})
  //http.Handle("/", m)*/
  	//fs := http.FileServer(http.Dir("template"))
	//m.Handlers(http.Handle("/templates", http.StripPrefix("/templates/", fs)))
	/*m.Get("/template", func (w http.ResponseWriter, r *http.Request) {
     http.Handle("/template", http.StripPrefix("/template/", http.FileServer(http.Dir("template")))) //http.FileServer(http.Dir("template")) //http.ServeFile(w, r, "template")
	})*/ 
	///sso/:client/:id
	m.Get("/createPDF", routes.CreatePDF)
	//m.Get("/callRoute/:client/:id", routes.CallRoute)
	m.Post("/saveMembersForTest", routes.SaveMembersForTest)
	m.Post("/saveTestsForMember", routes.SaveTestsForMember)
	m.Get("/getThisTestResults", routes.GetThisTestResults)
	m.Get("/getAllTestResults", routes.GetAllTestResults)
	m.Get("/getSubjectMatterCands", routes.GetSubjectMatterCands)
	m.Get("/getSubjectMatter", routes.GetSubjectMatter)
	m.Get("/getSubjectMatterTest", routes.GetSubjectMatterTest)
	m.Get("/getTestCandidates", routes.GetTestCandidates)
	m.Get("/getEnterpriseMembers", routes.GetEnterpriseMembers)
	m.Get("/getEnterpriseTests", routes.GetEnterpriseTests)
	m.Post("/saveManagedTests", routes.SaveManagedTests)
	m.Get("/dashboardBreakingFigures", routes.DashboardBreakingFigures)
	m.Get("/featuredTests", routes.FeaturedTests)
	// m.Post("/doRegister", routes.DoRegister)
	m.Get("/listTests", routes.ListTests)
	m.Get("/listMyTests", routes.ListMyTests)
	// m.Post("/doLogin", routes.DoLogin)
	m.Get("/listTestResults", routes.ListTestResults)
	m.Get("/prepareTestResults", routes.PrepareTestResults)
	//m.Run()
	m.RunOnAddr("178.128.251.254:5800")
	
	//log.Println("Listening...")
  	//http.ListenAndServe(":5000", nil)
}
