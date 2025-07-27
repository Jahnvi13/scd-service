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

}
