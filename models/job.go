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

func (Job) TableName() string {
	return "jobs"
}
