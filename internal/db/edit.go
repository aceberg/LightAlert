package db

import (
	"fmt"

	"github.com/aceberg/LightAlert/internal/models"
)

// Create - create table if not exists
func Create(path string) {

	sqlStatement := `CREATE TABLE IF NOT EXISTS records (
		"ID"		INTEGER PRIMARY KEY,
		"DATE"		TEXT,
		"NAME"		TEXT,
		"HASH"		TEXT,
		"IP"		TEXT,
		"AGENT"		TEXT
	);`
	exec(path, sqlStatement)
}

// Insert - insert one rec into DB
func Insert(path string, rec models.Record) {

	sqlStatement := `INSERT INTO records (DATE, NAME, HASH, IP, AGENT) 
	VALUES ('%s','%s','%s','%s','%s');`

	rec.Name = quoteStr(rec.Name)
	rec.Agent = quoteStr(rec.Agent)

	sqlStatement = fmt.Sprintf(sqlStatement, rec.Date, rec.Name, rec.Hash, rec.IP, rec.Agent)

	exec(path, sqlStatement)
}

// Delete - delete one record
func Delete(path string, id int) {

	sqlStatement := `DELETE FROM records WHERE ID='%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	exec(path, sqlStatement)
}

// Clear - delete all records from table
func Clear(path string) {
	sqlStatement := `DELETE FROM records;`
	exec(path, sqlStatement)
}
