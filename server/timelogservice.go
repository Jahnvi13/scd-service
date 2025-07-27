package server

import (
	"context"
	"fmt"
	"scd-service/db"
	"scd-service/models"
	"scd-service/proto"
	"strconv"

	"github.com/google/uuid"
)

func GetLatestTimelogsByJobUID(jobUID string) ([]models.Timelog, error) {
	var timelogs []models.Timelog
	err := GetLatestVersionQuery(db.DB, &models.Timelog{}, "id").
		Where("job_uid = ?", jobUID).
		Find(&timelogs).Error
	return timelogs, err
}

func (s *server) UpdateTimelog(ctx context.Context, req *proto.UpdateTimelogRequest) (*proto.Timelog, error) {
	var latest models.Timelog

	err := db.DB.Where("uid = ?", req.Id).Order("version DESC").First(&latest).Error
	if err != nil {
		return nil, err
	}

	newLog := latest
	newLog.ID = uuid.NewString()
	newLog.Version = latest.Version + 1

	for field, value := range req.UpdatedFields {
		switch field {
		case "uid":
			newLog.UID = value
		case "duration":
			duration, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid duration: %v", err)
			}
			newLog.Duration = duration
		case "time_start":
			timeStart, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid time_start: %v", err)
			}
			newLog.TimeStart = timeStart
		case "time_end":
			timeEnd, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid time_end: %v", err)
			}
			newLog.TimeEnd = timeEnd
		case "type":
			newLog.Type = value
		case "job_uid":
			newLog.JobUID = value
		}
	}

	err = db.DB.Create(&newLog).Error
	if err != nil {
		return nil, err
	}

	return &proto.Timelog{
		Id:        newLog.ID,
		Version:   newLog.Version,
		Uid:       newLog.UID,
		Duration:  newLog.Duration,
		TimeStart: newLog.TimeStart,
		TimeEnd:   newLog.TimeEnd,
		Type:      newLog.Type,
		JobUid:    newLog.JobUID,
	}, nil
}
