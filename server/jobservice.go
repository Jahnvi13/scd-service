package server

import (
	"context"
	"fmt"
	"log"
	"scd-service/db"
	"scd-service/models"
	"scd-service/proto"
	"strconv"

	"github.com/google/uuid"
)

type server struct {
	proto.UnimplementedSCDServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetLatestJobs(ctx context.Context, req *proto.GetLatestJobsRequest) (*proto.GetLatestJobsResponse, error) {
	var jobs []models.Job
	tx := GetLatestVersionQuery(db.DB, &models.Job{}, "jobs.id")
	if req.StatusFilter != "" {
		tx = tx.Where("status = ?", req.StatusFilter)
	}
	err := tx.Find(&jobs).Error
	if err != nil {
		return nil, err
	}
	log.Printf("Fetched %d jobs from DB", len(jobs))
	protoJobs := make([]*proto.Job, 0, len(jobs))
	for _, j := range jobs {
		protoJobs = append(protoJobs, &proto.Job{
			Id:           j.ID,
			Version:      j.Version,
			Uid:          j.UID,
			Status:       j.Status,
			Rate:         j.Rate,
			Title:        j.Title,
			CompanyId:    j.CompanyID,
			ContractorId: j.ContractorID,
		})
	}

	return &proto.GetLatestJobsResponse{Jobs: protoJobs}, nil
}

func (s *server) UpdateJob(ctx context.Context, req *proto.UpdateJobRequest) (*proto.Job, error) {
	var latest models.Job

	err := db.DB.Where("uid = ?", req.Id).Order("version DESC").First(&latest).Error
	if err != nil {
		return nil, err
	}

	newJob := latest
	newJob.ID = uuid.NewString()
	newJob.Version = latest.Version + 1

	for field, value := range req.UpdatedFields {
		switch field {
		case "uid":
			newJob.UID = value
		case "status":
			newJob.Status = value
		case "rate":
			r, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid rate: %v", err)
			}
			newJob.Rate = float32(r)
		case "title":
			newJob.Title = value
		case "company_id":
			newJob.CompanyID = value
		case "contractor_id":
			newJob.ContractorID = value
		}
	}

	err = db.DB.Create(&newJob).Error
	if err != nil {
		return nil, err
	}

	return &proto.Job{
		Id:           newJob.ID,
		Version:      newJob.Version,
		Uid:          newJob.UID,
		Status:       newJob.Status,
		Rate:         float32(newJob.Rate),
		Title:        newJob.Title,
		CompanyId:    newJob.CompanyID,
		ContractorId: newJob.ContractorID,
	}, nil
}
