package manager

import (
	"GO-Payment/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Init() error
	Deinit()
	GetDB() *sql.DB
}

type infra struct {
	db  *sql.DB
	cfg config.AppConfig
}

func NewInfraManager(cfg config.AppConfig) InfraManager {
	infra := infra{
		cfg: cfg,
	}
	infra.Init()
	return &infra
}

func (self *infra) Init() error {
	if err := self.dbOpen(); err != nil {
		return err
	}

	return nil
}

func (self *infra) Deinit() {
	if self.db != nil {
		self.db.Close()
	}
}

func (self *infra) GetDB() *sql.DB {
	return self.db
}

// private
func (self *infra) dbOpen() error {
	var err error
	var c = &self.cfg.DbConfig
	var dbParams = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.SslMode,
	)
	self.db = nil
	self.db, err = sql.Open("postgres", dbParams)
	if err != nil {
		return err
	}

	// test db connection
	if err = self.db.Ping(); err != nil {
		return fmt.Errorf("sql.Open: %s", err.Error())
	}

	return nil
}
