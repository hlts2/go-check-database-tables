package models

import (
	"database/sql"
)

//Table represents a database table
type Table struct {
	Name string `db:"table_name"`
}

//Tables is Table slice
type Tables []Table

//DescribeTable represents a attributes of database table
type DescribeTable struct {
	Field   string         `db:"Field"`
	Type    string         `db:"Type"`
	Null    sql.NullString `db:"Null"`
	Key     sql.NullString `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra   sql.NullString `db:"Extra"`
}

//FieldName returns struct Field Name
func (d DescribeTable) FieldName() []string {
	return []string{"Field", "Type", "Null", "Key", "Default", "Extra"}
}

//FieldValue returns struct Field Value
func (d DescribeTable) FieldValue() []string {
	return []string{d.Field, d.Type, d.Null.String, d.Key.String, d.Default.String, d.Extra.String}
}

//DescribeTables is DescribeTable slice
type DescribeTables []DescribeTable
