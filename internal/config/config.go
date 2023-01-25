package config

type Config struct {
	DBDriver string `envconfig:"DB_DRIVER" default:"mysql"`
	DBUser   string `envconfig:"DB_USER" default:"eep"`
	DBPass   string `envconfig:"DB_PASS" default:"1903"`
	DBName   string `envconfig:"DB_NAME" default:"RESTfulAPI_todos"`
}
