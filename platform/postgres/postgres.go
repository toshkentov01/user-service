package postgres

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //pq for connection
	"github.com/toshkentov01/alif-tech-task/user-service/pkg/utils"
)

var (
	instance *sqlx.DB
	once     sync.Once
)

//DB ...
func DB() *sqlx.DB {

	once.Do(func() {
		dsn, err := utils.ConnectionURLBuilder("postgres")
		if err != nil {
			panic(err)
		}
		instance = sqlx.MustConnect("postgres", dsn)
	})

	return instance
}
