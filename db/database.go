package db

import "database/sql"

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "hai:Antihacking@1234@/faq")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
