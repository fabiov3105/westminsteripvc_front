package models

import (
	"github.com/fabiov3105/westminsteripvc_front/drivers/database"
)

type People struct {
	// Id              int
	Name            string
	FinancialNumber int
}

func CreateNewPeopleSimple(name string, financialnumberConver int) {
	database.ConectDB()
	DefineDados := People{name: name, financial_number: financialnumberConver}
	database.DB.Create(&DefineDados)
}
