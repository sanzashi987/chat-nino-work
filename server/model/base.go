package model

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null;index"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

func (model BaseModel) GetStringID() string {
	return strconv.FormatUint(model.ID, 10)
}

func (model BaseModel) GetCreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

func (model BaseModel) GetUpdatedDate() string {
	return model.UpdatedAt.Format("2006-01-02")
}

var DBRef *gorm.DB

func SetupDB(db *gorm.DB) {
	DBRef = db
}
