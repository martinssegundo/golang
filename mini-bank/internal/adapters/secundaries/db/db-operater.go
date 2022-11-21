package db

type Dboperater[T any] interface {
	createConnection()
	close()
	save(t T) int
	update(t T)
	list() []T
}
