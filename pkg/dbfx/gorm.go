package dbfx

import (
	"fmt"
	"github.com/noxymon/ecomm-go/domain/pkg/dao/query"
	"github.com/noxymon/ecomm-go/pkg/commonfx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newGorm(common *commonfx.Common) (*query.Query, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		common.Viper.GetString("database.host"),
		common.Viper.GetString("database.username"),
		common.Viper.GetString("database.password"),
		common.Viper.GetString("database.schema"),
		common.Viper.GetInt("database.port"),
	)

	postgresConnection := postgres.New(
		postgres.Config{
			DSN: dsn,
		},
	)

	db, err := gorm.Open(postgresConnection, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return query.Use(db), nil
}
