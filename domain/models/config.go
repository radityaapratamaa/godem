package models

import (
	"github.com/kodekoding/phastos/go/database"
	"github.com/kodekoding/phastos/go/notifications"
	"github.com/kodekoding/phastos/go/server"
)

type Config struct {
	Server        server.Config        `yaml:"server"`
	JWT           JWTConfig            `yaml:"jwt"`
	Databases     database.SQLs        `yaml:"databases"`
	Redis         RedisCfg             `yaml:"redis"`
	Notifications notifications.Config `yaml:"notifications"`
}

type ServerConfig struct {
	Port    string `yaml:"port"`
	Timeout struct {
		Read  int `yaml:"read"`
		Write int `yaml:"write"`
	} `yaml:"timeout"`
}

type JWTConfig struct {
	PublicKey  string `yaml:"public_key"`
	SigningKey string `yaml:"signing_key"`
}

type RedisCfg struct {
	Address   string `yaml:"address"`
	Timeout   int    `yaml:"timeout"`
	MaxIdle   int    `yaml:"max_iddle"`
	MaxActive int    `yaml:"max_active"`
}

type DatabasesConfig struct {
	Master   DatabaseConfig `yaml:"master"`
	Follower DatabaseConfig `yaml:"follower"`
}

type DatabaseConfig struct {
	Driver           string `yaml:"driver"`
	ConnString       string `yaml:"conn_string"`
	ConnMaxLifetime  int    `yaml:"conn_max_lifetime"`
	ConnMaxIddleTime int    `yaml:"conn_max_iddle_time"`
	MaxOpenConn      int    `yaml:"max_open_conn"`
	MaxIddleConn     int    `yaml:"max_iddle_conn"`
}
