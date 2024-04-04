package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/fabiov3105/westminsteripvc_front/drivers/database"
	"github.com/fabiov3105/westminsteripvc_front/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/peoplesimple", peoplesimple)
	http.HandleFunc("/newpeoplesimple", newpeoplesimple)
	http.ListenAndServe(":8001", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

type People struct {
	Id              int
	Name            string
	FinancialNumber int
}

func peoplesimple(w http.ResponseWriter, r *http.Request) {
	database.ConectDB()
	var p []People
	database.DB.Limit(5).Order("name asc").Find(&p)

	temp.ExecuteTemplate(w, "PeopleSimple", p)
}

func newpeoplesimple(w http.ResponseWriter, r *http.Request) {
	database.ConectDB()
	var p []People
	database.DB.Select("financial_number+1 as financial_number").Limit(1).Order("financial_number desc").Find(&p)

	temp.ExecuteTemplate(w, "NewPeopleSimple", p)
}

func InsertNewPeopleSimple(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		financialnumber := r.FormValue("financialnumber")

		financialnumberConver, err := strconv.Atoi(financialnumber)
		if err != nil {
			log.Println("Erro na conversao do Financial Number", err)
		}

		models.CreateNewPeopleSimple(name, financialnumberConver)

		http.Redirect(w, r, "/newpeoplesimple", 301)

	}
}
