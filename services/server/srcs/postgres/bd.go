package my_postgres

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

const USERNAME = "postgres"
const PASSWORD = "foobar"
const URI = "db"
const PORT = "5432"
const DB = "daily"

// urlExample := "postgres://username:password@localhost:5432/database_name"
var DATABASE_URL string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", USERNAME, PASSWORD, URI, PORT, DB)

// If there's an error, panic
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// It connects to a database, checks if the day is 6, and if it is not, it gets the day solution and
// updates the day result table.
func Connection() string {
	fmt.Printf("Connecting to data base `%s`...\n", DB)
	db, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[-] Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("[+] Connected succesfully!")
	}
	defer db.Close(context.Background())

	daySolution := GetDaySolution(db)
	return daySolution
}

// It returns the day of the month of the last time the solution was updated
func GetDaySolution(db *pgx.Conn) string {
	type Row struct {
		id       int
		solution string
		dt       time.Time
	}
	r := Row{}
	_ = db.QueryRow(context.Background(), "SELECT * FROM day_solution WHERE ID=1").Scan(&r.id, &r.solution, &r.dt)
	if r.dt.Day() != time.Now().Day() {
		r.solution = GetNewDaySolution(db)
		UpdateDaySolutionTable(db, r.solution)
	}
	return r.solution
}

// It connects to the database, queries the table, and returns a random solution
func GetNewDaySolution(db *pgx.Conn) string {
	rows, err := db.Query(context.Background(), "SELECT * FROM equations")
	checkErr(err)
	type Row struct {
		id       int
		equation string
	}
	var rowSlice []Row
	for rows.Next() {
		r := Row{}
		err := rows.Scan(&r.id, &r.equation)
		checkErr(err)
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	nResults := len(rowSlice)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(nResults)
	return rowSlice[index].equation
}

// It updates the dayresult table with the solution of the day
func UpdateDaySolutionTable(db *pgx.Conn, daySolution string) {
	sqlStatement := fmt.Sprintf(`UPDATE day_solution SET solution='%s', dt=NOW() WHERE id=1`, daySolution)
	_, err := db.Exec(context.Background(), sqlStatement)
	if err != nil {
		fmt.Printf("[-] Table update failed due `%s`\n", err)
	} else {
		fmt.Println("[+] Table updated succesfully")
	}
}
