package postgres

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	sysdb "github.com/alehano/gobootstrap/sys/db"
	"github.com/alehano/gobootstrap/config"
	"sync"
	"log"
	"fmt"
)

var dbInstance = struct {
	mu sync.Mutex
	db *sqlx.DB
}{
	mu: sync.Mutex{},
	db: nil,
}

func GetDB(tries ...int) *sqlx.DB {
	dbInstance.mu.Lock()
	defer dbInstance.mu.Unlock()
	if dbInstance.db == nil {
		db, err := sqlx.Connect("postgres",
			fmt.Sprintf("host=%s password=%s user=%s dbname=%s sslmode=%s",
				config.Get().PostgresHost, config.Get().PostgresPassword,
				config.Get().PostgresUser, config.Get().PostgresDatabase, config.Get().PostgresSSLMode))
		if err != nil {
			log.Printf("Postgres connection error: %s", err)
			return GetDB(sysdb.ReconnectCounter(tries...))
		}
		dbInstance.db = db
	}
	return dbInstance.db
}
