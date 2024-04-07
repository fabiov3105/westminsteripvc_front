package models

import (
	// "fmt"

	"github.com/fabiov3105/westminsteripvc_front/drivers/database"
	_ "github.com/lib/pq"
)

type PeopleSimple struct {
	Id              int
	Name            string
	FinancialNumber int
	Location        string
	Enabled         bool
}

// All People
func SelectPeopleSimple() []PeopleSimple {
	db := database.ConectDB()
	// selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples order by financial_number desc limit 5;")
	selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples order by name asc;")
	if err != nil {
		panic(err.Error())
	}

	p := PeopleSimple{}
	peoples := []PeopleSimple{}

	for selectPeopleSimple.Next() {
		var id, financialNumber int
		var name, location string
		var enabled bool

		err = selectPeopleSimple.Scan(&id, &name, &financialNumber, &location, &enabled)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Id = id
		p.FinancialNumber = financialNumber
		p.Location = location
		p.Enabled = enabled

		peoples = append(peoples, p)
	}
	defer db.Close()
	return peoples
}

// People Enabled
func SelectPeopleSimpleenabled() []PeopleSimple {
	db := database.ConectDB()
	// selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples where enabled='true' order by financial_number desc limit 5;")
	selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples where enabled='true' order by name asc;")
	if err != nil {
		panic(err.Error())
	}

	p := PeopleSimple{}
	peoples := []PeopleSimple{}

	for selectPeopleSimple.Next() {
		var id, financialNumber int
		var name, location string
		var enabled bool

		err = selectPeopleSimple.Scan(&id, &name, &financialNumber, &location, &enabled)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Id = id
		p.FinancialNumber = financialNumber
		p.Location = location
		p.Enabled = enabled

		peoples = append(peoples, p)
	}
	defer db.Close()
	return peoples
}

// People Disabled
func SelectPeopleSimpledisabled() []PeopleSimple {
	db := database.ConectDB()
	// selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples where enabled='false' order by financial_number desc limit 5;")
	selectPeopleSimple, err := db.Query("select id, name, financial_number, location, enabled from peoples where enabled='false' order by name asc;")
	if err != nil {
		panic(err.Error())
	}

	p := PeopleSimple{}
	peoples := []PeopleSimple{}

	for selectPeopleSimple.Next() {
		var id, financialNumber int
		var name, location string
		var enabled bool

		err = selectPeopleSimple.Scan(&id, &name, &financialNumber, &location, &enabled)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Id = id
		p.FinancialNumber = financialNumber
		p.Location = location
		p.Enabled = enabled

		peoples = append(peoples, p)
	}
	defer db.Close()
	return peoples
}

func SelectOnePeopleSimple() []PeopleSimple {
	db := database.ConectDB()

	selectPeopleSimple, err := db.Query("select financial_number+1 from peoples order by financial_number desc limit 1;")
	if err != nil {
		panic(err.Error())
	}

	p := PeopleSimple{}
	peoples := []PeopleSimple{}

	for selectPeopleSimple.Next() {
		var financialNumber int

		err = selectPeopleSimple.Scan(&financialNumber)
		if err != nil {
			panic(err.Error())
		}
		p.FinancialNumber = financialNumber

		peoples = append(peoples, p)
	}
	defer db.Close()
	return peoples
}

// Seleciona uma pessoa para edição
func SelectOnePeopleSimpleUpdate(id string) []PeopleSimple {
	db := database.ConectDB()
	selectPeopleSimple, err := db.Query("select id, name, location, financial_number from peoples where id = " + id + " limit 1")
	if err != nil {
		panic(err.Error())
	}

	p := PeopleSimple{}
	peoples := []PeopleSimple{}

	for selectPeopleSimple.Next() {
		var id, financialNumber int
		var location, name string

		err = selectPeopleSimple.Scan(&id, &name, &location, &financialNumber)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Location = location
		p.Name = name
		p.FinancialNumber = financialNumber

		peoples = append(peoples, p)
	}
	defer db.Close()
	return peoples

}

// Executa o UPdate Simples de uma pessoa

func UpdatePeopleSimpleExec(id, financialnumber int, name, location string) {
	db := database.ConectDB()
	UpdatePeopleSimple, err := db.Prepare("update peoples set name=$1, location=$2, financial_number=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	UpdatePeopleSimple.Exec(name, location, financialnumber, id)
	defer db.Close()
}

// Criando uma nova pessoa

func CreateNewPeopleSimple(name string, financialnumber int, location string) {
	db := database.ConectDB()

	InsertNewPeopleSimple, err := db.Prepare("insert into peoples (name, financial_number, created_at, access_level, password, enabled, location) values ($1, $2, now(), 999, '25F9E794323B453885F5181F1B624D0B', true, $3)")

	if err != nil {
		panic(err.Error())
	}
	InsertNewPeopleSimple.Exec(name, financialnumber, location)
	defer db.Close()
}

func DeletaPeople(id string) {
	db := database.ConectDB()

	DeletaPeople, err := db.Prepare("delete from peoples where id=$1")

	if err != nil {
		panic(err.Error())
	}
	DeletaPeople.Exec(id)
	defer db.Close()
}

func DisabledPeople(id string) {
	db := database.ConectDB()

	DisabledPeople, err := db.Prepare("update peoples set enabled=false where id=$1")

	if err != nil {
		panic(err.Error())
	}
	DisabledPeople.Exec(id)
	defer db.Close()
}

func EnabledPeople(id string) {
	db := database.ConectDB()

	EnabledPeople, err := db.Prepare("update peoples set enabled=true where id=$1")
	if err != nil {
		panic(err.Error())
	}
	EnabledPeople.Exec(id)
	defer db.Close()
}
