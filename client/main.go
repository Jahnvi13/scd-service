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
}
