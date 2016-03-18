package db

import (
	"database/sql"
	"github.com/Congenital/config"
	"github.com/Congenital/log/v0.2/log"
)

type BaseModel struct {
	db *DB
}

func (this *BaseModel) Get() *DB{
	return this.db
}

func (this *BaseModel) GetDB() *sql.DB {
	return this.db.Get()
}

func (this *BaseModel) GetConfig() *DBConfig {
	return this.db.GetConfig()
}


func NewBaseModel(configPath string) (*BaseModel, error) {
	dbConfig := NewDBConfig()
	err := config.Init(dbConfig, configPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &BaseModel{db : NewDB(dbConfig)}, nil
}
