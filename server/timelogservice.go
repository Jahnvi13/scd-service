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

func (s *server) GetLatestTimelogs(ctx context.Context, req *proto.GetLatestTimelogsRequest) (*proto.GetLatestTimelogsResponse, error) {
	var latestVersions []models.Timelog
	subQuery := db.DB.
		Table("timelogs").
		Select("uid, MAX(version) as max_version").
		Group("uid")

	query := db.DB.
		Joins("JOIN (?) as latest ON timelogs.uid = latest.uid AND timelogs.version = latest.max_version", subQuery)

	if req.TypeFilter != "" {
		query = query.Where("timelogs.type = ?", req.TypeFilter)
	}

	err := query.Find(&latestVersions).Error
	if err != nil {
		return nil, err
	}

	var result []*proto.Timelog
	for _, tl := range latestVersions {
		result = append(result, &proto.Timelog{
			Id:        tl.ID,
			Version:   int32(tl.Version),
			Uid:       tl.UID,
			Duration:  tl.Duration,
			TimeStart: tl.TimeStart,
			TimeEnd:   tl.TimeEnd,
			Type:      tl.Type,
			JobUid:    tl.JobUID,
		})
	}

	return &proto.GetLatestTimelogsResponse{Timelogs: result}, nil
}

func (s *server) UpdateTimelog(ctx context.Context, req *proto.UpdateTimelogRequest) (*proto.Timelog, error) {
	var latest models.Timelog

	err := db.DB.Where("uid = ?", req.Id).Order("version DESC").First(&latest).Error
	if err != nil {
		return nil, err
	}

	newLog := latest
	newLog.UID = uuid.NewString()
	newLog.Version = latest.Version + 1

	for field, value := range req.UpdatedFields {
		switch field {
		case "id":
			newLog.ID = value
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
