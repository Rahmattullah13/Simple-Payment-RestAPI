package util

import (
	"fmt"
	"simple-payment/config"
	"strings"

	"github.com/jmoiron/sqlx"
)

type SaveLog struct{}

func (sl *SaveLog) Write(p []byte) (n int, err error) {
	conf := config.NewConfig()

	db, err := sqlx.Connect("postgres", conf.DataSourceName)
	if err != nil {
		fmt.Println("initDb failed", err.Error())
		panic(err)
	}
	if errConf := db.Ping(); errConf != nil {
		panic(errConf)
	}
	if strings.Contains(string(p), "api") {
		_, _ = db.Exec(CREATE_LOG, string(p))
	}

	n = len(p)
	return n, err
}
