package db

import (
	"scd-service/models"
)

func SeedJobs() {
	jobs := []models.Job{
		{ID: "job1", Version: 1, UID: "uid1", Status: "open", Rate: 10, Title: "Dev", CompanyID: "comp1", ContractorID: "cont1"},
		{ID: "job2", Version: 2, UID: "uid2", Status: "closed", Rate: 15, Title: "QA", CompanyID: "comp2", ContractorID: "cont2"},
	}
	for _, job := range jobs {
		DB.Create(&job)
	}
}
