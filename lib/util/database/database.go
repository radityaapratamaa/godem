package database

import (
	"godem/domain/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DB struct {
	Master
	Follower
	master   *sqlx.DB
	follower *sqlx.DB
	timeout  time.Duration
}

func newDB(master, follower *sqlx.DB) *DB {
	return &DB{
		Master:   master,
		Follower: follower,
		master:   master,
		follower: follower,
		timeout:  3 * time.Second,
	}
}

func Connect(cfg *models.Config) (*DB, error) {

	masterDB, err := connectDB(&cfg.Databases.Master)
	if err != nil {
		return nil, errors.Wrap(err, "lib.util.database.ConnectMaster")
	}

	followerDB, err := connectDB(&cfg.Databases.Follower)
	if err != nil {
		return nil, errors.Wrap(err, "lib.util.database.ConnectSlave")
	}

	db := newDB(masterDB, followerDB)
	return db, nil
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
