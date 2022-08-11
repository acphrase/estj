package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// TODO: profile 설정 진행 후 변경 예정.
const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "gouser"
	password = "gouser"
	dbname   = "godb"
)

// singleton 객체값(pointer)
var dbInstance *sqlx.DB

func init() {
	initDB()
}

func GetDB() *sqlx.DB {
	if dbInstance == nil {
		initDB()
	}
	return dbInstance
}

func initDB() {
	if dbInstance == nil {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		// db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
		db, err := sqlx.Connect("postgres", psqlInfo)
		// db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
		}
		//defer func(db *sql.DB) {
		//	err := db.Close()
		//	if err != nil {
		//		panic(err)
		//	}
		//}(dbInstance)
		dbInstance = db
	}
}
