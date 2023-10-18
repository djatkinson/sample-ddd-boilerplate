package oracle

import (
	"ddd-boilerplate/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ConnectDB(config *config.OracleDBConfig) *sqlx.DB {
	dsn := fmt.Sprintf("%s/%s@%s:%d/%s", config.Username, config.Password, config.HostPrimary, config.Port, config.Name)

	fmt.Println(dsn)

	var err error
	db, err := sqlx.Connect("godror", dsn)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}

	return db
}
