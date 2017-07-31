package postgres

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/alehano/gobootstrap/utils/pause"
	"github.com/alehano/gobootstrap/config"
	"sync"
	"log"
	"fmt"
	"time"
)

const (
	maxReconnect      = 10
	reconnectInterval = 5 * time.Second
)

var dbInstance = struct {
	mu sync.Mutex
	db *sqlx.DB
}{
	mu: sync.Mutex{},
	db: nil,
}

func GetDB() *sqlx.DB {
	dbInstance.mu.Lock()
	defer dbInstance.mu.Unlock()
	if dbInstance.db == nil {
		for pause.New(maxReconnect, reconnectInterval).Do() {
			db, err := sqlx.Connect("postgres",
				fmt.Sprintf("host=%s password=%s user=%s dbname=%s sslmode=%s",
					config.Get().PostgresHost, config.Get().PostgresPassword,
					config.Get().PostgresUser, config.Get().PostgresDatabase, config.Get().PostgresSSLMode))
			if err != nil {
				log.Printf("Postgres connection error: %s", err)
			} else {
				dbInstance.db = db
				return dbInstance.db
			}
		}
		return &sqlx.DB{}
	}
	return dbInstance.db
}
