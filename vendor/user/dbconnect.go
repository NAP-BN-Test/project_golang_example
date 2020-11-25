package user

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

var server = "localhost"
var port = 1433
var user = "sa"
var password = "1234"
var database = "TX_EquipmentManager"
var (
	ctx context.Context
	db  *sql.DB
)

func connectdb(w http.ResponseWriter, r *http.Request) {
	var err error

	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Fprintf(w, "Error creating connection pool: "+err.Error())
	}
	fmt.Print("Connected!\n")
	// Close the database connection pool after program executes
	// defer db.Close()

	// Use background context
	ctx := context.Background()

	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT Username from dbo.Account where ID = 3").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
	// SelectVersion()
}

func SelectVersion() {
	// go run . để chạy toàn bộ file go
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	// Run query and scan for result
	err = db.QueryRowContext(ctx, "SELECT 1").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("%s\n", result)
}
