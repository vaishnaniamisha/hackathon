package config

//ServerConfiguration to store server configuration
type ServerConfiguration struct {
	Port     string
	APIPath  string
	LogLevel string
}

//DBConfiguration to store database configuration
type DBConfiguration struct {
	DBHost     string
	DBName     string
	DBUserName string
	DBPassword string
	DBPort     string
}

//Configuration to read and store config file
type Configuration struct {
	Server   *ServerConfiguration
	Database *DBConfiguration
}
