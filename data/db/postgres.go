package db

import (
	"fmt"
	"gin-api-blog/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbClient *sqlx.DB

func InitDB(cfg *config.Config) error {
	var err error

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)

	db := sqlx.MustOpen("postgres", conn)
	if err = db.Ping(); err != nil {
		return err
	}

	db.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)

	dbClient = db
	log.Println("Postgres connection established")
	return nil
}

func GetDB() *sqlx.DB {
	return dbClient
}

func CloseDB() {
	dbClient.DB.Close()
}
