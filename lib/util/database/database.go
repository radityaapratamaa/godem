package database

import (
	"bcg-test/domain/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func Connect(cfg *models.Config) (*models.Database, error) {

	masterDB, err := connectDB(&cfg.Databases.Master)
	if err != nil {
		return nil, errors.Wrap(err, "lib.util.database.ConnectMaster")
	}

	slaveDB, err := connectDB(&cfg.Databases.Slave)
	if err != nil {
		return nil, errors.Wrap(err, "lib.util.database.ConnectSlave")
	}

	database := &models.Database{
		Master: masterDB,
		Slave:  slaveDB,
	}
	return database, nil
}

func connectDB(cfg *models.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect(cfg.Driver, cfg.ConnString)
	if err != nil {
		return nil, errors.Wrap(err, "lib.util.database.Connect")
	}

	maxLifetime := time.Duration(cfg.ConnMaxLifetime) * time.Second
	db.SetConnMaxLifetime(maxLifetime)
	maxIddleTime := time.Duration(cfg.ConnMaxIddleTime) * time.Second
	db.SetConnMaxIdleTime(maxIddleTime)

	db.SetMaxOpenConns(cfg.ConnMaxLifetime)
	db.SetMaxIdleConns(cfg.MaxOpenConn)
	return db, nil
}
