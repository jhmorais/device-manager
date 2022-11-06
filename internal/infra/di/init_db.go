package di

import (
	"fmt"

	"github.com/jhmorais/device-manager/config"
	"github.com/jhmorais/device-manager/internal/domain/entities"
	"github.com/jhmorais/device-manager/sample"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormMysqlDB() (*gorm.DB, error) {
	config.LoadServerEnvironmentVars()

	dsn := fmt.Sprintf("%s:%s@%s", config.GetMysqlUser(), config.GetMysqlPassword(), config.GetMysqlConnectionString())
	fmt.Printf("%s \n", dsn)

	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	mysqlDb.AutoMigrate(&entities.Device{})

	sample.DBSeed(mysqlDb)

	return mysqlDb, err
}
