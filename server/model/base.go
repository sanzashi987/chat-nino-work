package model

import (
	"strconv"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

func (model BaseModel) GetStringID() string {
	return strconv.FormatUint(model.ID, 10)
}

func (model BaseModel) CreatedAtDate() string {
	return model.CreatedAt.Format("2006-01-02")
}

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("chat-nino-work.db"), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	return DB
}
