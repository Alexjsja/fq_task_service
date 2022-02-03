package db

import (
	"fmt"
	"log"

	"github.com/StanDenisov/fq_utils/confstruct"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cnf confstruct.ConfStruct) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.PgDBHost, cnf.PgDBPort, cnf.PgDBUserName, cnf.PgDBPassword, cnf.PgDBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
