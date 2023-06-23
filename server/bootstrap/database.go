package bootstrap

import (
	"fmt"

	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/model/message"
	"github.com/cza14h/chat-nino-work/model/user"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

func SetupDB() {

	db := model.ConnectDB()
	migrate(db)

}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&user.UserModel{}, &message.MessageModal{})
	if err != nil {
		fmt.Printf("error from database migrate: %f", err.Error())
	}
}

func shard(db *gorm.DB) {
	db.Use(sharding.Register{
		sharding.Config{},
	})
}
