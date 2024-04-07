package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/fabiov3105/westminsteripvc_front/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Peoplesimple(w http.ResponseWriter, r *http.Request) {
	AllPeople := models.SelectPeopleSimple()
	temp.ExecuteTemplate(w, "PeopleSimple", AllPeople)
}

func Peoplesimpleenabled(w http.ResponseWriter, r *http.Request) {
	AllPeople := models.SelectPeopleSimpleenabled()
	temp.ExecuteTemplate(w, "PeopleSimple", AllPeople)
}

func Peoplesimpledisabled(w http.ResponseWriter, r *http.Request) {
	AllPeople := models.SelectPeopleSimpledisabled()
	temp.ExecuteTemplate(w, "PeopleSimple", AllPeople)
}

func Newpeoplesimple(w http.ResponseWriter, r *http.Request) {
	OnePeople := models.SelectOnePeopleSimple()
	temp.ExecuteTemplate(w, "NewPeopleSimple", OnePeople)
}

func Updatepeoplesimple(w http.ResponseWriter, r *http.Request) {
	IdPeople := r.URL.Query().Get("id")
	OnePeople :=  models.SelectOnePeopleSimpleUpdate(IdPeople)
	temp.ExecuteTemplate(w, "UpdatePeopleSimple", OnePeople)
}

func Updatepeoplesimpleexec(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		financialnumber := r.FormValue("financialnumber")
		location := r.FormValue("location")

		idConver, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do Id Number", err)
		}
		financialnumberConver, err := strconv.Atoi(financialnumber)
		if err != nil {
			log.Println("Erro na conversao do Financial Number", err)
		}

		models.UpdatePeopleSimpleExec(idConver, financialnumberConver, name, location)
	}
	http.Redirect(w, r, "/peoplesimple", 301)
}

func InsertNewPeopleSimple(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		financialnumber := r.FormValue("financialnumber")
		location := r.FormValue("location")

		financialnumberConver, err := strconv.Atoi(financialnumber)
		if err != nil {
			log.Println("Erro na conversao do Financial Number", err)
		}
		models.CreateNewPeopleSimple(name, financialnumberConver, location)
		http.Redirect(w, r, "/peoplesimple", 301)
	}
}

func Deletepeoplesimple(w http.ResponseWriter, r *http.Request) {
	IdPeople := r.URL.Query().Get("id")
	models.DeletaPeople(IdPeople)
	http.Redirect(w, r, "/peoplesimple", 301)
}

func DisabledPeopleSimple(w http.ResponseWriter, r *http.Request) {
	IdPeople := r.URL.Query().Get("id")
	models.DisabledPeople(IdPeople)
	http.Redirect(w, r, "/peoplesimple", 301)
}

func EnabledPeopleSimple(w http.ResponseWriter, r *http.Request) {
	IdPeople := r.URL.Query().Get("id")
	models.EnabledPeople(IdPeople)
	http.Redirect(w, r, "/peoplesimple", 301)
}
