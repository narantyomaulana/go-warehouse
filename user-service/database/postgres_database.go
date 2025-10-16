package database

import (
	"fmt"
	"micro-warehouse/user-service/configs"
	"micro-warehouse/user-service/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func ConnectionPostgres(cfg configs.Config) (*Postgres, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.SqlDB.User, cfg.SqlDB.Password, cfg.SqlDB.Host, cfg.SqlDB.Port, cfg.SqlDB.DBName)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Errorf("[Postgres] ConnectionPostgres - 1: %v", err)
		return nil, err
	}

	db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRole{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("[Postgres] ConnectionPostgres - 2: %v", err)
		return nil, err
	}

	SeedRole(db)
	SeedManager(db)

	sqlDB.SetMaxIdleConns(cfg.SqlDB.DBMaxOpenConns)
	sqlDB.SetMaxOpenConns(cfg.SqlDB.DBMaxIdleConns)

	return &Postgres{DB: db}, nil
}
