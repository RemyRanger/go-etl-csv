package mysql

import (
	"time"

	"github.com/RemyRanger/go-etl-csv/internal/common/config"
	"github.com/RemyRanger/go-etl-csv/internal/common/logs"
	"github.com/rs/zerolog/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New : create client to connect to database
func New() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Conf.Mysql.DSN), &gorm.Config{
		Logger:                 logs.NewMysqlLogger(),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating mysql client")
	}
	PingDB(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating mysql pool")
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(30 * time.Second)

	return db
}

// PingDB Ping Database by runing SELECT VERSION()
func PingDB(db *gorm.DB) {
	var version string
	db.Raw("SELECT VERSION()").Scan(&version)
	log.Info().Msgf("Mysql Version : %v", version)
}
