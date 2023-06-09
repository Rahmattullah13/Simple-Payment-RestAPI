package manager

import (
	"fmt"
	"os"
	"simple-payment/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	SqlDB() *sqlx.DB
}

type infraManager struct {
	dbConnect *sqlx.DB
	config    config.Config
}

func (i *infraManager) SqlDB() *sqlx.DB {
	return i.dbConnect
}

func (i *infraManager) initDB() {
	os.Setenv("PGSSLMODE", "disable")
	db, err := sqlx.Connect("postgres", i.config.DataSourceName)

	if err != nil {
		fmt.Println("initDB failed", err.Error())
		panic(err)
	}
	if errConf := db.Ping(); errConf != nil {
		panic(errConf)
	}

	fmt.Println("Connecting DB Success!!!")

	i.dbConnect = db
}

func NewInfraManager(configParam config.Config) InfraManager {
	infra := new(infraManager)
	infra.config = configParam
	infra.initDB()
	return infra
}
