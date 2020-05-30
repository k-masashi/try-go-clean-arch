package db

import "database/sql"

type SQLHandler struct {
	Connection *sql.DB
}

