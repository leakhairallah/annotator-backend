package mysql

import (
	"annotator-backend/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MaxOpenConn     = 60
	ConnMaxLifetime = 120
	MaxIdleConn     = 30
	ConMaxIdleTime  = 20
)

func NewMySqlDB(c *config.Config) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.MySql.User,
		c.MySql.Password,
		c.MySql.Host,
		c.MySql.Port,
		c.MySql.DbName,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(ConnMaxLifetime)
	db.SetMaxOpenConns(MaxOpenConn)
	db.SetMaxIdleConns(MaxIdleConn)
	db.SetConnMaxIdleTime(ConMaxIdleTime)

	return db, nil
}
