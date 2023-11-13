package database

import (
	"fmt"
	"github.com/app-dictionary/pkg/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func SetupDatabase() {
	var err error

	p := env.GetEnv("DB_PORT", "3306")
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.GetEnv("DB_USER", "root"),
		env.GetEnv("DB_PASSWORD", ""),
		env.GetEnv("DB_HOST", "0.0.0.0"),
		port,
		env.GetEnv("DB_NAME", ""),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
