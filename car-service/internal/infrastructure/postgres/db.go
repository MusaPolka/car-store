package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDBConn(host, user, password, dbname string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", psqlInfo)
}
