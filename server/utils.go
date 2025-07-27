package server

import "gorm.io/gorm"

func GetLatestVersionQuery(db *gorm.DB, model interface{}, idColumn string) *gorm.DB {
	sub := db.
		Model(model).
		Select("id, MAX(version) as max_version").
		Group("id")

	return db.
		Model(model).
		Joins("JOIN (?) AS latest ON latest.id = "+idColumn+" AND latest.max_version = version", sub)
}
