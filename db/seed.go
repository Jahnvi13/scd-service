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

func SeedTimelogs() {
	timelogs := []models.Timelog{
		{ID: "tl_1", Version: 1, UID: "tl_uid_1", Duration: 1000, TimeStart: 1735164463798, TimeEnd: 1735165961010, Type: "captured", JobUID: "job_uid_ywij5sh1tvfp5nkq7azav"},
		{ID: "tl_1", Version: 2, UID: "tl_uid_2", Duration: 1050, TimeStart: 1735164463798, TimeEnd: 1735165961010, Type: "adjusted", JobUID: "job_uid_ywij5sh1tvfp5nkq7azav"},
	}
	for _, tl := range timelogs {
		DB.Create(&tl)
	}
}

func SeedPaymentLineItems() {
	items := []models.PaymentLineItem{
		{ID: "li_1", Version: 1, UID: "li_uid_1", JobUID: "job_uid_ywij5sh1tvfp5nkq7azav", TimelogUID: "tl_uid_2", Amount: 35, Status: "not-paid"},
		{ID: "li_1", Version: 2, UID: "li_uid_2", JobUID: "job_uid_ywij5sh1tvfp5nkq7azav", TimelogUID: "tl_uid_2", Amount: 35, Status: "paid"},
	}
	for _, item := range items {
		DB.Create(&item)
	}
}
