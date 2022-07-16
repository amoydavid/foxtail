package config

type ServerConfig struct {
	Name        string       `mapstructure:"name"`
	Port        int          `mapstructure:"port"`
	Mysqlinfo   MysqlConfig  `mapstructure:"mysql"`
	RedisInfo   RedisConfig  `mapstructure:"redis"`
	LogsAddress string       `mapstructure:"logsAddress"`
	Sqliteinfo  SqliteConfig `mapstructure:"sqlite"`
}

type SqliteConfig struct {
	Path string `mapstructure:"path"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
