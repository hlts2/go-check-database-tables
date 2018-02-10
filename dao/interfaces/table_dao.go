package interfaces

import (
	"github.com/hlts2/go-check-database-tables/models"
)

//TableDao is database table access interface
type TableDao interface {
	GetTables() (models.Tables, error)
	GetTableDescribe(tableName string) (*models.DescribeTable, error)
}
