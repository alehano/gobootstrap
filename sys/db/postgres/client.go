package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	sysdb "github.com/alehano/gobootstrap/sys/db"
	"github.com/alehano/gobootstrap/config"
	"sync"
	"log"
	"fmt"
)

type dbStruct struct {
	mu sync.Mutex
	db *sql.DB
}

var dbInstance = dbStruct{
	mu: sync.Mutex{},
	db: nil,
}

func GetDB(tries ...int) *sql.DB {
	dbInstance.mu.Lock()
	defer dbInstance.mu.Unlock()
	if dbInstance.db == nil {
		db, err := sql.Open("postgres",
			fmt.Sprintf("host=%s password=%s user=%s dbname=%s sslmode=%s",
				config.Get().PostgresHost, config.Get().PostgresPassword,
				config.Get().PostgresUser, config.Get().PostgresDatabase, config.Get().PostgresSSLMode))
		if err != nil {
			log.Printf("Postgres open error: %s", err)
			return GetDB(sysdb.ReconnectCounter(tries...))
		} else if err = db.Ping(); err != nil {
			log.Printf("Postgres ping error: %s", err)
			return GetDB(sysdb.ReconnectCounter(tries...))
		}
		dbInstance.db = db
	}
	return dbInstance.db
}
