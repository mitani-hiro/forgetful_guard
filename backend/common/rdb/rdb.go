package rdb

import (
	"context"
	"database/sql"
	"fmt"
	"forgetful-guard/common/logger"
	"os"

	_ "time/tzdata"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
)

var DB *sql.DB

// InitDB DB初期化.
// TODO 接続情報は別で管理.
func InitDB() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&Asia%%2FTokyo", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return errors.Wrap(err, "sql.Open error")
	}

	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "db.Ping error")
	}

	logger.Info("DB connection successful")
	DB = db
	boil.SetDB(DB)
	return nil
}

func GetDBConnection() *sql.DB {
	return DB
}

// Tx トランザクション処理.
func Tx(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := DB.Begin()
	if err != nil {
		return errors.Wrap(err, "DB.Begin error")
	}

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.WithMessage(err, rerr.Error())
		}
		return errors.Wrap(err, "function error")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit error")
	}

	return nil
}
