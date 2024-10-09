package settings

type Database struct {
	Name     string `envconfig:"product-postgres"`
	Host     string `envconfig:"default=localhost"`
	Port     string `envconfig:"default=5432"`
	Username string `envconfig:"default=product-admin"`
	Password string `envconfig:"default=123qwe"`
}

type Settings struct {
	Server struct {
		Port    string `envconfig:"default=8080"`
		Context string `envconfig:"default=product-manager"`
	}
}
