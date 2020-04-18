package routes

import (
  "github.com/martini-contrib/render"
  //"github.com/go-martini/martini"
  "net/http"
  "fmt"
  "log"
  "strings"
  wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)  


func CreatePDF(r render.Render, res http.ResponseWriter, req *http.Request) {
  //parse request parameters
  //req.ParseForm()

  pdfg, err :=  wkhtml.NewPDFGenerator()
  
  if err != nil{
    r.JSON(200, "{'status' : 'Failed To Create The PDF Instance.'}")
    return
  }
  
  htmlStr := `<html><body><h1 style="color:red;">This is an html
  from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
  code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

  pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))


  // Create PDF document in internal buffer
  err = pdfg.Create()
  if err != nil {
    log.Fatal(err)
    r.JSON(200, "{'status' : 'Failed To Create File'}")
    return
  }

  //Your Pdf Name
  err = pdfg.WriteFile("./Your_pdfname.pdf")
  if err != nil {
    log.Fatal(err)
    r.JSON(200, "{'status' : 'Failed To Write File'}")
    return
  }

  fmt.Println("Done")

  r.JSON(200, "{'status' : 'OK'}")
  return
}