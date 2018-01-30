package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mrizkip/backend-finding-dosen/env"
	"gopkg.in/gorp.v2"
)

var Dbm *gorp.DbMap

// initialize dbMap instance
func init() {

	dbName := env.Getenv("DB_NAME", "findingdosen")
	dbHost := env.Getenv("DB_HOST", "127.0.0.1")
	dbUsername := env.Getenv("DB_USERNAME", "root")
	dbPort := env.Getenv("DB_PORT", "3306")
	dbPassword := env.Getenv("DB_PASSWORD", "")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	log.Println(dbUrl)

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	Dbm = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	Dbm.TraceOn("[gorm]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
	Dbm.AddTableWithName(User{}, "users").SetKeys(true, "ID").AddIndex("EmailIdentitasIndex", "Btree", []string{"email", "no_identitas"}).SetUnique(true)
	Dbm.AddTableWithName(Status{}, "status").SetKeys(true, "ID").AddIndex("UserIDIndex", "Btree", []string{"user_id"}).SetUnique(true)
	Dbm.TraceOff()

}

// Create tables
func CreateTables() error {
	if err := Dbm.DropTablesIfExists(); err != nil {
		return err
	}

	if err := Dbm.CreateTablesIfNotExists(); err != nil {
		return err
	}
	if err := Dbm.CreateIndex(); err != nil {
		return err
	}

	return nil
}
