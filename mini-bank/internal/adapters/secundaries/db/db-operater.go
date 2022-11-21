package db

type Dboperater interface {
	createConnection()
	close()
}
