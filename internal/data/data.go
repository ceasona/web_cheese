package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"web_cheese/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGormDB, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *gorm.DB
}

func NewGormDB(c *conf.Data) (*gorm.DB, error) {
	fmt.Println("NewGormDB1")
	dsn := c.Database.Source
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(150)
	sqlDB.SetConnMaxLifetime(time.Second * 25)
	fmt.Println("NewGormDB2")
	return db, err
}

// NewData .
func NewData(logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	fmt.Println("NewData1")
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	fmt.Println("NewData2")
	return &Data{db: db}, cleanup, nil
}
