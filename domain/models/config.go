package models

type Config struct {
	Databases DatabasesConfig `yaml:"databases"`
	Redis     RedisCfg        `yaml:"redis"`
}

type RedisCfg struct {
	Address   string `yaml:"address"`
	Timeout   int    `yaml:"timeout"`
	MaxIdle   int    `yaml:"max_iddle"`
	MaxActive int    `yaml:"max_active"`
}

type DatabasesConfig struct {
	Master DatabaseConfig `yaml:"master"`
	Slave  DatabaseConfig `yaml:"slave"`
}

type DatabaseConfig struct {
	Driver           string `yaml:"driver"`
	ConnString       string `yaml:"conn_string"`
	ConnMaxLifetime  int    `yaml:"conn_max_lifetime"`
	ConnMaxIddleTime int    `yaml:"conn_max_iddle_time"`
	MaxOpenConn      int    `yaml:"max_open_conn"`
	MaxIddleConn     int    `yaml:"max_iddle_conn"`
}
