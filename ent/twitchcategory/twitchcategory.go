// Code generated by ent, DO NOT EDIT.

package twitchcategory

import (
	"time"
)

const (
	// Label holds the string label denoting the twitchcategory type in the database.
	Label = "twitch_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBoxArtURL holds the string denoting the box_art_url field in the database.
	FieldBoxArtURL = "box_art_url"
	// FieldIgdbID holds the string denoting the igdb_id field in the database.
	FieldIgdbID = "igdb_id"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the twitchcategory in the database.
	Table = "twitch_categories"
)

// Columns holds all SQL columns for twitchcategory fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBoxArtURL,
	FieldIgdbID,
	FieldUpdatedAt,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)