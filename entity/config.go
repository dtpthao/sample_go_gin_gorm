package entity

type MysqlConfig struct {
	Host     string `env:"DB_HOST" env-default:"127.0.0.1"`
	Port     int    `env:"DB_PORT" env-default:"3306"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}
