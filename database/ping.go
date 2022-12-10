package database

func PingDB() (result bool, error error) {
	var (
		db     = Connect()
		status bool
	)

	// We use the Ping function to make a request to the database
	if error = db.Ping(); error != nil {
		status = false
	} else {
		status = true
	}
	return status, error
}
