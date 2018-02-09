package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/models"
)

type TableDaoImpl struct {
	config.DBConfig
}

//GetTables returns all tables of database
func (impl TableDaoImpl) GetTables() (models.Tables, error) {
	_, err := sql.Open("mysql", impl.DSN())
	if err != nil {
		return nil, err
	}
	return nil, nil
}
