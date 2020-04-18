package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"regexp"
	"sso/data"
	"sso/model"
	"sso/util"
	"log"
	"strings"
	"fmt"
)

func DoLogin(r render.Render, res http.ResponseWriter, req *http.Request) {
	log.Println("in Login")
	//parse request parameters
	req.ParseForm()
	service := req.FormValue("service")
	username := req.FormValue("username")
	password := req.FormValue("password")

	log.Println("parsed params:service:" + service + ",username:" + username + ",password:" + password)
	fmt.Println("parsed params:service:" + service + ",username:" + username + ",password:" + password)

	//TODO Handle User Auth

	tgt := data.GrantTicketGrantingTicket(username, "")
	if tgt == nil {
		log.Fatalln("error grant tgt")
		fmt.Println("error grant tgt")
		r.HTML(200, "error", "")
		return
	}
	log.Println("tgt granted:" + tgt.Tgt)
	fmt.Println("tgt granted:" + tgt.Tgt)

	cookie := http.Cookie{Name: "CASTGC", Value: tgt.Tgt, Path: "/", Domain: util.COOKIE_DOMAIN, MaxAge: util.TICKET_GRANTING_TICKET_TIME_TO_LIVE}
	http.SetCookie(res, &cookie)

	st := data.GrantServiceTicket(tgt.Tgt, service)
	log.Println("st granted:" + st.St)
	fmt.Println("st granted:" + st.St)
	acookie := http.Cookie{Name: "GPIGC", Value: username, Path: "/", Domain: "127.0.0.1", MaxAge: util.TICKET_GRANTING_TICKET_TIME_TO_LIVE}
	http.SetCookie(res, &acookie)
	if st == nil {
		log.Fatalln("error grant st")
		fmt.Println("error grant st")
		r.HTML(200, "error", "")
		return
	}
	data.AddSTToTGT(tgt, st)
	fmt.Println("redirecting to... service")
	redirectToServiced(r, st, username)
	return

}

func validateServiced(service string) bool {
	if service == "" {
		return true
	}
	reg := regexp.MustCompile(`^(https|http)://.*`)
	return reg.MatchString(service)
}

func redirectToServiced(r render.Render, st *model.ServiceTicket, name string) {
	needAnd := strings.Contains(st.Service, "?")
	sep := "?"
	if needAnd {
		sep = "&"
	}
	redirectUrl := st.Service + sep + "ticket=" + st.St + "&parsedobj="+name
	log.Println("redirect to service:" + redirectUrl)
	fmt.Println("redirect to service:" + redirectUrl)
	r.Redirect(redirectUrl)
	//r.Redirect("callRoute/"+redirectUrl)
}
