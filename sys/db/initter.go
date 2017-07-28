package db

var dbInitStore []DBInitter

/*
Interface to init DB
 */
type DBInitter interface {
	DBInit() error
}

// Register DBInitter to init by command
func RegisterDBInitter(dbi DBInitter) {
	dbInitStore = append(dbInitStore, dbi)
}

// Init all registered DBInitters
func InitAllDBs() error {
	for _, initter := range dbInitStore {
		if err := initter.DBInit(); err != nil {
			return err
		}
	}
	return nil
}
