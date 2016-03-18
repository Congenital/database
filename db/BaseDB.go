package db

import (
	"database/sql"
	"github.com/Congenital/log/v0.2/log"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	config *DBConfig
	db *sql.DB
}

func NewConfigDB(driverName, protocol, ip, port, username, password, databaseName string) *DB {
	config := &DBConfig{
			driverName : driverName,
			protocol : protocol,
			ip : ip,
			port : port,
			username : username,
			password : password,
			databaseName : databaseName}

			return NewDB(config)
}

func NewDB(config *DBConfig) *DB {
	return &DB{
		db : nil,
		config : config,
	}
}

func (this *DB) Init(){
	if (this.config == nil) {
		log.Error("Err - Config is ", this.config)
		return
	}

	db, err := sql.Open(this.config.driverName, this.config.username + ":" + this.config.password + "@" + this.config.protocol + "(" + this.config.ip + ":" + this.config.port + ")/" + this.config.databaseName)
	if err != nil {
		log.Error(err)
		return
	}

	this.Set(db)
}

func (this *DB) Get() *sql.DB{
	return this.db
}

func (this *DB) Set(db *sql.DB) {
	this.db = db
}

func (this *DB) GetConfig() *DBConfig{
	return this.config
}

type DBConfig struct {
	driverName string
	protocol string
	ip string
	port string
	username string
	password string
	databaseName string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}
