package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	// "regexp"
	// "sso/data"
	// "sso/model"
	// //"sso/util"
	// "log"
	// "strings"
	// "fmt"
)

func Register(r render.Render, res http.ResponseWriter, req *http.Request) {
	r.HTML(200, "register", "gpitest")
	return
}