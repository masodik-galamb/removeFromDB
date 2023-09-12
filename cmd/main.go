package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const schema = `
CREATE TABLE IF NOT EXISTS person (
    first_name text,
    last_name text,
    email text,
	created_at timestamp null
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	CreatedAt string `db:"created_at"`
}

const driverName = "postgres"

func main() {
	db, err := connectToDB()
	if err != nil {
		log.Fatalln(err)
	}
	persons := createTestPersons()
	err = createTestData(persons, db)
	if err != nil {
		log.Fatalln(err)
	}

	// Remove from table 'person' requested value

	_, err = db.NamedExec(`DELETE FROM person WHERE created_at = :created_at;`, map[string]interface{}{
		"created_at": "2022-11-29 10:17:41",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Row is cleared")

}

func connectToDB() (*sqlx.DB, error) {
	dataSourceName := os.Getenv("DATA_SOURCE_NAME")

	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createTestData(persons []Person, db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	for _, s := range persons {

		_, err = tx.NamedExec("INSERT INTO person (first_name, last_name, email, created_at) VALUES (:first_name, :last_name, :email, :created_at)", &s)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func createTestPersons() []Person {

	return []Person{
		{FirstName: "Jason", LastName: "Moiron", Email: "jmoiron@jmoiron.net", CreatedAt: "2022-11-29 10:17:41"},
		{FirstName: "John", LastName: "Doe", Email: "johndoeDNE@gmail.net", CreatedAt: "2022-11-29 13:17:41"},
		{FirstName: "Jane", LastName: "Citizen", Email: "jane.citzen@example.com", CreatedAt: "2022-11-29 16:17:41"},
	}
}
