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
	Field   string         `db:"field"`
	Type    string         `db:"type"`
	Null    string         `db:"null"`
	Key     sql.NullString `db:"key"`
	Default sql.NullString `db:"default"`
	Extra   sql.NullString `db:"extra"`
}

//FieldName returns struct Field Name
func (d DescribeTable) FieldName() []string {
	return []string{"Field", "Type", "Null", "Key", "Default", "Extra"}
}

//FieldValue returns struct Field Value
func (d DescribeTable) FieldValue() []string {
	return []string{d.Field, d.Type, d.Null, d.Key.String, d.Default.String, d.Extra.String}
}

//DescribeTables is DescribeTable slice
type DescribeTables []DescribeTable
