package database

//DBClientInterface interface to include database interection method
type DBClientInterface interface {
	DBConnect() error
}
