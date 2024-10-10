package settings

type Database struct {
	Name     string `envconfig:"DATABASE_NAME" default:"products_api"`
	Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port     string `envconfig:"DATABASE_PORT" default:"5432"`
	Username string `envconfig:"DATABASE_USERNAME" default:"products_admin"`
	Password string `envconfig:"DATABASE_PASSWORD" default:"123qwe"`
}

type Settings struct {
	Server struct {
		Host    string `envconfig:"SERVER_HOST" default:"localhost"`
		Port    string `envconfig:"SERVER_PORT" default:"8080"`
		Context string `envconfig:"SERVER_CONTEXT" default:"products-api"`
	}
}
