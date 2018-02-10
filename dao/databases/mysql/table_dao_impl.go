package mysql

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/models"
)

type TableDaoImpl struct {
	config.DBConfig
}

//GetTables returns all tables of database
func (impl TableDaoImpl) GetTables() (models.Tables, error) {
	db, err := sql.Open("mysql", impl.DSN())
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()

	var tables models.Tables
	_, err = dbmap.Select(&tables, fmt.Sprintf("SELECT table_name FROM information_schema.tables WHERE table_schema = '%s'", impl.DatabaseName()))
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return tables, nil
}
