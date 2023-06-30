package bootstrap

import (
	"fmt"

	"github.com/cza14h/chat-nino-work/model"
	"github.com/cza14h/chat-nino-work/model/completion"
	"github.com/cza14h/chat-nino-work/model/user"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/sharding"
)

func setupDB() {

	db, err := gorm.Open(sqlite.Open("chat-nino-work.db"), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	model.SetupDB(db)
	shard(db)
	migrate(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&user.UserModel{}, &completion.MessageModel{}, &completion.DialogModel{})
	if err != nil {
		fmt.Printf("error from database migrate: %f", err.Error())
	}
}

func shard(db *gorm.DB) {
	db.Use(sharding.Register(
		sharding.Config{
			ShardingKey:         "id",
			NumberOfShards:      64,
			PrimaryKeyGenerator: sharding.PKSnowflake,
		}, completion.MessageModel{},
	))
}
