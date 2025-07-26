package server

import (
	"context"
	"log"
	"scd-service/db"
	"scd-service/models"
	"scd-service/proto"
)

type server struct {
	proto.UnimplementedSCDServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetLatestJobs(ctx context.Context, req *proto.GetLatestJobsRequest) (*proto.GetLatestJobsResponse, error) {
	var jobs []models.Job
	tx := db.DB
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
	log.Printf("Received UpdateJob request: id=%s fields=%v", req.Id, req.UpdatedFields)

	// dummy updated job
	job := &proto.Job{
		Id:      req.Id,
		Version: 2, // incremented version
		Status:  req.UpdatedFields["status"],
		Title:   req.UpdatedFields["title"],
	}

	return job, nil
}
