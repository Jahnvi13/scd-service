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

// func GetLatestPaymentLineItemsByJobUID(jobUID string) ([]models.PaymentLineItem, error) {
// 	var items []models.PaymentLineItem
// 	err := GetLatestVersionQuery(db.DB, &models.PaymentLineItem{}, "id").
// 		Where("job_uid = ?", jobUID).
// 		Find(&items).Error
// 	return items, err
// }

func (s *server) GetLatestPaymentLineItems(ctx context.Context, req *proto.GetLatestPaymentLineItemsRequest) (*proto.GetLatestPaymentLineItemsResponse, error) {
	var latestItems []models.PaymentLineItem
	subQuery := db.DB.
		Table("paymentLineItems").
		Select("uid, MAX(version) as max_version").
		Group("uid")

	query := db.DB.
		Joins("JOIN (?) as latest ON paymentLineItems.uid = latest.uid AND paymentLineItems.version = latest.max_version", subQuery)

	if req.StatusFilter != "" {
		query = query.Where("paymentLineItems.status = ?", req.StatusFilter)
	}

	err := query.Find(&latestItems).Error
	if err != nil {
		return nil, err
	}

	var result []*proto.PaymentLineItem
	for _, item := range latestItems {
		result = append(result, &proto.PaymentLineItem{
			Id:         item.ID,
			Version:    int32(item.Version),
			Uid:        item.UID,
			JobUid:     item.JobUID,
			TimelogUid: item.TimelogUID,
			Amount:     item.Amount,
			Status:     item.Status,
		})
	}

	return &proto.GetLatestPaymentLineItemsResponse{Items: result}, nil
}

func (s *server) UpdatePaymentLineItem(ctx context.Context, req *proto.UpdatePaymentLineItemRequest) (*proto.PaymentLineItem, error) {
	var latest models.PaymentLineItem

	err := db.DB.Where("uid = ?", req.Id).Order("version DESC").First(&latest).Error
	if err != nil {
		return nil, err
	}

	newItem := latest
	newItem.UID = uuid.NewString()
	newItem.Version = latest.Version + 1

	for field, value := range req.UpdatedFields {
		switch field {
		case "id":
			newItem.ID = value
		case "job_uid":
			newItem.JobUID = value
		case "timelog_uid":
			newItem.TimelogUID = value
		case "amount":
			amount, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid amount: %v", err)
			}
			newItem.Amount = float32(amount)
		case "status":
			newItem.Status = value
		}
	}

	err = db.DB.Create(&newItem).Error
	if err != nil {
		return nil, err
	}
	return &proto.PaymentLineItem{
		Id:         newItem.ID,
		Version:    newItem.Version,
		Uid:        newItem.UID,
		JobUid:     newItem.JobUID,
		TimelogUid: newItem.TimelogUID,
		Amount:     newItem.Amount,
		Status:     newItem.Status,
	}, nil
}
