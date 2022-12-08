package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Credential struct {
	Username string
	Password string
}

type GormConfig struct {
	Dsn        string
	Credential Credential
}

type ConnectionFactory interface {
	GetConnection(config GormConfig) (*gorm.DB, error)
}

type SQLiteConnectionFactory struct{}

func (f *SQLiteConnectionFactory) GetConnection(config GormConfig) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(config.Dsn), &gorm.Config{})
}

func OpenConnection(dsn string, factory ConnectionFactory) (*gorm.DB, error) {
	db, err := factory.GetConnection(GormConfig{Dsn: dsn})
	return db, err
}
