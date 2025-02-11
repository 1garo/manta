package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Schema struct {
	Person []Person `json:"person"`
}

type Table struct {
	Rows  []interface{} `json:"rows"`
	Count int           `json:"count"`
}

// Method to append a row
func (t *Table) Append(row interface{}) {
	t.Rows = append(t.Rows, row)
	t.Count++ // Increase count to track the number of elements
}

type Database struct {
	filename string           `json:"-"`
	schema   Schema           `json:"-"`
	Tables   map[string]Table `json:"tables"`
}

func ReadDB(filename string) (Database, error) {
	dbFile, err := os.Open(filename)
	if err != nil {
		return Database{}, fmt.Errorf("could not find file: %v", err)
	}
	defer dbFile.Close()

	var data Database
	err = json.NewDecoder(dbFile).Decode(&data)
	if err != nil {
		return Database{}, fmt.Errorf("could not decode file content: %v", err)
	}

	return data, nil
}

func NewDatabase(filename string, schema Schema, tables map[string]Table) Database {
	return Database{
		filename: filename,
		schema:   schema,
		Tables:   tables,
	}
}

func (db *Database) tableExists(table string) bool {
	if _, ok := db.Tables[table]; ok {
		return true
	} else {
		return false
	}
}
func (db *Database) updateDB() error {
	fmt.Printf("updatedb: %+v\n", db.Tables)

	output, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return fmt.Errorf("cannot marshal data: %v", err)
	}
	if err := os.WriteFile(db.filename, output, 0644); err != nil {
		return fmt.Errorf("Error writing file: %v", err)
	}

	return nil
}

func (db *Database) CreateOne(table string, content any) error {
	if !db.tableExists(table) {
		return fmt.Errorf("table do not exist")
	}

	tables := db.Tables[table]
	tables.Append(content)
	fmt.Println("before tables: ", tables, db.Tables)
	db.Tables[table] = tables
	fmt.Println("after tables: ", tables, db.Tables)
	defer db.updateDB()

	return nil
}

func main() {
	// TODO: Do not leak internal specifics to the user

	var schema Schema
	filename := "db.json"
	values, err := ReadDB(filename)
	if err != nil {
		log.Fatalf("cannot read db: %v", err)
	}

	db := NewDatabase(filename, schema, values.Tables)
	fmt.Printf("%+v\n", db.Tables)
	person := Person{
		Name: "alexandre",
		Age:  10,
	}
	if err := db.CreateOne("person", person); err != nil {
		log.Fatalf("cannot create: %v", err)
	}
}
