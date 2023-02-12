package database

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	createTables(db)
}

func createTables(db *gorm.DB) {
	tables := []interface{}{}
	tables = addNewTable(db, &User{}, tables)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}
