package models

//Table represents a database table
type Table struct {
	Name string `db:"table_name"`
}

//Tables is Table Slice
type Tables []Table
