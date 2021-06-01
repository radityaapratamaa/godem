package models

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	JWT       JWTConfig       `yaml:"jwt"`
	Databases DatabasesConfig `yaml:"databases"`
	Redis     RedisCfg        `yaml:"redis"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type JWTConfig struct {
	PublicKey string `yaml:"public_key"`
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
