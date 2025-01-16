package repository

import "database/sql"

type ProductRepository struct {
	connection *sql.DB
}
