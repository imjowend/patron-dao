package factory

import (
	"log"

	"github.com/imjowend/patron-dao/interfaces"
	"github.com/imjowend/patron-dao/mysql"
	"github.com/imjowend/patron-dao/psql"
)

func FactoryDAO(engine string) interfaces.UserDAO {
	var i interfaces.UserDAO
	switch engine {
	case "postgres":
		i = psql.UserImplPsql{}
	case "mysql":
		i = mysql.UserImplMysql{}
	default:
		log.Fatalf("el motor: %s, no esta implementado", engine)
		return nil
	}
	return i
}
