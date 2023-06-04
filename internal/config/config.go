package config

type Config struct {
	DBDriver string `envconfig:"DB_DRIVER" default:"postgres"`
	DBUser   string `envconfig:"DB_USER" default:"eep"`
	DBPass   string `envconfig:"DB_PASS" default:"1903"`
	DBName   string `envconfig:"DB_NAME" default:"restfulapi_todos"`
	DBHost   string `envconfig:"DB_HOST" default:"localhost"`
	DBPort   int    `envconfig:"DB_PORT" default:"5432"`
}
