package manager

import (
	"github.com/jmoiron/sqlx"
	"github.com/rickyhidayatt/do-app/config"
)

type InfraManager interface {
	SqlDb() *sqlx.DB
}

type infraManager struct {
	db     *sqlx.DB
	config config.Config
}

func (i *infraManager) SqlDb() *sqlx.DB {
	return i.db
}

func (i *infraManager) initDb() {
	connectDb, err := sqlx.Connect("postgres", i.config.Dsn)

	if err != nil {
		panic(err.Error())
	}

	if checkErrConfig := connectDb.Ping(); checkErrConfig != nil {
		panic(checkErrConfig)
	}

	i.db = connectDb
}

func NewInfraManager(configArg config.Config) InfraManager {
	infra := infraManager{
		config: configArg,
	}
	infra.initDb()
	return &infra
}
