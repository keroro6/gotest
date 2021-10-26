package interfacetest

type DB interface {
	Get(key string) (int, error)
	Add(key string, value int) error
}

func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key); err == nil {
		return v
	}
	return -1
}
