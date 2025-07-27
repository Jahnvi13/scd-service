package models

type Job struct {
	ID           string `gorm:"primaryKey"`
	Version      int32
	UID          string
	Status       string
	Rate         float32
	Title        string
	CompanyID    string
	ContractorID string
}

type Timelog struct {
	ID        string
	Version   int32
	UID       string `gorm:"primaryKey"`
	Duration  int64
	TimeStart int64
	TimeEnd   int64
	Type      string
	JobUID    string
}

type PaymentLineItem struct {
	ID         string
	Version    int32
	UID        string `gorm:"primaryKey"`
	JobUID     string
	TimelogUID string
	Amount     float32
	Status     string
}

func (Job) TableName() string {
	return "jobs"
}

func (Timelog) TableName() string {
	return "timelogs"
}

func (PaymentLineItem) TableName() string {
	return "paymentLineItems"
}
