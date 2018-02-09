package factories

import (
	"github.com/hlts2/go-check-database-tables/dao/databases/config"
	"github.com/hlts2/go-check-database-tables/dao/databases/mysql"
	"github.com/hlts2/go-check-database-tables/dao/interfaces"
)

//FactoryTableDao create table dao
func FactoryTableDao(s string, c config.Config) interfaces.TableDao {
	var i interfaces.TableDao

	switch s {
	case "mysql":
		i = mysql.TableDaoImpl{
			mysql.GetMysqlConfig(c),
		}
	default:
		break
	}
	return i
}
