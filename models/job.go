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
	ID        string `gorm:"primaryKey"`
	Version   int32  `gorm:"primaryKey"`
	UID       string `gorm:"unique"`
	Duration  int64
	TimeStart int64
	TimeEnd   int64
	Type      string
	JobUID    string
}

type PaymentLineItem struct {
	ID         string `gorm:"primaryKey"`
	Version    int32  `gorm:"primaryKey"`
	UID        string `gorm:"unique"`
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
