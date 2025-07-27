package main

import (
	"fmt"
	"log"

	"scd-service/client/client"
)

func main() {
	c, err := client.NewClient("localhost:8000")
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer c.Close()

	// Test GetLatestJobs
	jobs, err := c.GetLatestJobs("")
	if err != nil {
		log.Fatalf("error calling GetLatestJobs: %v", err)
	}
	fmt.Println("Jobs received:")
	for _, job := range jobs {
		fmt.Printf("- ID: %s, Title: %s, Status: %s, Version: %d\n", job.Id, job.Title, job.Status, job.Version)
	}

	updatedJob, err := c.UpdateJob("uid1", map[string]string{
		"status": "in-progress",
		"title":  "Backend Developer",
		"rate":   "25.5",
	})
	if err != nil {
		log.Fatalf("UpdateJob failed: %v", err)
	}
	fmt.Printf("Updated Job Version: %d\n", updatedJob.Version)
	fmt.Printf("Updated Job Fields: %+v\n", updatedJob)

	timelog, err := c.GetTimelog("tl_uid_2")
	if err != nil {
		log.Fatalf("GetTimelog failed: %v", err)
	}
	fmt.Printf("Got Timelog: %+v\n", timelog)

	pli, err := c.GetPaymentLineItem("li_uid_2")
	if err != nil {
		log.Fatalf("GetPaymentLineItem failed: %v", err)
	}
	fmt.Printf("Got PaymentLineItem: %+v\n", pli)

	updatedPLI, err := c.UpdatePaymentLineItem("li_uid_2", map[string]string{
		"status": "paid",
		"amount": "50",
	})
	if err != nil {
		log.Fatalf("UpdatePaymentLineItem failed: %v", err)
	}
	fmt.Printf("Updated PaymentLineItem Version: %d\n", updatedPLI.Version)
	fmt.Printf("Updated PaymentLineItem Fields: %+v\n", updatedPLI)

	updatedTimelog, err := c.UpdateTimelog("tl_uid_2", map[string]string{
		"duration": "1800",
		"type":     "adjusted",
	})
	if err != nil {
		log.Fatalf("UpdateTimelog failed: %v", err)
	}
	fmt.Printf("Updated Timelog Version: %d\n", updatedTimelog.Version)
	fmt.Printf("Updated Timelog Fields: %+v\n", updatedTimelog)
}
