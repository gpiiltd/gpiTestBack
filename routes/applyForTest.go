package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"gpitest/data"
	//"gpitest/model"
	//"strings"
	//"log"
)

func ApplyForTest(r render.Render, res http.ResponseWriter, req *http.Request) {
	AT := data.ApplyForTest()
	
	r.JSON(200, AT)
	return
}